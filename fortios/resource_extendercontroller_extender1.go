// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Extender controller configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceExtenderControllerExtender1() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtenderControllerExtender1Create,
		Read:   resourceExtenderControllerExtender1Read,
		Update: resourceExtenderControllerExtender1Update,
		Delete: resourceExtenderControllerExtender1Delete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Required:     true,
				ForceNew:     true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"authorized": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ext_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"login_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"controller_report": &schema.Schema{
				Type:     schema.TypeList,
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
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"modem1": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ifname": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
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
							Computed:     true,
						},
						"auto_switch": &schema.Schema{
							Type:     schema.TypeList,
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
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
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
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ifname": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
							Computed:     true,
						},
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
							Computed:     true,
						},
						"auto_switch": &schema.Schema{
							Type:     schema.TypeList,
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
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
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
	}
}

func resourceExtenderControllerExtender1Create(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectExtenderControllerExtender1(d)
	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtender1 resource while getting object: %v", err)
	}

	o, err := c.CreateExtenderControllerExtender1(obj)

	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtender1 resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerExtender1")
	}

	return resourceExtenderControllerExtender1Read(d, m)
}

func resourceExtenderControllerExtender1Update(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectExtenderControllerExtender1(d)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtender1 resource while getting object: %v", err)
	}

	o, err := c.UpdateExtenderControllerExtender1(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtender1 resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerExtender1")
	}

	return resourceExtenderControllerExtender1Read(d, m)
}

