// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Action for automation stitches.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func dataSourceSystemAutomationAction() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemAutomationActionRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"name": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"description": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"action_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"tls_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_to": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"email_from": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_subject": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"email_body": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"minimum_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"delay": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"required": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_api_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_api_stage": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_api_path": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"aws_api_key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"azure_app": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"azure_function": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"azure_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"azure_function_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"azure_api_key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"gcp_function_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gcp_project": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gcp_function_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gcp_function": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alicloud_account_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alicloud_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alicloud_function_domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alicloud_version": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alicloud_service": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alicloud_function": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alicloud_function_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alicloud_access_key_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alicloud_access_key_secret": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"message_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"message": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"replacement_message": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"protocol": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"method": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"http_body": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"http_headers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"value": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"headers": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"verify_host_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"script": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"output_size": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"timeout": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"execute_security_fabric": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"accprofile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"security_tag": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"sdn_connector": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemAutomationActionRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := ""

	t := d.Get("name")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemAutomationAction: type error")
	}

	o, err := c.ReadSystemAutomationAction(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutomationAction: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemAutomationAction(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemAutomationAction from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemAutomationActionName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionDescription(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionActionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionTlsCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionEmailTo(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemAutomationActionEmailToName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAutomationActionEmailToName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionEmailFrom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionEmailSubject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionEmailBody(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionMinimumInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionDelay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionRequired(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAwsApiId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAwsRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAwsDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAwsApiStage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAwsApiPath(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAwsApiKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAzureApp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAzureFunction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAzureDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAzureFunctionAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAzureApiKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionGcpFunctionRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionGcpProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionGcpFunctionDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionGcpFunction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAlicloudAccountId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAlicloudRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAlicloudFunctionDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAlicloudVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAlicloudService(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAlicloudFunction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAlicloudFunctionAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAlicloudAccessKeyId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAlicloudAccessKeySecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionMessageType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionReplacementMessage(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionProtocol(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionMethod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionUri(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionHttpBody(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionHttpHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemAutomationActionHttpHeadersId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := i["key"]; ok {
			tmp["key"] = dataSourceFlattenSystemAutomationActionHttpHeadersKey(i["key"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := i["value"]; ok {
			tmp["value"] = dataSourceFlattenSystemAutomationActionHttpHeadersValue(i["value"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAutomationActionHttpHeadersId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionHttpHeadersKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionHttpHeadersValue(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionHeaders(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["header"] = dataSourceFlattenSystemAutomationActionHeadersHeader(i["header"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAutomationActionHeadersHeader(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionVerifyHostCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionScript(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionOutputSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionExecuteSecurityFabric(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionSecurityTag(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemAutomationActionSdnConnector(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemAutomationActionSdnConnectorName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemAutomationActionSdnConnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemAutomationAction(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemAutomationActionName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", dataSourceFlattenSystemAutomationActionDescription(o["description"], d, "description")); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("action_type", dataSourceFlattenSystemAutomationActionActionType(o["action-type"], d, "action_type")); err != nil {
		if !fortiAPIPatch(o["action-type"]) {
			return fmt.Errorf("Error reading action_type: %v", err)
		}
	}

	if err = d.Set("tls_certificate", dataSourceFlattenSystemAutomationActionTlsCertificate(o["tls-certificate"], d, "tls_certificate")); err != nil {
		if !fortiAPIPatch(o["tls-certificate"]) {
			return fmt.Errorf("Error reading tls_certificate: %v", err)
		}
	}

	if err = d.Set("email_to", dataSourceFlattenSystemAutomationActionEmailTo(o["email-to"], d, "email_to")); err != nil {
		if !fortiAPIPatch(o["email-to"]) {
			return fmt.Errorf("Error reading email_to: %v", err)
		}
	}

	if err = d.Set("email_from", dataSourceFlattenSystemAutomationActionEmailFrom(o["email-from"], d, "email_from")); err != nil {
		if !fortiAPIPatch(o["email-from"]) {
			return fmt.Errorf("Error reading email_from: %v", err)
		}
	}

	if err = d.Set("email_subject", dataSourceFlattenSystemAutomationActionEmailSubject(o["email-subject"], d, "email_subject")); err != nil {
		if !fortiAPIPatch(o["email-subject"]) {
			return fmt.Errorf("Error reading email_subject: %v", err)
		}
	}

	if err = d.Set("email_body", dataSourceFlattenSystemAutomationActionEmailBody(o["email-body"], d, "email_body")); err != nil {
		if !fortiAPIPatch(o["email-body"]) {
			return fmt.Errorf("Error reading email_body: %v", err)
		}
	}

	if err = d.Set("minimum_interval", dataSourceFlattenSystemAutomationActionMinimumInterval(o["minimum-interval"], d, "minimum_interval")); err != nil {
		if !fortiAPIPatch(o["minimum-interval"]) {
			return fmt.Errorf("Error reading minimum_interval: %v", err)
		}
	}

	if err = d.Set("delay", dataSourceFlattenSystemAutomationActionDelay(o["delay"], d, "delay")); err != nil {
		if !fortiAPIPatch(o["delay"]) {
			return fmt.Errorf("Error reading delay: %v", err)
		}
	}

	if err = d.Set("required", dataSourceFlattenSystemAutomationActionRequired(o["required"], d, "required")); err != nil {
		if !fortiAPIPatch(o["required"]) {
			return fmt.Errorf("Error reading required: %v", err)
		}
	}

	if err = d.Set("aws_api_id", dataSourceFlattenSystemAutomationActionAwsApiId(o["aws-api-id"], d, "aws_api_id")); err != nil {
		if !fortiAPIPatch(o["aws-api-id"]) {
			return fmt.Errorf("Error reading aws_api_id: %v", err)
		}
	}

	if err = d.Set("aws_region", dataSourceFlattenSystemAutomationActionAwsRegion(o["aws-region"], d, "aws_region")); err != nil {
		if !fortiAPIPatch(o["aws-region"]) {
			return fmt.Errorf("Error reading aws_region: %v", err)
		}
	}

	if err = d.Set("aws_domain", dataSourceFlattenSystemAutomationActionAwsDomain(o["aws-domain"], d, "aws_domain")); err != nil {
		if !fortiAPIPatch(o["aws-domain"]) {
			return fmt.Errorf("Error reading aws_domain: %v", err)
		}
	}

	if err = d.Set("aws_api_stage", dataSourceFlattenSystemAutomationActionAwsApiStage(o["aws-api-stage"], d, "aws_api_stage")); err != nil {
		if !fortiAPIPatch(o["aws-api-stage"]) {
			return fmt.Errorf("Error reading aws_api_stage: %v", err)
		}
	}

	if err = d.Set("aws_api_path", dataSourceFlattenSystemAutomationActionAwsApiPath(o["aws-api-path"], d, "aws_api_path")); err != nil {
		if !fortiAPIPatch(o["aws-api-path"]) {
			return fmt.Errorf("Error reading aws_api_path: %v", err)
		}
	}

	if err = d.Set("azure_app", dataSourceFlattenSystemAutomationActionAzureApp(o["azure-app"], d, "azure_app")); err != nil {
		if !fortiAPIPatch(o["azure-app"]) {
			return fmt.Errorf("Error reading azure_app: %v", err)
		}
	}

	if err = d.Set("azure_function", dataSourceFlattenSystemAutomationActionAzureFunction(o["azure-function"], d, "azure_function")); err != nil {
		if !fortiAPIPatch(o["azure-function"]) {
			return fmt.Errorf("Error reading azure_function: %v", err)
		}
	}

	if err = d.Set("azure_domain", dataSourceFlattenSystemAutomationActionAzureDomain(o["azure-domain"], d, "azure_domain")); err != nil {
		if !fortiAPIPatch(o["azure-domain"]) {
			return fmt.Errorf("Error reading azure_domain: %v", err)
		}
	}

	if err = d.Set("azure_function_authorization", dataSourceFlattenSystemAutomationActionAzureFunctionAuthorization(o["azure-function-authorization"], d, "azure_function_authorization")); err != nil {
		if !fortiAPIPatch(o["azure-function-authorization"]) {
			return fmt.Errorf("Error reading azure_function_authorization: %v", err)
		}
	}

	if err = d.Set("gcp_function_region", dataSourceFlattenSystemAutomationActionGcpFunctionRegion(o["gcp-function-region"], d, "gcp_function_region")); err != nil {
		if !fortiAPIPatch(o["gcp-function-region"]) {
			return fmt.Errorf("Error reading gcp_function_region: %v", err)
		}
	}

	if err = d.Set("gcp_project", dataSourceFlattenSystemAutomationActionGcpProject(o["gcp-project"], d, "gcp_project")); err != nil {
		if !fortiAPIPatch(o["gcp-project"]) {
			return fmt.Errorf("Error reading gcp_project: %v", err)
		}
	}

	if err = d.Set("gcp_function_domain", dataSourceFlattenSystemAutomationActionGcpFunctionDomain(o["gcp-function-domain"], d, "gcp_function_domain")); err != nil {
		if !fortiAPIPatch(o["gcp-function-domain"]) {
			return fmt.Errorf("Error reading gcp_function_domain: %v", err)
		}
	}

	if err = d.Set("gcp_function", dataSourceFlattenSystemAutomationActionGcpFunction(o["gcp-function"], d, "gcp_function")); err != nil {
		if !fortiAPIPatch(o["gcp-function"]) {
			return fmt.Errorf("Error reading gcp_function: %v", err)
		}
	}

	if err = d.Set("alicloud_account_id", dataSourceFlattenSystemAutomationActionAlicloudAccountId(o["alicloud-account-id"], d, "alicloud_account_id")); err != nil {
		if !fortiAPIPatch(o["alicloud-account-id"]) {
			return fmt.Errorf("Error reading alicloud_account_id: %v", err)
		}
	}

	if err = d.Set("alicloud_region", dataSourceFlattenSystemAutomationActionAlicloudRegion(o["alicloud-region"], d, "alicloud_region")); err != nil {
		if !fortiAPIPatch(o["alicloud-region"]) {
			return fmt.Errorf("Error reading alicloud_region: %v", err)
		}
	}

	if err = d.Set("alicloud_function_domain", dataSourceFlattenSystemAutomationActionAlicloudFunctionDomain(o["alicloud-function-domain"], d, "alicloud_function_domain")); err != nil {
		if !fortiAPIPatch(o["alicloud-function-domain"]) {
			return fmt.Errorf("Error reading alicloud_function_domain: %v", err)
		}
	}

	if err = d.Set("alicloud_version", dataSourceFlattenSystemAutomationActionAlicloudVersion(o["alicloud-version"], d, "alicloud_version")); err != nil {
		if !fortiAPIPatch(o["alicloud-version"]) {
			return fmt.Errorf("Error reading alicloud_version: %v", err)
		}
	}

	if err = d.Set("alicloud_service", dataSourceFlattenSystemAutomationActionAlicloudService(o["alicloud-service"], d, "alicloud_service")); err != nil {
		if !fortiAPIPatch(o["alicloud-service"]) {
			return fmt.Errorf("Error reading alicloud_service: %v", err)
		}
	}

	if err = d.Set("alicloud_function", dataSourceFlattenSystemAutomationActionAlicloudFunction(o["alicloud-function"], d, "alicloud_function")); err != nil {
		if !fortiAPIPatch(o["alicloud-function"]) {
			return fmt.Errorf("Error reading alicloud_function: %v", err)
		}
	}

	if err = d.Set("alicloud_function_authorization", dataSourceFlattenSystemAutomationActionAlicloudFunctionAuthorization(o["alicloud-function-authorization"], d, "alicloud_function_authorization")); err != nil {
		if !fortiAPIPatch(o["alicloud-function-authorization"]) {
			return fmt.Errorf("Error reading alicloud_function_authorization: %v", err)
		}
	}

	if err = d.Set("alicloud_access_key_id", dataSourceFlattenSystemAutomationActionAlicloudAccessKeyId(o["alicloud-access-key-id"], d, "alicloud_access_key_id")); err != nil {
		if !fortiAPIPatch(o["alicloud-access-key-id"]) {
			return fmt.Errorf("Error reading alicloud_access_key_id: %v", err)
		}
	}

	if err = d.Set("message_type", dataSourceFlattenSystemAutomationActionMessageType(o["message-type"], d, "message_type")); err != nil {
		if !fortiAPIPatch(o["message-type"]) {
			return fmt.Errorf("Error reading message_type: %v", err)
		}
	}

	if err = d.Set("message", dataSourceFlattenSystemAutomationActionMessage(o["message"], d, "message")); err != nil {
		if !fortiAPIPatch(o["message"]) {
			return fmt.Errorf("Error reading message: %v", err)
		}
	}

	if err = d.Set("replacement_message", dataSourceFlattenSystemAutomationActionReplacementMessage(o["replacement-message"], d, "replacement_message")); err != nil {
		if !fortiAPIPatch(o["replacement-message"]) {
			return fmt.Errorf("Error reading replacement_message: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", dataSourceFlattenSystemAutomationActionReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group")); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("protocol", dataSourceFlattenSystemAutomationActionProtocol(o["protocol"], d, "protocol")); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("method", dataSourceFlattenSystemAutomationActionMethod(o["method"], d, "method")); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("uri", dataSourceFlattenSystemAutomationActionUri(o["uri"], d, "uri")); err != nil {
		if !fortiAPIPatch(o["uri"]) {
			return fmt.Errorf("Error reading uri: %v", err)
		}
	}

	if err = d.Set("http_body", dataSourceFlattenSystemAutomationActionHttpBody(o["http-body"], d, "http_body")); err != nil {
		if !fortiAPIPatch(o["http-body"]) {
			return fmt.Errorf("Error reading http_body: %v", err)
		}
	}

	if err = d.Set("port", dataSourceFlattenSystemAutomationActionPort(o["port"], d, "port")); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if err = d.Set("http_headers", dataSourceFlattenSystemAutomationActionHttpHeaders(o["http-headers"], d, "http_headers")); err != nil {
		if !fortiAPIPatch(o["http-headers"]) {
			return fmt.Errorf("Error reading http_headers: %v", err)
		}
	}

	if err = d.Set("headers", dataSourceFlattenSystemAutomationActionHeaders(o["headers"], d, "headers")); err != nil {
		if !fortiAPIPatch(o["headers"]) {
			return fmt.Errorf("Error reading headers: %v", err)
		}
	}

	if err = d.Set("verify_host_cert", dataSourceFlattenSystemAutomationActionVerifyHostCert(o["verify-host-cert"], d, "verify_host_cert")); err != nil {
		if !fortiAPIPatch(o["verify-host-cert"]) {
			return fmt.Errorf("Error reading verify_host_cert: %v", err)
		}
	}

	if err = d.Set("script", dataSourceFlattenSystemAutomationActionScript(o["script"], d, "script")); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("output_size", dataSourceFlattenSystemAutomationActionOutputSize(o["output-size"], d, "output_size")); err != nil {
		if !fortiAPIPatch(o["output-size"]) {
			return fmt.Errorf("Error reading output_size: %v", err)
		}
	}

	if err = d.Set("timeout", dataSourceFlattenSystemAutomationActionTimeout(o["timeout"], d, "timeout")); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("execute_security_fabric", dataSourceFlattenSystemAutomationActionExecuteSecurityFabric(o["execute-security-fabric"], d, "execute_security_fabric")); err != nil {
		if !fortiAPIPatch(o["execute-security-fabric"]) {
			return fmt.Errorf("Error reading execute_security_fabric: %v", err)
		}
	}

	if err = d.Set("accprofile", dataSourceFlattenSystemAutomationActionAccprofile(o["accprofile"], d, "accprofile")); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("security_tag", dataSourceFlattenSystemAutomationActionSecurityTag(o["security-tag"], d, "security_tag")); err != nil {
		if !fortiAPIPatch(o["security-tag"]) {
			return fmt.Errorf("Error reading security_tag: %v", err)
		}
	}

	if err = d.Set("sdn_connector", dataSourceFlattenSystemAutomationActionSdnConnector(o["sdn-connector"], d, "sdn_connector")); err != nil {
		if !fortiAPIPatch(o["sdn-connector"]) {
			return fmt.Errorf("Error reading sdn_connector: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemAutomationActionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
