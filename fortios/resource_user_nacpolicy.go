// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure NAC policy matching pattern to identify matching NAC devices.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceUserNacPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceUserNacPolicyCreate,
		Read:   resourceUserNacPolicyRead,
		Update: resourceUserNacPolicyUpdate,
		Delete: resourceUserNacPolicyDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"category": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"mac": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hw_vendor": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"type": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"family": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"os": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"hw_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"sw_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"host": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"user": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
				Computed:     true,
			},
			"src": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"user_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ems_tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
				Computed:     true,
			},
			"switch_fortilink": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"switch_scope": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"switch_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"switch_auto_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"switch_port_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"switch_mac_policy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceUserNacPolicyCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserNacPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating UserNacPolicy resource while getting object: %v", err)
	}

	o, err := c.CreateUserNacPolicy(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating UserNacPolicy resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserNacPolicy")
	}

	return resourceUserNacPolicyRead(d, m)
}

func resourceUserNacPolicyUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectUserNacPolicy(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating UserNacPolicy resource while getting object: %v", err)
	}

	o, err := c.UpdateUserNacPolicy(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating UserNacPolicy resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("UserNacPolicy")
	}

	return resourceUserNacPolicyRead(d, m)
}

func resourceUserNacPolicyDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteUserNacPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting UserNacPolicy resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceUserNacPolicyRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadUserNacPolicy(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading UserNacPolicy resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectUserNacPolicy(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading UserNacPolicy resource from API: %v", err)
	}
	return nil
}

