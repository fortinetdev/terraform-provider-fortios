// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure 802.1x MAC Authentication Bypass (MAB) policies.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerSecurityPolicy8021X() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerSecurityPolicy8021XCreate,
		Read:   resourceSwitchControllerSecurityPolicy8021XRead,
		Update: resourceSwitchControllerSecurityPolicy8021XUpdate,
		Delete: resourceSwitchControllerSecurityPolicy8021XDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"security_mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"user_group": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"mac_auth_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"open_auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"eap_passthru": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"guest_vlanid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"guest_vlan_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"guest_auth_delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 900),
				Optional:     true,
				Computed:     true,
			},
			"auth_fail_vlan": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth_fail_vlanid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"auth_fail_vlan_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"framevid_apply": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_timeout_overwrite": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"policy_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerSecurityPolicy8021XCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerSecurityPolicy8021X(d)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicy8021X resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerSecurityPolicy8021X(obj)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSecurityPolicy8021X")
	}

	return resourceSwitchControllerSecurityPolicy8021XRead(d, m)
}

func resourceSwitchControllerSecurityPolicy8021XUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSwitchControllerSecurityPolicy8021X(d)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicy8021X resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerSecurityPolicy8021X(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerSecurityPolicy8021X")
	}

	return resourceSwitchControllerSecurityPolicy8021XRead(d, m)
}

