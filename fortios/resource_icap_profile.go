// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure ICAP profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceIcapProfile() *schema.Resource {
	return &schema.Resource{
		Create: resourceIcapProfileCreate,
		Read:   resourceIcapProfileRead,
		Update: resourceIcapProfileUpdate,
		Delete: resourceIcapProfileDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"request": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"streaming_content_bypass": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preview": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"preview_data_length": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 4096),
				Optional:     true,
				Computed:     true,
			},
			"request_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"response_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"request_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_failure": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"response_path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"methods": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_req_hdr": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"respmod_default_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_block_log": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"chunk_encap": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"extension_feature": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"scan_progress_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(5, 30),
				Optional:     true,
				Computed:     true,
			},
			"icap_headers": &schema.Schema{
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
					},
				},
			},
			"respmod_forward_rules": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
							Computed:     true,
						},
						"host": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional:     true,
							Computed:     true,
						},
						"header_group": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"id": &schema.Schema{
										Type:     schema.TypeInt,
										Optional: true,
										Computed: true,
									},
									"header_name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 79),
										Optional:     true,
										Computed:     true,
									},
									"header": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 255),
										Optional:     true,
										Computed:     true,
									},
									"case_sensitivity": &schema.Schema{
										Type:     schema.TypeString,
										Optional: true,
										Computed: true,
									},
								},
							},
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"http_resp_status_code": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"code": &schema.Schema{
										Type:         schema.TypeInt,
										ValidateFunc: validation.IntBetween(100, 599),
										Optional:     true,
										Computed:     true,
									},
								},
							},
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

func resourceIcapProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIcapProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating IcapProfile resource while getting object: %v", err)
	}

	o, err := c.CreateIcapProfile(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating IcapProfile resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IcapProfile")
	}

	return resourceIcapProfileRead(d, m)
}

func resourceIcapProfileUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectIcapProfile(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating IcapProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateIcapProfile(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating IcapProfile resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("IcapProfile")
	}

	return resourceIcapProfileRead(d, m)
}

func resourceIcapProfileDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteIcapProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting IcapProfile resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceIcapProfileRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadIcapProfile(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading IcapProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapProfile(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading IcapProfile resource from API: %v", err)
	}
	return nil
}

func flattenIcapProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRequest(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileResponse(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfilePreview(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfilePreviewDataLength(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRequestServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileResponseServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRequestFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileResponseFailure(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRequestPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileResponsePath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileMethods(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileResponseReqHdr(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRespmodDefaultAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileIcapBlockLog(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileChunkEncap(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileExtensionFeature(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileScanProgressInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileIcapHeaders(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenIcapProfileIcapHeadersId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {

			tmp["name"] = flattenIcapProfileIcapHeadersName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := i["content"]; ok {

			tmp["content"] = flattenIcapProfileIcapHeadersContent(i["content"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := i["base64-encoding"]; ok {

			tmp["base64_encoding"] = flattenIcapProfileIcapHeadersBase64Encoding(i["base64-encoding"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenIcapProfileIcapHeadersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersContent(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersBase64Encoding(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRules(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenIcapProfileRespmodForwardRulesName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := i["host"]; ok {

			tmp["host"] = flattenIcapProfileRespmodForwardRulesHost(i["host"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_group"
		if _, ok := i["header-group"]; ok {

			tmp["header_group"] = flattenIcapProfileRespmodForwardRulesHeaderGroup(i["header-group"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {

			tmp["action"] = flattenIcapProfileRespmodForwardRulesAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_resp_status_code"
		if _, ok := i["http-resp-status-code"]; ok {

			tmp["http_resp_status_code"] = flattenIcapProfileRespmodForwardRulesHttpRespStatusCode(i["http-resp-status-code"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenIcapProfileRespmodForwardRulesName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHost(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHeaderGroup(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
		if _, ok := i["id"]; ok {

			tmp["id"] = flattenIcapProfileRespmodForwardRulesHeaderGroupId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := i["header-name"]; ok {

			tmp["header_name"] = flattenIcapProfileRespmodForwardRulesHeaderGroupHeaderName(i["header-name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {

			tmp["header"] = flattenIcapProfileRespmodForwardRulesHeaderGroupHeader(i["header"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := i["case-sensitivity"]; ok {

			tmp["case_sensitivity"] = flattenIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(i["case-sensitivity"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenIcapProfileRespmodForwardRulesHeaderGroupId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHeaderGroupHeaderName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHeaderGroupHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenIcapProfileRespmodForwardRulesHttpRespStatusCode(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := i["code"]; ok {

			tmp["code"] = flattenIcapProfileRespmodForwardRulesHttpRespStatusCodeCode(i["code"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenIcapProfileRespmodForwardRulesHttpRespStatusCodeCode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectIcapProfile(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("replacemsg_group", flattenIcapProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("name", flattenIcapProfileName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("request", flattenIcapProfileRequest(o["request"], d, "request", sv)); err != nil {
		if !fortiAPIPatch(o["request"]) {
			return fmt.Errorf("Error reading request: %v", err)
		}
	}

	if err = d.Set("response", flattenIcapProfileResponse(o["response"], d, "response", sv)); err != nil {
		if !fortiAPIPatch(o["response"]) {
			return fmt.Errorf("Error reading response: %v", err)
		}
	}

	if err = d.Set("streaming_content_bypass", flattenIcapProfileStreamingContentBypass(o["streaming-content-bypass"], d, "streaming_content_bypass", sv)); err != nil {
		if !fortiAPIPatch(o["streaming-content-bypass"]) {
			return fmt.Errorf("Error reading streaming_content_bypass: %v", err)
		}
	}

	if err = d.Set("preview", flattenIcapProfilePreview(o["preview"], d, "preview", sv)); err != nil {
		if !fortiAPIPatch(o["preview"]) {
			return fmt.Errorf("Error reading preview: %v", err)
		}
	}

	if err = d.Set("preview_data_length", flattenIcapProfilePreviewDataLength(o["preview-data-length"], d, "preview_data_length", sv)); err != nil {
		if !fortiAPIPatch(o["preview-data-length"]) {
			return fmt.Errorf("Error reading preview_data_length: %v", err)
		}
	}

	if err = d.Set("request_server", flattenIcapProfileRequestServer(o["request-server"], d, "request_server", sv)); err != nil {
		if !fortiAPIPatch(o["request-server"]) {
			return fmt.Errorf("Error reading request_server: %v", err)
		}
	}

	if err = d.Set("response_server", flattenIcapProfileResponseServer(o["response-server"], d, "response_server", sv)); err != nil {
		if !fortiAPIPatch(o["response-server"]) {
			return fmt.Errorf("Error reading response_server: %v", err)
		}
	}

	if err = d.Set("request_failure", flattenIcapProfileRequestFailure(o["request-failure"], d, "request_failure", sv)); err != nil {
		if !fortiAPIPatch(o["request-failure"]) {
			return fmt.Errorf("Error reading request_failure: %v", err)
		}
	}

	if err = d.Set("response_failure", flattenIcapProfileResponseFailure(o["response-failure"], d, "response_failure", sv)); err != nil {
		if !fortiAPIPatch(o["response-failure"]) {
			return fmt.Errorf("Error reading response_failure: %v", err)
		}
	}

	if err = d.Set("request_path", flattenIcapProfileRequestPath(o["request-path"], d, "request_path", sv)); err != nil {
		if !fortiAPIPatch(o["request-path"]) {
			return fmt.Errorf("Error reading request_path: %v", err)
		}
	}

	if err = d.Set("response_path", flattenIcapProfileResponsePath(o["response-path"], d, "response_path", sv)); err != nil {
		if !fortiAPIPatch(o["response-path"]) {
			return fmt.Errorf("Error reading response_path: %v", err)
		}
	}

	if err = d.Set("methods", flattenIcapProfileMethods(o["methods"], d, "methods", sv)); err != nil {
		if !fortiAPIPatch(o["methods"]) {
			return fmt.Errorf("Error reading methods: %v", err)
		}
	}

	if err = d.Set("response_req_hdr", flattenIcapProfileResponseReqHdr(o["response-req-hdr"], d, "response_req_hdr", sv)); err != nil {
		if !fortiAPIPatch(o["response-req-hdr"]) {
			return fmt.Errorf("Error reading response_req_hdr: %v", err)
		}
	}

	if err = d.Set("respmod_default_action", flattenIcapProfileRespmodDefaultAction(o["respmod-default-action"], d, "respmod_default_action", sv)); err != nil {
		if !fortiAPIPatch(o["respmod-default-action"]) {
			return fmt.Errorf("Error reading respmod_default_action: %v", err)
		}
	}

	if err = d.Set("icap_block_log", flattenIcapProfileIcapBlockLog(o["icap-block-log"], d, "icap_block_log", sv)); err != nil {
		if !fortiAPIPatch(o["icap-block-log"]) {
			return fmt.Errorf("Error reading icap_block_log: %v", err)
		}
	}

	if err = d.Set("chunk_encap", flattenIcapProfileChunkEncap(o["chunk-encap"], d, "chunk_encap", sv)); err != nil {
		if !fortiAPIPatch(o["chunk-encap"]) {
			return fmt.Errorf("Error reading chunk_encap: %v", err)
		}
	}

	if err = d.Set("extension_feature", flattenIcapProfileExtensionFeature(o["extension-feature"], d, "extension_feature", sv)); err != nil {
		if !fortiAPIPatch(o["extension-feature"]) {
			return fmt.Errorf("Error reading extension_feature: %v", err)
		}
	}

	if err = d.Set("scan_progress_interval", flattenIcapProfileScanProgressInterval(o["scan-progress-interval"], d, "scan_progress_interval", sv)); err != nil {
		if !fortiAPIPatch(o["scan-progress-interval"]) {
			return fmt.Errorf("Error reading scan_progress_interval: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("icap_headers", flattenIcapProfileIcapHeaders(o["icap-headers"], d, "icap_headers", sv)); err != nil {
			if !fortiAPIPatch(o["icap-headers"]) {
				return fmt.Errorf("Error reading icap_headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("icap_headers"); ok {
			if err = d.Set("icap_headers", flattenIcapProfileIcapHeaders(o["icap-headers"], d, "icap_headers", sv)); err != nil {
				if !fortiAPIPatch(o["icap-headers"]) {
					return fmt.Errorf("Error reading icap_headers: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("respmod_forward_rules", flattenIcapProfileRespmodForwardRules(o["respmod-forward-rules"], d, "respmod_forward_rules", sv)); err != nil {
			if !fortiAPIPatch(o["respmod-forward-rules"]) {
				return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("respmod_forward_rules"); ok {
			if err = d.Set("respmod_forward_rules", flattenIcapProfileRespmodForwardRules(o["respmod-forward-rules"], d, "respmod_forward_rules", sv)); err != nil {
				if !fortiAPIPatch(o["respmod-forward-rules"]) {
					return fmt.Errorf("Error reading respmod_forward_rules: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenIcapProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandIcapProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequest(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponse(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileStreamingContentBypass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfilePreview(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfilePreviewDataLength(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequestServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponseServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequestFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponseFailure(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequestPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponsePath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileMethods(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponseReqHdr(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodDefaultAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapBlockLog(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileChunkEncap(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileExtensionFeature(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileScanProgressInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeaders(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandIcapProfileIcapHeadersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["name"], _ = expandIcapProfileIcapHeadersName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["content"], _ = expandIcapProfileIcapHeadersContent(d, i["content"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["base64-encoding"], _ = expandIcapProfileIcapHeadersBase64Encoding(d, i["base64_encoding"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIcapProfileIcapHeadersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersContent(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersBase64Encoding(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRules(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandIcapProfileRespmodForwardRulesName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "host"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["host"], _ = expandIcapProfileRespmodForwardRulesHost(d, i["host"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_group"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["header-group"], _ = expandIcapProfileRespmodForwardRulesHeaderGroup(d, i["header_group"], pre_append, sv)
		} else {
			tmp["header-group"] = make([]string, 0)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["action"], _ = expandIcapProfileRespmodForwardRulesAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "http_resp_status_code"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["http-resp-status-code"], _ = expandIcapProfileRespmodForwardRulesHttpRespStatusCode(d, i["http_resp_status_code"], pre_append, sv)
		} else {
			tmp["http-resp-status-code"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIcapProfileRespmodForwardRulesName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHost(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["id"], _ = expandIcapProfileRespmodForwardRulesHeaderGroupId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header_name"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["header-name"], _ = expandIcapProfileRespmodForwardRulesHeaderGroupHeaderName(d, i["header_name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["header"], _ = expandIcapProfileRespmodForwardRulesHeaderGroupHeader(d, i["header"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "case_sensitivity"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["case-sensitivity"], _ = expandIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(d, i["case_sensitivity"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroupId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroupHeaderName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroupHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHeaderGroupCaseSensitivity(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRespmodForwardRulesHttpRespStatusCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "code"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["code"], _ = expandIcapProfileRespmodForwardRulesHttpRespStatusCodeCode(d, i["code"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIcapProfileRespmodForwardRulesHttpRespStatusCodeCode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectIcapProfile(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("replacemsg_group"); ok {

		t, err := expandIcapProfileReplacemsgGroup(d, v, "replacemsg_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {

		t, err := expandIcapProfileName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("request"); ok {

		t, err := expandIcapProfileRequest(d, v, "request", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request"] = t
		}
	}

	if v, ok := d.GetOk("response"); ok {

		t, err := expandIcapProfileResponse(d, v, "response", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response"] = t
		}
	}

	if v, ok := d.GetOk("streaming_content_bypass"); ok {

		t, err := expandIcapProfileStreamingContentBypass(d, v, "streaming_content_bypass", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["streaming-content-bypass"] = t
		}
	}

	if v, ok := d.GetOk("preview"); ok {

		t, err := expandIcapProfilePreview(d, v, "preview", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preview"] = t
		}
	}

	if v, ok := d.GetOkExists("preview_data_length"); ok {

		t, err := expandIcapProfilePreviewDataLength(d, v, "preview_data_length", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preview-data-length"] = t
		}
	}

	if v, ok := d.GetOk("request_server"); ok {

		t, err := expandIcapProfileRequestServer(d, v, "request_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-server"] = t
		}
	}

	if v, ok := d.GetOk("response_server"); ok {

		t, err := expandIcapProfileResponseServer(d, v, "response_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-server"] = t
		}
	}

	if v, ok := d.GetOk("request_failure"); ok {

		t, err := expandIcapProfileRequestFailure(d, v, "request_failure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-failure"] = t
		}
	}

	if v, ok := d.GetOk("response_failure"); ok {

		t, err := expandIcapProfileResponseFailure(d, v, "response_failure", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-failure"] = t
		}
	}

	if v, ok := d.GetOk("request_path"); ok {

		t, err := expandIcapProfileRequestPath(d, v, "request_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-path"] = t
		}
	}

	if v, ok := d.GetOk("response_path"); ok {

		t, err := expandIcapProfileResponsePath(d, v, "response_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-path"] = t
		}
	}

	if v, ok := d.GetOk("methods"); ok {

		t, err := expandIcapProfileMethods(d, v, "methods", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["methods"] = t
		}
	}

	if v, ok := d.GetOk("response_req_hdr"); ok {

		t, err := expandIcapProfileResponseReqHdr(d, v, "response_req_hdr", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-req-hdr"] = t
		}
	}

	if v, ok := d.GetOk("respmod_default_action"); ok {

		t, err := expandIcapProfileRespmodDefaultAction(d, v, "respmod_default_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["respmod-default-action"] = t
		}
	}

	if v, ok := d.GetOk("icap_block_log"); ok {

		t, err := expandIcapProfileIcapBlockLog(d, v, "icap_block_log", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-block-log"] = t
		}
	}

	if v, ok := d.GetOk("chunk_encap"); ok {

		t, err := expandIcapProfileChunkEncap(d, v, "chunk_encap", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["chunk-encap"] = t
		}
	}

	if v, ok := d.GetOk("extension_feature"); ok {

		t, err := expandIcapProfileExtensionFeature(d, v, "extension_feature", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["extension-feature"] = t
		}
	}

	if v, ok := d.GetOk("scan_progress_interval"); ok {

		t, err := expandIcapProfileScanProgressInterval(d, v, "scan_progress_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["scan-progress-interval"] = t
		}
	}

	if v, ok := d.GetOk("icap_headers"); ok {

		t, err := expandIcapProfileIcapHeaders(d, v, "icap_headers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-headers"] = t
		}
	}

	if v, ok := d.GetOk("respmod_forward_rules"); ok {

		t, err := expandIcapProfileRespmodForwardRules(d, v, "respmod_forward_rules", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["respmod-forward-rules"] = t
		}
	}

	return &obj, nil
}