func resourceExtenderControllerExtender1Delete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteExtenderControllerExtender1(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting ExtenderControllerExtender1 resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerExtender1Read(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadExtenderControllerExtender1(mkey)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtender1 resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtenderControllerExtender1(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtender1 resource from API: %v", err)
	}
	return nil
}

func flattenExtenderControllerExtender1Name(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Id(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Authorized(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1ExtName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Description(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Vdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1LoginPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1ControllerReport(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenExtenderControllerExtender1ControllerReportStatus(i["status"], d, pre_append)
	}

	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {
		result["interval"] = flattenExtenderControllerExtender1ControllerReportInterval(i["interval"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {
		result["signal_threshold"] = flattenExtenderControllerExtender1ControllerReportSignalThreshold(i["signal-threshold"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtender1ControllerReportStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1ControllerReportInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1ControllerReportSignalThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ifname"
	if _, ok := i["ifname"]; ok {
		result["ifname"] = flattenExtenderControllerExtender1Modem1Ifname(i["ifname"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtenderControllerExtender1Modem1RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtenderControllerExtender1Modem1RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtenderControllerExtender1Modem1ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtenderControllerExtender1Modem1DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtenderControllerExtender1Modem1Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtenderControllerExtender1Modem1Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtenderControllerExtender1Modem1Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		result["sim1_pin_code"] = flattenExtenderControllerExtender1Modem1Sim1PinCode(i["sim1-pin-code"], d, pre_append)
		c := d.Get(pre_append).(string)
		if c != "" {
			result["sim1_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenExtenderControllerExtender1Modem1Sim2PinCode(i["sim2-pin-code"], d, pre_append)
		c := d.Get(pre_append).(string)
		if c != "" {
			result["sim2_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtenderControllerExtender1Modem1PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtenderControllerExtender1Modem1AutoSwitch(i["auto-switch"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtender1Modem1Ifname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1Sim1PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1Sim2PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtenderControllerExtender1Modem1AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtenderControllerExtender1Modem1AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtenderControllerExtender1Modem1AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtenderControllerExtender1Modem1AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtenderControllerExtender1Modem1AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtenderControllerExtender1Modem1AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtenderControllerExtender1Modem1AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtenderControllerExtender1Modem1AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtender1Modem1AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem1AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ifname"
	if _, ok := i["ifname"]; ok {
		result["ifname"] = flattenExtenderControllerExtender1Modem2Ifname(i["ifname"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {
		result["redundant_mode"] = flattenExtenderControllerExtender1Modem2RedundantMode(i["redundant-mode"], d, pre_append)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {
		result["redundant_intf"] = flattenExtenderControllerExtender1Modem2RedundantIntf(i["redundant-intf"], d, pre_append)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {
		result["conn_status"] = flattenExtenderControllerExtender1Modem2ConnStatus(i["conn-status"], d, pre_append)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {
		result["default_sim"] = flattenExtenderControllerExtender1Modem2DefaultSim(i["default-sim"], d, pre_append)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {
		result["gps"] = flattenExtenderControllerExtender1Modem2Gps(i["gps"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {
		result["sim1_pin"] = flattenExtenderControllerExtender1Modem2Sim1Pin(i["sim1-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {
		result["sim2_pin"] = flattenExtenderControllerExtender1Modem2Sim2Pin(i["sim2-pin"], d, pre_append)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {
		result["sim1_pin_code"] = flattenExtenderControllerExtender1Modem2Sim1PinCode(i["sim1-pin-code"], d, pre_append)
		c := d.Get(pre_append).(string)
		if c != "" {
			result["sim1_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {
		result["sim2_pin_code"] = flattenExtenderControllerExtender1Modem2Sim2PinCode(i["sim2-pin-code"], d, pre_append)
		c := d.Get(pre_append).(string)
		if c != "" {
			result["sim2_pin_code"] = c
		}
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {
		result["preferred_carrier"] = flattenExtenderControllerExtender1Modem2PreferredCarrier(i["preferred-carrier"], d, pre_append)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {
		result["auto_switch"] = flattenExtenderControllerExtender1Modem2AutoSwitch(i["auto-switch"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtender1Modem2Ifname(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2RedundantMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2RedundantIntf(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2ConnStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2DefaultSim(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2Gps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2Sim1Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2Sim2Pin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2Sim1PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2Sim2PinCode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2PreferredCarrier(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2AutoSwitch(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {
		result["disconnect"] = flattenExtenderControllerExtender1Modem2AutoSwitchDisconnect(i["disconnect"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {
		result["disconnect_threshold"] = flattenExtenderControllerExtender1Modem2AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {
		result["disconnect_period"] = flattenExtenderControllerExtender1Modem2AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {
		result["signal"] = flattenExtenderControllerExtender1Modem2AutoSwitchSignal(i["signal"], d, pre_append)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {
		result["dataplan"] = flattenExtenderControllerExtender1Modem2AutoSwitchDataplan(i["dataplan"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {
		result["switch_back"] = flattenExtenderControllerExtender1Modem2AutoSwitchSwitchBack(i["switch-back"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {
		result["switch_back_time"] = flattenExtenderControllerExtender1Modem2AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {
		result["switch_back_timer"] = flattenExtenderControllerExtender1Modem2AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtender1Modem2AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenExtenderControllerExtender1Modem2AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectExtenderControllerExtender1(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenExtenderControllerExtender1Name(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtenderControllerExtender1Id(o["id"], d, "fosid")); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("authorized", flattenExtenderControllerExtender1Authorized(o["authorized"], d, "authorized")); err != nil {
		if !fortiAPIPatch(o["authorized"]) {
			return fmt.Errorf("Error reading authorized: %v", err)
		}
	}

	if err = d.Set("ext_name", flattenExtenderControllerExtender1ExtName(o["ext-name"], d, "ext_name")); err != nil {
		if !fortiAPIPatch(o["ext-name"]) {
			return fmt.Errorf("Error reading ext_name: %v", err)
		}
	}

	if err = d.Set("description", flattenExtenderControllerExtender1Description(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("vdom", flattenExtenderControllerExtender1Vdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("controller_report", flattenExtenderControllerExtender1ControllerReport(o["controller-report"], d, "controller_report")); err != nil {
			if !fortiAPIPatch(o["controller-report"]) {
				return fmt.Errorf("Error reading controller_report: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("controller_report"); ok {
			if err = d.Set("controller_report", flattenExtenderControllerExtender1ControllerReport(o["controller-report"], d, "controller_report")); err != nil {
				if !fortiAPIPatch(o["controller-report"]) {
					return fmt.Errorf("Error reading controller_report: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("modem1", flattenExtenderControllerExtender1Modem1(o["modem1"], d, "modem1")); err != nil {
			if !fortiAPIPatch(o["modem1"]) {
				return fmt.Errorf("Error reading modem1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem1"); ok {
			if err = d.Set("modem1", flattenExtenderControllerExtender1Modem1(o["modem1"], d, "modem1")); err != nil {
				if !fortiAPIPatch(o["modem1"]) {
					return fmt.Errorf("Error reading modem1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("modem2", flattenExtenderControllerExtender1Modem2(o["modem2"], d, "modem2")); err != nil {
			if !fortiAPIPatch(o["modem2"]) {
				return fmt.Errorf("Error reading modem2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem2"); ok {
			if err = d.Set("modem2", flattenExtenderControllerExtender1Modem2(o["modem2"], d, "modem2")); err != nil {
				if !fortiAPIPatch(o["modem2"]) {
					return fmt.Errorf("Error reading modem2: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenExtenderControllerExtender1FortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandExtenderControllerExtender1Name(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Id(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Authorized(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1ExtName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Description(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Vdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1LoginPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1ControllerReport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandExtenderControllerExtender1ControllerReportStatus(d, i["status"], pre_append)
	}
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok {
		result["interval"], _ = expandExtenderControllerExtender1ControllerReportInterval(d, i["interval"], pre_append)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["signal-threshold"], _ = expandExtenderControllerExtender1ControllerReportSignalThreshold(d, i["signal_threshold"], pre_append)
	}

	return result, nil
}

func expandExtenderControllerExtender1ControllerReportStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1ControllerReportInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1ControllerReportSignalThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ifname"
	if _, ok := d.GetOk(pre_append); ok {
		result["ifname"], _ = expandExtenderControllerExtender1Modem1Ifname(d, i["ifname"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-mode"], _ = expandExtenderControllerExtender1Modem1RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-intf"], _ = expandExtenderControllerExtender1Modem1RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok {
		result["conn-status"], _ = expandExtenderControllerExtender1Modem1ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok {
		result["default-sim"], _ = expandExtenderControllerExtender1Modem1DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok {
		result["gps"], _ = expandExtenderControllerExtender1Modem1Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin"], _ = expandExtenderControllerExtender1Modem1Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin"], _ = expandExtenderControllerExtender1Modem1Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin-code"], _ = expandExtenderControllerExtender1Modem1Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin-code"], _ = expandExtenderControllerExtender1Modem1Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok {
		result["preferred-carrier"], _ = expandExtenderControllerExtender1Modem1PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-switch"], _ = expandExtenderControllerExtender1Modem1AutoSwitch(d, i["auto_switch"], pre_append)
	} else {
		result["auto-switch"] = make([]string, 0)
	}

	return result, nil
}

func expandExtenderControllerExtender1Modem1Ifname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect"], _ = expandExtenderControllerExtender1Modem1AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-threshold"], _ = expandExtenderControllerExtender1Modem1AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-period"], _ = expandExtenderControllerExtender1Modem1AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok {
		result["signal"], _ = expandExtenderControllerExtender1Modem1AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["dataplan"], _ = expandExtenderControllerExtender1Modem1AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back"], _ = expandExtenderControllerExtender1Modem1AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-time"], _ = expandExtenderControllerExtender1Modem1AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-timer"], _ = expandExtenderControllerExtender1Modem1AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandExtenderControllerExtender1Modem1AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem1AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ifname"
	if _, ok := d.GetOk(pre_append); ok {
		result["ifname"], _ = expandExtenderControllerExtender1Modem2Ifname(d, i["ifname"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-mode"], _ = expandExtenderControllerExtender1Modem2RedundantMode(d, i["redundant_mode"], pre_append)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok {
		result["redundant-intf"], _ = expandExtenderControllerExtender1Modem2RedundantIntf(d, i["redundant_intf"], pre_append)
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok {
		result["conn-status"], _ = expandExtenderControllerExtender1Modem2ConnStatus(d, i["conn_status"], pre_append)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok {
		result["default-sim"], _ = expandExtenderControllerExtender1Modem2DefaultSim(d, i["default_sim"], pre_append)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok {
		result["gps"], _ = expandExtenderControllerExtender1Modem2Gps(d, i["gps"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin"], _ = expandExtenderControllerExtender1Modem2Sim1Pin(d, i["sim1_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin"], _ = expandExtenderControllerExtender1Modem2Sim2Pin(d, i["sim2_pin"], pre_append)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim1-pin-code"], _ = expandExtenderControllerExtender1Modem2Sim1PinCode(d, i["sim1_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok {
		result["sim2-pin-code"], _ = expandExtenderControllerExtender1Modem2Sim2PinCode(d, i["sim2_pin_code"], pre_append)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok {
		result["preferred-carrier"], _ = expandExtenderControllerExtender1Modem2PreferredCarrier(d, i["preferred_carrier"], pre_append)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok {
		result["auto-switch"], _ = expandExtenderControllerExtender1Modem2AutoSwitch(d, i["auto_switch"], pre_append)
	} else {
		result["auto-switch"] = make([]string, 0)
	}

	return result, nil
}

func expandExtenderControllerExtender1Modem2Ifname(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2RedundantMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2RedundantIntf(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2ConnStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2DefaultSim(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2Gps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2Sim1Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2Sim2Pin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2Sim1PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2Sim2PinCode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2PreferredCarrier(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect"], _ = expandExtenderControllerExtender1Modem2AutoSwitchDisconnect(d, i["disconnect"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-threshold"], _ = expandExtenderControllerExtender1Modem2AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok {
		result["disconnect-period"], _ = expandExtenderControllerExtender1Modem2AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok {
		result["signal"], _ = expandExtenderControllerExtender1Modem2AutoSwitchSignal(d, i["signal"], pre_append)
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {
		result["dataplan"], _ = expandExtenderControllerExtender1Modem2AutoSwitchDataplan(d, i["dataplan"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back"], _ = expandExtenderControllerExtender1Modem2AutoSwitchSwitchBack(d, i["switch_back"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-time"], _ = expandExtenderControllerExtender1Modem2AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok {
		result["switch-back-timer"], _ = expandExtenderControllerExtender1Modem2AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append)
	}

	return result, nil
}

func expandExtenderControllerExtender1Modem2AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtender1Modem2AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectExtenderControllerExtender1(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandExtenderControllerExtender1Name(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandExtenderControllerExtender1Id(d, v, "fosid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("authorized"); ok {
		t, err := expandExtenderControllerExtender1Authorized(d, v, "authorized")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized"] = t
		}
	}

	if v, ok := d.GetOk("ext_name"); ok {
		t, err := expandExtenderControllerExtender1ExtName(d, v, "ext_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandExtenderControllerExtender1Description(d, v, "description")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOkExists("vdom"); ok {
		t, err := expandExtenderControllerExtender1Vdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok {
		t, err := expandExtenderControllerExtender1LoginPassword(d, v, "login_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	}

	if v, ok := d.GetOk("controller_report"); ok {
		t, err := expandExtenderControllerExtender1ControllerReport(d, v, "controller_report")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["controller-report"] = t
		}
	}

	if v, ok := d.GetOk("modem1"); ok {
		t, err := expandExtenderControllerExtender1Modem1(d, v, "modem1")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem1"] = t
		}
	}

	if v, ok := d.GetOk("modem2"); ok {
		t, err := expandExtenderControllerExtender1Modem2(d, v, "modem2")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem2"] = t
		}
	}

	return &obj, nil
}