func flattenUserNacPolicyName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyCategory(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyMac(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyHwVendor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyFamily(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyOs(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyHwVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicySwVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicySrc(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyUserGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicyEmsTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicySwitchFortilink(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicySwitchScope(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := i["switch-id"]; ok {

			tmp["switch_id"] = flattenUserNacPolicySwitchScopeSwitchId(i["switch-id"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "switch_id", d)
	return result
}

func flattenUserNacPolicySwitchScopeSwitchId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicySwitchAutoAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicySwitchPortPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenUserNacPolicySwitchMacPolicy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectUserNacPolicy(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenUserNacPolicyName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenUserNacPolicyDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("category", flattenUserNacPolicyCategory(o["category"], d, "category", sv)); err != nil {
		if !fortiAPIPatch(o["category"]) {
			return fmt.Errorf("Error reading category: %v", err)
		}
	}

	if err = d.Set("status", flattenUserNacPolicyStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("mac", flattenUserNacPolicyMac(o["mac"], d, "mac", sv)); err != nil {
		if !fortiAPIPatch(o["mac"]) {
			return fmt.Errorf("Error reading mac: %v", err)
		}
	}

	if err = d.Set("hw_vendor", flattenUserNacPolicyHwVendor(o["hw-vendor"], d, "hw_vendor", sv)); err != nil {
		if !fortiAPIPatch(o["hw-vendor"]) {
			return fmt.Errorf("Error reading hw_vendor: %v", err)
		}
	}

	if err = d.Set("type", flattenUserNacPolicyType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("family", flattenUserNacPolicyFamily(o["family"], d, "family", sv)); err != nil {
		if !fortiAPIPatch(o["family"]) {
			return fmt.Errorf("Error reading family: %v", err)
		}
	}

	if err = d.Set("os", flattenUserNacPolicyOs(o["os"], d, "os", sv)); err != nil {
		if !fortiAPIPatch(o["os"]) {
			return fmt.Errorf("Error reading os: %v", err)
		}
	}

	if err = d.Set("hw_version", flattenUserNacPolicyHwVersion(o["hw-version"], d, "hw_version", sv)); err != nil {
		if !fortiAPIPatch(o["hw-version"]) {
			return fmt.Errorf("Error reading hw_version: %v", err)
		}
	}

	if err = d.Set("sw_version", flattenUserNacPolicySwVersion(o["sw-version"], d, "sw_version", sv)); err != nil {
		if !fortiAPIPatch(o["sw-version"]) {
			return fmt.Errorf("Error reading sw_version: %v", err)
		}
	}

	if err = d.Set("host", flattenUserNacPolicyHost(o["host"], d, "host", sv)); err != nil {
		if !fortiAPIPatch(o["host"]) {
			return fmt.Errorf("Error reading host: %v", err)
		}
	}

	if err = d.Set("user", flattenUserNacPolicyUser(o["user"], d, "user", sv)); err != nil {
		if !fortiAPIPatch(o["user"]) {
			return fmt.Errorf("Error reading user: %v", err)
		}
	}

	if err = d.Set("src", flattenUserNacPolicySrc(o["src"], d, "src", sv)); err != nil {
		if !fortiAPIPatch(o["src"]) {
			return fmt.Errorf("Error reading src: %v", err)
		}
	}

	if err = d.Set("user_group", flattenUserNacPolicyUserGroup(o["user-group"], d, "user_group", sv)); err != nil {
		if !fortiAPIPatch(o["user-group"]) {
			return fmt.Errorf("Error reading user_group: %v", err)
		}
	}

	if err = d.Set("ems_tag", flattenUserNacPolicyEmsTag(o["ems-tag"], d, "ems_tag", sv)); err != nil {
		if !fortiAPIPatch(o["ems-tag"]) {
			return fmt.Errorf("Error reading ems_tag: %v", err)
		}
	}

	if err = d.Set("switch_fortilink", flattenUserNacPolicySwitchFortilink(o["switch-fortilink"], d, "switch_fortilink", sv)); err != nil {
		if !fortiAPIPatch(o["switch-fortilink"]) {
			return fmt.Errorf("Error reading switch_fortilink: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("switch_scope", flattenUserNacPolicySwitchScope(o["switch-scope"], d, "switch_scope", sv)); err != nil {
			if !fortiAPIPatch(o["switch-scope"]) {
				return fmt.Errorf("Error reading switch_scope: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("switch_scope"); ok {
			if err = d.Set("switch_scope", flattenUserNacPolicySwitchScope(o["switch-scope"], d, "switch_scope", sv)); err != nil {
				if !fortiAPIPatch(o["switch-scope"]) {
					return fmt.Errorf("Error reading switch_scope: %v", err)
				}
			}
		}
	}

	if err = d.Set("switch_auto_auth", flattenUserNacPolicySwitchAutoAuth(o["switch-auto-auth"], d, "switch_auto_auth", sv)); err != nil {
		if !fortiAPIPatch(o["switch-auto-auth"]) {
			return fmt.Errorf("Error reading switch_auto_auth: %v", err)
		}
	}

	if err = d.Set("switch_port_policy", flattenUserNacPolicySwitchPortPolicy(o["switch-port-policy"], d, "switch_port_policy", sv)); err != nil {
		if !fortiAPIPatch(o["switch-port-policy"]) {
			return fmt.Errorf("Error reading switch_port_policy: %v", err)
		}
	}

	if err = d.Set("switch_mac_policy", flattenUserNacPolicySwitchMacPolicy(o["switch-mac-policy"], d, "switch_mac_policy", sv)); err != nil {
		if !fortiAPIPatch(o["switch-mac-policy"]) {
			return fmt.Errorf("Error reading switch_mac_policy: %v", err)
		}
	}

	return nil
}

func flattenUserNacPolicyFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandUserNacPolicyName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyCategory(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyMac(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyHwVendor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyFamily(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyOs(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyHwVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySwVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySrc(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyUserGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicyEmsTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySwitchFortilink(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySwitchScope(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "switch_id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["switch-id"], _ = expandUserNacPolicySwitchScopeSwitchId(d, i["switch_id"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandUserNacPolicySwitchScopeSwitchId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySwitchAutoAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySwitchPortPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandUserNacPolicySwitchMacPolicy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectUserNacPolicy(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandUserNacPolicyName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {

		t, err := expandUserNacPolicyDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOk("category"); ok {

		t, err := expandUserNacPolicyCategory(d, v, "category", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["category"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandUserNacPolicyStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("mac"); ok {

		t, err := expandUserNacPolicyMac(d, v, "mac", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac"] = t
		}
	}

	if v, ok := d.GetOk("hw_vendor"); ok {

		t, err := expandUserNacPolicyHwVendor(d, v, "hw_vendor", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-vendor"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {

		t, err := expandUserNacPolicyType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("family"); ok {

		t, err := expandUserNacPolicyFamily(d, v, "family", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["family"] = t
		}
	}

	if v, ok := d.GetOk("os"); ok {

		t, err := expandUserNacPolicyOs(d, v, "os", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["os"] = t
		}
	}

	if v, ok := d.GetOk("hw_version"); ok {

		t, err := expandUserNacPolicyHwVersion(d, v, "hw_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hw-version"] = t
		}
	}

	if v, ok := d.GetOk("sw_version"); ok {

		t, err := expandUserNacPolicySwVersion(d, v, "sw_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sw-version"] = t
		}
	}

	if v, ok := d.GetOk("host"); ok {

		t, err := expandUserNacPolicyHost(d, v, "host", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["host"] = t
		}
	}

	if v, ok := d.GetOk("user"); ok {

		t, err := expandUserNacPolicyUser(d, v, "user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user"] = t
		}
	}

	if v, ok := d.GetOk("src"); ok {

		t, err := expandUserNacPolicySrc(d, v, "src", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["src"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok {

		t, err := expandUserNacPolicyUserGroup(d, v, "user_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	if v, ok := d.GetOk("ems_tag"); ok {

		t, err := expandUserNacPolicyEmsTag(d, v, "ems_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ems-tag"] = t
		}
	}

	if v, ok := d.GetOk("switch_fortilink"); ok {

		t, err := expandUserNacPolicySwitchFortilink(d, v, "switch_fortilink", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-fortilink"] = t
		}
	}

	if v, ok := d.GetOk("switch_scope"); ok {

		t, err := expandUserNacPolicySwitchScope(d, v, "switch_scope", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-scope"] = t
		}
	}

	if v, ok := d.GetOk("switch_auto_auth"); ok {

		t, err := expandUserNacPolicySwitchAutoAuth(d, v, "switch_auto_auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-auto-auth"] = t
		}
	}

	if v, ok := d.GetOk("switch_port_policy"); ok {

		t, err := expandUserNacPolicySwitchPortPolicy(d, v, "switch_port_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-port-policy"] = t
		}
	}

	if v, ok := d.GetOk("switch_mac_policy"); ok {

		t, err := expandUserNacPolicySwitchMacPolicy(d, v, "switch_mac_policy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["switch-mac-policy"] = t
		}
	}

	return &obj, nil
}
