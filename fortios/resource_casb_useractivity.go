// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure CASB user activity.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbUserActivity() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbUserActivityCreate,
		Read:   resourceCasbUserActivityRead,
		Update: resourceCasbUserActivityUpdate,
		Delete: resourceCasbUserActivityDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 79),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"uuid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 36),
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"casb_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"application": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match_strategy": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"match": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"strategy": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"rules": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"domains": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"domain": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 127),
													Optional:     true,
												},
											},
										},
									},
									"methods": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"method": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
												},
											},
										},
									},
									"match_pattern": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"match_value": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 1023),
										Optional:     true,
									},
									"header_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
									"body_type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"jq": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
									"case_sensitive": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"negate": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"tenant_extraction": &schema.Schema{
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
									"type": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"jq": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 1023),
										Optional:     true,
									},
									"filters": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"direction": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"place": &schema.Schema{
													Type:     schema.TypeString,
													Optional: true,
													Computed: true,
												},
												"header_name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 255),
													Optional:     true,
												},
												"body_type": &schema.Schema{
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
					},
				},
			},
			"control_options": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"operations": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
									"target": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"direction": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"header_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
									},
									"search_pattern": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"search_key": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 1023),
										Optional:     true,
									},
									"case_sensitive": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value_from_input": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"value_name_from_input": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
									"values": &schema.Schema{
										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"value": &schema.Schema{
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

func resourceCasbUserActivityCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbUserActivity(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CasbUserActivity resource while getting object: %v", err)
	}

	o, err := c.CreateCasbUserActivity(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating CasbUserActivity resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CasbUserActivity")
	}

	return resourceCasbUserActivityRead(d, m)
}

func resourceCasbUserActivityUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbUserActivity(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivity resource while getting object: %v", err)
	}

	o, err := c.UpdateCasbUserActivity(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating CasbUserActivity resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CasbUserActivity")
	}

	return resourceCasbUserActivityRead(d, m)
}

func resourceCasbUserActivityDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCasbUserActivity(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting CasbUserActivity resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbUserActivityRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbUserActivity(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CasbUserActivity resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbUserActivity(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CasbUserActivity resource from API: %v", err)
	}
	return nil
}

func flattenCasbUserActivityName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityUuid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityCasbName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityApplication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchStrategyU(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatch(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenCasbUserActivityMatchId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["strategy"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "strategy"
			}
			tmp["strategy"] = flattenCasbUserActivityMatchStrategy(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["rules"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "rules"
			}
			tmp["rules"] = flattenCasbUserActivityMatchRules(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["tenant-extraction"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_extraction"
			}
			tmp["tenant_extraction"] = flattenCasbUserActivityMatchTenantExtraction(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenCasbUserActivityMatchId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCasbUserActivityMatchStrategy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenCasbUserActivityMatchRulesId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
			}
			tmp["type"] = flattenCasbUserActivityMatchRulesType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["domains"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
			}
			tmp["domains"] = flattenCasbUserActivityMatchRulesDomains(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["methods"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "methods"
			}
			tmp["methods"] = flattenCasbUserActivityMatchRulesMethods(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["match-pattern"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
			}
			tmp["match_pattern"] = flattenCasbUserActivityMatchRulesMatchPattern(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["match-value"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
			}
			tmp["match_value"] = flattenCasbUserActivityMatchRulesMatchValue(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["header-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
			}
			tmp["header_name"] = flattenCasbUserActivityMatchRulesHeaderName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["body-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
			}
			tmp["body_type"] = flattenCasbUserActivityMatchRulesBodyType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["jq"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "jq"
			}
			tmp["jq"] = flattenCasbUserActivityMatchRulesJq(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["case-sensitive"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
			}
			tmp["case_sensitive"] = flattenCasbUserActivityMatchRulesCaseSensitive(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["negate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
			}
			tmp["negate"] = flattenCasbUserActivityMatchRulesNegate(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenCasbUserActivityMatchRulesId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCasbUserActivityMatchRulesType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesDomains(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "domain", "domain")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["domain"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "domain"
			}
			tmp["domain"] = flattenCasbUserActivityMatchRulesDomainsDomain(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "domain", d)
	return result
}

func flattenCasbUserActivityMatchRulesDomainsDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMethods(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "method", "method")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["method"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "method"
			}
			tmp["method"] = flattenCasbUserActivityMatchRulesMethodsMethod(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "method", d)
	return result
}

func flattenCasbUserActivityMatchRulesMethodsMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMatchPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesMatchValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesHeaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesBodyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesJq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesCaseSensitive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchRulesNegate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtraction(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenCasbUserActivityMatchTenantExtractionStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "type"
	if _, ok := i["type"]; ok {
		result["type"] = flattenCasbUserActivityMatchTenantExtractionType(i["type"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "jq"
	if _, ok := i["jq"]; ok {
		result["jq"] = flattenCasbUserActivityMatchTenantExtractionJq(i["jq"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "filters"
	if _, ok := i["filters"]; ok {
		result["filters"] = flattenCasbUserActivityMatchTenantExtractionFilters(i["filters"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenCasbUserActivityMatchTenantExtractionStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionJq(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFilters(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "id", "id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
			}
			tmp["id"] = flattenCasbUserActivityMatchTenantExtractionFiltersId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["direction"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
			}
			tmp["direction"] = flattenCasbUserActivityMatchTenantExtractionFiltersDirection(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["place"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
			}
			tmp["place"] = flattenCasbUserActivityMatchTenantExtractionFiltersPlace(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["header-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
			}
			tmp["header_name"] = flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["body-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
			}
			tmp["body_type"] = flattenCasbUserActivityMatchTenantExtractionFiltersBodyType(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenCasbUserActivityMatchTenantExtractionFiltersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCasbUserActivityMatchTenantExtractionFiltersDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersPlace(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersHeaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityMatchTenantExtractionFiltersBodyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptions(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenCasbUserActivityControlOptionsName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["status"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
			}
			tmp["status"] = flattenCasbUserActivityControlOptionsStatus(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["operations"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "operations"
			}
			tmp["operations"] = flattenCasbUserActivityControlOptionsOperations(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbUserActivityControlOptionsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperations(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenCasbUserActivityControlOptionsOperationsName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["target"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
			}
			tmp["target"] = flattenCasbUserActivityControlOptionsOperationsTarget(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["action"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
			}
			tmp["action"] = flattenCasbUserActivityControlOptionsOperationsAction(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["direction"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
			}
			tmp["direction"] = flattenCasbUserActivityControlOptionsOperationsDirection(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["header-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
			}
			tmp["header_name"] = flattenCasbUserActivityControlOptionsOperationsHeaderName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["search-pattern"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "search_pattern"
			}
			tmp["search_pattern"] = flattenCasbUserActivityControlOptionsOperationsSearchPattern(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["search-key"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "search_key"
			}
			tmp["search_key"] = flattenCasbUserActivityControlOptionsOperationsSearchKey(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["case-sensitive"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
			}
			tmp["case_sensitive"] = flattenCasbUserActivityControlOptionsOperationsCaseSensitive(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["value-from-input"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "value_from_input"
			}
			tmp["value_from_input"] = flattenCasbUserActivityControlOptionsOperationsValueFromInput(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["value-name-from-input"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "value_name_from_input"
			}
			tmp["value_name_from_input"] = flattenCasbUserActivityControlOptionsOperationsValueNameFromInput(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["values"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
			}
			tmp["values"] = flattenCasbUserActivityControlOptionsOperationsValues(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbUserActivityControlOptionsOperationsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsDirection(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsHeaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsSearchPattern(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsSearchKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsCaseSensitive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsValueFromInput(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsValueNameFromInput(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbUserActivityControlOptionsOperationsValues(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "value", "value")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["value"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
			}
			tmp["value"] = flattenCasbUserActivityControlOptionsOperationsValuesValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "value", d)
	return result
}

func flattenCasbUserActivityControlOptionsOperationsValuesValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCasbUserActivity(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenCasbUserActivityName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("uuid", flattenCasbUserActivityUuid(o["uuid"], d, "uuid", sv)); err != nil {
		if !fortiAPIPatch(o["uuid"]) {
			return fmt.Errorf("Error reading uuid: %v", err)
		}
	}

	if err = d.Set("status", flattenCasbUserActivityStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("description", flattenCasbUserActivityDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("type", flattenCasbUserActivityType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("casb_name", flattenCasbUserActivityCasbName(o["casb-name"], d, "casb_name", sv)); err != nil {
		if !fortiAPIPatch(o["casb-name"]) {
			return fmt.Errorf("Error reading casb_name: %v", err)
		}
	}

	if err = d.Set("application", flattenCasbUserActivityApplication(o["application"], d, "application", sv)); err != nil {
		if !fortiAPIPatch(o["application"]) {
			return fmt.Errorf("Error reading application: %v", err)
		}
	}

	if err = d.Set("category", flattenCasbUserActivityCategory(o["category"], d, "category", sv)); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("match_strategy", flattenCasbUserActivityMatchStrategyU(o["match-strategy"], d, "match_strategy", sv)); err != nil {
		if !fortiAPIPatch(o["match-strategy"]) {
			return fmt.Errorf("Error reading match_strategy: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("match", flattenCasbUserActivityMatch(o["match"], d, "match", sv)); err != nil {
			if !fortiAPIPatch(o["match"]) {
				return fmt.Errorf("Error reading match: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("match"); ok {
			if err = d.Set("match", flattenCasbUserActivityMatch(o["match"], d, "match", sv)); err != nil {
				if !fortiAPIPatch(o["match"]) {
					return fmt.Errorf("Error reading match: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("control_options", flattenCasbUserActivityControlOptions(o["control-options"], d, "control_options", sv)); err != nil {
			if !fortiAPIPatch(o["control-options"]) {
				return fmt.Errorf("Error reading control_options: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("control_options"); ok {
			if err = d.Set("control_options", flattenCasbUserActivityControlOptions(o["control-options"], d, "control_options", sv)); err != nil {
				if !fortiAPIPatch(o["control-options"]) {
					return fmt.Errorf("Error reading control_options: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenCasbUserActivityFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCasbUserActivityName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityUuid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityCasbName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchStrategyU(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandCasbUserActivityMatchId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "strategy"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["strategy"], _ = expandCasbUserActivityMatchStrategy(d, i["strategy"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rules"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rules"], _ = expandCasbUserActivityMatchRules(d, i["rules"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rules"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_extraction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tenant-extraction"], _ = expandCasbUserActivityMatchTenantExtraction(d, i["tenant_extraction"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["tenant-extraction"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchStrategy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandCasbUserActivityMatchRulesId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["type"], _ = expandCasbUserActivityMatchRulesType(d, i["type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domains"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domains"], _ = expandCasbUserActivityMatchRulesDomains(d, i["domains"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["domains"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "methods"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["methods"], _ = expandCasbUserActivityMatchRulesMethods(d, i["methods"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["methods"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-pattern"], _ = expandCasbUserActivityMatchRulesMatchPattern(d, i["match_pattern"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "match_value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["match-value"], _ = expandCasbUserActivityMatchRulesMatchValue(d, i["match_value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["match-value"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header-name"], _ = expandCasbUserActivityMatchRulesHeaderName(d, i["header_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["header-name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["body-type"], _ = expandCasbUserActivityMatchRulesBodyType(d, i["body_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "jq"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["jq"], _ = expandCasbUserActivityMatchRulesJq(d, i["jq"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["jq"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["case-sensitive"], _ = expandCasbUserActivityMatchRulesCaseSensitive(d, i["case_sensitive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "negate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["negate"], _ = expandCasbUserActivityMatchRulesNegate(d, i["negate"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchRulesId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesDomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["domain"], _ = expandCasbUserActivityMatchRulesDomainsDomain(d, i["domain"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchRulesDomainsDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMethods(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["method"], _ = expandCasbUserActivityMatchRulesMethodsMethod(d, i["method"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchRulesMethodsMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMatchPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesMatchValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesHeaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesBodyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesJq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesCaseSensitive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchRulesNegate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtraction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandCasbUserActivityMatchTenantExtractionStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "type"
	if _, ok := d.GetOk(pre_append); ok {
		result["type"], _ = expandCasbUserActivityMatchTenantExtractionType(d, i["type"], pre_append, sv)
	}
	pre_append = pre + ".0." + "jq"
	if _, ok := d.GetOk(pre_append); ok {
		result["jq"], _ = expandCasbUserActivityMatchTenantExtractionJq(d, i["jq"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["jq"] = nil
	}
	pre_append = pre + ".0." + "filters"
	if _, ok := d.GetOk(pre_append); ok {
		result["filters"], _ = expandCasbUserActivityMatchTenantExtractionFilters(d, i["filters"], pre_append, sv)
	} else {
		result["filters"] = make([]string, 0)
	}

	return result, nil
}

func expandCasbUserActivityMatchTenantExtractionStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionJq(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFilters(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandCasbUserActivityMatchTenantExtractionFiltersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["direction"], _ = expandCasbUserActivityMatchTenantExtractionFiltersDirection(d, i["direction"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "place"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["place"], _ = expandCasbUserActivityMatchTenantExtractionFiltersPlace(d, i["place"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header-name"], _ = expandCasbUserActivityMatchTenantExtractionFiltersHeaderName(d, i["header_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["header-name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "body_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["body-type"], _ = expandCasbUserActivityMatchTenantExtractionFiltersBodyType(d, i["body_type"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersPlace(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersHeaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityMatchTenantExtractionFiltersBodyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptions(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbUserActivityControlOptionsName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandCasbUserActivityControlOptionsStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "operations"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["operations"], _ = expandCasbUserActivityControlOptionsOperations(d, i["operations"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["operations"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityControlOptionsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperations(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbUserActivityControlOptionsOperationsName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["target"], _ = expandCasbUserActivityControlOptionsOperationsTarget(d, i["target"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandCasbUserActivityControlOptionsOperationsAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "direction"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["direction"], _ = expandCasbUserActivityControlOptionsOperationsDirection(d, i["direction"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header-name"], _ = expandCasbUserActivityControlOptionsOperationsHeaderName(d, i["header_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["header-name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_pattern"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["search-pattern"], _ = expandCasbUserActivityControlOptionsOperationsSearchPattern(d, i["search_pattern"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "search_key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["search-key"], _ = expandCasbUserActivityControlOptionsOperationsSearchKey(d, i["search_key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["search-key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitive"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["case-sensitive"], _ = expandCasbUserActivityControlOptionsOperationsCaseSensitive(d, i["case_sensitive"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_from_input"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value-from-input"], _ = expandCasbUserActivityControlOptionsOperationsValueFromInput(d, i["value_from_input"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value_name_from_input"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value-name-from-input"], _ = expandCasbUserActivityControlOptionsOperationsValueNameFromInput(d, i["value_name_from_input"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value-name-from-input"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "values"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["values"], _ = expandCasbUserActivityControlOptionsOperationsValues(d, i["values"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["values"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityControlOptionsOperationsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsDirection(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsHeaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsSearchPattern(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsSearchKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsCaseSensitive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsValueFromInput(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsValueNameFromInput(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbUserActivityControlOptionsOperationsValues(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["value"], _ = expandCasbUserActivityControlOptionsOperationsValuesValue(d, i["value"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbUserActivityControlOptionsOperationsValuesValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCasbUserActivity(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandCasbUserActivityName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("uuid"); ok {
		t, err := expandCasbUserActivityUuid(d, v, "uuid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uuid"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandCasbUserActivityStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandCasbUserActivityDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandCasbUserActivityType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("casb_name"); ok {
		t, err := expandCasbUserActivityCasbName(d, v, "casb_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["casb-name"] = t
		}
	} else if d.HasChange("casb_name") {
		obj["casb-name"] = nil
	}

	if v, ok := d.GetOk("application"); ok {
		t, err := expandCasbUserActivityApplication(d, v, "application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["application"] = t
		}
	} else if d.HasChange("application") {
		obj["application"] = nil
	}

	if v, ok := d.GetOk("category"); ok {
		t, err := expandCasbUserActivityCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("match_strategy"); ok {
		t, err := expandCasbUserActivityMatchStrategyU(d, v, "match_strategy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match-strategy"] = t
		}
	}

	if v, ok := d.GetOk("match"); ok || d.HasChange("match") {
		t, err := expandCasbUserActivityMatch(d, v, "match", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["match"] = t
		}
	}

	if v, ok := d.GetOk("control_options"); ok || d.HasChange("control_options") {
		t, err := expandCasbUserActivityControlOptions(d, v, "control_options", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["control-options"] = t
		}
	}

	return &obj, nil
}
