// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure VLAN policy to be applied on the managed FortiSwitch ports through port-policy.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSwitchControllerVlanPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerVlanPolicyCreate,
		Read:   resourceSwitchControllerVlanPolicyRead,
		Update: resourceSwitchControllerVlanPolicyUpdate,
		Delete: resourceSwitchControllerVlanPolicyDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"fortilink": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"vlan": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
			},
			"allowed_vlans": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"untagged_vlans": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"vlan_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
						},
					},
				},
			},
			"allowed_vlans_all": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"discard_mode": &schema.Schema{
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

func resourceSwitchControllerVlanPolicyCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerVlanPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerVlanPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerVlanPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerVlanPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerVlanPolicy")
	}

	return resourceSwitchControllerVlanPolicyRead(d, m)
}

func resourceSwitchControllerVlanPolicyUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSwitchControllerVlanPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerVlanPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerVlanPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerVlanPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerVlanPolicy")
	}

	return resourceSwitchControllerVlanPolicyRead(d, m)
}

func resourceSwitchControllerVlanPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerVlanPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerVlanPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerVlanPolicyRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSwitchControllerVlanPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerVlanPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerVlanPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerVlanPolicy resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerVlanPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyFortilink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyVlan(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyAllowedVlans(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "vlan-name", "vlan_name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["vlan-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
			}
			tmp["vlan_name"] = flattenSwitchControllerVlanPolicyAllowedVlansVlanName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vlan_name", d)
	return result
}

func flattenSwitchControllerVlanPolicyAllowedVlansVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyUntaggedVlans(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

	parsed_list := mergeBlock(tf_list, l, "vlan-name", "vlan_name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["vlan-name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vlan_name"
			}
			tmp["vlan_name"] = flattenSwitchControllerVlanPolicyUntaggedVlansVlanName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "vlan_name", d)
	return result
}

func flattenSwitchControllerVlanPolicyUntaggedVlansVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyAllowedVlansAll(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPolicyDiscardMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerVlanPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSwitchControllerVlanPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSwitchControllerVlanPolicyDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("fortilink", flattenSwitchControllerVlanPolicyFortilink(o["fortilink"], d, "fortilink", sv)); err != nil {
		if !fortiAPIPatch(o["fortilink"]) {
			return fmt.Errorf("Error reading fortilink: %v", err)
		}
	}

	if err = d.Set("vlan", flattenSwitchControllerVlanPolicyVlan(o["vlan"], d, "vlan", sv)); err != nil {
		if !fortiAPIPatch(o["vlan"]) {
			return fmt.Errorf("Error reading vlan: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("allowed_vlans", flattenSwitchControllerVlanPolicyAllowedVlans(o["allowed-vlans"], d, "allowed_vlans", sv)); err != nil {
			if !fortiAPIPatch(o["allowed-vlans"]) {
				return fmt.Errorf("Error reading allowed_vlans: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("allowed_vlans"); ok {
			if err = d.Set("allowed_vlans", flattenSwitchControllerVlanPolicyAllowedVlans(o["allowed-vlans"], d, "allowed_vlans", sv)); err != nil {
				if !fortiAPIPatch(o["allowed-vlans"]) {
					return fmt.Errorf("Error reading allowed_vlans: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("untagged_vlans", flattenSwitchControllerVlanPolicyUntaggedVlans(o["untagged-vlans"], d, "untagged_vlans", sv)); err != nil {
			if !fortiAPIPatch(o["untagged-vlans"]) {
				return fmt.Errorf("Error reading untagged_vlans: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("untagged_vlans"); ok {
			if err = d.Set("untagged_vlans", flattenSwitchControllerVlanPolicyUntaggedVlans(o["untagged-vlans"], d, "untagged_vlans", sv)); err != nil {
				if !fortiAPIPatch(o["untagged-vlans"]) {
					return fmt.Errorf("Error reading untagged_vlans: %v", err)
				}
			}
		}
	}

	if err = d.Set("allowed_vlans_all", flattenSwitchControllerVlanPolicyAllowedVlansAll(o["allowed-vlans-all"], d, "allowed_vlans_all", sv)); err != nil {
		if !fortiAPIPatch(o["allowed-vlans-all"]) {
			return fmt.Errorf("Error reading allowed_vlans_all: %v", err)
		}
	}

	if err = d.Set("discard_mode", flattenSwitchControllerVlanPolicyDiscardMode(o["discard-mode"], d, "discard_mode", sv)); err != nil {
		if !fortiAPIPatch(o["discard-mode"]) {
			return fmt.Errorf("Error reading discard_mode: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerVlanPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerVlanPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyFortilink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyVlan(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyAllowedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vlan-name"], _ = expandSwitchControllerVlanPolicyAllowedVlansVlanName(d, i["vlan_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerVlanPolicyAllowedVlansVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyUntaggedVlans(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["vlan-name"], _ = expandSwitchControllerVlanPolicyUntaggedVlansVlanName(d, i["vlan_name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerVlanPolicyUntaggedVlansVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyAllowedVlansAll(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPolicyDiscardMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerVlanPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerVlanPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSwitchControllerVlanPolicyDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("fortilink"); ok {
		t, err := expandSwitchControllerVlanPolicyFortilink(d, v, "fortilink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fortilink"] = t
		}
	} else if d.HasChange("fortilink") {
		obj["fortilink"] = nil
	}

	if v, ok := d.GetOk("vlan"); ok {
		t, err := expandSwitchControllerVlanPolicyVlan(d, v, "vlan", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlan"] = t
		}
	} else if d.HasChange("vlan") {
		obj["vlan"] = nil
	}

	if v, ok := d.GetOk("allowed_vlans"); ok || d.HasChange("allowed_vlans") {
		t, err := expandSwitchControllerVlanPolicyAllowedVlans(d, v, "allowed_vlans", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-vlans"] = t
		}
	}

	if v, ok := d.GetOk("untagged_vlans"); ok || d.HasChange("untagged_vlans") {
		t, err := expandSwitchControllerVlanPolicyUntaggedVlans(d, v, "untagged_vlans", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["untagged-vlans"] = t
		}
	}

	if v, ok := d.GetOk("allowed_vlans_all"); ok {
		t, err := expandSwitchControllerVlanPolicyAllowedVlansAll(d, v, "allowed_vlans_all", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowed-vlans-all"] = t
		}
	}

	if v, ok := d.GetOk("discard_mode"); ok {
		t, err := expandSwitchControllerVlanPolicyDiscardMode(d, v, "discard_mode", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["discard-mode"] = t
		}
	}

	return &obj, nil
}
