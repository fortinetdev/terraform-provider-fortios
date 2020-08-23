// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure web proxy profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceWebProxyProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceWebProxyProfileCreate,
		Read:   resourceWebProxyProfileRead,
		Update: resourceWebProxyProfileUpdate,
		Delete: resourceWebProxyProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"header_client_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_via_request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_via_response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_forwarded_for": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_front_end_https": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_authenticated_user": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"header_x_authenticated_groups": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"strip_encoding": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_header_change": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional:     true,
							Computed:     true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
						"base64_encoding": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"add_option": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"protocol": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceWebProxyProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebProxyProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWebProxyProfile(obj)

	if err != nil {
		return fmt.Errorf("Error creating WebProxyProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyProfile")
	}

	return resourceWebProxyProfileRead(d, m)
}

func resourceWebProxyProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectWebProxyProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyProfile(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("WebProxyProfile")
	}

	return resourceWebProxyProfileRead(d, m)
}

func resourceWebProxyProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteWebProxyProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting WebProxyProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceWebProxyProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadWebProxyProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyProfile resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderClientIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderViaRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderViaResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXForwardedFor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderFrontEndHttps(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXAuthenticatedUser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXAuthenticatedGroups(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileStripEncoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileLogHeaderChange(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := i["id"]; ok {
			tmp["id"] = flattenWebProxyProfileHeadersId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenWebProxyProfileHeadersName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenWebProxyProfileHeadersAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := i["content"]; ok {
			tmp["content"] = flattenWebProxyProfileHeadersContent(i["content"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := i["base64-encoding"]; ok {
			tmp["base64_encoding"] = flattenWebProxyProfileHeadersBase64Encoding(i["base64-encoding"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_option"
		if _, ok := i["add-option"]; ok {
			tmp["add_option"] = flattenWebProxyProfileHeadersAddOption(i["add-option"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := i["protocol"]; ok {
			tmp["protocol"] = flattenWebProxyProfileHeadersProtocol(i["protocol"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenWebProxyProfileHeadersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersBase64Encoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersAddOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectWebProxyProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenWebProxyProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("header_client_ip", flattenWebProxyProfileHeaderClientIp(o["header-client-ip"], d, "header_client_ip")); err != nil {
		if !fortiAPIPatch(o["header-client-ip"]) {
			return fmt.Errorf("Error reading header_client_ip: %v", err)
		}
	}

	if err = d.Set("header_via_request", flattenWebProxyProfileHeaderViaRequest(o["header-via-request"], d, "header_via_request")); err != nil {
		if !fortiAPIPatch(o["header-via-request"]) {
			return fmt.Errorf("Error reading header_via_request: %v", err)
		}
	}

	if err = d.Set("header_via_response", flattenWebProxyProfileHeaderViaResponse(o["header-via-response"], d, "header_via_response")); err != nil {
		if !fortiAPIPatch(o["header-via-response"]) {
			return fmt.Errorf("Error reading header_via_response: %v", err)
		}
	}

	if err = d.Set("header_x_forwarded_for", flattenWebProxyProfileHeaderXForwardedFor(o["header-x-forwarded-for"], d, "header_x_forwarded_for")); err != nil {
		if !fortiAPIPatch(o["header-x-forwarded-for"]) {
			return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
		}
	}

	if err = d.Set("header_front_end_https", flattenWebProxyProfileHeaderFrontEndHttps(o["header-front-end-https"], d, "header_front_end_https")); err != nil {
		if !fortiAPIPatch(o["header-front-end-https"]) {
			return fmt.Errorf("Error reading header_front_end_https: %v", err)
		}
	}

	if err = d.Set("header_x_authenticated_user", flattenWebProxyProfileHeaderXAuthenticatedUser(o["header-x-authenticated-user"], d, "header_x_authenticated_user")); err != nil {
		if !fortiAPIPatch(o["header-x-authenticated-user"]) {
			return fmt.Errorf("Error reading header_x_authenticated_user: %v", err)
		}
	}

	if err = d.Set("header_x_authenticated_groups", flattenWebProxyProfileHeaderXAuthenticatedGroups(o["header-x-authenticated-groups"], d, "header_x_authenticated_groups")); err != nil {
		if !fortiAPIPatch(o["header-x-authenticated-groups"]) {
			return fmt.Errorf("Error reading header_x_authenticated_groups: %v", err)
		}
	}

	if err = d.Set("strip_encoding", flattenWebProxyProfileStripEncoding(o["strip-encoding"], d, "strip_encoding")); err != nil {
		if !fortiAPIPatch(o["strip-encoding"]) {
			return fmt.Errorf("Error reading strip_encoding: %v", err)
		}
	}

	if err = d.Set("log_header_change", flattenWebProxyProfileLogHeaderChange(o["log-header-change"], d, "log_header_change")); err != nil {
		if !fortiAPIPatch(o["log-header-change"]) {
			return fmt.Errorf("Error reading log_header_change: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("headers", flattenWebProxyProfileHeaders(o["headers"], d, "headers")); err != nil {
			if !fortiAPIPatch(o["headers"]) {
				return fmt.Errorf("Error reading headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("headers"); ok {
			if err = d.Set("headers", flattenWebProxyProfileHeaders(o["headers"], d, "headers")); err != nil {
				if !fortiAPIPatch(o["headers"]) {
					return fmt.Errorf("Error reading headers: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenWebProxyProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandWebProxyProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderClientIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderViaRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderViaResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXForwardedFor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderFrontEndHttps(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXAuthenticatedUser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXAuthenticatedGroups(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileStripEncoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileLogHeaderChange(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandWebProxyProfileHeadersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandWebProxyProfileHeadersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWebProxyProfileHeadersAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content"], _ = expandWebProxyProfileHeadersContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["base64-encoding"], _ = expandWebProxyProfileHeadersBase64Encoding(d, i["base64_encoding"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_option"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["add-option"], _ = expandWebProxyProfileHeadersAddOption(d, i["add_option"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandWebProxyProfileHeadersProtocol(d, i["protocol"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyProfileHeadersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersBase64Encoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersAddOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebProxyProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("header_client_ip"); ok {
		t, err := expandWebProxyProfileHeaderClientIp(d, v, "header_client_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-client-ip"] = t
		}
	}

	if v, ok := d.GetOk("header_via_request"); ok {
		t, err := expandWebProxyProfileHeaderViaRequest(d, v, "header_via_request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-via-request"] = t
		}
	}

	if v, ok := d.GetOk("header_via_response"); ok {
		t, err := expandWebProxyProfileHeaderViaResponse(d, v, "header_via_response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-via-response"] = t
		}
	}

	if v, ok := d.GetOk("header_x_forwarded_for"); ok {
		t, err := expandWebProxyProfileHeaderXForwardedFor(d, v, "header_x_forwarded_for")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-forwarded-for"] = t
		}
	}

	if v, ok := d.GetOk("header_front_end_https"); ok {
		t, err := expandWebProxyProfileHeaderFrontEndHttps(d, v, "header_front_end_https")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-front-end-https"] = t
		}
	}

	if v, ok := d.GetOk("header_x_authenticated_user"); ok {
		t, err := expandWebProxyProfileHeaderXAuthenticatedUser(d, v, "header_x_authenticated_user")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-authenticated-user"] = t
		}
	}

	if v, ok := d.GetOk("header_x_authenticated_groups"); ok {
		t, err := expandWebProxyProfileHeaderXAuthenticatedGroups(d, v, "header_x_authenticated_groups")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-authenticated-groups"] = t
		}
	}

	if v, ok := d.GetOk("strip_encoding"); ok {
		t, err := expandWebProxyProfileStripEncoding(d, v, "strip_encoding")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-encoding"] = t
		}
	}

	if v, ok := d.GetOk("log_header_change"); ok {
		t, err := expandWebProxyProfileLogHeaderChange(d, v, "log_header_change")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-header-change"] = t
		}
	}

	if v, ok := d.GetOk("headers"); ok {
		t, err := expandWebProxyProfileHeaders(d, v, "headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["headers"] = t
		}
	}

	return &obj, nil
}
