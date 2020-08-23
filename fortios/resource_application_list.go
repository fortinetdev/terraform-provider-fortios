// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure application control lists.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required: true,
			},
			"comment": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional: true,
			},
			"replacemsg_group": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"extended_log": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"other_application_action": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"app_replacemsg": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"other_application_log": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_default_app_port": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_application_action": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"unknown_application_log": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"p2p_black_list": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"deep_app_inspection": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"options": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"entries": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"risk": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"level": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"category": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"sub_category": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"application": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"protocols": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vendor": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"technology": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"behavior": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"popularity": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"parameters": &schema.Schema{
							Type: schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type: schema.TypeInt,
										ValidateFunc: validation.IntBetween(0, 4294967295),
										Optional: true,
										Computed: true,
									},
									"value": &schema.Schema{
										Type: schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"action": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"log_packet": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_count": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional: true,
							Computed: true,
						},
						"rate_duration": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional: true,
							Computed: true,
						},
						"rate_mode": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rate_track": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"session_ttl": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"shaper": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"shaper_reverse": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"per_ip_shaper": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional: true,
							Computed: true,
						},
						"quarantine": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_expiry": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"quarantine_log": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceApplicationListCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectApplicationList(d)
	if err != nil {
		return fmt.Errorf("Error creating ApplicationList resource while getting object: %v", err)
	}

	o, err := c.CreateApplicationList(obj)

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

	obj, err := getObjectApplicationList(d)
	if err != nil {
		return fmt.Errorf("Error updating ApplicationList resource while getting object: %v", err)
	}

	o, err := c.UpdateApplicationList(obj, mkey)
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

	err := c.DeleteApplicationList(mkey)
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

	o, err := c.ReadApplicationList(mkey)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectApplicationList(d, o)
	if err != nil {
		return fmt.Errorf("Error reading ApplicationList resource from API: %v", err)
	}
	return nil
}


func flattenApplicationListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListComment(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListExtendedLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListOtherApplicationAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListAppReplacemsg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListOtherApplicationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEnforceDefaultAppPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListUnknownApplicationAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListUnknownApplicationLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListP2PBlackList(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListDeepAppInspection(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListOptions(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntries(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk"
		if _, ok := i["risk"]; ok {
			tmp["risk"] = flattenApplicationListEntriesRisk(i["risk"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := i["category"]; ok {
			tmp["category"] = flattenApplicationListEntriesCategory(i["category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_category"
		if _, ok := i["sub-category"]; ok {
			tmp["sub_category"] = flattenApplicationListEntriesSubCategory(i["sub-category"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := i["application"]; ok {
			tmp["application"] = flattenApplicationListEntriesApplication(i["application"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocols"
		if _, ok := i["protocols"]; ok {
			tmp["protocols"] = flattenApplicationListEntriesProtocols(i["protocols"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := i["vendor"]; ok {
			tmp["vendor"] = flattenApplicationListEntriesVendor(i["vendor"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "technology"
		if _, ok := i["technology"]; ok {
			tmp["technology"] = flattenApplicationListEntriesTechnology(i["technology"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "behavior"
		if _, ok := i["behavior"]; ok {
			tmp["behavior"] = flattenApplicationListEntriesBehavior(i["behavior"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "popularity"
		if _, ok := i["popularity"]; ok {
			tmp["popularity"] = flattenApplicationListEntriesPopularity(i["popularity"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := i["parameters"]; ok {
			tmp["parameters"] = flattenApplicationListEntriesParameters(i["parameters"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenApplicationListEntriesAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := i["log"]; ok {
			tmp["log"] = flattenApplicationListEntriesLog(i["log"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := i["log-packet"]; ok {
			tmp["log_packet"] = flattenApplicationListEntriesLogPacket(i["log-packet"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := i["rate-count"]; ok {
			tmp["rate_count"] = flattenApplicationListEntriesRateCount(i["rate-count"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := i["rate-duration"]; ok {
			tmp["rate_duration"] = flattenApplicationListEntriesRateDuration(i["rate-duration"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := i["rate-mode"]; ok {
			tmp["rate_mode"] = flattenApplicationListEntriesRateMode(i["rate-mode"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := i["rate-track"]; ok {
			tmp["rate_track"] = flattenApplicationListEntriesRateTrack(i["rate-track"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_ttl"
		if _, ok := i["session-ttl"]; ok {
			tmp["session_ttl"] = flattenApplicationListEntriesSessionTtl(i["session-ttl"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper"
		if _, ok := i["shaper"]; ok {
			tmp["shaper"] = flattenApplicationListEntriesShaper(i["shaper"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper_reverse"
		if _, ok := i["shaper-reverse"]; ok {
			tmp["shaper_reverse"] = flattenApplicationListEntriesShaperReverse(i["shaper-reverse"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "per_ip_shaper"
		if _, ok := i["per-ip-shaper"]; ok {
			tmp["per_ip_shaper"] = flattenApplicationListEntriesPerIpShaper(i["per-ip-shaper"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := i["quarantine"]; ok {
			tmp["quarantine"] = flattenApplicationListEntriesQuarantine(i["quarantine"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := i["quarantine-expiry"]; ok {
			tmp["quarantine_expiry"] = flattenApplicationListEntriesQuarantineExpiry(i["quarantine-expiry"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := i["quarantine-log"]; ok {
			tmp["quarantine_log"] = flattenApplicationListEntriesQuarantineLog(i["quarantine-log"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenApplicationListEntriesId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRisk(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := i["level"]; ok {
			tmp["level"] = flattenApplicationListEntriesRiskLevel(i["level"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenApplicationListEntriesRiskLevel(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesCategoryId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenApplicationListEntriesCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesSubCategory(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesSubCategoryId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenApplicationListEntriesSubCategoryId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesApplication(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesApplicationId(i["id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenApplicationListEntriesApplicationId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesProtocols(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesVendor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesTechnology(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesBehavior(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesPopularity(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParameters(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenApplicationListEntriesParametersId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = flattenApplicationListEntriesParametersValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenApplicationListEntriesParametersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesParametersValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesLogPacket(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateCount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateDuration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesRateTrack(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesSessionTtl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesShaperReverse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesPerIpShaper(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantine(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantineExpiry(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenApplicationListEntriesQuarantineLog(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectApplicationList(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("name", flattenApplicationListName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenApplicationListComment(o["comment"], d, "comment")); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenApplicationListReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("extended_log", flattenApplicationListExtendedLog(o["extended-log"], d, "extended_log")); err != nil {
		if !fortiAPIPatch(o["extended-log"]) {
			return fmt.Errorf("Error reading extended_log: %v", err)
		}
	}

	if err = d.Set("other_application_action", flattenApplicationListOtherApplicationAction(o["other-application-action"], d, "other_application_action")); err != nil {
		if !fortiAPIPatch(o["other-application-action"]) {
			return fmt.Errorf("Error reading other_application_action: %v", err)
		}
	}

	if err = d.Set("app_replacemsg", flattenApplicationListAppReplacemsg(o["app-replacemsg"], d, "app_replacemsg")); err != nil {
		if !fortiAPIPatch(o["app-replacemsg"]) {
			return fmt.Errorf("Error reading app_replacemsg: %v", err)
		}
	}

	if err = d.Set("other_application_log", flattenApplicationListOtherApplicationLog(o["other-application-log"], d, "other_application_log")); err != nil {
		if !fortiAPIPatch(o["other-application-log"]) {
			return fmt.Errorf("Error reading other_application_log: %v", err)
		}
	}

	if err = d.Set("enforce_default_app_port", flattenApplicationListEnforceDefaultAppPort(o["enforce-default-app-port"], d, "enforce_default_app_port")); err != nil {
		if !fortiAPIPatch(o["enforce-default-app-port"]) {
			return fmt.Errorf("Error reading enforce_default_app_port: %v", err)
		}
	}

	if err = d.Set("unknown_application_action", flattenApplicationListUnknownApplicationAction(o["unknown-application-action"], d, "unknown_application_action")); err != nil {
		if !fortiAPIPatch(o["unknown-application-action"]) {
			return fmt.Errorf("Error reading unknown_application_action: %v", err)
		}
	}

	if err = d.Set("unknown_application_log", flattenApplicationListUnknownApplicationLog(o["unknown-application-log"], d, "unknown_application_log")); err != nil {
		if !fortiAPIPatch(o["unknown-application-log"]) {
			return fmt.Errorf("Error reading unknown_application_log: %v", err)
		}
	}

	if err = d.Set("p2p_black_list", flattenApplicationListP2PBlackList(o["p2p-black-list"], d, "p2p_black_list")); err != nil {
		if !fortiAPIPatch(o["p2p-black-list"]) {
			return fmt.Errorf("Error reading p2p_black_list: %v", err)
		}
	}

	if err = d.Set("deep_app_inspection", flattenApplicationListDeepAppInspection(o["deep-app-inspection"], d, "deep_app_inspection")); err != nil {
		if !fortiAPIPatch(o["deep-app-inspection"]) {
			return fmt.Errorf("Error reading deep_app_inspection: %v", err)
		}
	}

	if err = d.Set("options", flattenApplicationListOptions(o["options"], d, "options")); err != nil {
		if !fortiAPIPatch(o["options"]) {
			return fmt.Errorf("Error reading options: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("entries", flattenApplicationListEntries(o["entries"], d, "entries")); err != nil {
            if !fortiAPIPatch(o["entries"]) {
                return fmt.Errorf("Error reading entries: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("entries"); ok {
            if err = d.Set("entries", flattenApplicationListEntries(o["entries"], d, "entries")); err != nil {
                if !fortiAPIPatch(o["entries"]) {
                    return fmt.Errorf("Error reading entries: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenApplicationListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandApplicationListName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListComment(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListExtendedLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListOtherApplicationAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListAppReplacemsg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListOtherApplicationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEnforceDefaultAppPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListUnknownApplicationAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListUnknownApplicationLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListP2PBlackList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListDeepAppInspection(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListOptions(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntries(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandApplicationListEntriesId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "risk"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["risk"], _ = expandApplicationListEntriesRisk(d, i["risk"], pre_append)
		} else {
			tmp["risk"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["category"], _ = expandApplicationListEntriesCategory(d, i["category"], pre_append)
		} else {
			tmp["category"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "sub_category"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["sub-category"], _ = expandApplicationListEntriesSubCategory(d, i["sub_category"], pre_append)
		} else {
			tmp["sub-category"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "application"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["application"], _ = expandApplicationListEntriesApplication(d, i["application"], pre_append)
		} else {
			tmp["application"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocols"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocols"], _ = expandApplicationListEntriesProtocols(d, i["protocols"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vendor"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vendor"], _ = expandApplicationListEntriesVendor(d, i["vendor"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "technology"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["technology"], _ = expandApplicationListEntriesTechnology(d, i["technology"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "behavior"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["behavior"], _ = expandApplicationListEntriesBehavior(d, i["behavior"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "popularity"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["popularity"], _ = expandApplicationListEntriesPopularity(d, i["popularity"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "parameters"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["parameters"], _ = expandApplicationListEntriesParameters(d, i["parameters"], pre_append)
		} else {
			tmp["parameters"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandApplicationListEntriesAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandApplicationListEntriesLog(d, i["log"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log_packet"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log-packet"], _ = expandApplicationListEntriesLogPacket(d, i["log_packet"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_count"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-count"], _ = expandApplicationListEntriesRateCount(d, i["rate_count"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_duration"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-duration"], _ = expandApplicationListEntriesRateDuration(d, i["rate_duration"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-mode"], _ = expandApplicationListEntriesRateMode(d, i["rate_mode"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rate_track"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rate-track"], _ = expandApplicationListEntriesRateTrack(d, i["rate_track"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "session_ttl"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["session-ttl"], _ = expandApplicationListEntriesSessionTtl(d, i["session_ttl"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["shaper"], _ = expandApplicationListEntriesShaper(d, i["shaper"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "shaper_reverse"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["shaper-reverse"], _ = expandApplicationListEntriesShaperReverse(d, i["shaper_reverse"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "per_ip_shaper"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["per-ip-shaper"], _ = expandApplicationListEntriesPerIpShaper(d, i["per_ip_shaper"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine"], _ = expandApplicationListEntriesQuarantine(d, i["quarantine"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_expiry"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-expiry"], _ = expandApplicationListEntriesQuarantineExpiry(d, i["quarantine_expiry"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "quarantine_log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["quarantine-log"], _ = expandApplicationListEntriesQuarantineLog(d, i["quarantine_log"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRisk(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "level"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["level"], _ = expandApplicationListEntriesRiskLevel(d, i["level"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesRiskLevel(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandApplicationListEntriesCategoryId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesCategoryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesSubCategory(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandApplicationListEntriesSubCategoryId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesSubCategoryId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesApplication(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandApplicationListEntriesApplicationId(d, i["id"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesApplicationId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesProtocols(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesVendor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesTechnology(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesBehavior(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesPopularity(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParameters(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandApplicationListEntriesParametersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandApplicationListEntriesParametersValue(d, i["value"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandApplicationListEntriesParametersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesParametersValue(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesLogPacket(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateCount(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateDuration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesRateTrack(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesSessionTtl(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesShaperReverse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesPerIpShaper(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantine(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantineExpiry(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandApplicationListEntriesQuarantineLog(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectApplicationList(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("name"); ok {
		t, err := expandApplicationListName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandApplicationListComment(d, v, "comment")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandApplicationListReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("extended_log"); ok {
		t, err := expandApplicationListExtendedLog(d, v, "extended_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extended-log"] = t
		}
	}

	if v, ok := d.GetOk("other_application_action"); ok {
		t, err := expandApplicationListOtherApplicationAction(d, v, "other_application_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-application-action"] = t
		}
	}

	if v, ok := d.GetOk("app_replacemsg"); ok {
		t, err := expandApplicationListAppReplacemsg(d, v, "app_replacemsg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["app-replacemsg"] = t
		}
	}

	if v, ok := d.GetOk("other_application_log"); ok {
		t, err := expandApplicationListOtherApplicationLog(d, v, "other_application_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["other-application-log"] = t
		}
	}

	if v, ok := d.GetOk("enforce_default_app_port"); ok {
		t, err := expandApplicationListEnforceDefaultAppPort(d, v, "enforce_default_app_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-default-app-port"] = t
		}
	}

	if v, ok := d.GetOk("unknown_application_action"); ok {
		t, err := expandApplicationListUnknownApplicationAction(d, v, "unknown_application_action")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-application-action"] = t
		}
	}

	if v, ok := d.GetOk("unknown_application_log"); ok {
		t, err := expandApplicationListUnknownApplicationLog(d, v, "unknown_application_log")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-application-log"] = t
		}
	}

	if v, ok := d.GetOk("p2p_black_list"); ok {
		t, err := expandApplicationListP2PBlackList(d, v, "p2p_black_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["p2p-black-list"] = t
		}
	}

	if v, ok := d.GetOk("deep_app_inspection"); ok {
		t, err := expandApplicationListDeepAppInspection(d, v, "deep_app_inspection")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["deep-app-inspection"] = t
		}
	}

	if v, ok := d.GetOk("options"); ok {
		t, err := expandApplicationListOptions(d, v, "options")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["options"] = t
		}
	}

	if v, ok := d.GetOk("entries"); ok {
		t, err := expandApplicationListEntries(d, v, "entries")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["entries"] = t
		}
	}


	return &obj, nil
}

