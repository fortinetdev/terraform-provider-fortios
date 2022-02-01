// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Extender controller configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceExtenderControllerExtender() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtenderControllerExtenderCreate,
		Read:   resourceExtenderControllerExtenderRead,
		Update: resourceExtenderControllerExtenderUpdate,
		Delete: resourceExtenderControllerExtenderDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				ForceNew:     true,
				Required:     true,
			},
			"authorized": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"admin": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ifname": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"extension_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"override_login_password_change": &schema.Schema{
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
			"override_enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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
			"wan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"modem1_extension": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"modem2_extension": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
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
			"role": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dial_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"redial": &schema.Schema{
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
			"dial_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"conn_status": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
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
			"quota_limit_mb": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 10485760),
				Optional:     true,
				Computed:     true,
			},
			"billing_start_day": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 28),
				Optional:     true,
				Computed:     true,
			},
			"at_dial_script": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"modem_passwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"initiated_update": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"modem_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"ppp_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"ppp_auth_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ppp_echo_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"wimax_carrier": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"wimax_realm": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"wimax_auth_protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"sim_pin": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"access_point_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"multi_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roaming": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cdma_nai": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"aaa_shared_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"ha_shared_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
				Sensitive:    true,
			},
			"primary_ha": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"secondary_ha": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"cdma_aaa_spi": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"cdma_ha_spi": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceExtenderControllerExtenderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectExtenderControllerExtender(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtender resource while getting object: %v", err)
	}

	o, err := c.CreateExtenderControllerExtender(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ExtenderControllerExtender resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerExtender")
	}

	return resourceExtenderControllerExtenderRead(d, m)
}

func resourceExtenderControllerExtenderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectExtenderControllerExtender(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtender resource while getting object: %v", err)
	}

	o, err := c.UpdateExtenderControllerExtender(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ExtenderControllerExtender resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtenderControllerExtender")
	}

	return resourceExtenderControllerExtenderRead(d, m)
}

func resourceExtenderControllerExtenderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteExtenderControllerExtender(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ExtenderControllerExtender resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtenderControllerExtenderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadExtenderControllerExtender(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtender resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtenderControllerExtender(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ExtenderControllerExtender resource from API: %v", err)
	}
	return nil
}

func flattenExtenderControllerExtenderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAuthorized(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAdmin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderIfname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDeviceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderExtensionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderOverrideAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderOverrideLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderLoginPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderOverrideEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderBandwidthLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWanExtension(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "modem1_extension"
	if _, ok := i["modem1-extension"]; ok {

		result["modem1_extension"] = flattenExtenderControllerExtenderWanExtensionModem1Extension(i["modem1-extension"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "modem2_extension"
	if _, ok := i["modem2-extension"]; ok {

		result["modem2_extension"] = flattenExtenderControllerExtenderWanExtensionModem2Extension(i["modem2-extension"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderWanExtensionModem1Extension(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWanExtensionModem2Extension(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderControllerReport(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {

		result["status"] = flattenExtenderControllerExtenderControllerReportStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "interval"
	if _, ok := i["interval"]; ok {

		result["interval"] = flattenExtenderControllerExtenderControllerReportInterval(i["interval"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := i["signal-threshold"]; ok {

		result["signal_threshold"] = flattenExtenderControllerExtenderControllerReportSignalThreshold(i["signal-threshold"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderControllerReportStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderControllerReportInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderControllerReportSignalThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ifname"
	if _, ok := i["ifname"]; ok {

		result["ifname"] = flattenExtenderControllerExtenderModem1Ifname(i["ifname"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {

		result["redundant_mode"] = flattenExtenderControllerExtenderModem1RedundantMode(i["redundant-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {

		result["redundant_intf"] = flattenExtenderControllerExtenderModem1RedundantIntf(i["redundant-intf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {

		result["conn_status"] = flattenExtenderControllerExtenderModem1ConnStatus(i["conn-status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {

		result["default_sim"] = flattenExtenderControllerExtenderModem1DefaultSim(i["default-sim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {

		result["gps"] = flattenExtenderControllerExtenderModem1Gps(i["gps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {

		result["sim1_pin"] = flattenExtenderControllerExtenderModem1Sim1Pin(i["sim1-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {

		result["sim2_pin"] = flattenExtenderControllerExtenderModem1Sim2Pin(i["sim2-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {

		result["sim1_pin_code"] = flattenExtenderControllerExtenderModem1Sim1PinCode(i["sim1-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {

		result["sim2_pin_code"] = flattenExtenderControllerExtenderModem1Sim2PinCode(i["sim2-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {

		result["preferred_carrier"] = flattenExtenderControllerExtenderModem1PreferredCarrier(i["preferred-carrier"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {

		result["auto_switch"] = flattenExtenderControllerExtenderModem1AutoSwitch(i["auto-switch"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderModem1Ifname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1RedundantMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1RedundantIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1ConnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1DefaultSim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Gps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Sim1Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Sim2Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Sim1PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1Sim2PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1PreferredCarrier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {

		result["disconnect"] = flattenExtenderControllerExtenderModem1AutoSwitchDisconnect(i["disconnect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {

		result["disconnect_threshold"] = flattenExtenderControllerExtenderModem1AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {

		result["disconnect_period"] = flattenExtenderControllerExtenderModem1AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {

		result["signal"] = flattenExtenderControllerExtenderModem1AutoSwitchSignal(i["signal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {

		result["dataplan"] = flattenExtenderControllerExtenderModem1AutoSwitchDataplan(i["dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {

		result["switch_back"] = flattenExtenderControllerExtenderModem1AutoSwitchSwitchBack(i["switch-back"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {

		result["switch_back_time"] = flattenExtenderControllerExtenderModem1AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {

		result["switch_back_timer"] = flattenExtenderControllerExtenderModem1AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderModem1AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem1AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ifname"
	if _, ok := i["ifname"]; ok {

		result["ifname"] = flattenExtenderControllerExtenderModem2Ifname(i["ifname"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := i["redundant-mode"]; ok {

		result["redundant_mode"] = flattenExtenderControllerExtenderModem2RedundantMode(i["redundant-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := i["redundant-intf"]; ok {

		result["redundant_intf"] = flattenExtenderControllerExtenderModem2RedundantIntf(i["redundant-intf"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "conn_status"
	if _, ok := i["conn-status"]; ok {

		result["conn_status"] = flattenExtenderControllerExtenderModem2ConnStatus(i["conn-status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "default_sim"
	if _, ok := i["default-sim"]; ok {

		result["default_sim"] = flattenExtenderControllerExtenderModem2DefaultSim(i["default-sim"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "gps"
	if _, ok := i["gps"]; ok {

		result["gps"] = flattenExtenderControllerExtenderModem2Gps(i["gps"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := i["sim1-pin"]; ok {

		result["sim1_pin"] = flattenExtenderControllerExtenderModem2Sim1Pin(i["sim1-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := i["sim2-pin"]; ok {

		result["sim2_pin"] = flattenExtenderControllerExtenderModem2Sim2Pin(i["sim2-pin"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := i["sim1-pin-code"]; ok {

		result["sim1_pin_code"] = flattenExtenderControllerExtenderModem2Sim1PinCode(i["sim1-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := i["sim2-pin-code"]; ok {

		result["sim2_pin_code"] = flattenExtenderControllerExtenderModem2Sim2PinCode(i["sim2-pin-code"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := i["preferred-carrier"]; ok {

		result["preferred_carrier"] = flattenExtenderControllerExtenderModem2PreferredCarrier(i["preferred-carrier"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auto_switch"
	if _, ok := i["auto-switch"]; ok {

		result["auto_switch"] = flattenExtenderControllerExtenderModem2AutoSwitch(i["auto-switch"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderModem2Ifname(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2RedundantMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2RedundantIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2ConnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2DefaultSim(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Gps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Sim1Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Sim2Pin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Sim1PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2Sim2PinCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2PreferredCarrier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitch(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := i["disconnect"]; ok {

		result["disconnect"] = flattenExtenderControllerExtenderModem2AutoSwitchDisconnect(i["disconnect"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := i["disconnect-threshold"]; ok {

		result["disconnect_threshold"] = flattenExtenderControllerExtenderModem2AutoSwitchDisconnectThreshold(i["disconnect-threshold"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := i["disconnect-period"]; ok {

		result["disconnect_period"] = flattenExtenderControllerExtenderModem2AutoSwitchDisconnectPeriod(i["disconnect-period"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "signal"
	if _, ok := i["signal"]; ok {

		result["signal"] = flattenExtenderControllerExtenderModem2AutoSwitchSignal(i["signal"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "dataplan"
	if _, ok := i["dataplan"]; ok {

		result["dataplan"] = flattenExtenderControllerExtenderModem2AutoSwitchDataplan(i["dataplan"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back"
	if _, ok := i["switch-back"]; ok {

		result["switch_back"] = flattenExtenderControllerExtenderModem2AutoSwitchSwitchBack(i["switch-back"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := i["switch-back-time"]; ok {

		result["switch_back_time"] = flattenExtenderControllerExtenderModem2AutoSwitchSwitchBackTime(i["switch-back-time"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := i["switch-back-timer"]; ok {

		result["switch_back_timer"] = flattenExtenderControllerExtenderModem2AutoSwitchSwitchBackTimer(i["switch-back-timer"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtenderControllerExtenderModem2AutoSwitchDisconnect(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchDisconnectThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchDisconnectPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchSignal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchDataplan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchSwitchBack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchSwitchBackTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModem2AutoSwitchSwitchBackTimer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRole(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDialMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRedial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRedundantIntf(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDialStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderConnStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderExtName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderQuotaLimitMb(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderBillingStartDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAtDialScript(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModemPasswd(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderInitiatedUpdate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderModemType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppAuthProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPppEchoRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWimaxCarrier(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWimaxRealm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderWimaxAuthProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderSimPin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAccessPointName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderMultiMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderRoaming(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderCdmaNai(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderAaaSharedSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderHaSharedSecret(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderPrimaryHa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderSecondaryHa(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderCdmaAaaSpi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtenderControllerExtenderCdmaHaSpi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectExtenderControllerExtender(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenExtenderControllerExtenderName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtenderControllerExtenderId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("authorized", flattenExtenderControllerExtenderAuthorized(o["authorized"], d, "authorized", sv)); err != nil {
		if !fortiAPIPatch(o["authorized"]) {
			return fmt.Errorf("Error reading authorized: %v", err)
		}
	}

	if err = d.Set("admin", flattenExtenderControllerExtenderAdmin(o["admin"], d, "admin", sv)); err != nil {
		if !fortiAPIPatch(o["admin"]) {
			return fmt.Errorf("Error reading admin: %v", err)
		}
	}

	if err = d.Set("ifname", flattenExtenderControllerExtenderIfname(o["ifname"], d, "ifname", sv)); err != nil {
		if !fortiAPIPatch(o["ifname"]) {
			return fmt.Errorf("Error reading ifname: %v", err)
		}
	}

	if err = d.Set("vdom", flattenExtenderControllerExtenderVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("device_id", flattenExtenderControllerExtenderDeviceId(o["device-id"], d, "device_id", sv)); err != nil {
		if !fortiAPIPatch(o["device-id"]) {
			return fmt.Errorf("Error reading device_id: %v", err)
		}
	}

	if err = d.Set("extension_type", flattenExtenderControllerExtenderExtensionType(o["extension-type"], d, "extension_type", sv)); err != nil {
		if !fortiAPIPatch(o["extension-type"]) {
			return fmt.Errorf("Error reading extension_type: %v", err)
		}
	}

	if err = d.Set("override_allowaccess", flattenExtenderControllerExtenderOverrideAllowaccess(o["override-allowaccess"], d, "override_allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["override-allowaccess"]) {
			return fmt.Errorf("Error reading override_allowaccess: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenExtenderControllerExtenderAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("override_login_password_change", flattenExtenderControllerExtenderOverrideLoginPasswordChange(o["override-login-password-change"], d, "override_login_password_change", sv)); err != nil {
		if !fortiAPIPatch(o["override-login-password-change"]) {
			return fmt.Errorf("Error reading override_login_password_change: %v", err)
		}
	}

	if err = d.Set("login_password_change", flattenExtenderControllerExtenderLoginPasswordChange(o["login-password-change"], d, "login_password_change", sv)); err != nil {
		if !fortiAPIPatch(o["login-password-change"]) {
			return fmt.Errorf("Error reading login_password_change: %v", err)
		}
	}

	if err = d.Set("login_password", flattenExtenderControllerExtenderLoginPassword(o["login-password"], d, "login_password", sv)); err != nil {
		if !fortiAPIPatch(o["login-password"]) {
			return fmt.Errorf("Error reading login_password: %v", err)
		}
	}

	if err = d.Set("override_enforce_bandwidth", flattenExtenderControllerExtenderOverrideEnforceBandwidth(o["override-enforce-bandwidth"], d, "override_enforce_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["override-enforce-bandwidth"]) {
			return fmt.Errorf("Error reading override_enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("enforce_bandwidth", flattenExtenderControllerExtenderEnforceBandwidth(o["enforce-bandwidth"], d, "enforce_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-bandwidth"]) {
			return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_limit", flattenExtenderControllerExtenderBandwidthLimit(o["bandwidth-limit"], d, "bandwidth_limit", sv)); err != nil {
		if !fortiAPIPatch(o["bandwidth-limit"]) {
			return fmt.Errorf("Error reading bandwidth_limit: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("wan_extension", flattenExtenderControllerExtenderWanExtension(o["wan-extension"], d, "wan_extension", sv)); err != nil {
			if !fortiAPIPatch(o["wan-extension"]) {
				return fmt.Errorf("Error reading wan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wan_extension"); ok {
			if err = d.Set("wan_extension", flattenExtenderControllerExtenderWanExtension(o["wan-extension"], d, "wan_extension", sv)); err != nil {
				if !fortiAPIPatch(o["wan-extension"]) {
					return fmt.Errorf("Error reading wan_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("profile", flattenExtenderControllerExtenderProfile(o["profile"], d, "profile", sv)); err != nil {
		if !fortiAPIPatch(o["profile"]) {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("controller_report", flattenExtenderControllerExtenderControllerReport(o["controller-report"], d, "controller_report", sv)); err != nil {
			if !fortiAPIPatch(o["controller-report"]) {
				return fmt.Errorf("Error reading controller_report: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("controller_report"); ok {
			if err = d.Set("controller_report", flattenExtenderControllerExtenderControllerReport(o["controller-report"], d, "controller_report", sv)); err != nil {
				if !fortiAPIPatch(o["controller-report"]) {
					return fmt.Errorf("Error reading controller_report: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("modem1", flattenExtenderControllerExtenderModem1(o["modem1"], d, "modem1", sv)); err != nil {
			if !fortiAPIPatch(o["modem1"]) {
				return fmt.Errorf("Error reading modem1: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem1"); ok {
			if err = d.Set("modem1", flattenExtenderControllerExtenderModem1(o["modem1"], d, "modem1", sv)); err != nil {
				if !fortiAPIPatch(o["modem1"]) {
					return fmt.Errorf("Error reading modem1: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("modem2", flattenExtenderControllerExtenderModem2(o["modem2"], d, "modem2", sv)); err != nil {
			if !fortiAPIPatch(o["modem2"]) {
				return fmt.Errorf("Error reading modem2: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("modem2"); ok {
			if err = d.Set("modem2", flattenExtenderControllerExtenderModem2(o["modem2"], d, "modem2", sv)); err != nil {
				if !fortiAPIPatch(o["modem2"]) {
					return fmt.Errorf("Error reading modem2: %v", err)
				}
			}
		}
	}

	if err = d.Set("role", flattenExtenderControllerExtenderRole(o["role"], d, "role", sv)); err != nil {
		if !fortiAPIPatch(o["role"]) {
			return fmt.Errorf("Error reading role: %v", err)
		}
	}

	if err = d.Set("mode", flattenExtenderControllerExtenderMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("dial_mode", flattenExtenderControllerExtenderDialMode(o["dial-mode"], d, "dial_mode", sv)); err != nil {
		if !fortiAPIPatch(o["dial-mode"]) {
			return fmt.Errorf("Error reading dial_mode: %v", err)
		}
	}

	if err = d.Set("redial", flattenExtenderControllerExtenderRedial(o["redial"], d, "redial", sv)); err != nil {
		if !fortiAPIPatch(o["redial"]) {
			return fmt.Errorf("Error reading redial: %v", err)
		}
	}

	if err = d.Set("redundant_intf", flattenExtenderControllerExtenderRedundantIntf(o["redundant-intf"], d, "redundant_intf", sv)); err != nil {
		if !fortiAPIPatch(o["redundant-intf"]) {
			return fmt.Errorf("Error reading redundant_intf: %v", err)
		}
	}

	if err = d.Set("dial_status", flattenExtenderControllerExtenderDialStatus(o["dial-status"], d, "dial_status", sv)); err != nil {
		if !fortiAPIPatch(o["dial-status"]) {
			return fmt.Errorf("Error reading dial_status: %v", err)
		}
	}

	if err = d.Set("conn_status", flattenExtenderControllerExtenderConnStatus(o["conn-status"], d, "conn_status", sv)); err != nil {
		if !fortiAPIPatch(o["conn-status"]) {
			return fmt.Errorf("Error reading conn_status: %v", err)
		}
	}

	if err = d.Set("ext_name", flattenExtenderControllerExtenderExtName(o["ext-name"], d, "ext_name", sv)); err != nil {
		if !fortiAPIPatch(o["ext-name"]) {
			return fmt.Errorf("Error reading ext_name: %v", err)
		}
	}

	if err = d.Set("description", flattenExtenderControllerExtenderDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("quota_limit_mb", flattenExtenderControllerExtenderQuotaLimitMb(o["quota-limit-mb"], d, "quota_limit_mb", sv)); err != nil {
		if !fortiAPIPatch(o["quota-limit-mb"]) {
			return fmt.Errorf("Error reading quota_limit_mb: %v", err)
		}
	}

	if err = d.Set("billing_start_day", flattenExtenderControllerExtenderBillingStartDay(o["billing-start-day"], d, "billing_start_day", sv)); err != nil {
		if !fortiAPIPatch(o["billing-start-day"]) {
			return fmt.Errorf("Error reading billing_start_day: %v", err)
		}
	}

	if err = d.Set("at_dial_script", flattenExtenderControllerExtenderAtDialScript(o["at-dial-script"], d, "at_dial_script", sv)); err != nil {
		if !fortiAPIPatch(o["at-dial-script"]) {
			return fmt.Errorf("Error reading at_dial_script: %v", err)
		}
	}

	if err = d.Set("initiated_update", flattenExtenderControllerExtenderInitiatedUpdate(o["initiated-update"], d, "initiated_update", sv)); err != nil {
		if !fortiAPIPatch(o["initiated-update"]) {
			return fmt.Errorf("Error reading initiated_update: %v", err)
		}
	}

	if err = d.Set("modem_type", flattenExtenderControllerExtenderModemType(o["modem-type"], d, "modem_type", sv)); err != nil {
		if !fortiAPIPatch(o["modem-type"]) {
			return fmt.Errorf("Error reading modem_type: %v", err)
		}
	}

	if err = d.Set("ppp_username", flattenExtenderControllerExtenderPppUsername(o["ppp-username"], d, "ppp_username", sv)); err != nil {
		if !fortiAPIPatch(o["ppp-username"]) {
			return fmt.Errorf("Error reading ppp_username: %v", err)
		}
	}

	if err = d.Set("ppp_auth_protocol", flattenExtenderControllerExtenderPppAuthProtocol(o["ppp-auth-protocol"], d, "ppp_auth_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["ppp-auth-protocol"]) {
			return fmt.Errorf("Error reading ppp_auth_protocol: %v", err)
		}
	}

	if err = d.Set("ppp_echo_request", flattenExtenderControllerExtenderPppEchoRequest(o["ppp-echo-request"], d, "ppp_echo_request", sv)); err != nil {
		if !fortiAPIPatch(o["ppp-echo-request"]) {
			return fmt.Errorf("Error reading ppp_echo_request: %v", err)
		}
	}

	if err = d.Set("wimax_carrier", flattenExtenderControllerExtenderWimaxCarrier(o["wimax-carrier"], d, "wimax_carrier", sv)); err != nil {
		if !fortiAPIPatch(o["wimax-carrier"]) {
			return fmt.Errorf("Error reading wimax_carrier: %v", err)
		}
	}

	if err = d.Set("wimax_realm", flattenExtenderControllerExtenderWimaxRealm(o["wimax-realm"], d, "wimax_realm", sv)); err != nil {
		if !fortiAPIPatch(o["wimax-realm"]) {
			return fmt.Errorf("Error reading wimax_realm: %v", err)
		}
	}

	if err = d.Set("wimax_auth_protocol", flattenExtenderControllerExtenderWimaxAuthProtocol(o["wimax-auth-protocol"], d, "wimax_auth_protocol", sv)); err != nil {
		if !fortiAPIPatch(o["wimax-auth-protocol"]) {
			return fmt.Errorf("Error reading wimax_auth_protocol: %v", err)
		}
	}

	if err = d.Set("access_point_name", flattenExtenderControllerExtenderAccessPointName(o["access-point-name"], d, "access_point_name", sv)); err != nil {
		if !fortiAPIPatch(o["access-point-name"]) {
			return fmt.Errorf("Error reading access_point_name: %v", err)
		}
	}

	if err = d.Set("multi_mode", flattenExtenderControllerExtenderMultiMode(o["multi-mode"], d, "multi_mode", sv)); err != nil {
		if !fortiAPIPatch(o["multi-mode"]) {
			return fmt.Errorf("Error reading multi_mode: %v", err)
		}
	}

	if err = d.Set("roaming", flattenExtenderControllerExtenderRoaming(o["roaming"], d, "roaming", sv)); err != nil {
		if !fortiAPIPatch(o["roaming"]) {
			return fmt.Errorf("Error reading roaming: %v", err)
		}
	}

	if err = d.Set("cdma_nai", flattenExtenderControllerExtenderCdmaNai(o["cdma-nai"], d, "cdma_nai", sv)); err != nil {
		if !fortiAPIPatch(o["cdma-nai"]) {
			return fmt.Errorf("Error reading cdma_nai: %v", err)
		}
	}

	if err = d.Set("primary_ha", flattenExtenderControllerExtenderPrimaryHa(o["primary-ha"], d, "primary_ha", sv)); err != nil {
		if !fortiAPIPatch(o["primary-ha"]) {
			return fmt.Errorf("Error reading primary_ha: %v", err)
		}
	}

	if err = d.Set("secondary_ha", flattenExtenderControllerExtenderSecondaryHa(o["secondary-ha"], d, "secondary_ha", sv)); err != nil {
		if !fortiAPIPatch(o["secondary-ha"]) {
			return fmt.Errorf("Error reading secondary_ha: %v", err)
		}
	}

	if err = d.Set("cdma_aaa_spi", flattenExtenderControllerExtenderCdmaAaaSpi(o["cdma-aaa-spi"], d, "cdma_aaa_spi", sv)); err != nil {
		if !fortiAPIPatch(o["cdma-aaa-spi"]) {
			return fmt.Errorf("Error reading cdma_aaa_spi: %v", err)
		}
	}

	if err = d.Set("cdma_ha_spi", flattenExtenderControllerExtenderCdmaHaSpi(o["cdma-ha-spi"], d, "cdma_ha_spi", sv)); err != nil {
		if !fortiAPIPatch(o["cdma-ha-spi"]) {
			return fmt.Errorf("Error reading cdma_ha_spi: %v", err)
		}
	}

	return nil
}

func flattenExtenderControllerExtenderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandExtenderControllerExtenderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAuthorized(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAdmin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderIfname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDeviceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderExtensionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderOverrideAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderOverrideLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderLoginPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderOverrideEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderBandwidthLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWanExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "modem1_extension"
	if _, ok := d.GetOk(pre_append); ok {

		result["modem1-extension"], _ = expandExtenderControllerExtenderWanExtensionModem1Extension(d, i["modem1_extension"], pre_append, sv)
	}
	pre_append = pre + ".0." + "modem2_extension"
	if _, ok := d.GetOk(pre_append); ok {

		result["modem2-extension"], _ = expandExtenderControllerExtenderWanExtensionModem2Extension(d, i["modem2_extension"], pre_append, sv)
	}

	return result, nil
}

func expandExtenderControllerExtenderWanExtensionModem1Extension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWanExtensionModem2Extension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderControllerReport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {

		result["status"], _ = expandExtenderControllerExtenderControllerReportStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "interval"
	if _, ok := d.GetOk(pre_append); ok {

		result["interval"], _ = expandExtenderControllerExtenderControllerReportInterval(d, i["interval"], pre_append, sv)
	}
	pre_append = pre + ".0." + "signal_threshold"
	if _, ok := d.GetOk(pre_append); ok {

		result["signal-threshold"], _ = expandExtenderControllerExtenderControllerReportSignalThreshold(d, i["signal_threshold"], pre_append, sv)
	}

	return result, nil
}

func expandExtenderControllerExtenderControllerReportStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderControllerReportInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderControllerReportSignalThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ifname"
	if _, ok := d.GetOk(pre_append); ok {

		result["ifname"], _ = expandExtenderControllerExtenderModem1Ifname(d, i["ifname"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["redundant-mode"], _ = expandExtenderControllerExtenderModem1RedundantMode(d, i["redundant_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok {

		result["redundant-intf"], _ = expandExtenderControllerExtenderModem1RedundantIntf(d, i["redundant_intf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok {

		result["conn-status"], _ = expandExtenderControllerExtenderModem1ConnStatus(d, i["conn_status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok {

		result["default-sim"], _ = expandExtenderControllerExtenderModem1DefaultSim(d, i["default_sim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok {

		result["gps"], _ = expandExtenderControllerExtenderModem1Gps(d, i["gps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok {

		result["sim1-pin"], _ = expandExtenderControllerExtenderModem1Sim1Pin(d, i["sim1_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok {

		result["sim2-pin"], _ = expandExtenderControllerExtenderModem1Sim2Pin(d, i["sim2_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok {

		result["sim1-pin-code"], _ = expandExtenderControllerExtenderModem1Sim1PinCode(d, i["sim1_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok {

		result["sim2-pin-code"], _ = expandExtenderControllerExtenderModem1Sim2PinCode(d, i["sim2_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok {

		result["preferred-carrier"], _ = expandExtenderControllerExtenderModem1PreferredCarrier(d, i["preferred_carrier"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-switch"], _ = expandExtenderControllerExtenderModem1AutoSwitch(d, i["auto_switch"], pre_append, sv)
	} else {
		result["auto-switch"] = make([]string, 0)
	}

	return result, nil
}

func expandExtenderControllerExtenderModem1Ifname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1RedundantMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1RedundantIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1ConnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1DefaultSim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Gps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Sim1Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Sim2Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Sim1PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1Sim2PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1PreferredCarrier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok {

		result["disconnect"], _ = expandExtenderControllerExtenderModem1AutoSwitchDisconnect(d, i["disconnect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok {

		result["disconnect-threshold"], _ = expandExtenderControllerExtenderModem1AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok {

		result["disconnect-period"], _ = expandExtenderControllerExtenderModem1AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append, sv)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok {

		result["signal"], _ = expandExtenderControllerExtenderModem1AutoSwitchSignal(d, i["signal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {

		result["dataplan"], _ = expandExtenderControllerExtenderModem1AutoSwitchDataplan(d, i["dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok {

		result["switch-back"], _ = expandExtenderControllerExtenderModem1AutoSwitchSwitchBack(d, i["switch_back"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok {

		result["switch-back-time"], _ = expandExtenderControllerExtenderModem1AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok {

		result["switch-back-timer"], _ = expandExtenderControllerExtenderModem1AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append, sv)
	}

	return result, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem1AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "ifname"
	if _, ok := d.GetOk(pre_append); ok {

		result["ifname"], _ = expandExtenderControllerExtenderModem2Ifname(d, i["ifname"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redundant_mode"
	if _, ok := d.GetOk(pre_append); ok {

		result["redundant-mode"], _ = expandExtenderControllerExtenderModem2RedundantMode(d, i["redundant_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "redundant_intf"
	if _, ok := d.GetOk(pre_append); ok {

		result["redundant-intf"], _ = expandExtenderControllerExtenderModem2RedundantIntf(d, i["redundant_intf"], pre_append, sv)
	}
	pre_append = pre + ".0." + "conn_status"
	if _, ok := d.GetOk(pre_append); ok {

		result["conn-status"], _ = expandExtenderControllerExtenderModem2ConnStatus(d, i["conn_status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "default_sim"
	if _, ok := d.GetOk(pre_append); ok {

		result["default-sim"], _ = expandExtenderControllerExtenderModem2DefaultSim(d, i["default_sim"], pre_append, sv)
	}
	pre_append = pre + ".0." + "gps"
	if _, ok := d.GetOk(pre_append); ok {

		result["gps"], _ = expandExtenderControllerExtenderModem2Gps(d, i["gps"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin"
	if _, ok := d.GetOk(pre_append); ok {

		result["sim1-pin"], _ = expandExtenderControllerExtenderModem2Sim1Pin(d, i["sim1_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin"
	if _, ok := d.GetOk(pre_append); ok {

		result["sim2-pin"], _ = expandExtenderControllerExtenderModem2Sim2Pin(d, i["sim2_pin"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim1_pin_code"
	if _, ok := d.GetOk(pre_append); ok {

		result["sim1-pin-code"], _ = expandExtenderControllerExtenderModem2Sim1PinCode(d, i["sim1_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "sim2_pin_code"
	if _, ok := d.GetOk(pre_append); ok {

		result["sim2-pin-code"], _ = expandExtenderControllerExtenderModem2Sim2PinCode(d, i["sim2_pin_code"], pre_append, sv)
	}
	pre_append = pre + ".0." + "preferred_carrier"
	if _, ok := d.GetOk(pre_append); ok {

		result["preferred-carrier"], _ = expandExtenderControllerExtenderModem2PreferredCarrier(d, i["preferred_carrier"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auto_switch"
	if _, ok := d.GetOk(pre_append); ok {

		result["auto-switch"], _ = expandExtenderControllerExtenderModem2AutoSwitch(d, i["auto_switch"], pre_append, sv)
	} else {
		result["auto-switch"] = make([]string, 0)
	}

	return result, nil
}

func expandExtenderControllerExtenderModem2Ifname(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2RedundantMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2RedundantIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2ConnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2DefaultSim(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Gps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Sim1Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Sim2Pin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Sim1PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2Sim2PinCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2PreferredCarrier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "disconnect"
	if _, ok := d.GetOk(pre_append); ok {

		result["disconnect"], _ = expandExtenderControllerExtenderModem2AutoSwitchDisconnect(d, i["disconnect"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_threshold"
	if _, ok := d.GetOk(pre_append); ok {

		result["disconnect-threshold"], _ = expandExtenderControllerExtenderModem2AutoSwitchDisconnectThreshold(d, i["disconnect_threshold"], pre_append, sv)
	}
	pre_append = pre + ".0." + "disconnect_period"
	if _, ok := d.GetOk(pre_append); ok {

		result["disconnect-period"], _ = expandExtenderControllerExtenderModem2AutoSwitchDisconnectPeriod(d, i["disconnect_period"], pre_append, sv)
	}
	pre_append = pre + ".0." + "signal"
	if _, ok := d.GetOk(pre_append); ok {

		result["signal"], _ = expandExtenderControllerExtenderModem2AutoSwitchSignal(d, i["signal"], pre_append, sv)
	}
	pre_append = pre + ".0." + "dataplan"
	if _, ok := d.GetOk(pre_append); ok {

		result["dataplan"], _ = expandExtenderControllerExtenderModem2AutoSwitchDataplan(d, i["dataplan"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back"
	if _, ok := d.GetOk(pre_append); ok {

		result["switch-back"], _ = expandExtenderControllerExtenderModem2AutoSwitchSwitchBack(d, i["switch_back"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_time"
	if _, ok := d.GetOk(pre_append); ok {

		result["switch-back-time"], _ = expandExtenderControllerExtenderModem2AutoSwitchSwitchBackTime(d, i["switch_back_time"], pre_append, sv)
	}
	pre_append = pre + ".0." + "switch_back_timer"
	if _, ok := d.GetOk(pre_append); ok {

		result["switch-back-timer"], _ = expandExtenderControllerExtenderModem2AutoSwitchSwitchBackTimer(d, i["switch_back_timer"], pre_append, sv)
	}

	return result, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchDisconnect(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchDisconnectThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchDisconnectPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchSignal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchDataplan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchSwitchBack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchSwitchBackTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModem2AutoSwitchSwitchBackTimer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRole(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDialMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRedial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRedundantIntf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDialStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderConnStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderExtName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderQuotaLimitMb(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderBillingStartDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAtDialScript(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModemPasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderInitiatedUpdate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderModemType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppAuthProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPppEchoRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWimaxCarrier(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWimaxRealm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderWimaxAuthProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderSimPin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAccessPointName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderMultiMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderRoaming(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderCdmaNai(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderAaaSharedSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderHaSharedSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderPrimaryHa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderSecondaryHa(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderCdmaAaaSpi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtenderControllerExtenderCdmaHaSpi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectExtenderControllerExtender(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandExtenderControllerExtenderName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {

		t, err := expandExtenderControllerExtenderId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("authorized"); ok {

		t, err := expandExtenderControllerExtenderAuthorized(d, v, "authorized", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized"] = t
		}
	}

	if v, ok := d.GetOk("admin"); ok {

		t, err := expandExtenderControllerExtenderAdmin(d, v, "admin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["admin"] = t
		}
	}

	if v, ok := d.GetOk("ifname"); ok {

		t, err := expandExtenderControllerExtenderIfname(d, v, "ifname", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ifname"] = t
		}
	}

	if v, ok := d.GetOkExists("vdom"); ok {

		t, err := expandExtenderControllerExtenderVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOkExists("device_id"); ok {

		t, err := expandExtenderControllerExtenderDeviceId(d, v, "device_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-id"] = t
		}
	}

	if v, ok := d.GetOk("extension_type"); ok {

		t, err := expandExtenderControllerExtenderExtensionType(d, v, "extension_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-type"] = t
		}
	}

	if v, ok := d.GetOk("override_allowaccess"); ok {

		t, err := expandExtenderControllerExtenderOverrideAllowaccess(d, v, "override_allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {

		t, err := expandExtenderControllerExtenderAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("override_login_password_change"); ok {

		t, err := expandExtenderControllerExtenderOverrideLoginPasswordChange(d, v, "override_login_password_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("login_password_change"); ok {

		t, err := expandExtenderControllerExtenderLoginPasswordChange(d, v, "login_password_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok {

		t, err := expandExtenderControllerExtenderLoginPassword(d, v, "login_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	}

	if v, ok := d.GetOk("override_enforce_bandwidth"); ok {

		t, err := expandExtenderControllerExtenderOverrideEnforceBandwidth(d, v, "override_enforce_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("enforce_bandwidth"); ok {

		t, err := expandExtenderControllerExtenderEnforceBandwidth(d, v, "enforce_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_limit"); ok {

		t, err := expandExtenderControllerExtenderBandwidthLimit(d, v, "bandwidth_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-limit"] = t
		}
	}

	if v, ok := d.GetOk("wan_extension"); ok {

		t, err := expandExtenderControllerExtenderWanExtension(d, v, "wan_extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-extension"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok {

		t, err := expandExtenderControllerExtenderProfile(d, v, "profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("controller_report"); ok {

		t, err := expandExtenderControllerExtenderControllerReport(d, v, "controller_report", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["controller-report"] = t
		}
	}

	if v, ok := d.GetOk("modem1"); ok {

		t, err := expandExtenderControllerExtenderModem1(d, v, "modem1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem1"] = t
		}
	}

	if v, ok := d.GetOk("modem2"); ok {

		t, err := expandExtenderControllerExtenderModem2(d, v, "modem2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem2"] = t
		}
	}

	if v, ok := d.GetOk("role"); ok {

		t, err := expandExtenderControllerExtenderRole(d, v, "role", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["role"] = t
		}
	}

	if v, ok := d.GetOk("mode"); ok {

		t, err := expandExtenderControllerExtenderMode(d, v, "mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("dial_mode"); ok {

		t, err := expandExtenderControllerExtenderDialMode(d, v, "dial_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dial-mode"] = t
		}
	}

	if v, ok := d.GetOk("redial"); ok {

		t, err := expandExtenderControllerExtenderRedial(d, v, "redial", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redial"] = t
		}
	}

	if v, ok := d.GetOk("redundant_intf"); ok {

		t, err := expandExtenderControllerExtenderRedundantIntf(d, v, "redundant_intf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["redundant-intf"] = t
		}
	}

	if v, ok := d.GetOkExists("dial_status"); ok {

		t, err := expandExtenderControllerExtenderDialStatus(d, v, "dial_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dial-status"] = t
		}
	}

	if v, ok := d.GetOkExists("conn_status"); ok {

		t, err := expandExtenderControllerExtenderConnStatus(d, v, "conn_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-status"] = t
		}
	}

	if v, ok := d.GetOk("ext_name"); ok {

		t, err := expandExtenderControllerExtenderExtName(d, v, "ext_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandExtenderControllerExtenderDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOkExists("quota_limit_mb"); ok {

		t, err := expandExtenderControllerExtenderQuotaLimitMb(d, v, "quota_limit_mb", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["quota-limit-mb"] = t
		}
	}

	if v, ok := d.GetOk("billing_start_day"); ok {

		t, err := expandExtenderControllerExtenderBillingStartDay(d, v, "billing_start_day", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["billing-start-day"] = t
		}
	}

	if v, ok := d.GetOk("at_dial_script"); ok {

		t, err := expandExtenderControllerExtenderAtDialScript(d, v, "at_dial_script", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["at-dial-script"] = t
		}
	}

	if v, ok := d.GetOk("modem_passwd"); ok {

		t, err := expandExtenderControllerExtenderModemPasswd(d, v, "modem_passwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-passwd"] = t
		}
	}

	if v, ok := d.GetOk("initiated_update"); ok {

		t, err := expandExtenderControllerExtenderInitiatedUpdate(d, v, "initiated_update", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["initiated-update"] = t
		}
	}

	if v, ok := d.GetOk("modem_type"); ok {

		t, err := expandExtenderControllerExtenderModemType(d, v, "modem_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["modem-type"] = t
		}
	}

	if v, ok := d.GetOk("ppp_username"); ok {

		t, err := expandExtenderControllerExtenderPppUsername(d, v, "ppp_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-username"] = t
		}
	}

	if v, ok := d.GetOk("ppp_password"); ok {

		t, err := expandExtenderControllerExtenderPppPassword(d, v, "ppp_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-password"] = t
		}
	}

	if v, ok := d.GetOk("ppp_auth_protocol"); ok {

		t, err := expandExtenderControllerExtenderPppAuthProtocol(d, v, "ppp_auth_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-auth-protocol"] = t
		}
	}

	if v, ok := d.GetOk("ppp_echo_request"); ok {

		t, err := expandExtenderControllerExtenderPppEchoRequest(d, v, "ppp_echo_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ppp-echo-request"] = t
		}
	}

	if v, ok := d.GetOk("wimax_carrier"); ok {

		t, err := expandExtenderControllerExtenderWimaxCarrier(d, v, "wimax_carrier", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-carrier"] = t
		}
	}

	if v, ok := d.GetOk("wimax_realm"); ok {

		t, err := expandExtenderControllerExtenderWimaxRealm(d, v, "wimax_realm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-realm"] = t
		}
	}

	if v, ok := d.GetOk("wimax_auth_protocol"); ok {

		t, err := expandExtenderControllerExtenderWimaxAuthProtocol(d, v, "wimax_auth_protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wimax-auth-protocol"] = t
		}
	}

	if v, ok := d.GetOk("sim_pin"); ok {

		t, err := expandExtenderControllerExtenderSimPin(d, v, "sim_pin", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sim-pin"] = t
		}
	}

	if v, ok := d.GetOk("access_point_name"); ok {

		t, err := expandExtenderControllerExtenderAccessPointName(d, v, "access_point_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-point-name"] = t
		}
	}

	if v, ok := d.GetOk("multi_mode"); ok {

		t, err := expandExtenderControllerExtenderMultiMode(d, v, "multi_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["multi-mode"] = t
		}
	}

	if v, ok := d.GetOk("roaming"); ok {

		t, err := expandExtenderControllerExtenderRoaming(d, v, "roaming", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roaming"] = t
		}
	}

	if v, ok := d.GetOk("cdma_nai"); ok {

		t, err := expandExtenderControllerExtenderCdmaNai(d, v, "cdma_nai", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdma-nai"] = t
		}
	}

	if v, ok := d.GetOk("aaa_shared_secret"); ok {

		t, err := expandExtenderControllerExtenderAaaSharedSecret(d, v, "aaa_shared_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aaa-shared-secret"] = t
		}
	}

	if v, ok := d.GetOk("ha_shared_secret"); ok {

		t, err := expandExtenderControllerExtenderHaSharedSecret(d, v, "ha_shared_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-shared-secret"] = t
		}
	}

	if v, ok := d.GetOk("primary_ha"); ok {

		t, err := expandExtenderControllerExtenderPrimaryHa(d, v, "primary_ha", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["primary-ha"] = t
		}
	}

	if v, ok := d.GetOk("secondary_ha"); ok {

		t, err := expandExtenderControllerExtenderSecondaryHa(d, v, "secondary_ha", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secondary-ha"] = t
		}
	}

	if v, ok := d.GetOk("cdma_aaa_spi"); ok {

		t, err := expandExtenderControllerExtenderCdmaAaaSpi(d, v, "cdma_aaa_spi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdma-aaa-spi"] = t
		}
	}

	if v, ok := d.GetOk("cdma_ha_spi"); ok {

		t, err := expandExtenderControllerExtenderCdmaHaSpi(d, v, "cdma_ha_spi", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cdma-ha-spi"] = t
		}
	}

	return &obj, nil
}