func resourceSwitchControllerSecurityPolicy8021XDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSwitchControllerSecurityPolicy8021X(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerSecurityPolicy8021XRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSwitchControllerSecurityPolicy8021X(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicy8021X resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerSecurityPolicy8021X(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerSecurityPolicy8021X resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerSecurityPolicy8021XName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XSecurityMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XUserGroup(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSwitchControllerSecurityPolicy8021XUserGroupName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerSecurityPolicy8021XUserGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XMacAuthBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XOpenAuth(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XEapPassthru(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XGuestVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XGuestVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XGuestVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XGuestAuthDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XAuthFailVlan(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XAuthFailVlanid(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XAuthFailVlanId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XFramevidApply(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSwitchControllerSecurityPolicy8021XPolicyType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSwitchControllerSecurityPolicy8021X(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerSecurityPolicy8021XName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("security_mode", flattenSwitchControllerSecurityPolicy8021XSecurityMode(o["security-mode"], d, "security_mode")); err != nil {
		if !fortiAPIPatch(o["security-mode"]) {
			return fmt.Errorf("Error reading security_mode: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("user_group", flattenSwitchControllerSecurityPolicy8021XUserGroup(o["user-group"], d, "user_group")); err != nil {
			if !fortiAPIPatch(o["user-group"]) {
				return fmt.Errorf("Error reading user_group: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("user_group"); ok {
			if err = d.Set("user_group", flattenSwitchControllerSecurityPolicy8021XUserGroup(o["user-group"], d, "user_group")); err != nil {
				if !fortiAPIPatch(o["user-group"]) {
					return fmt.Errorf("Error reading user_group: %v", err)
				}
			}
		}
	}

	if err = d.Set("mac_auth_bypass", flattenSwitchControllerSecurityPolicy8021XMacAuthBypass(o["mac-auth-bypass"], d, "mac_auth_bypass")); err != nil {
		if !fortiAPIPatch(o["mac-auth-bypass"]) {
			return fmt.Errorf("Error reading mac_auth_bypass: %v", err)
		}
	}

	if err = d.Set("open_auth", flattenSwitchControllerSecurityPolicy8021XOpenAuth(o["open-auth"], d, "open_auth")); err != nil {
		if !fortiAPIPatch(o["open-auth"]) {
			return fmt.Errorf("Error reading open_auth: %v", err)
		}
	}

	if err = d.Set("eap_passthru", flattenSwitchControllerSecurityPolicy8021XEapPassthru(o["eap-passthru"], d, "eap_passthru")); err != nil {
		if !fortiAPIPatch(o["eap-passthru"]) {
			return fmt.Errorf("Error reading eap_passthru: %v", err)
		}
	}

	if err = d.Set("guest_vlan", flattenSwitchControllerSecurityPolicy8021XGuestVlan(o["guest-vlan"], d, "guest_vlan")); err != nil {
		if !fortiAPIPatch(o["guest-vlan"]) {
			return fmt.Errorf("Error reading guest_vlan: %v", err)
		}
	}

	if err = d.Set("guest_vlanid", flattenSwitchControllerSecurityPolicy8021XGuestVlanid(o["guest-vlanid"], d, "guest_vlanid")); err != nil {
		if !fortiAPIPatch(o["guest-vlanid"]) {
			return fmt.Errorf("Error reading guest_vlanid: %v", err)
		}
	}

	if err = d.Set("guest_vlan_id", flattenSwitchControllerSecurityPolicy8021XGuestVlanId(o["guest-vlan-id"], d, "guest_vlan_id")); err != nil {
		if !fortiAPIPatch(o["guest-vlan-id"]) {
			return fmt.Errorf("Error reading guest_vlan_id: %v", err)
		}
	}

	if err = d.Set("guest_auth_delay", flattenSwitchControllerSecurityPolicy8021XGuestAuthDelay(o["guest-auth-delay"], d, "guest_auth_delay")); err != nil {
		if !fortiAPIPatch(o["guest-auth-delay"]) {
			return fmt.Errorf("Error reading guest_auth_delay: %v", err)
		}
	}

	if err = d.Set("auth_fail_vlan", flattenSwitchControllerSecurityPolicy8021XAuthFailVlan(o["auth-fail-vlan"], d, "auth_fail_vlan")); err != nil {
		if !fortiAPIPatch(o["auth-fail-vlan"]) {
			return fmt.Errorf("Error reading auth_fail_vlan: %v", err)
		}
	}

	if err = d.Set("auth_fail_vlanid", flattenSwitchControllerSecurityPolicy8021XAuthFailVlanid(o["auth-fail-vlanid"], d, "auth_fail_vlanid")); err != nil {
		if !fortiAPIPatch(o["auth-fail-vlanid"]) {
			return fmt.Errorf("Error reading auth_fail_vlanid: %v", err)
		}
	}

	if err = d.Set("auth_fail_vlan_id", flattenSwitchControllerSecurityPolicy8021XAuthFailVlanId(o["auth-fail-vlan-id"], d, "auth_fail_vlan_id")); err != nil {
		if !fortiAPIPatch(o["auth-fail-vlan-id"]) {
			return fmt.Errorf("Error reading auth_fail_vlan_id: %v", err)
		}
	}

	if err = d.Set("framevid_apply", flattenSwitchControllerSecurityPolicy8021XFramevidApply(o["framevid-apply"], d, "framevid_apply")); err != nil {
		if !fortiAPIPatch(o["framevid-apply"]) {
			return fmt.Errorf("Error reading framevid_apply: %v", err)
		}
	}

	if err = d.Set("radius_timeout_overwrite", flattenSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(o["radius-timeout-overwrite"], d, "radius_timeout_overwrite")); err != nil {
		if !fortiAPIPatch(o["radius-timeout-overwrite"]) {
			return fmt.Errorf("Error reading radius_timeout_overwrite: %v", err)
		}
	}

	if err = d.Set("policy_type", flattenSwitchControllerSecurityPolicy8021XPolicyType(o["policy-type"], d, "policy_type")); err != nil {
		if !fortiAPIPatch(o["policy-type"]) {
			return fmt.Errorf("Error reading policy_type: %v", err)
		}
	}

	return nil
}

func flattenSwitchControllerSecurityPolicy8021XFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSwitchControllerSecurityPolicy8021XName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XSecurityMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XUserGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSwitchControllerSecurityPolicy8021XUserGroupName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerSecurityPolicy8021XUserGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XMacAuthBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XOpenAuth(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XEapPassthru(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XGuestVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XGuestVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XGuestVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XGuestAuthDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XAuthFailVlan(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XAuthFailVlanid(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XAuthFailVlanId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XFramevidApply(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerSecurityPolicy8021XPolicyType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerSecurityPolicy8021X(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("security_mode"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XSecurityMode(d, v, "security_mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-mode"] = t
		}
	}

	if v, ok := d.GetOk("user_group"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XUserGroup(d, v, "user_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-group"] = t
		}
	}

	if v, ok := d.GetOk("mac_auth_bypass"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XMacAuthBypass(d, v, "mac_auth_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mac-auth-bypass"] = t
		}
	}

	if v, ok := d.GetOk("open_auth"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XOpenAuth(d, v, "open_auth")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["open-auth"] = t
		}
	}

	if v, ok := d.GetOk("eap_passthru"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XEapPassthru(d, v, "eap_passthru")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["eap-passthru"] = t
		}
	}

	if v, ok := d.GetOk("guest_vlan"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XGuestVlan(d, v, "guest_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-vlan"] = t
		}
	}

	if v, ok := d.GetOkExists("guest_vlanid"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XGuestVlanid(d, v, "guest_vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-vlanid"] = t
		}
	}

	if v, ok := d.GetOk("guest_vlan_id"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XGuestVlanId(d, v, "guest_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("guest_auth_delay"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XGuestAuthDelay(d, v, "guest_auth_delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["guest-auth-delay"] = t
		}
	}

	if v, ok := d.GetOk("auth_fail_vlan"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthFailVlan(d, v, "auth_fail_vlan")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-fail-vlan"] = t
		}
	}

	if v, ok := d.GetOkExists("auth_fail_vlanid"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthFailVlanid(d, v, "auth_fail_vlanid")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-fail-vlanid"] = t
		}
	}

	if v, ok := d.GetOk("auth_fail_vlan_id"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XAuthFailVlanId(d, v, "auth_fail_vlan_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth-fail-vlan-id"] = t
		}
	}

	if v, ok := d.GetOk("framevid_apply"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XFramevidApply(d, v, "framevid_apply")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["framevid-apply"] = t
		}
	}

	if v, ok := d.GetOk("radius_timeout_overwrite"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XRadiusTimeoutOverwrite(d, v, "radius_timeout_overwrite")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-timeout-overwrite"] = t
		}
	}

	if v, ok := d.GetOk("policy_type"); ok {
		t, err := expandSwitchControllerSecurityPolicy8021XPolicyType(d, v, "policy_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["policy-type"] = t
		}
	}

	return &obj, nil
}
