// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure FortiCloud SSO admin users.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSsoFortigateCloudAdmin() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSsoFortigateCloudAdminCreate,
		Read:   resourceSystemSsoFortigateCloudAdminRead,
		Update: resourceSystemSsoFortigateCloudAdminUpdate,
		Delete: resourceSystemSsoFortigateCloudAdminDelete,

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
				Optional:     true,
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

func resourceSystemSsoFortigateCloudAdminCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSsoFortigateCloudAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSsoFortigateCloudAdmin resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadSystemSsoFortigateCloudAdmin(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateSystemSsoFortigateCloudAdmin(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating SystemSsoFortigateCloudAdmin resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateSystemSsoFortigateCloudAdmin(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating SystemSsoFortigateCloudAdmin resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSsoFortigateCloudAdmin")
	}

	return resourceSystemSsoFortigateCloudAdminRead(d, m)
}

func resourceSystemSsoFortigateCloudAdminUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSsoFortigateCloudAdmin(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoFortigateCloudAdmin resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSsoFortigateCloudAdmin(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSsoFortigateCloudAdmin resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSsoFortigateCloudAdmin")
	}

	return resourceSystemSsoFortigateCloudAdminRead(d, m)
}

func resourceSystemSsoFortigateCloudAdminDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSsoFortigateCloudAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSsoFortigateCloudAdmin resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSsoFortigateCloudAdminRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSsoFortigateCloudAdmin(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoFortigateCloudAdmin resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSsoFortigateCloudAdmin(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSsoFortigateCloudAdmin resource from API: %v", err)
	}
	return nil
}

func flattenSystemSsoFortigateCloudAdminName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminGuiThemeType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminGuiTheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminGuiCustomTheme(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminGuiLlmProvider(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminOpenaiApiKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminOpenaiApiKeyPart2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminOpenaiModel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminOpenaiProjectId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminOpenaiOrgId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSsoFortigateCloudAdminVdom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSsoFortigateCloudAdminVdomName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSsoFortigateCloudAdminVdomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemSsoFortigateCloudAdmin(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemSsoFortigateCloudAdminName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("accprofile", flattenSystemSsoFortigateCloudAdminAccprofile(o["accprofile"], d, "accprofile", sv)); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("gui_theme_type", flattenSystemSsoFortigateCloudAdminGuiThemeType(o["gui-theme-type"], d, "gui_theme_type", sv)); err != nil {
		if !fortiAPIPatch(o["gui-theme-type"]) {
			return fmt.Errorf("Error reading gui_theme_type: %v", err)
		}
	}

	if err = d.Set("gui_theme", flattenSystemSsoFortigateCloudAdminGuiTheme(o["gui-theme"], d, "gui_theme", sv)); err != nil {
		if !fortiAPIPatch(o["gui-theme"]) {
			return fmt.Errorf("Error reading gui_theme: %v", err)
		}
	}

	if err = d.Set("gui_custom_theme", flattenSystemSsoFortigateCloudAdminGuiCustomTheme(o["gui-custom-theme"], d, "gui_custom_theme", sv)); err != nil {
		if !fortiAPIPatch(o["gui-custom-theme"]) {
			return fmt.Errorf("Error reading gui_custom_theme: %v", err)
		}
	}

	if err = d.Set("gui_llm_provider", flattenSystemSsoFortigateCloudAdminGuiLlmProvider(o["gui-llm-provider"], d, "gui_llm_provider", sv)); err != nil {
		if !fortiAPIPatch(o["gui-llm-provider"]) {
			return fmt.Errorf("Error reading gui_llm_provider: %v", err)
		}
	}

	if err = d.Set("openai_api_key", flattenSystemSsoFortigateCloudAdminOpenaiApiKey(o["openai-api-key"], d, "openai_api_key", sv)); err != nil {
		if !fortiAPIPatch(o["openai-api-key"]) {
			return fmt.Errorf("Error reading openai_api_key: %v", err)
		}
	}

	if err = d.Set("openai_api_key_part2", flattenSystemSsoFortigateCloudAdminOpenaiApiKeyPart2(o["openai-api-key-part2"], d, "openai_api_key_part2", sv)); err != nil {
		if !fortiAPIPatch(o["openai-api-key-part2"]) {
			return fmt.Errorf("Error reading openai_api_key_part2: %v", err)
		}
	}

	if err = d.Set("openai_model", flattenSystemSsoFortigateCloudAdminOpenaiModel(o["openai-model"], d, "openai_model", sv)); err != nil {
		if !fortiAPIPatch(o["openai-model"]) {
			return fmt.Errorf("Error reading openai_model: %v", err)
		}
	}

	if err = d.Set("openai_project_id", flattenSystemSsoFortigateCloudAdminOpenaiProjectId(o["openai-project-id"], d, "openai_project_id", sv)); err != nil {
		if !fortiAPIPatch(o["openai-project-id"]) {
			return fmt.Errorf("Error reading openai_project_id: %v", err)
		}
	}

	if err = d.Set("openai_org_id", flattenSystemSsoFortigateCloudAdminOpenaiOrgId(o["openai-org-id"], d, "openai_org_id", sv)); err != nil {
		if !fortiAPIPatch(o["openai-org-id"]) {
			return fmt.Errorf("Error reading openai_org_id: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("vdom", flattenSystemSsoFortigateCloudAdminVdom(o["vdom"], d, "vdom", sv)); err != nil {
			if !fortiAPIPatch(o["vdom"]) {
				return fmt.Errorf("Error reading vdom: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("vdom"); ok {
			if err = d.Set("vdom", flattenSystemSsoFortigateCloudAdminVdom(o["vdom"], d, "vdom", sv)); err != nil {
				if !fortiAPIPatch(o["vdom"]) {
					return fmt.Errorf("Error reading vdom: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemSsoFortigateCloudAdminFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSsoFortigateCloudAdminName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminAccprofile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminGuiThemeType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminGuiTheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminGuiCustomTheme(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminGuiLlmProvider(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminOpenaiApiKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminOpenaiApiKeyPart2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminOpenaiModel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminOpenaiProjectId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminOpenaiOrgId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSsoFortigateCloudAdminVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemSsoFortigateCloudAdminVdomName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSsoFortigateCloudAdminVdomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSsoFortigateCloudAdmin(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSsoFortigateCloudAdminName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("accprofile"); ok {
		t, err := expandSystemSsoFortigateCloudAdminAccprofile(d, v, "accprofile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	} else if d.HasChange("accprofile") {
		obj["accprofile"] = nil
	}

	if v, ok := d.GetOk("gui_theme_type"); ok {
		t, err := expandSystemSsoFortigateCloudAdminGuiThemeType(d, v, "gui_theme_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-theme-type"] = t
		}
	}

	if v, ok := d.GetOk("gui_theme"); ok {
		t, err := expandSystemSsoFortigateCloudAdminGuiTheme(d, v, "gui_theme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-theme"] = t
		}
	}

	if v, ok := d.GetOk("gui_custom_theme"); ok {
		t, err := expandSystemSsoFortigateCloudAdminGuiCustomTheme(d, v, "gui_custom_theme", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-custom-theme"] = t
		}
	} else if d.HasChange("gui_custom_theme") {
		obj["gui-custom-theme"] = nil
	}

	if v, ok := d.GetOk("gui_llm_provider"); ok {
		t, err := expandSystemSsoFortigateCloudAdminGuiLlmProvider(d, v, "gui_llm_provider", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gui-llm-provider"] = t
		}
	}

	if v, ok := d.GetOk("openai_api_key"); ok {
		t, err := expandSystemSsoFortigateCloudAdminOpenaiApiKey(d, v, "openai_api_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-api-key"] = t
		}
	} else if d.HasChange("openai_api_key") {
		obj["openai-api-key"] = nil
	}

	if v, ok := d.GetOk("openai_api_key_part2"); ok {
		t, err := expandSystemSsoFortigateCloudAdminOpenaiApiKeyPart2(d, v, "openai_api_key_part2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-api-key-part2"] = t
		}
	} else if d.HasChange("openai_api_key_part2") {
		obj["openai-api-key-part2"] = nil
	}

	if v, ok := d.GetOk("openai_model"); ok {
		t, err := expandSystemSsoFortigateCloudAdminOpenaiModel(d, v, "openai_model", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-model"] = t
		}
	} else if d.HasChange("openai_model") {
		obj["openai-model"] = nil
	}

	if v, ok := d.GetOk("openai_project_id"); ok {
		t, err := expandSystemSsoFortigateCloudAdminOpenaiProjectId(d, v, "openai_project_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-project-id"] = t
		}
	} else if d.HasChange("openai_project_id") {
		obj["openai-project-id"] = nil
	}

	if v, ok := d.GetOk("openai_org_id"); ok {
		t, err := expandSystemSsoFortigateCloudAdminOpenaiOrgId(d, v, "openai_org_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["openai-org-id"] = t
		}
	} else if d.HasChange("openai_org_id") {
		obj["openai-org-id"] = nil
	}

	if v, ok := d.GetOk("vdom"); ok || d.HasChange("vdom") {
		t, err := expandSystemSsoFortigateCloudAdminVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	return &obj, nil
}
