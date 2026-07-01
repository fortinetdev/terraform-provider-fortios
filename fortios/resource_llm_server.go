// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure LLM Proxy servers.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLlmServer() *schema.Resource {
	return &schema.Resource{
		Create: resourceLlmServerCreate,
		Read:   resourceLlmServerRead,
		Update: resourceLlmServerUpdate,
		Delete: resourceLlmServerDelete,

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
			"display_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"built_in_server": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"azure_resource_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"end_point": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"chat_completions_api": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"image_gen_api": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"anthropic_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"azure_api_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"models": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 79),
				Optional:     true,
			},
			"accept_custom_model": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"custom_model_allow_regex": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"custom_model_block_regex": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"verify_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"api_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
		},
	}
}

func resourceLlmServerCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LlmServer resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLlmServer(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLlmServer(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating LlmServer resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateLlmServer(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating LlmServer resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LlmServer")
	}

	return resourceLlmServerRead(d, m)
}

func resourceLlmServerUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmServer(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LlmServer resource while getting object: %v", err)
	}

	o, err := c.UpdateLlmServer(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LlmServer resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LlmServer")
	}

	return resourceLlmServerRead(d, m)
}

func resourceLlmServerDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteLlmServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting LlmServer resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLlmServerRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLlmServer(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LlmServer resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLlmServer(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LlmServer resource from API: %v", err)
	}
	return nil
}

