// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure VLANs for switch controller.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSwitchControllerVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceSwitchControllerVlanCreate,
		Read:   resourceSwitchControllerVlanRead,
		Update: resourceSwitchControllerVlanUpdate,
		Delete: resourceSwitchControllerVlanDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 15),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"vlanid": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4094),
				Optional:     true,
				Computed:     true,
			},
			"comments": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"color": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 32),
				Optional:     true,
				Computed:     true,
			},
			"security": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"auth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"radius_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"usergroup": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"portal_message_override_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"portal_message_overrides": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"auth_disclaimer_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"auth_reject_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"auth_login_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"auth_login_failed_page": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"selected_usergroups": &schema.Schema{
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
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSwitchControllerVlanCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerVlan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerVlan resource while getting object: %v", err)
	}

	o, err := c.CreateSwitchControllerVlan(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SwitchControllerVlan resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerVlan")
	}

	return resourceSwitchControllerVlanRead(d, m)
}

func resourceSwitchControllerVlanUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSwitchControllerVlan(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerVlan resource while getting object: %v", err)
	}

	o, err := c.UpdateSwitchControllerVlan(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SwitchControllerVlan resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SwitchControllerVlan")
	}

	return resourceSwitchControllerVlanRead(d, m)
}

func resourceSwitchControllerVlanDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSwitchControllerVlan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SwitchControllerVlan resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSwitchControllerVlanRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSwitchControllerVlan(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerVlan resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSwitchControllerVlan(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SwitchControllerVlan resource from API: %v", err)
	}
	return nil
}

