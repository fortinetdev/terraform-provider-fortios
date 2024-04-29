// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WAN optimization content delivery network rules.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWanoptContentDeliveryNetworkRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceWanoptContentDeliveryNetworkRuleCreate,
		Read:   resourceWanoptContentDeliveryNetworkRuleRead,
		Update: resourceWanoptContentDeliveryNetworkRuleUpdate,
		Delete: resourceWanoptContentDeliveryNetworkRuleDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"host_domain_name_suffix": &schema.Schema{
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
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_cache_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_cache_control": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_expires": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"text_response_vcache": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"updateserver": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"rules": &schema.Schema{
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
						"match_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"skip_rule_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"match_entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pattern": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"string": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
								},
							},
						},
						"skip_entries": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"pattern": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"string": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
													Computed:     true,
												},
											},
										},
									},
								},
							},
						},
						"content_id": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Optional: true,
							MaxItems: 1,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"start_str": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"start_skip": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"start_direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"end_str": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
									"end_skip": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"end_direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"range_str": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 35),
										Optional:     true,
										Computed:     true,
									},
								},
							},
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

func resourceWanoptContentDeliveryNetworkRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptContentDeliveryNetworkRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WanoptContentDeliveryNetworkRule resource while getting object: %v", err)
	}

	o, err := c.CreateWanoptContentDeliveryNetworkRule(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptContentDeliveryNetworkRule")
	}

	return resourceWanoptContentDeliveryNetworkRuleRead(d, m)
}

func resourceWanoptContentDeliveryNetworkRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWanoptContentDeliveryNetworkRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRule resource while getting object: %v", err)
	}

	o, err := c.UpdateWanoptContentDeliveryNetworkRule(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WanoptContentDeliveryNetworkRule")
	}

	return resourceWanoptContentDeliveryNetworkRuleRead(d, m)
}

func resourceWanoptContentDeliveryNetworkRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWanoptContentDeliveryNetworkRule(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWanoptContentDeliveryNetworkRuleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWanoptContentDeliveryNetworkRule(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWanoptContentDeliveryNetworkRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WanoptContentDeliveryNetworkRule resource from API: %v", err)
	}
	return nil
}

