// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure LLM Proxy profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceLlmProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceLlmProfileCreate,
		Read:   resourceLlmProfileRead,
		Update: resourceLlmProfileUpdate,
		Delete: resourceLlmProfileDelete,

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
			"unknown_api": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"chat": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_completion_tokens": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 10000000),
							Optional:     true,
						},
						"stream": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"max_req_len": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 65535),
							Optional:     true,
							Computed:     true,
						},
						"system_prompt_mode": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"system_prompt": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"image_gen": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"list_models": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Optional: true,
				MaxItems: 1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"status": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"get_all_tables": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceLlmProfileCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating LlmProfile resource while getting object: %v", err)
	}

	update_if_exist := getUpdateIfExist(c, d)
	mkey_tf, mkey_ok := d.GetOk("name")
	mkey := fmt.Sprint(mkey_tf)
	o := make(map[string]interface{})
	existing := false

	if update_if_exist && mkey_ok {
		// check existing
		o, err = c.ReadLlmProfile(mkey, vdomparam)
		if err == nil && o != nil {
			existing = true
			// update if existing
			o, err = c.UpdateLlmProfile(obj, mkey, vdomparam)
			if err != nil {
				return fmt.Errorf("Error updating LlmProfile resource: %v", err)
			}
		}
	}

	if !existing {
		o, err = c.CreateLlmProfile(obj, vdomparam)

		if err != nil {
			return fmt.Errorf("Error creating LlmProfile resource: %v", err)
		}
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LlmProfile")
	}

	return resourceLlmProfileRead(d, m)
}

func resourceLlmProfileUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLlmProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LlmProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateLlmProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LlmProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LlmProfile")
	}

	return resourceLlmProfileRead(d, m)
}

func resourceLlmProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteLlmProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting LlmProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLlmProfileRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLlmProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LlmProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLlmProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LlmProfile resource from API: %v", err)
	}
	return nil
}

func flattenLlmProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProfileUnknownApi(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProfileLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProfileChat(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenLlmProfileChatStatus(i["status"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_completion_tokens"
	if _, ok := i["max-completion-tokens"]; ok {
		result["max_completion_tokens"] = flattenLlmProfileChatMaxCompletionTokens(i["max-completion-tokens"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "stream"
	if _, ok := i["stream"]; ok {
		result["stream"] = flattenLlmProfileChatStream(i["stream"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "max_req_len"
	if _, ok := i["max-req-len"]; ok {
		result["max_req_len"] = flattenLlmProfileChatMaxReqLen(i["max-req-len"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "system_prompt_mode"
	if _, ok := i["system-prompt-mode"]; ok {
		result["system_prompt_mode"] = flattenLlmProfileChatSystemPromptMode(i["system-prompt-mode"], d, pre_append, sv)
	}

	pre_append = pre + ".0." + "system_prompt"
	if _, ok := i["system-prompt"]; ok {
		result["system_prompt"] = flattenLlmProfileChatSystemPrompt(i["system-prompt"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLlmProfileChatStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProfileChatMaxCompletionTokens(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLlmProfileChatStream(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProfileChatMaxReqLen(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenLlmProfileChatSystemPromptMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProfileChatSystemPrompt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProfileImageGen(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenLlmProfileImageGenStatus(i["status"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLlmProfileImageGenStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLlmProfileListModels(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
	if v == nil {
		return nil
	}

	i := v.(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := i["status"]; ok {
		result["status"] = flattenLlmProfileListModelsStatus(i["status"], d, pre_append, sv)
	}

	lastresult := []map[string]interface{}{result}
	return lastresult
}

func flattenLlmProfileListModelsStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLlmProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenLlmProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("unknown_api", flattenLlmProfileUnknownApi(o["unknown-api"], d, "unknown_api", sv)); err != nil {
		if !fortiAPIPatch(o["unknown-api"]) {
			return fmt.Errorf("Error reading unknown_api: %v", err)
		}
	}

	if err = d.Set("log", flattenLlmProfileLog(o["log"], d, "log", sv)); err != nil {
		if !fortiAPIPatch(o["log"]) {
			return fmt.Errorf("Error reading log: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("chat", flattenLlmProfileChat(o["chat"], d, "chat", sv)); err != nil {
			if !fortiAPIPatch(o["chat"]) {
				return fmt.Errorf("Error reading chat: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("chat"); ok {
			if err = d.Set("chat", flattenLlmProfileChat(o["chat"], d, "chat", sv)); err != nil {
				if !fortiAPIPatch(o["chat"]) {
					return fmt.Errorf("Error reading chat: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("image_gen", flattenLlmProfileImageGen(o["image-gen"], d, "image_gen", sv)); err != nil {
			if !fortiAPIPatch(o["image-gen"]) {
				return fmt.Errorf("Error reading image_gen: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("image_gen"); ok {
			if err = d.Set("image_gen", flattenLlmProfileImageGen(o["image-gen"], d, "image_gen", sv)); err != nil {
				if !fortiAPIPatch(o["image-gen"]) {
					return fmt.Errorf("Error reading image_gen: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("list_models", flattenLlmProfileListModels(o["list-models"], d, "list_models", sv)); err != nil {
			if !fortiAPIPatch(o["list-models"]) {
				return fmt.Errorf("Error reading list_models: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("list_models"); ok {
			if err = d.Set("list_models", flattenLlmProfileListModels(o["list-models"], d, "list_models", sv)); err != nil {
				if !fortiAPIPatch(o["list-models"]) {
					return fmt.Errorf("Error reading list_models: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenLlmProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLlmProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileUnknownApi(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChat(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandLlmProfileChatStatus(d, i["status"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_completion_tokens"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-completion-tokens"], _ = expandLlmProfileChatMaxCompletionTokens(d, i["max_completion_tokens"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["max-completion-tokens"] = nil
	}
	pre_append = pre + ".0." + "stream"
	if _, ok := d.GetOk(pre_append); ok {
		result["stream"], _ = expandLlmProfileChatStream(d, i["stream"], pre_append, sv)
	}
	pre_append = pre + ".0." + "max_req_len"
	if _, ok := d.GetOk(pre_append); ok {
		result["max-req-len"], _ = expandLlmProfileChatMaxReqLen(d, i["max_req_len"], pre_append, sv)
	}
	pre_append = pre + ".0." + "system_prompt_mode"
	if _, ok := d.GetOk(pre_append); ok {
		result["system-prompt-mode"], _ = expandLlmProfileChatSystemPromptMode(d, i["system_prompt_mode"], pre_append, sv)
	}
	pre_append = pre + ".0." + "system_prompt"
	if _, ok := d.GetOk(pre_append); ok {
		result["system-prompt"], _ = expandLlmProfileChatSystemPrompt(d, i["system_prompt"], pre_append, sv)
	} else if d.HasChange(pre_append) {
		result["system-prompt"] = nil
	}

	return result, nil
}

func expandLlmProfileChatStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatMaxCompletionTokens(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatStream(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatMaxReqLen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatSystemPromptMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileChatSystemPrompt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileImageGen(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandLlmProfileImageGenStatus(d, i["status"], pre_append, sv)
	}

	return result, nil
}

func expandLlmProfileImageGenStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLlmProfileListModels(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	i := l[0].(map[string]interface{})
	result := make(map[string]interface{})

	pre_append := "" // complex
	pre_append = pre + ".0." + "status"
	if _, ok := d.GetOk(pre_append); ok {
		result["status"], _ = expandLlmProfileListModelsStatus(d, i["status"], pre_append, sv)
	}

	return result, nil
}

func expandLlmProfileListModelsStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLlmProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandLlmProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("unknown_api"); ok {
		t, err := expandLlmProfileUnknownApi(d, v, "unknown_api", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["unknown-api"] = t
		}
	}

	if v, ok := d.GetOk("log"); ok {
		t, err := expandLlmProfileLog(d, v, "log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log"] = t
		}
	}

	if v, ok := d.GetOk("chat"); ok {
		t, err := expandLlmProfileChat(d, v, "chat", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chat"] = t
		}
	}

	if v, ok := d.GetOk("image_gen"); ok {
		t, err := expandLlmProfileImageGen(d, v, "image_gen", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["image-gen"] = t
		}
	}

	if v, ok := d.GetOk("list_models"); ok {
		t, err := expandLlmProfileListModels(d, v, "list_models", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["list-models"] = t
		}
	}

	return &obj, nil
}
