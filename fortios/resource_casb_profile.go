// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure CASB profile.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceCasbProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceCasbProfileCreate,
		Read:   resourceCasbProfileRead,
		Update: resourceCasbProfileUpdate,
		Delete: resourceCasbProfileDelete,

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
				Optional:     true,
				Computed:     true,
			},
			"comment": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"saas_application": &schema.Schema{
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
						"safe_search": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"safe_search_control": &schema.Schema{
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
						"tenant_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"tenant_control_tenants": &schema.Schema{
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
						"advanced_tenant_control": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
									"attribute": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
												},
												"input": &schema.Schema{
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
						"domain_control": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"domain_control_domains": &schema.Schema{
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
						"log": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"access_rule": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
									"action": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
									"bypass": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
									},
									"attribute_filter": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"attribute_match": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
												},
												"action": &schema.Schema{
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
						"custom_control": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
									},
									"option": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"name": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
												},
												"user_input": &schema.Schema{
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
									"attribute_filter": &schema.Schema{
										Type:     schema.TypeList,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{
												"id": &schema.Schema{
													Type:     schema.TypeInt,
													Optional: true,
													Computed: true,
												},
												"attribute_match": &schema.Schema{
													Type:         schema.TypeString,
													ValidateFunc: validation.StringLenBetween(0, 79),
													Optional:     true,
												},
												"action": &schema.Schema{
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

func resourceCasbProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CasbProfile resource while getting object: %v", err)
	}

	o, err := c.CreateCasbProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating CasbProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CasbProfile")
	}

	return resourceCasbProfileRead(d, m)
}

func resourceCasbProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectCasbProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateCasbProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating CasbProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("CasbProfile")
	}

	return resourceCasbProfileRead(d, m)
}

func resourceCasbProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCasbProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting CasbProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceCasbProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadCasbProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CasbProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCasbProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CasbProfile resource from API: %v", err)
	}
	return nil
}

func flattenCasbProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplication(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenCasbProfileSaasApplicationName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if cur_v, ok := i["status"]; ok {
			tmp["status"] = flattenCasbProfileSaasApplicationStatus(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search"
		if cur_v, ok := i["safe-search"]; ok {
			tmp["safe_search"] = flattenCasbProfileSaasApplicationSafeSearch(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search_control"
		if cur_v, ok := i["safe-search-control"]; ok {
			tmp["safe_search_control"] = flattenCasbProfileSaasApplicationSafeSearchControl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control"
		if cur_v, ok := i["tenant-control"]; ok {
			tmp["tenant_control"] = flattenCasbProfileSaasApplicationTenantControl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control_tenants"
		if cur_v, ok := i["tenant-control-tenants"]; ok {
			tmp["tenant_control_tenants"] = flattenCasbProfileSaasApplicationTenantControlTenants(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advanced_tenant_control"
		if cur_v, ok := i["advanced-tenant-control"]; ok {
			tmp["advanced_tenant_control"] = flattenCasbProfileSaasApplicationAdvancedTenantControl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control"
		if cur_v, ok := i["domain-control"]; ok {
			tmp["domain_control"] = flattenCasbProfileSaasApplicationDomainControl(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control_domains"
		if cur_v, ok := i["domain-control-domains"]; ok {
			tmp["domain_control_domains"] = flattenCasbProfileSaasApplicationDomainControlDomains(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if cur_v, ok := i["log"]; ok {
			tmp["log"] = flattenCasbProfileSaasApplicationLog(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_rule"
		if cur_v, ok := i["access-rule"]; ok {
			tmp["access_rule"] = flattenCasbProfileSaasApplicationAccessRule(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom_control"
		if cur_v, ok := i["custom-control"]; ok {
			tmp["custom_control"] = flattenCasbProfileSaasApplicationCustomControl(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbProfileSaasApplicationName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationSafeSearch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationSafeSearchControl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenCasbProfileSaasApplicationSafeSearchControlName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbProfileSaasApplicationSafeSearchControlName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationTenantControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationTenantControlTenants(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenCasbProfileSaasApplicationTenantControlTenantsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbProfileSaasApplicationTenantControlTenantsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAdvancedTenantControl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenCasbProfileSaasApplicationAdvancedTenantControlName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute"
		if cur_v, ok := i["attribute"]; ok {
			tmp["attribute"] = flattenCasbProfileSaasApplicationAdvancedTenantControlAttribute(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttribute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input"
		if cur_v, ok := i["input"]; ok {
			tmp["input"] = flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeInputValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "value", d)
	return result
}

func flattenCasbProfileSaasApplicationAdvancedTenantControlAttributeInputValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationDomainControl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationDomainControlDomains(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenCasbProfileSaasApplicationDomainControlDomainsName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbProfileSaasApplicationDomainControlDomainsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenCasbProfileSaasApplicationAccessRuleName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenCasbProfileSaasApplicationAccessRuleAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if cur_v, ok := i["bypass"]; ok {
			tmp["bypass"] = flattenCasbProfileSaasApplicationAccessRuleBypass(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if cur_v, ok := i["attribute-filter"]; ok {
			tmp["attribute_filter"] = flattenCasbProfileSaasApplicationAccessRuleAttributeFilter(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbProfileSaasApplicationAccessRuleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleBypass(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if cur_v, ok := i["attribute-match"]; ok {
			tmp["attribute_match"] = flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationAccessRuleAttributeFilterAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControl(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenCasbProfileSaasApplicationCustomControlName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if cur_v, ok := i["option"]; ok {
			tmp["option"] = flattenCasbProfileSaasApplicationCustomControlOption(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if cur_v, ok := i["attribute-filter"]; ok {
			tmp["attribute_filter"] = flattenCasbProfileSaasApplicationCustomControlAttributeFilter(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbProfileSaasApplicationCustomControlName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlOption(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenCasbProfileSaasApplicationCustomControlOptionName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if cur_v, ok := i["user-input"]; ok {
			tmp["user_input"] = flattenCasbProfileSaasApplicationCustomControlOptionUserInput(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenCasbProfileSaasApplicationCustomControlOptionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlOptionUserInput(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenCasbProfileSaasApplicationCustomControlOptionUserInputValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "value", d)
	return result
}

func flattenCasbProfileSaasApplicationCustomControlOptionUserInputValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilter(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenCasbProfileSaasApplicationCustomControlAttributeFilterId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if cur_v, ok := i["attribute-match"]; ok {
			tmp["attribute_match"] = flattenCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenCasbProfileSaasApplicationCustomControlAttributeFilterAction(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenCasbProfileSaasApplicationCustomControlAttributeFilterAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCasbProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenCasbProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenCasbProfileComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("saas_application", flattenCasbProfileSaasApplication(o["saas-application"], d, "saas_application", sv)); err != nil {
			if !fortiAPIPatch(o["saas-application"]) {
				return fmt.Errorf("Error reading saas_application: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("saas_application"); ok {
			if err = d.Set("saas_application", flattenCasbProfileSaasApplication(o["saas-application"], d, "saas_application", sv)); err != nil {
				if !fortiAPIPatch(o["saas-application"]) {
					return fmt.Errorf("Error reading saas_application: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenCasbProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandCasbProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbProfileSaasApplicationName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "status"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["status"], _ = expandCasbProfileSaasApplicationStatus(d, i["status"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["safe-search"], _ = expandCasbProfileSaasApplicationSafeSearch(d, i["safe_search"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "safe_search_control"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["safe-search-control"], _ = expandCasbProfileSaasApplicationSafeSearchControl(d, i["safe_search_control"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["safe-search-control"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tenant-control"], _ = expandCasbProfileSaasApplicationTenantControl(d, i["tenant_control"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "tenant_control_tenants"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["tenant-control-tenants"], _ = expandCasbProfileSaasApplicationTenantControlTenants(d, i["tenant_control_tenants"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["tenant-control-tenants"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "advanced_tenant_control"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["advanced-tenant-control"], _ = expandCasbProfileSaasApplicationAdvancedTenantControl(d, i["advanced_tenant_control"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["advanced-tenant-control"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain-control"], _ = expandCasbProfileSaasApplicationDomainControl(d, i["domain_control"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "domain_control_domains"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["domain-control-domains"], _ = expandCasbProfileSaasApplicationDomainControlDomains(d, i["domain_control_domains"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["domain-control-domains"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "log"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["log"], _ = expandCasbProfileSaasApplicationLog(d, i["log"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_rule"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["access-rule"], _ = expandCasbProfileSaasApplicationAccessRule(d, i["access_rule"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["access-rule"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "custom_control"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["custom-control"], _ = expandCasbProfileSaasApplicationCustomControl(d, i["custom_control"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["custom-control"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationSafeSearch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationSafeSearchControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandCasbProfileSaasApplicationSafeSearchControlName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationSafeSearchControlName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationTenantControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationTenantControlTenants(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandCasbProfileSaasApplicationTenantControlTenantsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationTenantControlTenantsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlAttribute(d, i["attribute"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttribute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlAttributeName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "input"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["input"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(d, i["input"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["input"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttributeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttributeInput(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["value"], _ = expandCasbProfileSaasApplicationAdvancedTenantControlAttributeInputValue(d, i["value"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAdvancedTenantControlAttributeInputValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationDomainControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationDomainControlDomains(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandCasbProfileSaasApplicationDomainControlDomainsName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationDomainControlDomainsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbProfileSaasApplicationAccessRuleName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandCasbProfileSaasApplicationAccessRuleAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "bypass"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["bypass"], _ = expandCasbProfileSaasApplicationAccessRuleBypass(d, i["bypass"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["bypass"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-filter"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilter(d, i["attribute_filter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-filter"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAccessRuleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleBypass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-match"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(d, i["attribute_match"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-match"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction(d, i["action"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAttributeMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationAccessRuleAttributeFilterAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbProfileSaasApplicationCustomControlName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "option"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["option"], _ = expandCasbProfileSaasApplicationCustomControlOption(d, i["option"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["option"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_filter"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-filter"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilter(d, i["attribute_filter"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-filter"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandCasbProfileSaasApplicationCustomControlOptionName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "user_input"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["user-input"], _ = expandCasbProfileSaasApplicationCustomControlOptionUserInput(d, i["user_input"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["user-input"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionUserInput(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["value"], _ = expandCasbProfileSaasApplicationCustomControlOptionUserInputValue(d, i["value"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlOptionUserInputValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilter(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "attribute_match"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["attribute-match"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(d, i["attribute_match"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["attribute-match"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandCasbProfileSaasApplicationCustomControlAttributeFilterAction(d, i["action"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterAttributeMatch(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCasbProfileSaasApplicationCustomControlAttributeFilterAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCasbProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandCasbProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandCasbProfileComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("saas_application"); ok || d.HasChange("saas_application") {
		t, err := expandCasbProfileSaasApplication(d, v, "saas_application", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["saas-application"] = t
		}
	}

	return &obj, nil
}