func flattenLlmServerName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerDisplayName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerBuiltInServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerAzureResourceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerEndPoint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerChatCompletionsApi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerImageGenApi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerAnthropicVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerAzureApiVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerModels(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerAcceptCustomModel(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerCustomModelAllowRegex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerCustomModelBlockRegex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerVerifyCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmServerApiKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLlmServer(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("name", flattenLlmServerName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("display_name", flattenLlmServerDisplayName(o["display-name"], d, "display_name", sv)); err != nil {
		if !fortiAPIPatch(o["display-name"]) {
			return fmt.Errorf("Error reading display_name: %v", err)
		}
	}

	if err = d.Set("type", flattenLlmServerType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("built_in_server", flattenLlmServerBuiltInServer(o["built-in-server"], d, "built_in_server", sv)); err != nil {
		if !fortiAPIPatch(o["built-in-server"]) {
			return fmt.Errorf("Error reading built_in_server: %v", err)
		}
	}

	if err = d.Set("azure_resource_name", flattenLlmServerAzureResourceName(o["azure-resource-name"], d, "azure_resource_name", sv)); err != nil {
		if !fortiAPIPatch(o["azure-resource-name"]) {
			return fmt.Errorf("Error reading azure_resource_name: %v", err)
		}
	}

	if err = d.Set("end_point", flattenLlmServerEndPoint(o["end-point"], d, "end_point", sv)); err != nil {
		if !fortiAPIPatch(o["end-point"]) {
			return fmt.Errorf("Error reading end_point: %v", err)
		}
	}

	if err = d.Set("chat_completions_api", flattenLlmServerChatCompletionsApi(o["chat-completions-api"], d, "chat_completions_api", sv)); err != nil {
		if !fortiAPIPatch(o["chat-completions-api"]) {
			return fmt.Errorf("Error reading chat_completions_api: %v", err)
		}
	}

	if err = d.Set("image_gen_api", flattenLlmServerImageGenApi(o["image-gen-api"], d, "image_gen_api", sv)); err != nil {
		if !fortiAPIPatch(o["image-gen-api"]) {
			return fmt.Errorf("Error reading image_gen_api: %v", err)
		}
	}

	if err = d.Set("anthropic_version", flattenLlmServerAnthropicVersion(o["anthropic-version"], d, "anthropic_version", sv)); err != nil {
		if !fortiAPIPatch(o["anthropic-version"]) {
			return fmt.Errorf("Error reading anthropic_version: %v", err)
		}
	}

	if err = d.Set("azure_api_version", flattenLlmServerAzureApiVersion(o["azure-api-version"], d, "azure_api_version", sv)); err != nil {
		if !fortiAPIPatch(o["azure-api-version"]) {
			return fmt.Errorf("Error reading azure_api_version: %v", err)
		}
	}

	if err = d.Set("models", flattenLlmServerModels(o["models"], d, "models", sv)); err != nil {
		if !fortiAPIPatch(o["models"]) {
			return fmt.Errorf("Error reading models: %v", err)
		}
	}

	if err = d.Set("accept_custom_model", flattenLlmServerAcceptCustomModel(o["accept-custom-model"], d, "accept_custom_model", sv)); err != nil {
		if !fortiAPIPatch(o["accept-custom-model"]) {
			return fmt.Errorf("Error reading accept_custom_model: %v", err)
		}
	}

	if err = d.Set("custom_model_allow_regex", flattenLlmServerCustomModelAllowRegex(o["custom-model-allow-regex"], d, "custom_model_allow_regex", sv)); err != nil {
		if !fortiAPIPatch(o["custom-model-allow-regex"]) {
			return fmt.Errorf("Error reading custom_model_allow_regex: %v", err)
		}
	}

	if err = d.Set("custom_model_block_regex", flattenLlmServerCustomModelBlockRegex(o["custom-model-block-regex"], d, "custom_model_block_regex", sv)); err != nil {
		if !fortiAPIPatch(o["custom-model-block-regex"]) {
			return fmt.Errorf("Error reading custom_model_block_regex: %v", err)
		}
	}

	if err = d.Set("verify_cert", flattenLlmServerVerifyCert(o["verify-cert"], d, "verify_cert", sv)); err != nil {
		if !fortiAPIPatch(o["verify-cert"]) {
			return fmt.Errorf("Error reading verify_cert: %v", err)
		}
	}

	if err = d.Set("api_key", flattenLlmServerApiKey(o["api-key"], d, "api_key", sv)); err != nil {
		if !fortiAPIPatch(o["api-key"]) {
			return fmt.Errorf("Error reading api_key: %v", err)
		}
	}

	return nil
}

func flattenLlmServerFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLlmServerName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerDisplayName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerBuiltInServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerAzureResourceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerEndPoint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerChatCompletionsApi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerImageGenApi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerAnthropicVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerAzureApiVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerModels(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerAcceptCustomModel(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerCustomModelAllowRegex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerCustomModelBlockRegex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerVerifyCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmServerApiKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLlmServer(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandLlmServerName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("display_name"); ok {
		t, err := expandLlmServerDisplayName(d, v, "display_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["display-name"] = t
		}
	} else if d.HasChange("display_name") {
		obj["display-name"] = nil
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandLlmServerType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("built_in_server"); ok {
		t, err := expandLlmServerBuiltInServer(d, v, "built_in_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["built-in-server"] = t
		}
	}

	if v, ok := d.GetOk("azure_resource_name"); ok {
		t, err := expandLlmServerAzureResourceName(d, v, "azure_resource_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-resource-name"] = t
		}
	} else if d.HasChange("azure_resource_name") {
		obj["azure-resource-name"] = nil
	}

	if v, ok := d.GetOk("end_point"); ok {
		t, err := expandLlmServerEndPoint(d, v, "end_point", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["end-point"] = t
		}
	} else if d.HasChange("end_point") {
		obj["end-point"] = nil
	}

	if v, ok := d.GetOk("chat_completions_api"); ok {
		t, err := expandLlmServerChatCompletionsApi(d, v, "chat_completions_api", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chat-completions-api"] = t
		}
	}

	if v, ok := d.GetOk("image_gen_api"); ok {
		t, err := expandLlmServerImageGenApi(d, v, "image_gen_api", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-gen-api"] = t
		}
	}

	if v, ok := d.GetOk("anthropic_version"); ok {
		t, err := expandLlmServerAnthropicVersion(d, v, "anthropic_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["anthropic-version"] = t
		}
	}

	if v, ok := d.GetOk("azure_api_version"); ok {
		t, err := expandLlmServerAzureApiVersion(d, v, "azure_api_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-api-version"] = t
		}
	}

	if v, ok := d.GetOk("models"); ok {
		t, err := expandLlmServerModels(d, v, "models", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["models"] = t
		}
	} else if d.HasChange("models") {
		obj["models"] = nil
	}

	if v, ok := d.GetOk("accept_custom_model"); ok {
		t, err := expandLlmServerAcceptCustomModel(d, v, "accept_custom_model", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accept-custom-model"] = t
		}
	}

	if v, ok := d.GetOk("custom_model_allow_regex"); ok {
		t, err := expandLlmServerCustomModelAllowRegex(d, v, "custom_model_allow_regex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-model-allow-regex"] = t
		}
	} else if d.HasChange("custom_model_allow_regex") {
		obj["custom-model-allow-regex"] = nil
	}

	if v, ok := d.GetOk("custom_model_block_regex"); ok {
		t, err := expandLlmServerCustomModelBlockRegex(d, v, "custom_model_block_regex", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["custom-model-block-regex"] = t
		}
	} else if d.HasChange("custom_model_block_regex") {
		obj["custom-model-block-regex"] = nil
	}

	if v, ok := d.GetOk("verify_cert"); ok {
		t, err := expandLlmServerVerifyCert(d, v, "verify_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-cert"] = t
		}
	}

	if v, ok := d.GetOk("api_key"); ok {
		t, err := expandLlmServerApiKey(d, v, "api_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-key"] = t
		}
	} else if d.HasChange("api_key") {
		obj["api-key"] = nil
	}

	return &obj, nil
}
