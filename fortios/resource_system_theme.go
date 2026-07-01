// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure custom gui themes.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemTheme() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemThemeCreate,
		Read:   resourceSystemThemeRead,
		Update: resourceSystemThemeUpdate,
		Delete: resourceSystemThemeDelete,

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
			"update_if_exist": &schema.Schema{
				Type:     schema.TypeBool,
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"theme_template": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"base_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nav_color": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nav_style": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"font": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"font_weight": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"table_style": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"border_radius": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_color": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
			},
			"selected_color": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
			},
			"call_to_action_color": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
			},
			"accent_color": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 7),
				Optional:     true,
			},
			"banner_msg": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"banner_msg_severity": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}

func resourceSystemThemeCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemTheme(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemTheme resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemTheme(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemTheme(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating SystemTheme resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateSystemTheme(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating SystemTheme resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemTheme")
	}

	return resourceSystemThemeRead(d, m)
}

func resourceSystemThemeUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemTheme(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemTheme resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemTheme(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemTheme resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemTheme")
	}

	return resourceSystemThemeRead(d, m)
}

func resourceSystemThemeDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemTheme(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemTheme resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemThemeRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemTheme(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemTheme resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemTheme(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemTheme resource from API: %v", err)
	}
	return nil
}

func flattenSystemThemeName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeThemeTemplate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeBaseTheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeNavColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeNavStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeFont(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeFontWeight(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeTableStyle(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeBorderRadius(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeHeaderColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeSelectedColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeCallToActionColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeAccentColor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeBannerMsg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemThemeBannerMsgSeverity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemTheme(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenSystemThemeName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("theme_template", flattenSystemThemeThemeTemplate(o["theme-template"], d, "theme_template", sv)); err != nil {
		if !fortiAPIPatch(o["theme-template"]) {
			return fmt.Errorf("Error reading theme_template: %v", err)
		}
	}

	if err = d.Set("base_theme", flattenSystemThemeBaseTheme(o["base-theme"], d, "base_theme", sv)); err != nil {
		if !fortiAPIPatch(o["base-theme"]) {
			return fmt.Errorf("Error reading base_theme: %v", err)
		}
	}

	if err = d.Set("nav_color", flattenSystemThemeNavColor(o["nav-color"], d, "nav_color", sv)); err != nil {
		if !fortiAPIPatch(o["nav-color"]) {
			return fmt.Errorf("Error reading nav_color: %v", err)
		}
	}

	if err = d.Set("nav_style", flattenSystemThemeNavStyle(o["nav-style"], d, "nav_style", sv)); err != nil {
		if !fortiAPIPatch(o["nav-style"]) {
			return fmt.Errorf("Error reading nav_style: %v", err)
		}
	}

	if err = d.Set("font", flattenSystemThemeFont(o["font"], d, "font", sv)); err != nil {
		if !fortiAPIPatch(o["font"]) {
			return fmt.Errorf("Error reading font: %v", err)
		}
	}

	if err = d.Set("font_weight", flattenSystemThemeFontWeight(o["font-weight"], d, "font_weight", sv)); err != nil {
		if !fortiAPIPatch(o["font-weight"]) {
			return fmt.Errorf("Error reading font_weight: %v", err)
		}
	}

	if err = d.Set("table_style", flattenSystemThemeTableStyle(o["table-style"], d, "table_style", sv)); err != nil {
		if !fortiAPIPatch(o["table-style"]) {
			return fmt.Errorf("Error reading table_style: %v", err)
		}
	}

	if err = d.Set("border_radius", flattenSystemThemeBorderRadius(o["border-radius"], d, "border_radius", sv)); err != nil {
		if !fortiAPIPatch(o["border-radius"]) {
			return fmt.Errorf("Error reading border_radius: %v", err)
		}
	}

	if err = d.Set("header_color", flattenSystemThemeHeaderColor(o["header-color"], d, "header_color", sv)); err != nil {
		if !fortiAPIPatch(o["header-color"]) {
			return fmt.Errorf("Error reading header_color: %v", err)
		}
	}

	if err = d.Set("selected_color", flattenSystemThemeSelectedColor(o["selected-color"], d, "selected_color", sv)); err != nil {
		if !fortiAPIPatch(o["selected-color"]) {
			return fmt.Errorf("Error reading selected_color: %v", err)
		}
	}

	if err = d.Set("call_to_action_color", flattenSystemThemeCallToActionColor(o["call-to-action-color"], d, "call_to_action_color", sv)); err != nil {
		if !fortiAPIPatch(o["call-to-action-color"]) {
			return fmt.Errorf("Error reading call_to_action_color: %v", err)
		}
	}

	if err = d.Set("accent_color", flattenSystemThemeAccentColor(o["accent-color"], d, "accent_color", sv)); err != nil {
		if !fortiAPIPatch(o["accent-color"]) {
			return fmt.Errorf("Error reading accent_color: %v", err)
		}
	}

	if err = d.Set("banner_msg", flattenSystemThemeBannerMsg(o["banner-msg"], d, "banner_msg", sv)); err != nil {
		if !fortiAPIPatch(o["banner-msg"]) {
			return fmt.Errorf("Error reading banner_msg: %v", err)
		}
	}

	if err = d.Set("banner_msg_severity", flattenSystemThemeBannerMsgSeverity(o["banner-msg-severity"], d, "banner_msg_severity", sv)); err != nil {
		if !fortiAPIPatch(o["banner-msg-severity"]) {
			return fmt.Errorf("Error reading banner_msg_severity: %v", err)
		}
	}

	return nil
}

func flattenSystemThemeFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemThemeName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeThemeTemplate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeBaseTheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeNavColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeNavStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeFont(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeFontWeight(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeTableStyle(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeBorderRadius(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeHeaderColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeSelectedColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeCallToActionColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeAccentColor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeBannerMsg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemThemeBannerMsgSeverity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemTheme(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemThemeName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("theme_template"); ok {
		t, err := expandSystemThemeThemeTemplate(d, v, "theme_template", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["theme-template"] = t
		}
	}

	if v, ok := d.GetOk("base_theme"); ok {
		t, err := expandSystemThemeBaseTheme(d, v, "base_theme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["base-theme"] = t
		}
	}

	if v, ok := d.GetOk("nav_color"); ok {
		t, err := expandSystemThemeNavColor(d, v, "nav_color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nav-color"] = t
		}
	}

	if v, ok := d.GetOk("nav_style"); ok {
		t, err := expandSystemThemeNavStyle(d, v, "nav_style", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nav-style"] = t
		}
	}

	if v, ok := d.GetOk("font"); ok {
		t, err := expandSystemThemeFont(d, v, "font", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font"] = t
		}
	}

	if v, ok := d.GetOk("font_weight"); ok {
		t, err := expandSystemThemeFontWeight(d, v, "font_weight", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["font-weight"] = t
		}
	}

	if v, ok := d.GetOk("table_style"); ok {
		t, err := expandSystemThemeTableStyle(d, v, "table_style", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["table-style"] = t
		}
	}

	if v, ok := d.GetOk("border_radius"); ok {
		t, err := expandSystemThemeBorderRadius(d, v, "border_radius", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["border-radius"] = t
		}
	}

	if v, ok := d.GetOk("header_color"); ok {
		t, err := expandSystemThemeHeaderColor(d, v, "header_color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-color"] = t
		}
	} else if d.HasChange("header_color") {
		obj["header-color"] = nil
	}

	if v, ok := d.GetOk("selected_color"); ok {
		t, err := expandSystemThemeSelectedColor(d, v, "selected_color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["selected-color"] = t
		}
	} else if d.HasChange("selected_color") {
		obj["selected-color"] = nil
	}

	if v, ok := d.GetOk("call_to_action_color"); ok {
		t, err := expandSystemThemeCallToActionColor(d, v, "call_to_action_color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["call-to-action-color"] = t
		}
	} else if d.HasChange("call_to_action_color") {
		obj["call-to-action-color"] = nil
	}

	if v, ok := d.GetOk("accent_color"); ok {
		t, err := expandSystemThemeAccentColor(d, v, "accent_color", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accent-color"] = t
		}
	} else if d.HasChange("accent_color") {
		obj["accent-color"] = nil
	}

	if v, ok := d.GetOk("banner_msg"); ok {
		t, err := expandSystemThemeBannerMsg(d, v, "banner_msg", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner-msg"] = t
		}
	} else if d.HasChange("banner_msg") {
		obj["banner-msg"] = nil
	}

	if v, ok := d.GetOk("banner_msg_severity"); ok {
		t, err := expandSystemThemeBannerMsgSeverity(d, v, "banner_msg_severity", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["banner-msg-severity"] = t
		}
	}

	return &obj, nil
}
