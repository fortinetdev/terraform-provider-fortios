// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure ICAP profiles.

package fortios

import (
	"fmt"
	"log"
	"strconv"

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
			"replacemsg_group": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"request": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"streaming_content_bypass": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_server": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"response_server": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional: true,
				Computed: true,
			},
			"request_failure": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_failure": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"request_path": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional: true,
				Computed: true,
			},
			"response_path": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional: true,
				Computed: true,
			},
			"methods": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"response_req_hdr": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"icap_headers": &schema.Schema{
				Type: schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type: schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 4294967295),
							Optional: true,
							Computed: true,
						},
						"name": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 79),
							Optional: true,
							Computed: true,
						},
						"content": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional: true,
							Computed: true,
						},
						"base64_encoding": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func resourceIcapProfileCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectIcapProfile(d)
	if err != nil {
		return fmt.Errorf("Error creating IcapProfile resource while getting object: %v", err)
	}

	o, err := c.CreateIcapProfile(obj)

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

	obj, err := getObjectIcapProfile(d)
	if err != nil {
		return fmt.Errorf("Error updating IcapProfile resource while getting object: %v", err)
	}

	o, err := c.UpdateIcapProfile(obj, mkey)
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

	err := c.DeleteIcapProfile(mkey)
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

	o, err := c.ReadIcapProfile(mkey)
	if err != nil {
		return fmt.Errorf("Error reading IcapProfile resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectIcapProfile(d, o)
	if err != nil {
		return fmt.Errorf("Error reading IcapProfile resource from API: %v", err)
	}
	return nil
}


func flattenIcapProfileReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRequest(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponse(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileStreamingContentBypass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRequestServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponseServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRequestFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponseFailure(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileRequestPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponsePath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileMethods(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileResponseReqHdr(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenIcapProfileIcapHeadersId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenIcapProfileIcapHeadersName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := i["content"]; ok {
			tmp["content"] = flattenIcapProfileIcapHeadersContent(i["content"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := i["base64-encoding"]; ok {
			tmp["base64_encoding"] = flattenIcapProfileIcapHeadersBase64Encoding(i["base64-encoding"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenIcapProfileIcapHeadersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersContent(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenIcapProfileIcapHeadersBase64Encoding(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectIcapProfile(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("replacemsg_group", flattenIcapProfileReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("name", flattenIcapProfileName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("request", flattenIcapProfileRequest(o["request"], d, "request")); err != nil {
		if !fortiAPIPatch(o["request"]) {
			return fmt.Errorf("Error reading request: %v", err)
		}
	}

	if err = d.Set("response", flattenIcapProfileResponse(o["response"], d, "response")); err != nil {
		if !fortiAPIPatch(o["response"]) {
			return fmt.Errorf("Error reading response: %v", err)
		}
	}

	if err = d.Set("streaming_content_bypass", flattenIcapProfileStreamingContentBypass(o["streaming-content-bypass"], d, "streaming_content_bypass")); err != nil {
		if !fortiAPIPatch(o["streaming-content-bypass"]) {
			return fmt.Errorf("Error reading streaming_content_bypass: %v", err)
		}
	}

	if err = d.Set("request_server", flattenIcapProfileRequestServer(o["request-server"], d, "request_server")); err != nil {
		if !fortiAPIPatch(o["request-server"]) {
			return fmt.Errorf("Error reading request_server: %v", err)
		}
	}

	if err = d.Set("response_server", flattenIcapProfileResponseServer(o["response-server"], d, "response_server")); err != nil {
		if !fortiAPIPatch(o["response-server"]) {
			return fmt.Errorf("Error reading response_server: %v", err)
		}
	}

	if err = d.Set("request_failure", flattenIcapProfileRequestFailure(o["request-failure"], d, "request_failure")); err != nil {
		if !fortiAPIPatch(o["request-failure"]) {
			return fmt.Errorf("Error reading request_failure: %v", err)
		}
	}

	if err = d.Set("response_failure", flattenIcapProfileResponseFailure(o["response-failure"], d, "response_failure")); err != nil {
		if !fortiAPIPatch(o["response-failure"]) {
			return fmt.Errorf("Error reading response_failure: %v", err)
		}
	}

	if err = d.Set("request_path", flattenIcapProfileRequestPath(o["request-path"], d, "request_path")); err != nil {
		if !fortiAPIPatch(o["request-path"]) {
			return fmt.Errorf("Error reading request_path: %v", err)
		}
	}

	if err = d.Set("response_path", flattenIcapProfileResponsePath(o["response-path"], d, "response_path")); err != nil {
		if !fortiAPIPatch(o["response-path"]) {
			return fmt.Errorf("Error reading response_path: %v", err)
		}
	}

	if err = d.Set("methods", flattenIcapProfileMethods(o["methods"], d, "methods")); err != nil {
		if !fortiAPIPatch(o["methods"]) {
			return fmt.Errorf("Error reading methods: %v", err)
		}
	}

	if err = d.Set("response_req_hdr", flattenIcapProfileResponseReqHdr(o["response-req-hdr"], d, "response_req_hdr")); err != nil {
		if !fortiAPIPatch(o["response-req-hdr"]) {
			return fmt.Errorf("Error reading response_req_hdr: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("icap_headers", flattenIcapProfileIcapHeaders(o["icap-headers"], d, "icap_headers")); err != nil {
            if !fortiAPIPatch(o["icap-headers"]) {
                return fmt.Errorf("Error reading icap_headers: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("icap_headers"); ok {
            if err = d.Set("icap_headers", flattenIcapProfileIcapHeaders(o["icap-headers"], d, "icap_headers")); err != nil {
                if !fortiAPIPatch(o["icap-headers"]) {
                    return fmt.Errorf("Error reading icap_headers: %v", err)
                }
            }
        }
    }


	return nil
}

func flattenIcapProfileFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandIcapProfileReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequest(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponse(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileStreamingContentBypass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequestServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponseServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequestFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponseFailure(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileRequestPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponsePath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileMethods(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileResponseReqHdr(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}

	result := make([]map[string]interface{}, 0, len(l))

	con := 0
	for _, r := range l {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		pre_append := ""  // table

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["id"], _ = expandIcapProfileIcapHeadersId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandIcapProfileIcapHeadersName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "content"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["content"], _ = expandIcapProfileIcapHeadersContent(d, i["content"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "base64_encoding"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["base64-encoding"], _ = expandIcapProfileIcapHeadersBase64Encoding(d, i["base64_encoding"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandIcapProfileIcapHeadersId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersContent(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandIcapProfileIcapHeadersBase64Encoding(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectIcapProfile(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandIcapProfileReplacemsgGroup(d, v, "replacemsg_group")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	}

	if v, ok := d.GetOk("name"); ok {
		t, err := expandIcapProfileName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("request"); ok {
		t, err := expandIcapProfileRequest(d, v, "request")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request"] = t
		}
	}

	if v, ok := d.GetOk("response"); ok {
		t, err := expandIcapProfileResponse(d, v, "response")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response"] = t
		}
	}

	if v, ok := d.GetOk("streaming_content_bypass"); ok {
		t, err := expandIcapProfileStreamingContentBypass(d, v, "streaming_content_bypass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["streaming-content-bypass"] = t
		}
	}

	if v, ok := d.GetOk("request_server"); ok {
		t, err := expandIcapProfileRequestServer(d, v, "request_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-server"] = t
		}
	}

	if v, ok := d.GetOk("response_server"); ok {
		t, err := expandIcapProfileResponseServer(d, v, "response_server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-server"] = t
		}
	}

	if v, ok := d.GetOk("request_failure"); ok {
		t, err := expandIcapProfileRequestFailure(d, v, "request_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-failure"] = t
		}
	}

	if v, ok := d.GetOk("response_failure"); ok {
		t, err := expandIcapProfileResponseFailure(d, v, "response_failure")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-failure"] = t
		}
	}

	if v, ok := d.GetOk("request_path"); ok {
		t, err := expandIcapProfileRequestPath(d, v, "request_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["request-path"] = t
		}
	}

	if v, ok := d.GetOk("response_path"); ok {
		t, err := expandIcapProfileResponsePath(d, v, "response_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-path"] = t
		}
	}

	if v, ok := d.GetOk("methods"); ok {
		t, err := expandIcapProfileMethods(d, v, "methods")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["methods"] = t
		}
	}

	if v, ok := d.GetOk("response_req_hdr"); ok {
		t, err := expandIcapProfileResponseReqHdr(d, v, "response_req_hdr")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["response-req-hdr"] = t
		}
	}

	if v, ok := d.GetOk("icap_headers"); ok {
		t, err := expandIcapProfileIcapHeaders(d, v, "icap_headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["icap-headers"] = t
		}
	}


	return &obj, nil
}

