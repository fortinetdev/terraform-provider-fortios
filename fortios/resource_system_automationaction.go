// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Action for automation stitches.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemAutomationAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemAutomationActionCreate,
		Read:   resourceSystemAutomationActionRead,
		Update: resourceSystemAutomationActionUpdate,
		Delete: resourceSystemAutomationActionDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Required:     true,
				ForceNew:     true,
			},
			"action_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email_to": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"email_subject": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
			},
			"minimum_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2592000),
				Optional:     true,
				Computed:     true,
			},
			"delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
				Computed:     true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"aws_api_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"aws_region": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"aws_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"aws_api_stage": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"aws_api_path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"aws_api_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Sensitive:    true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uri": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"http_body": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"security_tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
		},
	}
}

func resourceSystemAutomationActionCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutomationAction(d)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationAction resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAutomationAction(obj)

	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationAction resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationAction")
	}

	return resourceSystemAutomationActionRead(d, m)
}

func resourceSystemAutomationActionUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemAutomationAction(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationAction resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutomationAction(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationAction resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemAutomationAction")
	}

	return resourceSystemAutomationActionRead(d, m)
}

func resourceSystemAutomationActionDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemAutomationAction(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemAutomationAction resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemAutomationActionRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemAutomationAction(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationAction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationAction(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationAction resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationActionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionActionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailTo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemAutomationActionEmailToName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAutomationActionEmailToName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailSubject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionMinimumInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiStage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionHttpBody(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := i["header"]; ok {
			tmp["header"] = flattenSystemAutomationActionHeadersHeader(i["header"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAutomationActionHeadersHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionSecurityTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemAutomationActionSdnConnector(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := i["name"]; ok {
			tmp["name"] = flattenSystemAutomationActionSdnConnectorName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemAutomationActionSdnConnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemAutomationAction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", flattenSystemAutomationActionName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("action_type", flattenSystemAutomationActionActionType(o["action-type"], d, "action_type")); err != nil {
		if !fortiAPIPatch(o["action-type"]) {
			return fmt.Errorf("Error reading action_type: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("email_to", flattenSystemAutomationActionEmailTo(o["email-to"], d, "email_to")); err != nil {
			if !fortiAPIPatch(o["email-to"]) {
				return fmt.Errorf("Error reading email_to: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("email_to"); ok {
			if err = d.Set("email_to", flattenSystemAutomationActionEmailTo(o["email-to"], d, "email_to")); err != nil {
				if !fortiAPIPatch(o["email-to"]) {
					return fmt.Errorf("Error reading email_to: %v", err)
				}
			}
		}
	}

	if err = d.Set("email_subject", flattenSystemAutomationActionEmailSubject(o["email-subject"], d, "email_subject")); err != nil {
		if !fortiAPIPatch(o["email-subject"]) {
			return fmt.Errorf("Error reading email_subject: %v", err)
		}
	}

	if err = d.Set("minimum_interval", flattenSystemAutomationActionMinimumInterval(o["minimum-interval"], d, "minimum_interval")); err != nil {
		if !fortiAPIPatch(o["minimum-interval"]) {
			return fmt.Errorf("Error reading minimum_interval: %v", err)
		}
	}

	if err = d.Set("delay", flattenSystemAutomationActionDelay(o["delay"], d, "delay")); err != nil {
		if !fortiAPIPatch(o["delay"]) {
			return fmt.Errorf("Error reading delay: %v", err)
		}
	}

	if err = d.Set("required", flattenSystemAutomationActionRequired(o["required"], d, "required")); err != nil {
		if !fortiAPIPatch(o["required"]) {
			return fmt.Errorf("Error reading required: %v", err)
		}
	}

	if err = d.Set("aws_api_id", flattenSystemAutomationActionAwsApiId(o["aws-api-id"], d, "aws_api_id")); err != nil {
		if !fortiAPIPatch(o["aws-api-id"]) {
			return fmt.Errorf("Error reading aws_api_id: %v", err)
		}
	}

	if err = d.Set("aws_region", flattenSystemAutomationActionAwsRegion(o["aws-region"], d, "aws_region")); err != nil {
		if !fortiAPIPatch(o["aws-region"]) {
			return fmt.Errorf("Error reading aws_region: %v", err)
		}
	}

	if err = d.Set("aws_domain", flattenSystemAutomationActionAwsDomain(o["aws-domain"], d, "aws_domain")); err != nil {
		if !fortiAPIPatch(o["aws-domain"]) {
			return fmt.Errorf("Error reading aws_domain: %v", err)
		}
	}

	if err = d.Set("aws_api_stage", flattenSystemAutomationActionAwsApiStage(o["aws-api-stage"], d, "aws_api_stage")); err != nil {
		if !fortiAPIPatch(o["aws-api-stage"]) {
			return fmt.Errorf("Error reading aws_api_stage: %v", err)
		}
	}

	if err = d.Set("aws_api_path", flattenSystemAutomationActionAwsApiPath(o["aws-api-path"], d, "aws_api_path")); err != nil {
		if !fortiAPIPatch(o["aws-api-path"]) {
			return fmt.Errorf("Error reading aws_api_path: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemAutomationActionProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("method", flattenSystemAutomationActionMethod(o["method"], d, "method")); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("uri", flattenSystemAutomationActionUri(o["uri"], d, "uri")); err != nil {
		if !fortiAPIPatch(o["uri"]) {
			return fmt.Errorf("Error reading uri: %v", err)
		}
	}

	if err = d.Set("http_body", flattenSystemAutomationActionHttpBody(o["http-body"], d, "http_body")); err != nil {
		if !fortiAPIPatch(o["http-body"]) {
			return fmt.Errorf("Error reading http_body: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAutomationActionPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("headers", flattenSystemAutomationActionHeaders(o["headers"], d, "headers")); err != nil {
			if !fortiAPIPatch(o["headers"]) {
				return fmt.Errorf("Error reading headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("headers"); ok {
			if err = d.Set("headers", flattenSystemAutomationActionHeaders(o["headers"], d, "headers")); err != nil {
				if !fortiAPIPatch(o["headers"]) {
					return fmt.Errorf("Error reading headers: %v", err)
				}
			}
		}
	}

	if err = d.Set("security_tag", flattenSystemAutomationActionSecurityTag(o["security-tag"], d, "security_tag")); err != nil {
		if !fortiAPIPatch(o["security-tag"]) {
			return fmt.Errorf("Error reading security_tag: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("sdn_connector", flattenSystemAutomationActionSdnConnector(o["sdn-connector"], d, "sdn_connector")); err != nil {
			if !fortiAPIPatch(o["sdn-connector"]) {
				return fmt.Errorf("Error reading sdn_connector: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sdn_connector"); ok {
			if err = d.Set("sdn_connector", flattenSystemAutomationActionSdnConnector(o["sdn-connector"], d, "sdn_connector")); err != nil {
				if !fortiAPIPatch(o["sdn-connector"]) {
					return fmt.Errorf("Error reading sdn_connector: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemAutomationActionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandSystemAutomationActionName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionActionType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailTo(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemAutomationActionEmailToName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationActionEmailToName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailSubject(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMinimumInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionDelay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionRequired(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsRegion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsDomain(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiStage(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiPath(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionProtocol(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMethod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionUri(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpBody(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHeaders(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["header"], _ = expandSystemAutomationActionHeadersHeader(d, i["header"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationActionHeadersHeader(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionSecurityTag(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionSdnConnector(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemAutomationActionSdnConnectorName(d, i["name"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationActionSdnConnectorName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationAction(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAutomationActionName(d, v, "name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("action_type"); ok {
		t, err := expandSystemAutomationActionActionType(d, v, "action_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action-type"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok {
		t, err := expandSystemAutomationActionEmailTo(d, v, "email_to")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("email_subject"); ok {
		t, err := expandSystemAutomationActionEmailSubject(d, v, "email_subject")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-subject"] = t
		}
	}

	if v, ok := d.GetOkExists("minimum_interval"); ok {
		t, err := expandSystemAutomationActionMinimumInterval(d, v, "minimum_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-interval"] = t
		}
	}

	if v, ok := d.GetOkExists("delay"); ok {
		t, err := expandSystemAutomationActionDelay(d, v, "delay")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay"] = t
		}
	}

	if v, ok := d.GetOk("required"); ok {
		t, err := expandSystemAutomationActionRequired(d, v, "required")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["required"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_id"); ok {
		t, err := expandSystemAutomationActionAwsApiId(d, v, "aws_api_id")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-id"] = t
		}
	}

	if v, ok := d.GetOk("aws_region"); ok {
		t, err := expandSystemAutomationActionAwsRegion(d, v, "aws_region")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-region"] = t
		}
	}

	if v, ok := d.GetOk("aws_domain"); ok {
		t, err := expandSystemAutomationActionAwsDomain(d, v, "aws_domain")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-domain"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_stage"); ok {
		t, err := expandSystemAutomationActionAwsApiStage(d, v, "aws_api_stage")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-stage"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_path"); ok {
		t, err := expandSystemAutomationActionAwsApiPath(d, v, "aws_api_path")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-path"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_key"); ok {
		t, err := expandSystemAutomationActionAwsApiKey(d, v, "aws_api_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-key"] = t
		}
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandSystemAutomationActionProtocol(d, v, "protocol")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandSystemAutomationActionMethod(d, v, "method")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("uri"); ok {
		t, err := expandSystemAutomationActionUri(d, v, "uri")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uri"] = t
		}
	}

	if v, ok := d.GetOk("http_body"); ok {
		t, err := expandSystemAutomationActionHttpBody(d, v, "http_body")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-body"] = t
		}
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemAutomationActionPort(d, v, "port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	}

	if v, ok := d.GetOk("headers"); ok {
		t, err := expandSystemAutomationActionHeaders(d, v, "headers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["headers"] = t
		}
	}

	if v, ok := d.GetOk("security_tag"); ok {
		t, err := expandSystemAutomationActionSecurityTag(d, v, "security_tag")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-tag"] = t
		}
	}

	if v, ok := d.GetOk("sdn_connector"); ok {
		t, err := expandSystemAutomationActionSdnConnector(d, v, "sdn_connector")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-connector"] = t
		}
	}

	return &obj, nil
}