func flattenWanoptContentDeliveryNetworkRuleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWanoptContentDeliveryNetworkRuleHostDomainNameSuffixName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWanoptContentDeliveryNetworkRuleHostDomainNameSuffixName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRequestCacheControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleResponseCacheControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleResponseExpires(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleTextResponseVcache(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleUpdateserver(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenWanoptContentDeliveryNetworkRuleRulesName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_mode"
		if cur_v, ok := i["match-mode"]; ok {
			tmp["match_mode"] = flattenWanoptContentDeliveryNetworkRuleRulesMatchMode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "skip_rule_mode"
		if cur_v, ok := i["skip-rule-mode"]; ok {
			tmp["skip_rule_mode"] = flattenWanoptContentDeliveryNetworkRuleRulesSkipRuleMode(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_entries"
		if cur_v, ok := i["match-entries"]; ok {
			tmp["match_entries"] = flattenWanoptContentDeliveryNetworkRuleRulesMatchEntries(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "skip_entries"
		if cur_v, ok := i["skip-entries"]; ok {
			tmp["skip_entries"] = flattenWanoptContentDeliveryNetworkRuleRulesSkipEntries(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_id"
		if cur_v, ok := i["content-id"]; ok {
			tmp["content_id"] = flattenWanoptContentDeliveryNetworkRuleRulesContentId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipRuleMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if cur_v, ok := i["target"]; ok {
			tmp["target"] = flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if cur_v, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "string"
		if cur_v, ok := i["string"]; ok {
			tmp["string"] = flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPatternString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "string", d)
	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesMatchEntriesPatternString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntries(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if cur_v, ok := i["target"]; ok {
			tmp["target"] = flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if cur_v, ok := i["pattern"]; ok {
			tmp["pattern"] = flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "string"
		if cur_v, ok := i["string"]; ok {
			tmp["string"] = flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPatternString(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "string", d)
	return result
}

func flattenWanoptContentDeliveryNetworkRuleRulesSkipEntriesPatternString(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentId(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "target"
	if _, ok := i["target"]; ok {
		result["target"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdTarget(i["target"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "start_str"
	if _, ok := i["start-str"]; ok {
		result["start_str"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartStr(i["start-str"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "start_skip"
	if _, ok := i["start-skip"]; ok {
		result["start_skip"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip(i["start-skip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "start_direction"
	if _, ok := i["start-direction"]; ok {
		result["start_direction"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection(i["start-direction"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "end_str"
	if _, ok := i["end-str"]; ok {
		result["end_str"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndStr(i["end-str"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "end_skip"
	if _, ok := i["end-skip"]; ok {
		result["end_skip"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip(i["end-skip"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "end_direction"
	if _, ok := i["end-direction"]; ok {
		result["end_direction"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection(i["end-direction"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "range_str"
	if _, ok := i["range-str"]; ok {
		result["range_str"] = flattenWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr(i["range-str"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartStr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndStr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWanoptContentDeliveryNetworkRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWanoptContentDeliveryNetworkRuleName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWanoptContentDeliveryNetworkRuleComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if err = d.Set("status", flattenWanoptContentDeliveryNetworkRuleStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("host_domain_name_suffix", flattenWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(o["host-domain-name-suffix"], d, "host_domain_name_suffix", sv)); err != nil {
			if !fortiAPIPatch(o["host-domain-name-suffix"]) {
				return fmt.Errorf("Error reading host_domain_name_suffix: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("host_domain_name_suffix"); ok {
			if err = d.Set("host_domain_name_suffix", flattenWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(o["host-domain-name-suffix"], d, "host_domain_name_suffix", sv)); err != nil {
				if !fortiAPIPatch(o["host-domain-name-suffix"]) {
					return fmt.Errorf("Error reading host_domain_name_suffix: %v", err)
				}
			}
		}
	}

	if err = d.Set("category", flattenWanoptContentDeliveryNetworkRuleCategory(o["category"], d, "category", sv)); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("request_cache_control", flattenWanoptContentDeliveryNetworkRuleRequestCacheControl(o["request-cache-control"], d, "request_cache_control", sv)); err != nil {
		if !fortiAPIPatch(o["request-cache-control"]) {
			return fmt.Errorf("Error reading request_cache_control: %v", err)
		}
	}

	if err = d.Set("response_cache_control", flattenWanoptContentDeliveryNetworkRuleResponseCacheControl(o["response-cache-control"], d, "response_cache_control", sv)); err != nil {
		if !fortiAPIPatch(o["response-cache-control"]) {
			return fmt.Errorf("Error reading response_cache_control: %v", err)
		}
	}

	if err = d.Set("response_expires", flattenWanoptContentDeliveryNetworkRuleResponseExpires(o["response-expires"], d, "response_expires", sv)); err != nil {
		if !fortiAPIPatch(o["response-expires"]) {
			return fmt.Errorf("Error reading response_expires: %v", err)
		}
	}

	if err = d.Set("text_response_vcache", flattenWanoptContentDeliveryNetworkRuleTextResponseVcache(o["text-response-vcache"], d, "text_response_vcache", sv)); err != nil {
		if !fortiAPIPatch(o["text-response-vcache"]) {
			return fmt.Errorf("Error reading text_response_vcache: %v", err)
		}
	}

	if err = d.Set("updateserver", flattenWanoptContentDeliveryNetworkRuleUpdateserver(o["updateserver"], d, "updateserver", sv)); err != nil {
		if !fortiAPIPatch(o["updateserver"]) {
			return fmt.Errorf("Error reading updateserver: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("rules", flattenWanoptContentDeliveryNetworkRuleRules(o["rules"], d, "rules", sv)); err != nil {
			if !fortiAPIPatch(o["rules"]) {
				return fmt.Errorf("Error reading rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("rules"); ok {
			if err = d.Set("rules", flattenWanoptContentDeliveryNetworkRuleRules(o["rules"], d, "rules", sv)); err != nil {
				if !fortiAPIPatch(o["rules"]) {
					return fmt.Errorf("Error reading rules: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWanoptContentDeliveryNetworkRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWanoptContentDeliveryNetworkRuleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWanoptContentDeliveryNetworkRuleHostDomainNameSuffixName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleHostDomainNameSuffixName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRequestCacheControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleResponseCacheControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleResponseExpires(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleTextResponseVcache(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleUpdateserver(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandWanoptContentDeliveryNetworkRuleRulesName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-mode"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchMode(d, i["match_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "skip_rule_mode"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["skip-rule-mode"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipRuleMode(d, i["skip_rule_mode"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_entries"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["match-entries"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntries(d, i["match_entries"], pre_append, sv)
		} else {
			tmp["match-entries"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "skip_entries"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["skip-entries"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntries(d, i["skip_entries"], pre_append, sv)
		} else {
			tmp["skip-entries"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content-id"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentId(d, i["content_id"], pre_append, sv)
		} else {
			tmp["content-id"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipRuleMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["target"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget(d, i["target"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(d, i["pattern"], pre_append, sv)
		} else {
			tmp["pattern"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["string"], _ = expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPatternString(d, i["string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesMatchEntriesPatternString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntries(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["target"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget(d, i["target"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "pattern"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["pattern"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(d, i["pattern"], pre_append, sv)
		} else {
			tmp["pattern"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["string"], _ = expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPatternString(d, i["string"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesSkipEntriesPatternString(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "target"
	if _, ok := d.GetOk(pre_append); ok {
		result["target"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdTarget(d, i["target"], pre_append, sv)
	}
	pre_append = pre + ".0." + "start_str"
	if _, ok := d.GetOk(pre_append); ok {
		result["start-str"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdStartStr(d, i["start_str"], pre_append, sv)
	}
	pre_append = pre + ".0." + "start_skip"
	if _, ok := d.GetOk(pre_append); ok {
		result["start-skip"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip(d, i["start_skip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "start_direction"
	if _, ok := d.GetOk(pre_append); ok {
		result["start-direction"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection(d, i["start_direction"], pre_append, sv)
	}
	pre_append = pre + ".0." + "end_str"
	if _, ok := d.GetOk(pre_append); ok {
		result["end-str"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdEndStr(d, i["end_str"], pre_append, sv)
	}
	pre_append = pre + ".0." + "end_skip"
	if _, ok := d.GetOk(pre_append); ok {
		result["end-skip"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip(d, i["end_skip"], pre_append, sv)
	}
	pre_append = pre + ".0." + "end_direction"
	if _, ok := d.GetOk(pre_append); ok {
		result["end-direction"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection(d, i["end_direction"], pre_append, sv)
	}
	pre_append = pre + ".0." + "range_str"
	if _, ok := d.GetOk(pre_append); ok {
		result["range-str"], _ = expandWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr(d, i["range_str"], pre_append, sv)
	}

	return result, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartStr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartSkip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdStartDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndStr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndSkip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdEndDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWanoptContentDeliveryNetworkRuleRulesContentIdRangeStr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWanoptContentDeliveryNetworkRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWanoptContentDeliveryNetworkRuleName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWanoptContentDeliveryNetworkRuleComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandWanoptContentDeliveryNetworkRuleStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("host_domain_name_suffix"); ok || d.HasChange("host_domain_name_suffix") {
		t, err := expandWanoptContentDeliveryNetworkRuleHostDomainNameSuffix(d, v, "host_domain_name_suffix", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host-domain-name-suffix"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok {
		t, err := expandWanoptContentDeliveryNetworkRuleCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("request_cache_control"); ok {
		t, err := expandWanoptContentDeliveryNetworkRuleRequestCacheControl(d, v, "request_cache_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-cache-control"] = t
		}
	}

	if v, ok := d.GetOk("response_cache_control"); ok {
		t, err := expandWanoptContentDeliveryNetworkRuleResponseCacheControl(d, v, "response_cache_control", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-cache-control"] = t
		}
	}

	if v, ok := d.GetOk("response_expires"); ok {
		t, err := expandWanoptContentDeliveryNetworkRuleResponseExpires(d, v, "response_expires", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-expires"] = t
		}
	}

	if v, ok := d.GetOk("text_response_vcache"); ok {
		t, err := expandWanoptContentDeliveryNetworkRuleTextResponseVcache(d, v, "text_response_vcache", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["text-response-vcache"] = t
		}
	}

	if v, ok := d.GetOk("updateserver"); ok {
		t, err := expandWanoptContentDeliveryNetworkRuleUpdateserver(d, v, "updateserver", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["updateserver"] = t
		}
	}

	if v, ok := d.GetOk("rules"); ok || d.HasChange("rules") {
		t, err := expandWanoptContentDeliveryNetworkRuleRules(d, v, "rules", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["rules"] = t
		}
	}

	return &obj, nil
}
