// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure WiFi bridge access control list.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceWirelessControllerAccessControlList() *schema.Resource {
	return &schema.Resource{
		Create: resourceWirelessControllerAccessControlListCreate,
		Read:   resourceWirelessControllerAccessControlListRead,
		Update: resourceWirelessControllerAccessControlListUpdate,
		Delete: resourceWirelessControllerAccessControlListDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"layer3_ipv4_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"srcport": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstport": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},
			"layer3_ipv6_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
						},
						"comment": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"srcaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"srcport": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
						},
						"dstport": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
						},
						"protocol": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
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

func resourceWirelessControllerAccessControlListCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerAccessControlList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerAccessControlList resource while getting object: %v", err)
	}

	o, err := c.CreateWirelessControllerAccessControlList(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating WirelessControllerAccessControlList resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerAccessControlList")
	}

	return resourceWirelessControllerAccessControlListRead(d, m)
}

func resourceWirelessControllerAccessControlListUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectWirelessControllerAccessControlList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerAccessControlList resource while getting object: %v", err)
	}

	o, err := c.UpdateWirelessControllerAccessControlList(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating WirelessControllerAccessControlList resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WirelessControllerAccessControlList")
	}

	return resourceWirelessControllerAccessControlListRead(d, m)
}

func resourceWirelessControllerAccessControlListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWirelessControllerAccessControlList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting WirelessControllerAccessControlList resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWirelessControllerAccessControlListRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadWirelessControllerAccessControlList(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerAccessControlList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWirelessControllerAccessControlList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WirelessControllerAccessControlList resource from API: %v", err)
	}
	return nil
}

func flattenWirelessControllerAccessControlListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4Rules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "rule-id", "rule_id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["rule-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
			}
			tmp["rule_id"] = flattenWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["comment"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
			}
			tmp["comment"] = flattenWirelessControllerAccessControlListLayer3Ipv4RulesComment(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["srcaddr"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
			}
			tmp["srcaddr"] = flattenWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["srcport"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
			}
			tmp["srcport"] = flattenWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dstaddr"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
			}
			tmp["dstaddr"] = flattenWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dstport"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
			}
			tmp["dstport"] = flattenWirelessControllerAccessControlListLayer3Ipv4RulesDstport(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["action"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
			}
			tmp["action"] = flattenWirelessControllerAccessControlListLayer3Ipv4RulesAction(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "rule_id", d)
	return result
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesDstport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerAccessControlListLayer3Ipv4RulesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6Rules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "rule-id", "rule_id")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["rule-id"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
			}
			tmp["rule_id"] = flattenWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["comment"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
			}
			tmp["comment"] = flattenWirelessControllerAccessControlListLayer3Ipv6RulesComment(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["srcaddr"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
			}
			tmp["srcaddr"] = flattenWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["srcport"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
			}
			tmp["srcport"] = flattenWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dstaddr"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
			}
			tmp["dstaddr"] = flattenWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["dstport"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
			}
			tmp["dstport"] = flattenWirelessControllerAccessControlListLayer3Ipv6RulesDstport(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["protocol"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
			}
			tmp["protocol"] = flattenWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["action"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
			}
			tmp["action"] = flattenWirelessControllerAccessControlListLayer3Ipv6RulesAction(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "rule_id", d)
	return result
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesComment(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesDstport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenWirelessControllerAccessControlListLayer3Ipv6RulesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWirelessControllerAccessControlList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWirelessControllerAccessControlListName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("comment", flattenWirelessControllerAccessControlListComment(o["comment"], d, "comment", sv)); err != nil {
		if !fortiAPIPatch(o["comment"]) {
			return fmt.Errorf("Error reading comment: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("layer3_ipv4_rules", flattenWirelessControllerAccessControlListLayer3Ipv4Rules(o["layer3-ipv4-rules"], d, "layer3_ipv4_rules", sv)); err != nil {
			if !fortiAPIPatch(o["layer3-ipv4-rules"]) {
				return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("layer3_ipv4_rules"); ok {
			if err = d.Set("layer3_ipv4_rules", flattenWirelessControllerAccessControlListLayer3Ipv4Rules(o["layer3-ipv4-rules"], d, "layer3_ipv4_rules", sv)); err != nil {
				if !fortiAPIPatch(o["layer3-ipv4-rules"]) {
					return fmt.Errorf("Error reading layer3_ipv4_rules: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("layer3_ipv6_rules", flattenWirelessControllerAccessControlListLayer3Ipv6Rules(o["layer3-ipv6-rules"], d, "layer3_ipv6_rules", sv)); err != nil {
			if !fortiAPIPatch(o["layer3-ipv6-rules"]) {
				return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("layer3_ipv6_rules"); ok {
			if err = d.Set("layer3_ipv6_rules", flattenWirelessControllerAccessControlListLayer3Ipv6Rules(o["layer3-ipv6-rules"], d, "layer3_ipv6_rules", sv)); err != nil {
				if !fortiAPIPatch(o["layer3-ipv6-rules"]) {
					return fmt.Errorf("Error reading layer3_ipv6_rules: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWirelessControllerAccessControlListFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWirelessControllerAccessControlListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4Rules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rule-id"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(d, i["rule_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rule-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesComment(d, i["comment"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comment"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["srcaddr"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(d, i["srcaddr"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["srcaddr"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["srcport"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(d, i["srcport"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["srcport"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dstaddr"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(d, i["dstaddr"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dstaddr"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dstport"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesDstport(d, i["dstport"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dstport"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWirelessControllerAccessControlListLayer3Ipv4RulesAction(d, i["action"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["action"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesSrcport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesDstport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv4RulesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6Rules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rule-id"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(d, i["rule_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rule-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "comment"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["comment"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesComment(d, i["comment"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["comment"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcaddr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["srcaddr"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(d, i["srcaddr"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["srcaddr"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "srcport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["srcport"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(d, i["srcport"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["srcport"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dstaddr"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(d, i["dstaddr"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dstaddr"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstport"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["dstport"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesDstport(d, i["dstport"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["dstport"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(d, i["protocol"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWirelessControllerAccessControlListLayer3Ipv6RulesAction(d, i["action"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["action"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesRuleId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesComment(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesSrcaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesSrcport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesDstport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWirelessControllerAccessControlListLayer3Ipv6RulesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWirelessControllerAccessControlList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWirelessControllerAccessControlListName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("comment"); ok {
		t, err := expandWirelessControllerAccessControlListComment(d, v, "comment", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comment"] = t
		}
	} else if d.HasChange("comment") {
		obj["comment"] = nil
	}

	if v, ok := d.GetOk("layer3_ipv4_rules"); ok || d.HasChange("layer3_ipv4_rules") {
		t, err := expandWirelessControllerAccessControlListLayer3Ipv4Rules(d, v, "layer3_ipv4_rules", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["layer3-ipv4-rules"] = t
		}
	}

	if v, ok := d.GetOk("layer3_ipv6_rules"); ok || d.HasChange("layer3_ipv6_rules") {
		t, err := expandWirelessControllerAccessControlListLayer3Ipv6Rules(d, v, "layer3_ipv6_rules", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["layer3-ipv6-rules"] = t
		}
	}

	return &obj, nil
}
