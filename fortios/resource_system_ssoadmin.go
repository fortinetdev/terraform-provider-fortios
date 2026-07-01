// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure SSO admin users.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSsoAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSsoAdminCreate,
		Read:   resourceSystemSsoAdminRead,
		Update: resourceSystemSsoAdminUpdate,
		Delete: resourceSystemSsoAdminDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 64),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"accprofile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Required:     true,
			},
			"gui_theme_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_theme": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gui_custom_theme": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"gui_llm_provider": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"openai_api_key": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"openai_api_key_part2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"openai_model": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"openai_project_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"openai_org_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"vdom": &schema.Schema{
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
			"gui_ignore_release_overview_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 5),
				Optional:     true,
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

func resourceSystemSsoAdminCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSsoAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoAdmin resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSsoAdmin(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSsoAdmin(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating SystemSsoAdmin resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateSystemSsoAdmin(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating SystemSsoAdmin resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSsoAdmin")
	}

	return resourceSystemSsoAdminRead(d, m)
}

func resourceSystemSsoAdminUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSsoAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoAdmin resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSsoAdmin(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSsoAdmin")
	}

	return resourceSystemSsoAdminRead(d, m)
}

func resourceSystemSsoAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSsoAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSsoAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSsoAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSsoAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSsoAdmin(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemSsoAdminName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminGuiThemeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminGuiTheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminGuiCustomTheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminGuiLlmProvider(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminOpenaiApiKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminOpenaiApiKeyPart2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminOpenaiModel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminOpenaiProjectId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminOpenaiOrgId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminVdom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSsoAdminVdomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSsoAdminVdomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoAdminGuiIgnoreReleaseOverviewVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSsoAdmin(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemSsoAdminName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("accprofile", flattenSystemSsoAdminAccprofile(o["accprofile"], d, "accprofile", sv)); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("gui_theme_type", flattenSystemSsoAdminGuiThemeType(o["gui-theme-type"], d, "gui_theme_type", sv)); err != nil {
		if !fortiAPIPatch(o["gui-theme-type"]) {
			return fmt.Errorf("Error reading gui_theme_type: %v", err)
		}
	}

	if err = d.Set("gui_theme", flattenSystemSsoAdminGuiTheme(o["gui-theme"], d, "gui_theme", sv)); err != nil {
		if !fortiAPIPatch(o["gui-theme"]) {
			return fmt.Errorf("Error reading gui_theme: %v", err)
		}
	}

	if err = d.Set("gui_custom_theme", flattenSystemSsoAdminGuiCustomTheme(o["gui-custom-theme"], d, "gui_custom_theme", sv)); err != nil {
		if !fortiAPIPatch(o["gui-custom-theme"]) {
			return fmt.Errorf("Error reading gui_custom_theme: %v", err)
		}
	}

	if err = d.Set("gui_llm_provider", flattenSystemSsoAdminGuiLlmProvider(o["gui-llm-provider"], d, "gui_llm_provider", sv)); err != nil {
		if !fortiAPIPatch(o["gui-llm-provider"]) {
			return fmt.Errorf("Error reading gui_llm_provider: %v", err)
		}
	}

	if err = d.Set("openai_api_key", flattenSystemSsoAdminOpenaiApiKey(o["openai-api-key"], d, "openai_api_key", sv)); err != nil {
		if !fortiAPIPatch(o["openai-api-key"]) {
			return fmt.Errorf("Error reading openai_api_key: %v", err)
		}
	}

	if err = d.Set("openai_api_key_part2", flattenSystemSsoAdminOpenaiApiKeyPart2(o["openai-api-key-part2"], d, "openai_api_key_part2", sv)); err != nil {
		if !fortiAPIPatch(o["openai-api-key-part2"]) {
			return fmt.Errorf("Error reading openai_api_key_part2: %v", err)
		}
	}

	if err = d.Set("openai_model", flattenSystemSsoAdminOpenaiModel(o["openai-model"], d, "openai_model", sv)); err != nil {
		if !fortiAPIPatch(o["openai-model"]) {
			return fmt.Errorf("Error reading openai_model: %v", err)
		}
	}

	if err = d.Set("openai_project_id", flattenSystemSsoAdminOpenaiProjectId(o["openai-project-id"], d, "openai_project_id", sv)); err != nil {
		if !fortiAPIPatch(o["openai-project-id"]) {
			return fmt.Errorf("Error reading openai_project_id: %v", err)
		}
	}

	if err = d.Set("openai_org_id", flattenSystemSsoAdminOpenaiOrgId(o["openai-org-id"], d, "openai_org_id", sv)); err != nil {
		if !fortiAPIPatch(o["openai-org-id"]) {
			return fmt.Errorf("Error reading openai_org_id: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("vdom", flattenSystemSsoAdminVdom(o["vdom"], d, "vdom", sv)); err != nil {
			if !fortiAPIPatch(o["vdom"]) {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vdom"); ok {
			if err = d.Set("vdom", flattenSystemSsoAdminVdom(o["vdom"], d, "vdom", sv)); err != nil {
				if !fortiAPIPatch(o["vdom"]) {
					return fmt.Errorf("Error reading vdom: %v", err)
				}
			}
		}
	}

	if err = d.Set("gui_ignore_release_overview_version", flattenSystemSsoAdminGuiIgnoreReleaseOverviewVersion(o["gui-ignore-release-overview-version"], d, "gui_ignore_release_overview_version", sv)); err != nil {
		if !fortiAPIPatch(o["gui-ignore-release-overview-version"]) {
			return fmt.Errorf("Error reading gui_ignore_release_overview_version: %v", err)
		}
	}

	return nil
}

func flattenSystemSsoAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSsoAdminName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminAccprofile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminGuiThemeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminGuiTheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminGuiCustomTheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminGuiLlmProvider(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminOpenaiApiKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminOpenaiApiKeyPart2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminOpenaiModel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminOpenaiProjectId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminOpenaiOrgId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemSsoAdminVdomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSsoAdminVdomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoAdminGuiIgnoreReleaseOverviewVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSsoAdmin(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSsoAdminName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("accprofile"); ok {
		t, err := expandSystemSsoAdminAccprofile(d, v, "accprofile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	} else if d.HasChange("accprofile") {
		obj["accprofile"] = nil
	}

	if v, ok := d.GetOk("gui_theme_type"); ok {
		t, err := expandSystemSsoAdminGuiThemeType(d, v, "gui_theme_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-theme-type"] = t
		}
	}

	if v, ok := d.GetOk("gui_theme"); ok {
		t, err := expandSystemSsoAdminGuiTheme(d, v, "gui_theme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-theme"] = t
		}
	}

	if v, ok := d.GetOk("gui_custom_theme"); ok {
		t, err := expandSystemSsoAdminGuiCustomTheme(d, v, "gui_custom_theme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-custom-theme"] = t
		}
	} else if d.HasChange("gui_custom_theme") {
		obj["gui-custom-theme"] = nil
	}

	if v, ok := d.GetOk("gui_llm_provider"); ok {
		t, err := expandSystemSsoAdminGuiLlmProvider(d, v, "gui_llm_provider", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-llm-provider"] = t
		}
	}

	if v, ok := d.GetOk("openai_api_key"); ok {
		t, err := expandSystemSsoAdminOpenaiApiKey(d, v, "openai_api_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-api-key"] = t
		}
	} else if d.HasChange("openai_api_key") {
		obj["openai-api-key"] = nil
	}

	if v, ok := d.GetOk("openai_api_key_part2"); ok {
		t, err := expandSystemSsoAdminOpenaiApiKeyPart2(d, v, "openai_api_key_part2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-api-key-part2"] = t
		}
	} else if d.HasChange("openai_api_key_part2") {
		obj["openai-api-key-part2"] = nil
	}

	if v, ok := d.GetOk("openai_model"); ok {
		t, err := expandSystemSsoAdminOpenaiModel(d, v, "openai_model", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-model"] = t
		}
	} else if d.HasChange("openai_model") {
		obj["openai-model"] = nil
	}

	if v, ok := d.GetOk("openai_project_id"); ok {
		t, err := expandSystemSsoAdminOpenaiProjectId(d, v, "openai_project_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-project-id"] = t
		}
	} else if d.HasChange("openai_project_id") {
		obj["openai-project-id"] = nil
	}

	if v, ok := d.GetOk("openai_org_id"); ok {
		t, err := expandSystemSsoAdminOpenaiOrgId(d, v, "openai_org_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-org-id"] = t
		}
	} else if d.HasChange("openai_org_id") {
		obj["openai-org-id"] = nil
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemSsoAdminVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("gui_ignore_release_overview_version"); ok {
		t, err := expandSystemSsoAdminGuiIgnoreReleaseOverviewVersion(d, v, "gui_ignore_release_overview_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-ignore-release-overview-version"] = t
		}
	} else if d.HasChange("gui_ignore_release_overview_version") {
		obj["gui-ignore-release-overview-version"] = nil
	}

	return &obj, nil
}
