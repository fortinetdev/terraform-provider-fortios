// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure application control lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceApplicationList() *schema.Resource {
	return &schema.Resource{
		Create: resourceApplicationListCreate,
		Read:   resourceApplicationListRead,
		Update: resourceApplicationListUpdate,
		Delete: resourceApplicationListDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 47),
				ForceNew:     true,
				Required:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"replacemsg_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"extended_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"other_application_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_replacemsg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"other_application_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_default_app_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"force_inclusion_ssl_di_sigs": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_application_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_application_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"p2p_block_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"p2p_black_list": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"deep_app_inspection": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entries": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"risk": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"category": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"sub_category": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"application": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"protocols": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vendor": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"technology": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"behavior": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"popularity": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"exclusion": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
									},
								},
							},
						},
						"parameters": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"members": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 31),
													Optional:     true,
												},
												"value": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 199),
													Optional:     true,
												},
											},
										},
									},
									"value": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
									},
								},
							},
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_packet": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_count": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"rate_duration": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"rate_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_track": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"session_ttl": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"shaper": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"shaper_reverse": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"per_ip_shaper": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"quarantine": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"control_default_network_services": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"default_network_services": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
						},
						"port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"services": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"violation_action": &schema.Schema{
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
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceApplicationListCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationList resource while getting object: %v", err)
	}

	o, err := c.CreateApplicationList(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ApplicationList resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ApplicationList")
	}

	return resourceApplicationListRead(d, m)
}

func resourceApplicationListUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectApplicationList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationList resource while getting object: %v", err)
	}

	o, err := c.UpdateApplicationList(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ApplicationList")
	}

	return resourceApplicationListRead(d, m)
}

func resourceApplicationListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteApplicationList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ApplicationList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceApplicationListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadApplicationList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationList resource from API: %v", err)
	}
	return nil
}

func flattenApplicationListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListExtendedLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListOtherApplicationAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListAppReplacemsg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListOtherApplicationLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEnforceDefaultAppPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListForceInclusionSslDiSigs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListUnknownApplicationAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListUnknownApplicationLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListP2PBlockList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListP2PBlackList(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListDeepAppInspection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListOptions(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk"
		if cur_v, ok := i["risk"]; ok {
			tmp["risk"] = flattenApplicationListEntriesRisk(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if cur_v, ok := i["category"]; ok {
			tmp["category"] = flattenApplicationListEntriesCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_category"
		if cur_v, ok := i["sub-category"]; ok {
			tmp["sub_category"] = flattenApplicationListEntriesSubCategory(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if cur_v, ok := i["application"]; ok {
			tmp["application"] = flattenApplicationListEntriesApplication(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocols"
		if cur_v, ok := i["protocols"]; ok {
			tmp["protocols"] = flattenApplicationListEntriesProtocols(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if cur_v, ok := i["vendor"]; ok {
			tmp["vendor"] = flattenApplicationListEntriesVendor(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "technology"
		if cur_v, ok := i["technology"]; ok {
			tmp["technology"] = flattenApplicationListEntriesTechnology(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "behavior"
		if cur_v, ok := i["behavior"]; ok {
			tmp["behavior"] = flattenApplicationListEntriesBehavior(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "popularity"
		if cur_v, ok := i["popularity"]; ok {
			tmp["popularity"] = flattenApplicationListEntriesPopularity(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusion"
		if cur_v, ok := i["exclusion"]; ok {
			tmp["exclusion"] = flattenApplicationListEntriesExclusion(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if cur_v, ok := i["parameters"]; ok {
			tmp["parameters"] = flattenApplicationListEntriesParameters(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenApplicationListEntriesAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if cur_v, ok := i["log"]; ok {
			tmp["log"] = flattenApplicationListEntriesLog(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if cur_v, ok := i["log-packet"]; ok {
			tmp["log_packet"] = flattenApplicationListEntriesLogPacket(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if cur_v, ok := i["rate-count"]; ok {
			tmp["rate_count"] = flattenApplicationListEntriesRateCount(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if cur_v, ok := i["rate-duration"]; ok {
			tmp["rate_duration"] = flattenApplicationListEntriesRateDuration(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if cur_v, ok := i["rate-mode"]; ok {
			tmp["rate_mode"] = flattenApplicationListEntriesRateMode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if cur_v, ok := i["rate-track"]; ok {
			tmp["rate_track"] = flattenApplicationListEntriesRateTrack(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_ttl"
		if cur_v, ok := i["session-ttl"]; ok {
			tmp["session_ttl"] = flattenApplicationListEntriesSessionTtl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper"
		if cur_v, ok := i["shaper"]; ok {
			tmp["shaper"] = flattenApplicationListEntriesShaper(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper_reverse"
		if cur_v, ok := i["shaper-reverse"]; ok {
			tmp["shaper_reverse"] = flattenApplicationListEntriesShaperReverse(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "per_ip_shaper"
		if cur_v, ok := i["per-ip-shaper"]; ok {
			tmp["per_ip_shaper"] = flattenApplicationListEntriesPerIpShaper(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if cur_v, ok := i["quarantine"]; ok {
			tmp["quarantine"] = flattenApplicationListEntriesQuarantine(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if cur_v, ok := i["quarantine-expiry"]; ok {
			tmp["quarantine_expiry"] = flattenApplicationListEntriesQuarantineExpiry(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if cur_v, ok := i["quarantine-log"]; ok {
			tmp["quarantine_log"] = flattenApplicationListEntriesQuarantineLog(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationListEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesRisk(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if cur_v, ok := i["level"]; ok {
			tmp["level"] = flattenApplicationListEntriesRiskLevel(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "level", d)
	return result
}

func flattenApplicationListEntriesRiskLevel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesCategoryId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationListEntriesCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesSubCategory(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesSubCategoryId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationListEntriesSubCategoryId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesApplicationId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationListEntriesApplicationId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesProtocols(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesTechnology(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesBehavior(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesPopularity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesExclusion(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesExclusionId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationListEntriesExclusionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesParameters(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesParametersId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if cur_v, ok := i["members"]; ok {
			tmp["members"] = flattenApplicationListEntriesParametersMembers(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenApplicationListEntriesParametersValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationListEntriesParametersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesParametersMembers(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesParametersMembersId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenApplicationListEntriesParametersMembersName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenApplicationListEntriesParametersMembersValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationListEntriesParametersMembersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesParametersMembersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesParametersMembersValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesParametersValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesLogPacket(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesRateCount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesRateDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesRateMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesRateTrack(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesSessionTtl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListEntriesShaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesShaperReverse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesPerIpShaper(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantine(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantineLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListControlDefaultNetworkServices(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListDefaultNetworkServices(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListDefaultNetworkServicesId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if cur_v, ok := i["port"]; ok {
			tmp["port"] = flattenApplicationListDefaultNetworkServicesPort(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if cur_v, ok := i["services"]; ok {
			tmp["services"] = flattenApplicationListDefaultNetworkServicesServices(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "violation_action"
		if cur_v, ok := i["violation-action"]; ok {
			tmp["violation_action"] = flattenApplicationListDefaultNetworkServicesViolationAction(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenApplicationListDefaultNetworkServicesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListDefaultNetworkServicesPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenApplicationListDefaultNetworkServicesServices(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenApplicationListDefaultNetworkServicesViolationAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectApplicationList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenApplicationListName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenApplicationListComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenApplicationListReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenApplicationListExtendedLog(o["extended-log"], d, "extended_log", sv)); err != nil {
		if !fortiAPIPatch(o["extended-log"]) {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("other_application_action", flattenApplicationListOtherApplicationAction(o["other-application-action"], d, "other_application_action", sv)); err != nil {
		if !fortiAPIPatch(o["other-application-action"]) {
			return fmt.Errorf("Error reading other_application_action: %v", err)
		}
	}

	if err = d.Set("app_replacemsg", flattenApplicationListAppReplacemsg(o["app-replacemsg"], d, "app_replacemsg", sv)); err != nil {
		if !fortiAPIPatch(o["app-replacemsg"]) {
			return fmt.Errorf("Error reading app_replacemsg: %v", err)
		}
	}

	if err = d.Set("other_application_log", flattenApplicationListOtherApplicationLog(o["other-application-log"], d, "other_application_log", sv)); err != nil {
		if !fortiAPIPatch(o["other-application-log"]) {
			return fmt.Errorf("Error reading other_application_log: %v", err)
		}
	}

	if err = d.Set("enforce_default_app_port", flattenApplicationListEnforceDefaultAppPort(o["enforce-default-app-port"], d, "enforce_default_app_port", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-default-app-port"]) {
			return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
		}
	}

	if err = d.Set("force_inclusion_ssl_di_sigs", flattenApplicationListForceInclusionSslDiSigs(o["force-inclusion-ssl-di-sigs"], d, "force_inclusion_ssl_di_sigs", sv)); err != nil {
		if !fortiAPIPatch(o["force-inclusion-ssl-di-sigs"]) {
			return fmt.Errorf("Error reading force_inclusion_ssl_di_sigs: %v", err)
		}
	}

	if err = d.Set("unknown_application_action", flattenApplicationListUnknownApplicationAction(o["unknown-application-action"], d, "unknown_application_action", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-application-action"]) {
			return fmt.Errorf("Error reading unknown_application_action: %v", err)
		}
	}

	if err = d.Set("unknown_application_log", flattenApplicationListUnknownApplicationLog(o["unknown-application-log"], d, "unknown_application_log", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-application-log"]) {
			return fmt.Errorf("Error reading unknown_application_log: %v", err)
		}
	}

	if err = d.Set("p2p_block_list", flattenApplicationListP2PBlockList(o["p2p-block-list"], d, "p2p_block_list", sv)); err != nil {
		if !fortiAPIPatch(o["p2p-block-list"]) {
			return fmt.Errorf("Error reading p2p_block_list: %v", err)
		}
	}

	if err = d.Set("p2p_black_list", flattenApplicationListP2PBlackList(o["p2p-black-list"], d, "p2p_black_list", sv)); err != nil {
		if !fortiAPIPatch(o["p2p-black-list"]) {
			return fmt.Errorf("Error reading p2p_black_list: %v", err)
		}
	}

	if err = d.Set("deep_app_inspection", flattenApplicationListDeepAppInspection(o["deep-app-inspection"], d, "deep_app_inspection", sv)); err != nil {
		if !fortiAPIPatch(o["deep-app-inspection"]) {
			return fmt.Errorf("Error reading deep_app_inspection: %v", err)
		}
	}

	if err = d.Set("options", flattenApplicationListOptions(o["options"], d, "options", sv)); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("entries", flattenApplicationListEntries(o["entries"], d, "entries", sv)); err != nil {
			if !fortiAPIPatch(o["entries"]) {
				return fmt.Errorf("Error reading entries: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("entries"); ok {
			if err = d.Set("entries", flattenApplicationListEntries(o["entries"], d, "entries", sv)); err != nil {
				if !fortiAPIPatch(o["entries"]) {
					return fmt.Errorf("Error reading entries: %v", err)
				}
			}
		}
	}

	if err = d.Set("control_default_network_services", flattenApplicationListControlDefaultNetworkServices(o["control-default-network-services"], d, "control_default_network_services", sv)); err != nil {
		if !fortiAPIPatch(o["control-default-network-services"]) {
			return fmt.Errorf("Error reading control_default_network_services: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("default_network_services", flattenApplicationListDefaultNetworkServices(o["default-network-services"], d, "default_network_services", sv)); err != nil {
			if !fortiAPIPatch(o["default-network-services"]) {
				return fmt.Errorf("Error reading default_network_services: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("default_network_services"); ok {
			if err = d.Set("default_network_services", flattenApplicationListDefaultNetworkServices(o["default-network-services"], d, "default_network_services", sv)); err != nil {
				if !fortiAPIPatch(o["default-network-services"]) {
					return fmt.Errorf("Error reading default_network_services: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenApplicationListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandApplicationListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListExtendedLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListOtherApplicationAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListAppReplacemsg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListOtherApplicationLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEnforceDefaultAppPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListForceInclusionSslDiSigs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListUnknownApplicationAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListUnknownApplicationLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListP2PBlockList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListP2PBlackList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDeepAppInspection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandApplicationListEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["risk"], _ = expandApplicationListEntriesRisk(d, i["risk"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["risk"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandApplicationListEntriesCategory(d, i["category"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["category"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sub-category"], _ = expandApplicationListEntriesSubCategory(d, i["sub_category"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["sub-category"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["application"], _ = expandApplicationListEntriesApplication(d, i["application"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["application"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocols"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocols"], _ = expandApplicationListEntriesProtocols(d, i["protocols"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vendor"], _ = expandApplicationListEntriesVendor(d, i["vendor"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "technology"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["technology"], _ = expandApplicationListEntriesTechnology(d, i["technology"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "behavior"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["behavior"], _ = expandApplicationListEntriesBehavior(d, i["behavior"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "popularity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["popularity"], _ = expandApplicationListEntriesPopularity(d, i["popularity"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "exclusion"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["exclusion"], _ = expandApplicationListEntriesExclusion(d, i["exclusion"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["exclusion"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["parameters"], _ = expandApplicationListEntriesParameters(d, i["parameters"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["parameters"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandApplicationListEntriesAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandApplicationListEntriesLog(d, i["log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log-packet"], _ = expandApplicationListEntriesLogPacket(d, i["log_packet"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-count"], _ = expandApplicationListEntriesRateCount(d, i["rate_count"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rate-count"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-duration"], _ = expandApplicationListEntriesRateDuration(d, i["rate_duration"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-mode"], _ = expandApplicationListEntriesRateMode(d, i["rate_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-track"], _ = expandApplicationListEntriesRateTrack(d, i["rate_track"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_ttl"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["session-ttl"], _ = expandApplicationListEntriesSessionTtl(d, i["session_ttl"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["session-ttl"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["shaper"], _ = expandApplicationListEntriesShaper(d, i["shaper"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["shaper"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper_reverse"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["shaper-reverse"], _ = expandApplicationListEntriesShaperReverse(d, i["shaper_reverse"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["shaper-reverse"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "per_ip_shaper"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["per-ip-shaper"], _ = expandApplicationListEntriesPerIpShaper(d, i["per_ip_shaper"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["per-ip-shaper"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine"], _ = expandApplicationListEntriesQuarantine(d, i["quarantine"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-expiry"], _ = expandApplicationListEntriesQuarantineExpiry(d, i["quarantine_expiry"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-log"], _ = expandApplicationListEntriesQuarantineLog(d, i["quarantine_log"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRisk(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["level"], _ = expandApplicationListEntriesRiskLevel(d, i["level"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesRiskLevel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandApplicationListEntriesCategoryId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesSubCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandApplicationListEntriesSubCategoryId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesSubCategoryId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandApplicationListEntriesApplicationId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesApplicationId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesProtocols(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesTechnology(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesBehavior(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesPopularity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesExclusion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["id"], _ = expandApplicationListEntriesExclusionId(d, i["id"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesExclusionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParameters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandApplicationListEntriesParametersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "members"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["members"], _ = expandApplicationListEntriesParametersMembers(d, i["members"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["members"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandApplicationListEntriesParametersValue(d, i["value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesParametersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandApplicationListEntriesParametersMembersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandApplicationListEntriesParametersMembersName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandApplicationListEntriesParametersMembersValue(d, i["value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesParametersMembersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersMembersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersMembersValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesLogPacket(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateCount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateTrack(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesSessionTtl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesShaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesShaperReverse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesPerIpShaper(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantine(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantineLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListControlDefaultNetworkServices(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDefaultNetworkServices(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandApplicationListDefaultNetworkServicesId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["port"], _ = expandApplicationListDefaultNetworkServicesPort(d, i["port"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["port"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "services"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["services"], _ = expandApplicationListDefaultNetworkServicesServices(d, i["services"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["services"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "violation_action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["violation-action"], _ = expandApplicationListDefaultNetworkServicesViolationAction(d, i["violation_action"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListDefaultNetworkServicesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDefaultNetworkServicesPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDefaultNetworkServicesServices(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDefaultNetworkServicesViolationAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectApplicationList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandApplicationListName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandApplicationListComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandApplicationListReplacemsgGroup(d, v, "replacemsg_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	} else if d.HasChange("replacemsg_group") {
		obj["replacemsg-group"] = nil
	}

	if v, ok := d.GetOk("extended_log"); ok {
		t, err := expandApplicationListExtendedLog(d, v, "extended_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("other_application_action"); ok {
		t, err := expandApplicationListOtherApplicationAction(d, v, "other_application_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-application-action"] = t
		}
	}

	if v, ok := d.GetOk("app_replacemsg"); ok {
		t, err := expandApplicationListAppReplacemsg(d, v, "app_replacemsg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("other_application_log"); ok {
		t, err := expandApplicationListOtherApplicationLog(d, v, "other_application_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-application-log"] = t
		}
	}

	if v, ok := d.GetOk("enforce_default_app_port"); ok {
		t, err := expandApplicationListEnforceDefaultAppPort(d, v, "enforce_default_app_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-default-app-port"] = t
		}
	}

	if v, ok := d.GetOk("force_inclusion_ssl_di_sigs"); ok {
		t, err := expandApplicationListForceInclusionSslDiSigs(d, v, "force_inclusion_ssl_di_sigs", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["force-inclusion-ssl-di-sigs"] = t
		}
	}

	if v, ok := d.GetOk("unknown_application_action"); ok {
		t, err := expandApplicationListUnknownApplicationAction(d, v, "unknown_application_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-application-action"] = t
		}
	}

	if v, ok := d.GetOk("unknown_application_log"); ok {
		t, err := expandApplicationListUnknownApplicationLog(d, v, "unknown_application_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-application-log"] = t
		}
	}

	if v, ok := d.GetOk("p2p_block_list"); ok {
		t, err := expandApplicationListP2PBlockList(d, v, "p2p_block_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["p2p-block-list"] = t
		}
	} else if d.HasChange("p2p_block_list") {
		obj["p2p-block-list"] = nil
	}

	if v, ok := d.GetOk("p2p_black_list"); ok {
		t, err := expandApplicationListP2PBlackList(d, v, "p2p_black_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["p2p-black-list"] = t
		}
	} else if d.HasChange("p2p_black_list") {
		obj["p2p-black-list"] = nil
	}

	if v, ok := d.GetOk("deep_app_inspection"); ok {
		t, err := expandApplicationListDeepAppInspection(d, v, "deep_app_inspection", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-inspection"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandApplicationListOptions(d, v, "options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok || d.HasChange("entries") {
		t, err := expandApplicationListEntries(d, v, "entries", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}

	if v, ok := d.GetOk("control_default_network_services"); ok {
		t, err := expandApplicationListControlDefaultNetworkServices(d, v, "control_default_network_services", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-default-network-services"] = t
		}
	}

	if v, ok := d.GetOk("default_network_services"); ok || d.HasChange("default_network_services") {
		t, err := expandApplicationListDefaultNetworkServices(d, v, "default_network_services", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["default-network-services"] = t
		}
	}

	return &obj, nil
}