func flattenSwitchControllerVlanName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanVlanid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanComments(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanSecurity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanAuth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanRadiusServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanUsergroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPortalMessageOverrideGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPortalMessageOverrides(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := i["auth-disclaimer-page"]; ok {

		result["auth_disclaimer_page"] = flattenSwitchControllerVlanPortalMessageOverridesAuthDisclaimerPage(i["auth-disclaimer-page"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := i["auth-reject-page"]; ok {

		result["auth_reject_page"] = flattenSwitchControllerVlanPortalMessageOverridesAuthRejectPage(i["auth-reject-page"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := i["auth-login-page"]; ok {

		result["auth_login_page"] = flattenSwitchControllerVlanPortalMessageOverridesAuthLoginPage(i["auth-login-page"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := i["auth-login-failed-page"]; ok {

		result["auth_login_failed_page"] = flattenSwitchControllerVlanPortalMessageOverridesAuthLoginFailedPage(i["auth-login-failed-page"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenSwitchControllerVlanPortalMessageOverridesAuthDisclaimerPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPortalMessageOverridesAuthRejectPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPortalMessageOverridesAuthLoginPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanPortalMessageOverridesAuthLoginFailedPage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSwitchControllerVlanSelectedUsergroups(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenSwitchControllerVlanSelectedUsergroupsName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSwitchControllerVlanSelectedUsergroupsName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSwitchControllerVlan(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSwitchControllerVlanName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSwitchControllerVlanVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("vlanid", flattenSwitchControllerVlanVlanid(o["vlanid"], d, "vlanid", sv)); err != nil {
		if !fortiAPIPatch(o["vlanid"]) {
			return fmt.Errorf("Error reading vlanid: %v", err)
		}
	}

	if err = d.Set("comments", flattenSwitchControllerVlanComments(o["comments"], d, "comments", sv)); err != nil {
		if !fortiAPIPatch(o["comments"]) {
			return fmt.Errorf("Error reading comments: %v", err)
		}
	}

	if err = d.Set("color", flattenSwitchControllerVlanColor(o["color"], d, "color", sv)); err != nil {
		if !fortiAPIPatch(o["color"]) {
			return fmt.Errorf("Error reading color: %v", err)
		}
	}

	if err = d.Set("security", flattenSwitchControllerVlanSecurity(o["security"], d, "security", sv)); err != nil {
		if !fortiAPIPatch(o["security"]) {
			return fmt.Errorf("Error reading security: %v", err)
		}
	}

	if err = d.Set("auth", flattenSwitchControllerVlanAuth(o["auth"], d, "auth", sv)); err != nil {
		if !fortiAPIPatch(o["auth"]) {
			return fmt.Errorf("Error reading auth: %v", err)
		}
	}

	if err = d.Set("radius_server", flattenSwitchControllerVlanRadiusServer(o["radius-server"], d, "radius_server", sv)); err != nil {
		if !fortiAPIPatch(o["radius-server"]) {
			return fmt.Errorf("Error reading radius_server: %v", err)
		}
	}

	if err = d.Set("usergroup", flattenSwitchControllerVlanUsergroup(o["usergroup"], d, "usergroup", sv)); err != nil {
		if !fortiAPIPatch(o["usergroup"]) {
			return fmt.Errorf("Error reading usergroup: %v", err)
		}
	}

	if err = d.Set("portal_message_override_group", flattenSwitchControllerVlanPortalMessageOverrideGroup(o["portal-message-override-group"], d, "portal_message_override_group", sv)); err != nil {
		if !fortiAPIPatch(o["portal-message-override-group"]) {
			return fmt.Errorf("Error reading portal_message_override_group: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("portal_message_overrides", flattenSwitchControllerVlanPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides", sv)); err != nil {
			if !fortiAPIPatch(o["portal-message-overrides"]) {
				return fmt.Errorf("Error reading portal_message_overrides: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("portal_message_overrides"); ok {
			if err = d.Set("portal_message_overrides", flattenSwitchControllerVlanPortalMessageOverrides(o["portal-message-overrides"], d, "portal_message_overrides", sv)); err != nil {
				if !fortiAPIPatch(o["portal-message-overrides"]) {
					return fmt.Errorf("Error reading portal_message_overrides: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("selected_usergroups", flattenSwitchControllerVlanSelectedUsergroups(o["selected-usergroups"], d, "selected_usergroups", sv)); err != nil {
			if !fortiAPIPatch(o["selected-usergroups"]) {
				return fmt.Errorf("Error reading selected_usergroups: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("selected_usergroups"); ok {
			if err = d.Set("selected_usergroups", flattenSwitchControllerVlanSelectedUsergroups(o["selected-usergroups"], d, "selected_usergroups", sv)); err != nil {
				if !fortiAPIPatch(o["selected-usergroups"]) {
					return fmt.Errorf("Error reading selected_usergroups: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSwitchControllerVlanFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSwitchControllerVlanName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanVlanid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanComments(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanSecurity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanAuth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanRadiusServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanUsergroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPortalMessageOverrideGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPortalMessageOverrides(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "auth_disclaimer_page"
	if _, ok := d.GetOk(pre_append); ok {

		result["auth-disclaimer-page"], _ = expandSwitchControllerVlanPortalMessageOverridesAuthDisclaimerPage(d, i["auth_disclaimer_page"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_reject_page"
	if _, ok := d.GetOk(pre_append); ok {

		result["auth-reject-page"], _ = expandSwitchControllerVlanPortalMessageOverridesAuthRejectPage(d, i["auth_reject_page"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_login_page"
	if _, ok := d.GetOk(pre_append); ok {

		result["auth-login-page"], _ = expandSwitchControllerVlanPortalMessageOverridesAuthLoginPage(d, i["auth_login_page"], pre_append, sv)
	}
	pre_append = pre + ".0." + "auth_login_failed_page"
	if _, ok := d.GetOk(pre_append); ok {

		result["auth-login-failed-page"], _ = expandSwitchControllerVlanPortalMessageOverridesAuthLoginFailedPage(d, i["auth_login_failed_page"], pre_append, sv)
	}

	return result, nil
}

func expandSwitchControllerVlanPortalMessageOverridesAuthDisclaimerPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPortalMessageOverridesAuthRejectPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPortalMessageOverridesAuthLoginPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanPortalMessageOverridesAuthLoginFailedPage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSwitchControllerVlanSelectedUsergroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandSwitchControllerVlanSelectedUsergroupsName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSwitchControllerVlanSelectedUsergroupsName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSwitchControllerVlan(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {

		t, err := expandSwitchControllerVlanName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {

		t, err := expandSwitchControllerVlanVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("vlanid"); ok {

		t, err := expandSwitchControllerVlanVlanid(d, v, "vlanid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vlanid"] = t
		}
	}

	if v, ok := d.GetOk("comments"); ok {

		t, err := expandSwitchControllerVlanComments(d, v, "comments", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["comments"] = t
		}
	}

	if v, ok := d.GetOkExists("color"); ok {

		t, err := expandSwitchControllerVlanColor(d, v, "color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["color"] = t
		}
	}

	if v, ok := d.GetOk("security"); ok {

		t, err := expandSwitchControllerVlanSecurity(d, v, "security", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security"] = t
		}
	}

	if v, ok := d.GetOk("auth"); ok {

		t, err := expandSwitchControllerVlanAuth(d, v, "auth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["auth"] = t
		}
	}

	if v, ok := d.GetOk("radius_server"); ok {

		t, err := expandSwitchControllerVlanRadiusServer(d, v, "radius_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["radius-server"] = t
		}
	}

	if v, ok := d.GetOk("usergroup"); ok {

		t, err := expandSwitchControllerVlanUsergroup(d, v, "usergroup", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["usergroup"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_override_group"); ok {

		t, err := expandSwitchControllerVlanPortalMessageOverrideGroup(d, v, "portal_message_override_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-override-group"] = t
		}
	}

	if v, ok := d.GetOk("portal_message_overrides"); ok {

		t, err := expandSwitchControllerVlanPortalMessageOverrides(d, v, "portal_message_overrides", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["portal-message-overrides"] = t
		}
	}

	if v, ok := d.GetOk("selected_usergroups"); ok {

		t, err := expandSwitchControllerVlanSelectedUsergroups(d, v, "selected_usergroups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selected-usergroups"] = t
		}
	}

	return &obj, nil
}
