// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure Authentication Rules.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceAuthenticationRule() *schema.Resource {
	return &schema.Resource{
		Create: resourceAuthenticationRuleCreate,
		Read:   resourceAuthenticationRuleRead,
		Update: resourceAuthenticationRuleUpdate,
		Delete: resourceAuthenticationRuleDelete,

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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"srcintf": &schema.Schema{
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
			"srcaddr": &schema.Schema{
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
			"dstaddr": &schema.Schema{
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
			"srcaddr6": &schema.Schema{
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
			"dstaddr6": &schema.Schema{
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
			"ip_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"active_auth_method": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"sso_auth_method": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"web_auth_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cors_stateful": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"cors_depth": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 8),
				Optional:     true,
				Computed:     true,
			},
			"cert_auth_cookie": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"transaction_based": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"web_portal": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"session_logout": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceAuthenticationRuleCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating AuthenticationRule resource while getting object: %v", err)
	}

	o, err := c.CreateAuthenticationRule(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating AuthenticationRule resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AuthenticationRule")
	}

	return resourceAuthenticationRuleRead(d, m)
}

func resourceAuthenticationRuleUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectAuthenticationRule(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationRule resource while getting object: %v", err)
	}

	o, err := c.UpdateAuthenticationRule(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating AuthenticationRule resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("AuthenticationRule")
	}

	return resourceAuthenticationRuleRead(d, m)
}

func resourceAuthenticationRuleDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteAuthenticationRule(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting AuthenticationRule resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceAuthenticationRuleRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadAuthenticationRule(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationRule resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectAuthenticationRule(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading AuthenticationRule resource from API: %v", err)
	}
	return nil
}

func flattenAuthenticationRuleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleSrcintf(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenAuthenticationRuleSrcintfName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenAuthenticationRuleSrcintfName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleSrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenAuthenticationRuleSrcaddrName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenAuthenticationRuleSrcaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenAuthenticationRuleDstaddrName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenAuthenticationRuleDstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleSrcaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenAuthenticationRuleSrcaddr6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenAuthenticationRuleSrcaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleDstaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenAuthenticationRuleDstaddr6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenAuthenticationRuleDstaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleIpBased(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleActiveAuthMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleSsoAuthMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleWebAuthCookie(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleCorsStateful(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleCorsDepth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenAuthenticationRuleCertAuthCookie(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleTransactionBased(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleWebPortal(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenAuthenticationRuleSessionLogout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectAuthenticationRule(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenAuthenticationRuleName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenAuthenticationRuleStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("protocol", flattenAuthenticationRuleProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcintf", flattenAuthenticationRuleSrcintf(o["srcintf"], d, "srcintf", sv)); err != nil {
			if !fortiAPIPatch(o["srcintf"]) {
				return fmt.Errorf("Error reading srcintf: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcintf"); ok {
			if err = d.Set("srcintf", flattenAuthenticationRuleSrcintf(o["srcintf"], d, "srcintf", sv)); err != nil {
				if !fortiAPIPatch(o["srcintf"]) {
					return fmt.Errorf("Error reading srcintf: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcaddr", flattenAuthenticationRuleSrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr"]) {
				return fmt.Errorf("Error reading srcaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr"); ok {
			if err = d.Set("srcaddr", flattenAuthenticationRuleSrcaddr(o["srcaddr"], d, "srcaddr", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr"]) {
					return fmt.Errorf("Error reading srcaddr: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstaddr", flattenAuthenticationRuleDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr"]) {
				return fmt.Errorf("Error reading dstaddr: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr"); ok {
			if err = d.Set("dstaddr", flattenAuthenticationRuleDstaddr(o["dstaddr"], d, "dstaddr", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr"]) {
					return fmt.Errorf("Error reading dstaddr: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("srcaddr6", flattenAuthenticationRuleSrcaddr6(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["srcaddr6"]) {
				return fmt.Errorf("Error reading srcaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("srcaddr6"); ok {
			if err = d.Set("srcaddr6", flattenAuthenticationRuleSrcaddr6(o["srcaddr6"], d, "srcaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["srcaddr6"]) {
					return fmt.Errorf("Error reading srcaddr6: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("dstaddr6", flattenAuthenticationRuleDstaddr6(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
			if !fortiAPIPatch(o["dstaddr6"]) {
				return fmt.Errorf("Error reading dstaddr6: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("dstaddr6"); ok {
			if err = d.Set("dstaddr6", flattenAuthenticationRuleDstaddr6(o["dstaddr6"], d, "dstaddr6", sv)); err != nil {
				if !fortiAPIPatch(o["dstaddr6"]) {
					return fmt.Errorf("Error reading dstaddr6: %v", err)
				}
			}
		}
	}

	if err = d.Set("ip_based", flattenAuthenticationRuleIpBased(o["ip-based"], d, "ip_based", sv)); err != nil {
		if !fortiAPIPatch(o["ip-based"]) {
			return fmt.Errorf("Error reading ip_based: %v", err)
		}
	}

	if err = d.Set("active_auth_method", flattenAuthenticationRuleActiveAuthMethod(o["active-auth-method"], d, "active_auth_method", sv)); err != nil {
		if !fortiAPIPatch(o["active-auth-method"]) {
			return fmt.Errorf("Error reading active_auth_method: %v", err)
		}
	}

	if err = d.Set("sso_auth_method", flattenAuthenticationRuleSsoAuthMethod(o["sso-auth-method"], d, "sso_auth_method", sv)); err != nil {
		if !fortiAPIPatch(o["sso-auth-method"]) {
			return fmt.Errorf("Error reading sso_auth_method: %v", err)
		}
	}

	if err = d.Set("web_auth_cookie", flattenAuthenticationRuleWebAuthCookie(o["web-auth-cookie"], d, "web_auth_cookie", sv)); err != nil {
		if !fortiAPIPatch(o["web-auth-cookie"]) {
			return fmt.Errorf("Error reading web_auth_cookie: %v", err)
		}
	}

	if err = d.Set("cors_stateful", flattenAuthenticationRuleCorsStateful(o["cors-stateful"], d, "cors_stateful", sv)); err != nil {
		if !fortiAPIPatch(o["cors-stateful"]) {
			return fmt.Errorf("Error reading cors_stateful: %v", err)
		}
	}

	if err = d.Set("cors_depth", flattenAuthenticationRuleCorsDepth(o["cors-depth"], d, "cors_depth", sv)); err != nil {
		if !fortiAPIPatch(o["cors-depth"]) {
			return fmt.Errorf("Error reading cors_depth: %v", err)
		}
	}

	if err = d.Set("cert_auth_cookie", flattenAuthenticationRuleCertAuthCookie(o["cert-auth-cookie"], d, "cert_auth_cookie", sv)); err != nil {
		if !fortiAPIPatch(o["cert-auth-cookie"]) {
			return fmt.Errorf("Error reading cert_auth_cookie: %v", err)
		}
	}

	if err = d.Set("transaction_based", flattenAuthenticationRuleTransactionBased(o["transaction-based"], d, "transaction_based", sv)); err != nil {
		if !fortiAPIPatch(o["transaction-based"]) {
			return fmt.Errorf("Error reading transaction_based: %v", err)
		}
	}

	if err = d.Set("web_portal", flattenAuthenticationRuleWebPortal(o["web-portal"], d, "web_portal", sv)); err != nil {
		if !fortiAPIPatch(o["web-portal"]) {
			return fmt.Errorf("Error reading web_portal: %v", err)
		}
	}

	if err = d.Set("comments", flattenAuthenticationRuleComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("session_logout", flattenAuthenticationRuleSessionLogout(o["session-logout"], d, "session_logout", sv)); err != nil {
		if !fortiAPIPatch(o["session-logout"]) {
			return fmt.Errorf("Error reading session_logout: %v", err)
		}
	}

	return nil
}

func flattenAuthenticationRuleFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandAuthenticationRuleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleSrcintf(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandAuthenticationRuleSrcintfName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandAuthenticationRuleSrcintfName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandAuthenticationRuleSrcaddrName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandAuthenticationRuleSrcaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandAuthenticationRuleDstaddrName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandAuthenticationRuleDstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleSrcaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandAuthenticationRuleSrcaddr6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandAuthenticationRuleSrcaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandAuthenticationRuleDstaddr6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandAuthenticationRuleDstaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleIpBased(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleActiveAuthMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleSsoAuthMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleWebAuthCookie(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleCorsStateful(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleCorsDepth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleCertAuthCookie(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleTransactionBased(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleWebPortal(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandAuthenticationRuleSessionLogout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectAuthenticationRule(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandAuthenticationRuleName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandAuthenticationRuleStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandAuthenticationRuleProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("srcintf"); ok || d.HasChange("srcintf") {
		t, err := expandAuthenticationRuleSrcintf(d, v, "srcintf", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcintf"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr"); ok || d.HasChange("srcaddr") {
		t, err := expandAuthenticationRuleSrcaddr(d, v, "srcaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr"); ok || d.HasChange("dstaddr") {
		t, err := expandAuthenticationRuleDstaddr(d, v, "dstaddr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr"] = t
		}
	}

	if v, ok := d.GetOk("srcaddr6"); ok || d.HasChange("srcaddr6") {
		t, err := expandAuthenticationRuleSrcaddr6(d, v, "srcaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["srcaddr6"] = t
		}
	}

	if v, ok := d.GetOk("dstaddr6"); ok || d.HasChange("dstaddr6") {
		t, err := expandAuthenticationRuleDstaddr6(d, v, "dstaddr6", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dstaddr6"] = t
		}
	}

	if v, ok := d.GetOk("ip_based"); ok {
		t, err := expandAuthenticationRuleIpBased(d, v, "ip_based", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ip-based"] = t
		}
	}

	if v, ok := d.GetOk("active_auth_method"); ok {
		t, err := expandAuthenticationRuleActiveAuthMethod(d, v, "active_auth_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["active-auth-method"] = t
		}
	} else if d.HasChange("active_auth_method") {
		obj["active-auth-method"] = nil
	}

	if v, ok := d.GetOk("sso_auth_method"); ok {
		t, err := expandAuthenticationRuleSsoAuthMethod(d, v, "sso_auth_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sso-auth-method"] = t
		}
	} else if d.HasChange("sso_auth_method") {
		obj["sso-auth-method"] = nil
	}

	if v, ok := d.GetOk("web_auth_cookie"); ok {
		t, err := expandAuthenticationRuleWebAuthCookie(d, v, "web_auth_cookie", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-auth-cookie"] = t
		}
	}

	if v, ok := d.GetOk("cors_stateful"); ok {
		t, err := expandAuthenticationRuleCorsStateful(d, v, "cors_stateful", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors-stateful"] = t
		}
	}

	if v, ok := d.GetOk("cors_depth"); ok {
		t, err := expandAuthenticationRuleCorsDepth(d, v, "cors_depth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cors-depth"] = t
		}
	}

	if v, ok := d.GetOk("cert_auth_cookie"); ok {
		t, err := expandAuthenticationRuleCertAuthCookie(d, v, "cert_auth_cookie", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["cert-auth-cookie"] = t
		}
	}

	if v, ok := d.GetOk("transaction_based"); ok {
		t, err := expandAuthenticationRuleTransactionBased(d, v, "transaction_based", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["transaction-based"] = t
		}
	}

	if v, ok := d.GetOk("web_portal"); ok {
		t, err := expandAuthenticationRuleWebPortal(d, v, "web_portal", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["web-portal"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {
		t, err := expandAuthenticationRuleComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	} else if d.HasChange("comments") {
		obj["comments"] = nil
	}

	if v, ok := d.GetOk("session_logout"); ok {
		t, err := expandAuthenticationRuleSessionLogout(d, v, "session_logout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["session-logout"] = t
		}
	}

	return &obj, nil
}
