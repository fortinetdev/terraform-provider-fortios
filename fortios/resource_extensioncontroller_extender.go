// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Extender controller configuration.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceExtensionControllerExtender() *schema.Resource {
	return &schema.Resource{
		Create: resourceExtensionControllerExtenderCreate,
		Read:   resourceExtensionControllerExtenderRead,
		Update: resourceExtensionControllerExtenderUpdate,
		Delete: resourceExtensionControllerExtenderDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 19),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"fosid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 19),
				Optional:     true,
				Computed:     true,
			},
			"authorized": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ext_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"device_id": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"extension_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"profile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"override_allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allowaccess": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"override_login_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_password_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"login_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 27),
				Optional:     true,
			},
			"override_enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enforce_bandwidth": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"bandwidth_limit": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 16776000),
				Optional:     true,
				Computed:     true,
			},
			"wan_extension": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"modem1_extension": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
						"modem2_extension": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"firmware_provision_latest": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceExtensionControllerExtenderCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectExtensionControllerExtender(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtender resource while getting object: %v", err)
	}

	o, err := c.CreateExtensionControllerExtender(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating ExtensionControllerExtender resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerExtender")
	}

	return resourceExtensionControllerExtenderRead(d, m)
}

func resourceExtensionControllerExtenderUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectExtensionControllerExtender(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtender resource while getting object: %v", err)
	}

	o, err := c.UpdateExtensionControllerExtender(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating ExtensionControllerExtender resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("ExtensionControllerExtender")
	}

	return resourceExtensionControllerExtenderRead(d, m)
}

func resourceExtensionControllerExtenderDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteExtensionControllerExtender(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting ExtensionControllerExtender resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceExtensionControllerExtenderRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadExtensionControllerExtender(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtender resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectExtensionControllerExtender(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading ExtensionControllerExtender resource from API: %v", err)
	}
	return nil
}

func flattenExtensionControllerExtenderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderAuthorized(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderExtName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderDeviceId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderExtensionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderProfile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderOverrideAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderAllowaccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderOverrideLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderLoginPasswordChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderLoginPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderOverrideEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderEnforceBandwidth(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderBandwidthLimit(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderWanExtension(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "modem1_extension"
	if _, ok := i["modem1-extension"]; ok {
		result["modem1_extension"] = flattenExtensionControllerExtenderWanExtensionModem1Extension(i["modem1-extension"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "modem2_extension"
	if _, ok := i["modem2-extension"]; ok {
		result["modem2_extension"] = flattenExtensionControllerExtenderWanExtensionModem2Extension(i["modem2-extension"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenExtensionControllerExtenderWanExtensionModem1Extension(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderWanExtensionModem2Extension(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenExtensionControllerExtenderFirmwareProvisionLatest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectExtensionControllerExtender(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenExtensionControllerExtenderName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("fosid", flattenExtensionControllerExtenderId(o["id"], d, "fosid", sv)); err != nil {
		if !fortiAPIPatch(o["id"]) {
			return fmt.Errorf("Error reading fosid: %v", err)
		}
	}

	if err = d.Set("authorized", flattenExtensionControllerExtenderAuthorized(o["authorized"], d, "authorized", sv)); err != nil {
		if !fortiAPIPatch(o["authorized"]) {
			return fmt.Errorf("Error reading authorized: %v", err)
		}
	}

	if err = d.Set("ext_name", flattenExtensionControllerExtenderExtName(o["ext-name"], d, "ext_name", sv)); err != nil {
		if !fortiAPIPatch(o["ext-name"]) {
			return fmt.Errorf("Error reading ext_name: %v", err)
		}
	}

	if err = d.Set("description", flattenExtensionControllerExtenderDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("vdom", flattenExtensionControllerExtenderVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if err = d.Set("device_id", flattenExtensionControllerExtenderDeviceId(o["device-id"], d, "device_id", sv)); err != nil {
		if !fortiAPIPatch(o["device-id"]) {
			return fmt.Errorf("Error reading device_id: %v", err)
		}
	}

	if err = d.Set("extension_type", flattenExtensionControllerExtenderExtensionType(o["extension-type"], d, "extension_type", sv)); err != nil {
		if !fortiAPIPatch(o["extension-type"]) {
			return fmt.Errorf("Error reading extension_type: %v", err)
		}
	}

	if err = d.Set("profile", flattenExtensionControllerExtenderProfile(o["profile"], d, "profile", sv)); err != nil {
		if !fortiAPIPatch(o["profile"]) {
			return fmt.Errorf("Error reading profile: %v", err)
		}
	}

	if err = d.Set("override_allowaccess", flattenExtensionControllerExtenderOverrideAllowaccess(o["override-allowaccess"], d, "override_allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["override-allowaccess"]) {
			return fmt.Errorf("Error reading override_allowaccess: %v", err)
		}
	}

	if err = d.Set("allowaccess", flattenExtensionControllerExtenderAllowaccess(o["allowaccess"], d, "allowaccess", sv)); err != nil {
		if !fortiAPIPatch(o["allowaccess"]) {
			return fmt.Errorf("Error reading allowaccess: %v", err)
		}
	}

	if err = d.Set("override_login_password_change", flattenExtensionControllerExtenderOverrideLoginPasswordChange(o["override-login-password-change"], d, "override_login_password_change", sv)); err != nil {
		if !fortiAPIPatch(o["override-login-password-change"]) {
			return fmt.Errorf("Error reading override_login_password_change: %v", err)
		}
	}

	if err = d.Set("login_password_change", flattenExtensionControllerExtenderLoginPasswordChange(o["login-password-change"], d, "login_password_change", sv)); err != nil {
		if !fortiAPIPatch(o["login-password-change"]) {
			return fmt.Errorf("Error reading login_password_change: %v", err)
		}
	}

	if err = d.Set("login_password", flattenExtensionControllerExtenderLoginPassword(o["login-password"], d, "login_password", sv)); err != nil {
		if !fortiAPIPatch(o["login-password"]) {
			return fmt.Errorf("Error reading login_password: %v", err)
		}
	}

	if err = d.Set("override_enforce_bandwidth", flattenExtensionControllerExtenderOverrideEnforceBandwidth(o["override-enforce-bandwidth"], d, "override_enforce_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["override-enforce-bandwidth"]) {
			return fmt.Errorf("Error reading override_enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("enforce_bandwidth", flattenExtensionControllerExtenderEnforceBandwidth(o["enforce-bandwidth"], d, "enforce_bandwidth", sv)); err != nil {
		if !fortiAPIPatch(o["enforce-bandwidth"]) {
			return fmt.Errorf("Error reading enforce_bandwidth: %v", err)
		}
	}

	if err = d.Set("bandwidth_limit", flattenExtensionControllerExtenderBandwidthLimit(o["bandwidth-limit"], d, "bandwidth_limit", sv)); err != nil {
		if !fortiAPIPatch(o["bandwidth-limit"]) {
			return fmt.Errorf("Error reading bandwidth_limit: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("wan_extension", flattenExtensionControllerExtenderWanExtension(o["wan-extension"], d, "wan_extension", sv)); err != nil {
			if !fortiAPIPatch(o["wan-extension"]) {
				return fmt.Errorf("Error reading wan_extension: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("wan_extension"); ok {
			if err = d.Set("wan_extension", flattenExtensionControllerExtenderWanExtension(o["wan-extension"], d, "wan_extension", sv)); err != nil {
				if !fortiAPIPatch(o["wan-extension"]) {
					return fmt.Errorf("Error reading wan_extension: %v", err)
				}
			}
		}
	}

	if err = d.Set("firmware_provision_latest", flattenExtensionControllerExtenderFirmwareProvisionLatest(o["firmware-provision-latest"], d, "firmware_provision_latest", sv)); err != nil {
		if !fortiAPIPatch(o["firmware-provision-latest"]) {
			return fmt.Errorf("Error reading firmware_provision_latest: %v", err)
		}
	}

	return nil
}

func flattenExtensionControllerExtenderFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandExtensionControllerExtenderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderAuthorized(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderExtName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderDeviceId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderExtensionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderProfile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderOverrideAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderAllowaccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderOverrideLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderLoginPasswordChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderLoginPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderOverrideEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderEnforceBandwidth(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderBandwidthLimit(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderWanExtension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "modem1_extension"
	if _, ok := d.GetOk(pre_append); ok {
		result["modem1-extension"], _ = expandExtensionControllerExtenderWanExtensionModem1Extension(d, i["modem1_extension"], pre_append, sv)
	}
	pre_append = pre + ".0." + "modem2_extension"
	if _, ok := d.GetOk(pre_append); ok {
		result["modem2-extension"], _ = expandExtensionControllerExtenderWanExtensionModem2Extension(d, i["modem2_extension"], pre_append, sv)
	}

	return result, nil
}

func expandExtensionControllerExtenderWanExtensionModem1Extension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderWanExtensionModem2Extension(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandExtensionControllerExtenderFirmwareProvisionLatest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectExtensionControllerExtender(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandExtensionControllerExtenderName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("fosid"); ok {
		t, err := expandExtensionControllerExtenderId(d, v, "fosid", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["id"] = t
		}
	}

	if v, ok := d.GetOk("authorized"); ok {
		t, err := expandExtensionControllerExtenderAuthorized(d, v, "authorized", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["authorized"] = t
		}
	}

	if v, ok := d.GetOk("ext_name"); ok {
		t, err := expandExtensionControllerExtenderExtName(d, v, "ext_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ext-name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandExtensionControllerExtenderDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	}

	if v, ok := d.GetOkExists("vdom"); ok {
		t, err := expandExtensionControllerExtenderVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOkExists("device_id"); ok {
		t, err := expandExtensionControllerExtenderDeviceId(d, v, "device_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["device-id"] = t
		}
	}

	if v, ok := d.GetOk("extension_type"); ok {
		t, err := expandExtensionControllerExtenderExtensionType(d, v, "extension_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-type"] = t
		}
	}

	if v, ok := d.GetOk("profile"); ok {
		t, err := expandExtensionControllerExtenderProfile(d, v, "profile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["profile"] = t
		}
	}

	if v, ok := d.GetOk("override_allowaccess"); ok {
		t, err := expandExtensionControllerExtenderOverrideAllowaccess(d, v, "override_allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("allowaccess"); ok {
		t, err := expandExtensionControllerExtenderAllowaccess(d, v, "allowaccess", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allowaccess"] = t
		}
	}

	if v, ok := d.GetOk("override_login_password_change"); ok {
		t, err := expandExtensionControllerExtenderOverrideLoginPasswordChange(d, v, "override_login_password_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("login_password_change"); ok {
		t, err := expandExtensionControllerExtenderLoginPasswordChange(d, v, "login_password_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password-change"] = t
		}
	}

	if v, ok := d.GetOk("login_password"); ok {
		t, err := expandExtensionControllerExtenderLoginPassword(d, v, "login_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-password"] = t
		}
	}

	if v, ok := d.GetOk("override_enforce_bandwidth"); ok {
		t, err := expandExtensionControllerExtenderOverrideEnforceBandwidth(d, v, "override_enforce_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override-enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("enforce_bandwidth"); ok {
		t, err := expandExtensionControllerExtenderEnforceBandwidth(d, v, "enforce_bandwidth", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enforce-bandwidth"] = t
		}
	}

	if v, ok := d.GetOk("bandwidth_limit"); ok {
		t, err := expandExtensionControllerExtenderBandwidthLimit(d, v, "bandwidth_limit", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["bandwidth-limit"] = t
		}
	}

	if v, ok := d.GetOk("wan_extension"); ok {
		t, err := expandExtensionControllerExtenderWanExtension(d, v, "wan_extension", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["wan-extension"] = t
		}
	}

	if v, ok := d.GetOk("firmware_provision_latest"); ok {
		t, err := expandExtensionControllerExtenderFirmwareProvisionLatest(d, v, "firmware_provision_latest", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["firmware-provision-latest"] = t
		}
	}

	return &obj, nil
}
