// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure web proxy profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"header_x_forwarded_client_cert": &schema.Schema{
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
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"dstaddr": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"dstaddr6": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
								},
							},
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"content": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),
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

func resourceWebProxyProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebProxyProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating WebProxyProfile resource while getting object: %v", err)
	}

	o, err := c.CreateWebProxyProfile(obj, vdomparam)

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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectWebProxyProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating WebProxyProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateWebProxyProfile(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteWebProxyProfile(mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadWebProxyProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectWebProxyProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading WebProxyProfile resource from API: %v", err)
	}
	return nil
}

func flattenWebProxyProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderClientIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderViaRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderViaResponse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXForwardedFor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXForwardedClientCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderFrontEndHttps(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXAuthenticatedUser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeaderXAuthenticatedGroups(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileStripEncoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileLogHeaderChange(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeaders(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if cur_v, ok := i["id"]; ok {
			tmp["id"] = flattenWebProxyProfileHeadersId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenWebProxyProfileHeadersName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if cur_v, ok := i["dstaddr"]; ok {
			tmp["dstaddr"] = flattenWebProxyProfileHeadersDstaddr(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if cur_v, ok := i["dstaddr6"]; ok {
			tmp["dstaddr6"] = flattenWebProxyProfileHeadersDstaddr6(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if cur_v, ok := i["action"]; ok {
			tmp["action"] = flattenWebProxyProfileHeadersAction(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if cur_v, ok := i["content"]; ok {
			tmp["content"] = flattenWebProxyProfileHeadersContent(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if cur_v, ok := i["base64-encoding"]; ok {
			tmp["base64_encoding"] = flattenWebProxyProfileHeadersBase64Encoding(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_option"
		if cur_v, ok := i["add-option"]; ok {
			tmp["add_option"] = flattenWebProxyProfileHeadersAddOption(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if cur_v, ok := i["protocol"]; ok {
			tmp["protocol"] = flattenWebProxyProfileHeadersProtocol(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenWebProxyProfileHeadersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersDstaddr(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenWebProxyProfileHeadersDstaddrName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebProxyProfileHeadersDstaddrName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersDstaddr6(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if cur_v, ok := i["name"]; ok {
			tmp["name"] = flattenWebProxyProfileHeadersDstaddr6Name(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenWebProxyProfileHeadersDstaddr6Name(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersBase64Encoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersAddOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenWebProxyProfileHeadersProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectWebProxyProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenWebProxyProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("header_client_ip", flattenWebProxyProfileHeaderClientIp(o["header-client-ip"], d, "header_client_ip", sv)); err != nil {
		if !fortiAPIPatch(o["header-client-ip"]) {
			return fmt.Errorf("Error reading header_client_ip: %v", err)
		}
	}

	if err = d.Set("header_via_request", flattenWebProxyProfileHeaderViaRequest(o["header-via-request"], d, "header_via_request", sv)); err != nil {
		if !fortiAPIPatch(o["header-via-request"]) {
			return fmt.Errorf("Error reading header_via_request: %v", err)
		}
	}

	if err = d.Set("header_via_response", flattenWebProxyProfileHeaderViaResponse(o["header-via-response"], d, "header_via_response", sv)); err != nil {
		if !fortiAPIPatch(o["header-via-response"]) {
			return fmt.Errorf("Error reading header_via_response: %v", err)
		}
	}

	if err = d.Set("header_x_forwarded_for", flattenWebProxyProfileHeaderXForwardedFor(o["header-x-forwarded-for"], d, "header_x_forwarded_for", sv)); err != nil {
		if !fortiAPIPatch(o["header-x-forwarded-for"]) {
			return fmt.Errorf("Error reading header_x_forwarded_for: %v", err)
		}
	}

	if err = d.Set("header_x_forwarded_client_cert", flattenWebProxyProfileHeaderXForwardedClientCert(o["header-x-forwarded-client-cert"], d, "header_x_forwarded_client_cert", sv)); err != nil {
		if !fortiAPIPatch(o["header-x-forwarded-client-cert"]) {
			return fmt.Errorf("Error reading header_x_forwarded_client_cert: %v", err)
		}
	}

	if err = d.Set("header_front_end_https", flattenWebProxyProfileHeaderFrontEndHttps(o["header-front-end-https"], d, "header_front_end_https", sv)); err != nil {
		if !fortiAPIPatch(o["header-front-end-https"]) {
			return fmt.Errorf("Error reading header_front_end_https: %v", err)
		}
	}

	if err = d.Set("header_x_authenticated_user", flattenWebProxyProfileHeaderXAuthenticatedUser(o["header-x-authenticated-user"], d, "header_x_authenticated_user", sv)); err != nil {
		if !fortiAPIPatch(o["header-x-authenticated-user"]) {
			return fmt.Errorf("Error reading header_x_authenticated_user: %v", err)
		}
	}

	if err = d.Set("header_x_authenticated_groups", flattenWebProxyProfileHeaderXAuthenticatedGroups(o["header-x-authenticated-groups"], d, "header_x_authenticated_groups", sv)); err != nil {
		if !fortiAPIPatch(o["header-x-authenticated-groups"]) {
			return fmt.Errorf("Error reading header_x_authenticated_groups: %v", err)
		}
	}

	if err = d.Set("strip_encoding", flattenWebProxyProfileStripEncoding(o["strip-encoding"], d, "strip_encoding", sv)); err != nil {
		if !fortiAPIPatch(o["strip-encoding"]) {
			return fmt.Errorf("Error reading strip_encoding: %v", err)
		}
	}

	if err = d.Set("log_header_change", flattenWebProxyProfileLogHeaderChange(o["log-header-change"], d, "log_header_change", sv)); err != nil {
		if !fortiAPIPatch(o["log-header-change"]) {
			return fmt.Errorf("Error reading log_header_change: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("headers", flattenWebProxyProfileHeaders(o["headers"], d, "headers", sv)); err != nil {
			if !fortiAPIPatch(o["headers"]) {
				return fmt.Errorf("Error reading headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("headers"); ok {
			if err = d.Set("headers", flattenWebProxyProfileHeaders(o["headers"], d, "headers", sv)); err != nil {
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
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandWebProxyProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderClientIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderViaRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderViaResponse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXForwardedFor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXForwardedClientCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderFrontEndHttps(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXAuthenticatedUser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaderXAuthenticatedGroups(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileStripEncoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileLogHeaderChange(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeaders(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	l := v.([]interface{})
	result := make([]map[string]interface{}, 0, len(l))

	if len(l) == 0 || l[0] == nil {
		return result, nil
	}

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := "" // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandWebProxyProfileHeadersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandWebProxyProfileHeadersName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr"], _ = expandWebProxyProfileHeadersDstaddr(d, i["dstaddr"], pre_append, sv)
		} else {
			tmp["dstaddr"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "dstaddr6"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["dstaddr6"], _ = expandWebProxyProfileHeadersDstaddr6(d, i["dstaddr6"], pre_append, sv)
		} else {
			tmp["dstaddr6"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandWebProxyProfileHeadersAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content"], _ = expandWebProxyProfileHeadersContent(d, i["content"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["base64-encoding"], _ = expandWebProxyProfileHeadersBase64Encoding(d, i["base64_encoding"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "add_option"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["add-option"], _ = expandWebProxyProfileHeadersAddOption(d, i["add_option"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "protocol"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["protocol"], _ = expandWebProxyProfileHeadersProtocol(d, i["protocol"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyProfileHeadersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersDstaddr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWebProxyProfileHeadersDstaddrName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyProfileHeadersDstaddrName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersDstaddr6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandWebProxyProfileHeadersDstaddr6Name(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandWebProxyProfileHeadersDstaddr6Name(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersBase64Encoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersAddOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandWebProxyProfileHeadersProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectWebProxyProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandWebProxyProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("header_client_ip"); ok {
		t, err := expandWebProxyProfileHeaderClientIp(d, v, "header_client_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-client-ip"] = t
		}
	}

	if v, ok := d.GetOk("header_via_request"); ok {
		t, err := expandWebProxyProfileHeaderViaRequest(d, v, "header_via_request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-via-request"] = t
		}
	}

	if v, ok := d.GetOk("header_via_response"); ok {
		t, err := expandWebProxyProfileHeaderViaResponse(d, v, "header_via_response", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-via-response"] = t
		}
	}

	if v, ok := d.GetOk("header_x_forwarded_for"); ok {
		t, err := expandWebProxyProfileHeaderXForwardedFor(d, v, "header_x_forwarded_for", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-forwarded-for"] = t
		}
	}

	if v, ok := d.GetOk("header_x_forwarded_client_cert"); ok {
		t, err := expandWebProxyProfileHeaderXForwardedClientCert(d, v, "header_x_forwarded_client_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-forwarded-client-cert"] = t
		}
	}

	if v, ok := d.GetOk("header_front_end_https"); ok {
		t, err := expandWebProxyProfileHeaderFrontEndHttps(d, v, "header_front_end_https", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-front-end-https"] = t
		}
	}

	if v, ok := d.GetOk("header_x_authenticated_user"); ok {
		t, err := expandWebProxyProfileHeaderXAuthenticatedUser(d, v, "header_x_authenticated_user", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-authenticated-user"] = t
		}
	}

	if v, ok := d.GetOk("header_x_authenticated_groups"); ok {
		t, err := expandWebProxyProfileHeaderXAuthenticatedGroups(d, v, "header_x_authenticated_groups", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["header-x-authenticated-groups"] = t
		}
	}

	if v, ok := d.GetOk("strip_encoding"); ok {
		t, err := expandWebProxyProfileStripEncoding(d, v, "strip_encoding", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["strip-encoding"] = t
		}
	}

	if v, ok := d.GetOk("log_header_change"); ok {
		t, err := expandWebProxyProfileLogHeaderChange(d, v, "log_header_change", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-header-change"] = t
		}
	}

	if v, ok := d.GetOk("headers"); ok || d.HasChange("headers") {
		t, err := expandWebProxyProfileHeaders(d, v, "headers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["headers"] = t
		}
	}

	return &obj, nil
}
