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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
			"name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"description": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"action_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"system_action": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"tls_certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"forticare_email": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"email_to": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"email_from": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"email_subject": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 511),
				Optional:     true,
			},
			"email_body": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"minimum_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 2592000),
				Optional:     true,
			},
			"delay": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
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
			},
			"aws_region": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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
			},
			"aws_api_path": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"aws_api_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Sensitive:    true,
			},
			"azure_app": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"azure_function": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"azure_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"azure_function_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"azure_api_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 123),
				Optional:     true,
				Sensitive:    true,
			},
			"gcp_function_region": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"gcp_project": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"gcp_function_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"gcp_function": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"alicloud_account_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"alicloud_region": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"alicloud_function_domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"alicloud_version": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"alicloud_service": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"alicloud_function": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"alicloud_function_authorization": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"alicloud_access_key_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"alicloud_access_key_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 59),
				Optional:     true,
				Sensitive:    true,
			},
			"message_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"message": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4095),
				Optional:     true,
				Computed:     true,
			},
			"replacement_message": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"replacemsg_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"http_body": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 4095),
				Optional:     true,
			},
			"port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
			},
			"http_headers": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),
							Optional:     true,
						},
						"value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 4095),
							Optional:     true,
						},
					},
				},
			},
			"form_data": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"key": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1023),
							Optional:     true,
						},
						"value": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 4095),
							Optional:     true,
						},
					},
				},
			},
			"headers": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"header": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
						},
					},
				},
			},
			"verify_host_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"script": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"output_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 1024),
				Optional:     true,
				Computed:     true,
			},
			"timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 300),
				Optional:     true,
			},
			"duration": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 36000),
				Optional:     true,
				Computed:     true,
			},
			"output_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 36000),
				Optional:     true,
			},
			"file_only": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"execute_security_fabric": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"accprofile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"security_tag": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"sdn_connector": &schema.Schema{
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
			"regular_expression": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 1023),
				Optional:     true,
			},
			"log_debug_print": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
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

func resourceSystemAutomationActionCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemAutomationAction(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemAutomationAction resource while getting object: %v", err)
	}

	o, err := c.CreateSystemAutomationAction(obj, vdomparam)

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

	obj, err := getObjectSystemAutomationAction(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemAutomationAction resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemAutomationAction(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemAutomationAction(mkey, vdomparam)
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

	o, err := c.ReadSystemAutomationAction(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationAction resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemAutomationAction(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemAutomationAction resource from API: %v", err)
	}
	return nil
}

func flattenSystemAutomationActionName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionDescription(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionActionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionSystemAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionTlsCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionForticareEmail(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailTo(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemAutomationActionEmailToName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemAutomationActionEmailToName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailFrom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailSubject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionEmailBody(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionMinimumInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationActionDelay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationActionRequired(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiStage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAwsApiPath(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAzureApp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAzureFunction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAzureDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAzureFunctionAuthorization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionGcpFunctionRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionGcpProject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionGcpFunctionDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionGcpFunction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudAccountId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudFunctionDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudService(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudFunction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudFunctionAuthorization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAlicloudAccessKeyId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionMessageType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionMessage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionReplacementMessage(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionReplacemsgGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionProtocol(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionUri(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionHttpBody(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationActionHttpHeaders(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemAutomationActionHttpHeadersId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if cur_v, ok := i["key"]; ok {
			tmp["key"] = flattenSystemAutomationActionHttpHeadersKey(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenSystemAutomationActionHttpHeadersValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemAutomationActionHttpHeadersId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationActionHttpHeadersKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionHttpHeadersValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionFormData(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemAutomationActionFormDataId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if cur_v, ok := i["key"]; ok {
			tmp["key"] = flattenSystemAutomationActionFormDataKey(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if cur_v, ok := i["value"]; ok {
			tmp["value"] = flattenSystemAutomationActionFormDataValue(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemAutomationActionFormDataId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationActionFormDataKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionFormDataValue(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionHeaders(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "header"
		if cur_v, ok := i["header"]; ok {
			tmp["header"] = flattenSystemAutomationActionHeadersHeader(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "header", d)
	return result
}

func flattenSystemAutomationActionHeadersHeader(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionVerifyHostCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionScript(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionOutputSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationActionTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationActionDuration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationActionOutputInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemAutomationActionFileOnly(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionExecuteSecurityFabric(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionSecurityTag(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionSdnConnector(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemAutomationActionSdnConnectorName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemAutomationActionSdnConnectorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionRegularExpression(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemAutomationActionLogDebugPrint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemAutomationAction(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemAutomationActionName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("description", flattenSystemAutomationActionDescription(o["description"], d, "description", sv)); err != nil {
		if !fortiAPIPatch(o["description"]) {
			return fmt.Errorf("Error reading description: %v", err)
		}
	}

	if err = d.Set("action_type", flattenSystemAutomationActionActionType(o["action-type"], d, "action_type", sv)); err != nil {
		if !fortiAPIPatch(o["action-type"]) {
			return fmt.Errorf("Error reading action_type: %v", err)
		}
	}

	if err = d.Set("system_action", flattenSystemAutomationActionSystemAction(o["system-action"], d, "system_action", sv)); err != nil {
		if !fortiAPIPatch(o["system-action"]) {
			return fmt.Errorf("Error reading system_action: %v", err)
		}
	}

	if err = d.Set("tls_certificate", flattenSystemAutomationActionTlsCertificate(o["tls-certificate"], d, "tls_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["tls-certificate"]) {
			return fmt.Errorf("Error reading tls_certificate: %v", err)
		}
	}

	if err = d.Set("forticare_email", flattenSystemAutomationActionForticareEmail(o["forticare-email"], d, "forticare_email", sv)); err != nil {
		if !fortiAPIPatch(o["forticare-email"]) {
			return fmt.Errorf("Error reading forticare_email: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("email_to", flattenSystemAutomationActionEmailTo(o["email-to"], d, "email_to", sv)); err != nil {
			if !fortiAPIPatch(o["email-to"]) {
				return fmt.Errorf("Error reading email_to: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("email_to"); ok {
			if err = d.Set("email_to", flattenSystemAutomationActionEmailTo(o["email-to"], d, "email_to", sv)); err != nil {
				if !fortiAPIPatch(o["email-to"]) {
					return fmt.Errorf("Error reading email_to: %v", err)
				}
			}
		}
	}

	if err = d.Set("email_from", flattenSystemAutomationActionEmailFrom(o["email-from"], d, "email_from", sv)); err != nil {
		if !fortiAPIPatch(o["email-from"]) {
			return fmt.Errorf("Error reading email_from: %v", err)
		}
	}

	if err = d.Set("email_subject", flattenSystemAutomationActionEmailSubject(o["email-subject"], d, "email_subject", sv)); err != nil {
		if !fortiAPIPatch(o["email-subject"]) {
			return fmt.Errorf("Error reading email_subject: %v", err)
		}
	}

	if err = d.Set("email_body", flattenSystemAutomationActionEmailBody(o["email-body"], d, "email_body", sv)); err != nil {
		if !fortiAPIPatch(o["email-body"]) {
			return fmt.Errorf("Error reading email_body: %v", err)
		}
	}

	if err = d.Set("minimum_interval", flattenSystemAutomationActionMinimumInterval(o["minimum-interval"], d, "minimum_interval", sv)); err != nil {
		if !fortiAPIPatch(o["minimum-interval"]) {
			return fmt.Errorf("Error reading minimum_interval: %v", err)
		}
	}

	if err = d.Set("delay", flattenSystemAutomationActionDelay(o["delay"], d, "delay", sv)); err != nil {
		if !fortiAPIPatch(o["delay"]) {
			return fmt.Errorf("Error reading delay: %v", err)
		}
	}

	if err = d.Set("required", flattenSystemAutomationActionRequired(o["required"], d, "required", sv)); err != nil {
		if !fortiAPIPatch(o["required"]) {
			return fmt.Errorf("Error reading required: %v", err)
		}
	}

	if err = d.Set("aws_api_id", flattenSystemAutomationActionAwsApiId(o["aws-api-id"], d, "aws_api_id", sv)); err != nil {
		if !fortiAPIPatch(o["aws-api-id"]) {
			return fmt.Errorf("Error reading aws_api_id: %v", err)
		}
	}

	if err = d.Set("aws_region", flattenSystemAutomationActionAwsRegion(o["aws-region"], d, "aws_region", sv)); err != nil {
		if !fortiAPIPatch(o["aws-region"]) {
			return fmt.Errorf("Error reading aws_region: %v", err)
		}
	}

	if err = d.Set("aws_domain", flattenSystemAutomationActionAwsDomain(o["aws-domain"], d, "aws_domain", sv)); err != nil {
		if !fortiAPIPatch(o["aws-domain"]) {
			return fmt.Errorf("Error reading aws_domain: %v", err)
		}
	}

	if err = d.Set("aws_api_stage", flattenSystemAutomationActionAwsApiStage(o["aws-api-stage"], d, "aws_api_stage", sv)); err != nil {
		if !fortiAPIPatch(o["aws-api-stage"]) {
			return fmt.Errorf("Error reading aws_api_stage: %v", err)
		}
	}

	if err = d.Set("aws_api_path", flattenSystemAutomationActionAwsApiPath(o["aws-api-path"], d, "aws_api_path", sv)); err != nil {
		if !fortiAPIPatch(o["aws-api-path"]) {
			return fmt.Errorf("Error reading aws_api_path: %v", err)
		}
	}

	if err = d.Set("azure_app", flattenSystemAutomationActionAzureApp(o["azure-app"], d, "azure_app", sv)); err != nil {
		if !fortiAPIPatch(o["azure-app"]) {
			return fmt.Errorf("Error reading azure_app: %v", err)
		}
	}

	if err = d.Set("azure_function", flattenSystemAutomationActionAzureFunction(o["azure-function"], d, "azure_function", sv)); err != nil {
		if !fortiAPIPatch(o["azure-function"]) {
			return fmt.Errorf("Error reading azure_function: %v", err)
		}
	}

	if err = d.Set("azure_domain", flattenSystemAutomationActionAzureDomain(o["azure-domain"], d, "azure_domain", sv)); err != nil {
		if !fortiAPIPatch(o["azure-domain"]) {
			return fmt.Errorf("Error reading azure_domain: %v", err)
		}
	}

	if err = d.Set("azure_function_authorization", flattenSystemAutomationActionAzureFunctionAuthorization(o["azure-function-authorization"], d, "azure_function_authorization", sv)); err != nil {
		if !fortiAPIPatch(o["azure-function-authorization"]) {
			return fmt.Errorf("Error reading azure_function_authorization: %v", err)
		}
	}

	if err = d.Set("gcp_function_region", flattenSystemAutomationActionGcpFunctionRegion(o["gcp-function-region"], d, "gcp_function_region", sv)); err != nil {
		if !fortiAPIPatch(o["gcp-function-region"]) {
			return fmt.Errorf("Error reading gcp_function_region: %v", err)
		}
	}

	if err = d.Set("gcp_project", flattenSystemAutomationActionGcpProject(o["gcp-project"], d, "gcp_project", sv)); err != nil {
		if !fortiAPIPatch(o["gcp-project"]) {
			return fmt.Errorf("Error reading gcp_project: %v", err)
		}
	}

	if err = d.Set("gcp_function_domain", flattenSystemAutomationActionGcpFunctionDomain(o["gcp-function-domain"], d, "gcp_function_domain", sv)); err != nil {
		if !fortiAPIPatch(o["gcp-function-domain"]) {
			return fmt.Errorf("Error reading gcp_function_domain: %v", err)
		}
	}

	if err = d.Set("gcp_function", flattenSystemAutomationActionGcpFunction(o["gcp-function"], d, "gcp_function", sv)); err != nil {
		if !fortiAPIPatch(o["gcp-function"]) {
			return fmt.Errorf("Error reading gcp_function: %v", err)
		}
	}

	if err = d.Set("alicloud_account_id", flattenSystemAutomationActionAlicloudAccountId(o["alicloud-account-id"], d, "alicloud_account_id", sv)); err != nil {
		if !fortiAPIPatch(o["alicloud-account-id"]) {
			return fmt.Errorf("Error reading alicloud_account_id: %v", err)
		}
	}

	if err = d.Set("alicloud_region", flattenSystemAutomationActionAlicloudRegion(o["alicloud-region"], d, "alicloud_region", sv)); err != nil {
		if !fortiAPIPatch(o["alicloud-region"]) {
			return fmt.Errorf("Error reading alicloud_region: %v", err)
		}
	}

	if err = d.Set("alicloud_function_domain", flattenSystemAutomationActionAlicloudFunctionDomain(o["alicloud-function-domain"], d, "alicloud_function_domain", sv)); err != nil {
		if !fortiAPIPatch(o["alicloud-function-domain"]) {
			return fmt.Errorf("Error reading alicloud_function_domain: %v", err)
		}
	}

	if err = d.Set("alicloud_version", flattenSystemAutomationActionAlicloudVersion(o["alicloud-version"], d, "alicloud_version", sv)); err != nil {
		if !fortiAPIPatch(o["alicloud-version"]) {
			return fmt.Errorf("Error reading alicloud_version: %v", err)
		}
	}

	if err = d.Set("alicloud_service", flattenSystemAutomationActionAlicloudService(o["alicloud-service"], d, "alicloud_service", sv)); err != nil {
		if !fortiAPIPatch(o["alicloud-service"]) {
			return fmt.Errorf("Error reading alicloud_service: %v", err)
		}
	}

	if err = d.Set("alicloud_function", flattenSystemAutomationActionAlicloudFunction(o["alicloud-function"], d, "alicloud_function", sv)); err != nil {
		if !fortiAPIPatch(o["alicloud-function"]) {
			return fmt.Errorf("Error reading alicloud_function: %v", err)
		}
	}

	if err = d.Set("alicloud_function_authorization", flattenSystemAutomationActionAlicloudFunctionAuthorization(o["alicloud-function-authorization"], d, "alicloud_function_authorization", sv)); err != nil {
		if !fortiAPIPatch(o["alicloud-function-authorization"]) {
			return fmt.Errorf("Error reading alicloud_function_authorization: %v", err)
		}
	}

	if err = d.Set("alicloud_access_key_id", flattenSystemAutomationActionAlicloudAccessKeyId(o["alicloud-access-key-id"], d, "alicloud_access_key_id", sv)); err != nil {
		if !fortiAPIPatch(o["alicloud-access-key-id"]) {
			return fmt.Errorf("Error reading alicloud_access_key_id: %v", err)
		}
	}

	if err = d.Set("message_type", flattenSystemAutomationActionMessageType(o["message-type"], d, "message_type", sv)); err != nil {
		if !fortiAPIPatch(o["message-type"]) {
			return fmt.Errorf("Error reading message_type: %v", err)
		}
	}

	if err = d.Set("message", flattenSystemAutomationActionMessage(o["message"], d, "message", sv)); err != nil {
		if !fortiAPIPatch(o["message"]) {
			return fmt.Errorf("Error reading message: %v", err)
		}
	}

	if err = d.Set("replacement_message", flattenSystemAutomationActionReplacementMessage(o["replacement-message"], d, "replacement_message", sv)); err != nil {
		if !fortiAPIPatch(o["replacement-message"]) {
			return fmt.Errorf("Error reading replacement_message: %v", err)
		}
	}

	if err = d.Set("replacemsg_group", flattenSystemAutomationActionReplacemsgGroup(o["replacemsg-group"], d, "replacemsg_group", sv)); err != nil {
		if !fortiAPIPatch(o["replacemsg-group"]) {
			return fmt.Errorf("Error reading replacemsg_group: %v", err)
		}
	}

	if err = d.Set("protocol", flattenSystemAutomationActionProtocol(o["protocol"], d, "protocol", sv)); err != nil {
		if !fortiAPIPatch(o["protocol"]) {
			return fmt.Errorf("Error reading protocol: %v", err)
		}
	}

	if err = d.Set("method", flattenSystemAutomationActionMethod(o["method"], d, "method", sv)); err != nil {
		if !fortiAPIPatch(o["method"]) {
			return fmt.Errorf("Error reading method: %v", err)
		}
	}

	if err = d.Set("uri", flattenSystemAutomationActionUri(o["uri"], d, "uri", sv)); err != nil {
		if !fortiAPIPatch(o["uri"]) {
			return fmt.Errorf("Error reading uri: %v", err)
		}
	}

	if err = d.Set("http_body", flattenSystemAutomationActionHttpBody(o["http-body"], d, "http_body", sv)); err != nil {
		if !fortiAPIPatch(o["http-body"]) {
			return fmt.Errorf("Error reading http_body: %v", err)
		}
	}

	if err = d.Set("port", flattenSystemAutomationActionPort(o["port"], d, "port", sv)); err != nil {
		if !fortiAPIPatch(o["port"]) {
			return fmt.Errorf("Error reading port: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("http_headers", flattenSystemAutomationActionHttpHeaders(o["http-headers"], d, "http_headers", sv)); err != nil {
			if !fortiAPIPatch(o["http-headers"]) {
				return fmt.Errorf("Error reading http_headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("http_headers"); ok {
			if err = d.Set("http_headers", flattenSystemAutomationActionHttpHeaders(o["http-headers"], d, "http_headers", sv)); err != nil {
				if !fortiAPIPatch(o["http-headers"]) {
					return fmt.Errorf("Error reading http_headers: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("form_data", flattenSystemAutomationActionFormData(o["form-data"], d, "form_data", sv)); err != nil {
			if !fortiAPIPatch(o["form-data"]) {
				return fmt.Errorf("Error reading form_data: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("form_data"); ok {
			if err = d.Set("form_data", flattenSystemAutomationActionFormData(o["form-data"], d, "form_data", sv)); err != nil {
				if !fortiAPIPatch(o["form-data"]) {
					return fmt.Errorf("Error reading form_data: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("headers", flattenSystemAutomationActionHeaders(o["headers"], d, "headers", sv)); err != nil {
			if !fortiAPIPatch(o["headers"]) {
				return fmt.Errorf("Error reading headers: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("headers"); ok {
			if err = d.Set("headers", flattenSystemAutomationActionHeaders(o["headers"], d, "headers", sv)); err != nil {
				if !fortiAPIPatch(o["headers"]) {
					return fmt.Errorf("Error reading headers: %v", err)
				}
			}
		}
	}

	if err = d.Set("verify_host_cert", flattenSystemAutomationActionVerifyHostCert(o["verify-host-cert"], d, "verify_host_cert", sv)); err != nil {
		if !fortiAPIPatch(o["verify-host-cert"]) {
			return fmt.Errorf("Error reading verify_host_cert: %v", err)
		}
	}

	if err = d.Set("script", flattenSystemAutomationActionScript(o["script"], d, "script", sv)); err != nil {
		if !fortiAPIPatch(o["script"]) {
			return fmt.Errorf("Error reading script: %v", err)
		}
	}

	if err = d.Set("output_size", flattenSystemAutomationActionOutputSize(o["output-size"], d, "output_size", sv)); err != nil {
		if !fortiAPIPatch(o["output-size"]) {
			return fmt.Errorf("Error reading output_size: %v", err)
		}
	}

	if err = d.Set("timeout", flattenSystemAutomationActionTimeout(o["timeout"], d, "timeout", sv)); err != nil {
		if !fortiAPIPatch(o["timeout"]) {
			return fmt.Errorf("Error reading timeout: %v", err)
		}
	}

	if err = d.Set("duration", flattenSystemAutomationActionDuration(o["duration"], d, "duration", sv)); err != nil {
		if !fortiAPIPatch(o["duration"]) {
			return fmt.Errorf("Error reading duration: %v", err)
		}
	}

	if err = d.Set("output_interval", flattenSystemAutomationActionOutputInterval(o["output-interval"], d, "output_interval", sv)); err != nil {
		if !fortiAPIPatch(o["output-interval"]) {
			return fmt.Errorf("Error reading output_interval: %v", err)
		}
	}

	if err = d.Set("file_only", flattenSystemAutomationActionFileOnly(o["file-only"], d, "file_only", sv)); err != nil {
		if !fortiAPIPatch(o["file-only"]) {
			return fmt.Errorf("Error reading file_only: %v", err)
		}
	}

	if err = d.Set("execute_security_fabric", flattenSystemAutomationActionExecuteSecurityFabric(o["execute-security-fabric"], d, "execute_security_fabric", sv)); err != nil {
		if !fortiAPIPatch(o["execute-security-fabric"]) {
			return fmt.Errorf("Error reading execute_security_fabric: %v", err)
		}
	}

	if err = d.Set("accprofile", flattenSystemAutomationActionAccprofile(o["accprofile"], d, "accprofile", sv)); err != nil {
		if !fortiAPIPatch(o["accprofile"]) {
			return fmt.Errorf("Error reading accprofile: %v", err)
		}
	}

	if err = d.Set("security_tag", flattenSystemAutomationActionSecurityTag(o["security-tag"], d, "security_tag", sv)); err != nil {
		if !fortiAPIPatch(o["security-tag"]) {
			return fmt.Errorf("Error reading security_tag: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("sdn_connector", flattenSystemAutomationActionSdnConnector(o["sdn-connector"], d, "sdn_connector", sv)); err != nil {
			if !fortiAPIPatch(o["sdn-connector"]) {
				return fmt.Errorf("Error reading sdn_connector: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("sdn_connector"); ok {
			if err = d.Set("sdn_connector", flattenSystemAutomationActionSdnConnector(o["sdn-connector"], d, "sdn_connector", sv)); err != nil {
				if !fortiAPIPatch(o["sdn-connector"]) {
					return fmt.Errorf("Error reading sdn_connector: %v", err)
				}
			}
		}
	}

	if err = d.Set("regular_expression", flattenSystemAutomationActionRegularExpression(o["regular-expression"], d, "regular_expression", sv)); err != nil {
		if !fortiAPIPatch(o["regular-expression"]) {
			return fmt.Errorf("Error reading regular_expression: %v", err)
		}
	}

	if err = d.Set("log_debug_print", flattenSystemAutomationActionLogDebugPrint(o["log-debug-print"], d, "log_debug_print", sv)); err != nil {
		if !fortiAPIPatch(o["log-debug-print"]) {
			return fmt.Errorf("Error reading log_debug_print: %v", err)
		}
	}

	return nil
}

func flattenSystemAutomationActionFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemAutomationActionName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionDescription(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionActionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionSystemAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionTlsCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionForticareEmail(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailTo(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemAutomationActionEmailToName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationActionEmailToName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailFrom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailSubject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionEmailBody(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMinimumInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionDelay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionRequired(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiStage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiPath(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAwsApiKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAzureApp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAzureFunction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAzureDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAzureFunctionAuthorization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAzureApiKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionGcpFunctionRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionGcpProject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionGcpFunctionDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionGcpFunction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudAccountId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudFunctionDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudService(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudFunction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudFunctionAuthorization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudAccessKeyId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAlicloudAccessKeySecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMessageType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMessage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionReplacementMessage(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionReplacemsgGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionProtocol(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionUri(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpBody(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpHeaders(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAutomationActionHttpHeadersId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key"], _ = expandSystemAutomationActionHttpHeadersKey(d, i["key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandSystemAutomationActionHttpHeadersValue(d, i["value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationActionHttpHeadersId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpHeadersKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHttpHeadersValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionFormData(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemAutomationActionFormDataId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "key"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["key"], _ = expandSystemAutomationActionFormDataKey(d, i["key"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["key"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "value"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["value"], _ = expandSystemAutomationActionFormDataValue(d, i["value"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["value"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationActionFormDataId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionFormDataKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionFormDataValue(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionHeaders(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["header"], _ = expandSystemAutomationActionHeadersHeader(d, i["header"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationActionHeadersHeader(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionVerifyHostCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionScript(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionOutputSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionDuration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionOutputInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionFileOnly(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionExecuteSecurityFabric(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionAccprofile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionSecurityTag(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionSdnConnector(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemAutomationActionSdnConnectorName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemAutomationActionSdnConnectorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionRegularExpression(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemAutomationActionLogDebugPrint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemAutomationAction(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemAutomationActionName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("description"); ok {
		t, err := expandSystemAutomationActionDescription(d, v, "description", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["description"] = t
		}
	} else if d.HasChange("description") {
		obj["description"] = nil
	}

	if v, ok := d.GetOk("action_type"); ok {
		t, err := expandSystemAutomationActionActionType(d, v, "action_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["action-type"] = t
		}
	}

	if v, ok := d.GetOk("system_action"); ok {
		t, err := expandSystemAutomationActionSystemAction(d, v, "system_action", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["system-action"] = t
		}
	} else if d.HasChange("system_action") {
		obj["system-action"] = nil
	}

	if v, ok := d.GetOk("tls_certificate"); ok {
		t, err := expandSystemAutomationActionTlsCertificate(d, v, "tls_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tls-certificate"] = t
		}
	} else if d.HasChange("tls_certificate") {
		obj["tls-certificate"] = nil
	}

	if v, ok := d.GetOk("forticare_email"); ok {
		t, err := expandSystemAutomationActionForticareEmail(d, v, "forticare_email", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forticare-email"] = t
		}
	}

	if v, ok := d.GetOk("email_to"); ok || d.HasChange("email_to") {
		t, err := expandSystemAutomationActionEmailTo(d, v, "email_to", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-to"] = t
		}
	}

	if v, ok := d.GetOk("email_from"); ok {
		t, err := expandSystemAutomationActionEmailFrom(d, v, "email_from", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-from"] = t
		}
	} else if d.HasChange("email_from") {
		obj["email-from"] = nil
	}

	if v, ok := d.GetOk("email_subject"); ok {
		t, err := expandSystemAutomationActionEmailSubject(d, v, "email_subject", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-subject"] = t
		}
	} else if d.HasChange("email_subject") {
		obj["email-subject"] = nil
	}

	if v, ok := d.GetOk("email_body"); ok {
		t, err := expandSystemAutomationActionEmailBody(d, v, "email_body", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["email-body"] = t
		}
	} else if d.HasChange("email_body") {
		obj["email-body"] = nil
	}

	if v, ok := d.GetOkExists("minimum_interval"); ok {
		t, err := expandSystemAutomationActionMinimumInterval(d, v, "minimum_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["minimum-interval"] = t
		}
	} else if d.HasChange("minimum_interval") {
		obj["minimum-interval"] = nil
	}

	if v, ok := d.GetOkExists("delay"); ok {
		t, err := expandSystemAutomationActionDelay(d, v, "delay", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["delay"] = t
		}
	} else if d.HasChange("delay") {
		obj["delay"] = nil
	}

	if v, ok := d.GetOk("required"); ok {
		t, err := expandSystemAutomationActionRequired(d, v, "required", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["required"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_id"); ok {
		t, err := expandSystemAutomationActionAwsApiId(d, v, "aws_api_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-id"] = t
		}
	} else if d.HasChange("aws_api_id") {
		obj["aws-api-id"] = nil
	}

	if v, ok := d.GetOk("aws_region"); ok {
		t, err := expandSystemAutomationActionAwsRegion(d, v, "aws_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-region"] = t
		}
	} else if d.HasChange("aws_region") {
		obj["aws-region"] = nil
	}

	if v, ok := d.GetOk("aws_domain"); ok {
		t, err := expandSystemAutomationActionAwsDomain(d, v, "aws_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-domain"] = t
		}
	}

	if v, ok := d.GetOk("aws_api_stage"); ok {
		t, err := expandSystemAutomationActionAwsApiStage(d, v, "aws_api_stage", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-stage"] = t
		}
	} else if d.HasChange("aws_api_stage") {
		obj["aws-api-stage"] = nil
	}

	if v, ok := d.GetOk("aws_api_path"); ok {
		t, err := expandSystemAutomationActionAwsApiPath(d, v, "aws_api_path", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-path"] = t
		}
	} else if d.HasChange("aws_api_path") {
		obj["aws-api-path"] = nil
	}

	if v, ok := d.GetOk("aws_api_key"); ok {
		t, err := expandSystemAutomationActionAwsApiKey(d, v, "aws_api_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["aws-api-key"] = t
		}
	} else if d.HasChange("aws_api_key") {
		obj["aws-api-key"] = nil
	}

	if v, ok := d.GetOk("azure_app"); ok {
		t, err := expandSystemAutomationActionAzureApp(d, v, "azure_app", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-app"] = t
		}
	} else if d.HasChange("azure_app") {
		obj["azure-app"] = nil
	}

	if v, ok := d.GetOk("azure_function"); ok {
		t, err := expandSystemAutomationActionAzureFunction(d, v, "azure_function", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-function"] = t
		}
	} else if d.HasChange("azure_function") {
		obj["azure-function"] = nil
	}

	if v, ok := d.GetOk("azure_domain"); ok {
		t, err := expandSystemAutomationActionAzureDomain(d, v, "azure_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-domain"] = t
		}
	}

	if v, ok := d.GetOk("azure_function_authorization"); ok {
		t, err := expandSystemAutomationActionAzureFunctionAuthorization(d, v, "azure_function_authorization", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-function-authorization"] = t
		}
	}

	if v, ok := d.GetOk("azure_api_key"); ok {
		t, err := expandSystemAutomationActionAzureApiKey(d, v, "azure_api_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-api-key"] = t
		}
	} else if d.HasChange("azure_api_key") {
		obj["azure-api-key"] = nil
	}

	if v, ok := d.GetOk("gcp_function_region"); ok {
		t, err := expandSystemAutomationActionGcpFunctionRegion(d, v, "gcp_function_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-function-region"] = t
		}
	} else if d.HasChange("gcp_function_region") {
		obj["gcp-function-region"] = nil
	}

	if v, ok := d.GetOk("gcp_project"); ok {
		t, err := expandSystemAutomationActionGcpProject(d, v, "gcp_project", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-project"] = t
		}
	} else if d.HasChange("gcp_project") {
		obj["gcp-project"] = nil
	}

	if v, ok := d.GetOk("gcp_function_domain"); ok {
		t, err := expandSystemAutomationActionGcpFunctionDomain(d, v, "gcp_function_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-function-domain"] = t
		}
	}

	if v, ok := d.GetOk("gcp_function"); ok {
		t, err := expandSystemAutomationActionGcpFunction(d, v, "gcp_function", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-function"] = t
		}
	} else if d.HasChange("gcp_function") {
		obj["gcp-function"] = nil
	}

	if v, ok := d.GetOk("alicloud_account_id"); ok {
		t, err := expandSystemAutomationActionAlicloudAccountId(d, v, "alicloud_account_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-account-id"] = t
		}
	} else if d.HasChange("alicloud_account_id") {
		obj["alicloud-account-id"] = nil
	}

	if v, ok := d.GetOk("alicloud_region"); ok {
		t, err := expandSystemAutomationActionAlicloudRegion(d, v, "alicloud_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-region"] = t
		}
	} else if d.HasChange("alicloud_region") {
		obj["alicloud-region"] = nil
	}

	if v, ok := d.GetOk("alicloud_function_domain"); ok {
		t, err := expandSystemAutomationActionAlicloudFunctionDomain(d, v, "alicloud_function_domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-function-domain"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_version"); ok {
		t, err := expandSystemAutomationActionAlicloudVersion(d, v, "alicloud_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-version"] = t
		}
	} else if d.HasChange("alicloud_version") {
		obj["alicloud-version"] = nil
	}

	if v, ok := d.GetOk("alicloud_service"); ok {
		t, err := expandSystemAutomationActionAlicloudService(d, v, "alicloud_service", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-service"] = t
		}
	} else if d.HasChange("alicloud_service") {
		obj["alicloud-service"] = nil
	}

	if v, ok := d.GetOk("alicloud_function"); ok {
		t, err := expandSystemAutomationActionAlicloudFunction(d, v, "alicloud_function", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-function"] = t
		}
	} else if d.HasChange("alicloud_function") {
		obj["alicloud-function"] = nil
	}

	if v, ok := d.GetOk("alicloud_function_authorization"); ok {
		t, err := expandSystemAutomationActionAlicloudFunctionAuthorization(d, v, "alicloud_function_authorization", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-function-authorization"] = t
		}
	}

	if v, ok := d.GetOk("alicloud_access_key_id"); ok {
		t, err := expandSystemAutomationActionAlicloudAccessKeyId(d, v, "alicloud_access_key_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-access-key-id"] = t
		}
	} else if d.HasChange("alicloud_access_key_id") {
		obj["alicloud-access-key-id"] = nil
	}

	if v, ok := d.GetOk("alicloud_access_key_secret"); ok {
		t, err := expandSystemAutomationActionAlicloudAccessKeySecret(d, v, "alicloud_access_key_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alicloud-access-key-secret"] = t
		}
	} else if d.HasChange("alicloud_access_key_secret") {
		obj["alicloud-access-key-secret"] = nil
	}

	if v, ok := d.GetOk("message_type"); ok {
		t, err := expandSystemAutomationActionMessageType(d, v, "message_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-type"] = t
		}
	}

	if v, ok := d.GetOk("message"); ok {
		t, err := expandSystemAutomationActionMessage(d, v, "message", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message"] = t
		}
	}

	if v, ok := d.GetOk("replacement_message"); ok {
		t, err := expandSystemAutomationActionReplacementMessage(d, v, "replacement_message", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacement-message"] = t
		}
	}

	if v, ok := d.GetOk("replacemsg_group"); ok {
		t, err := expandSystemAutomationActionReplacemsgGroup(d, v, "replacemsg_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["replacemsg-group"] = t
		}
	} else if d.HasChange("replacemsg_group") {
		obj["replacemsg-group"] = nil
	}

	if v, ok := d.GetOk("protocol"); ok {
		t, err := expandSystemAutomationActionProtocol(d, v, "protocol", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["protocol"] = t
		}
	}

	if v, ok := d.GetOk("method"); ok {
		t, err := expandSystemAutomationActionMethod(d, v, "method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["method"] = t
		}
	}

	if v, ok := d.GetOk("uri"); ok {
		t, err := expandSystemAutomationActionUri(d, v, "uri", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uri"] = t
		}
	} else if d.HasChange("uri") {
		obj["uri"] = nil
	}

	if v, ok := d.GetOk("http_body"); ok {
		t, err := expandSystemAutomationActionHttpBody(d, v, "http_body", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-body"] = t
		}
	} else if d.HasChange("http_body") {
		obj["http-body"] = nil
	}

	if v, ok := d.GetOk("port"); ok {
		t, err := expandSystemAutomationActionPort(d, v, "port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["port"] = t
		}
	} else if d.HasChange("port") {
		obj["port"] = nil
	}

	if v, ok := d.GetOk("http_headers"); ok || d.HasChange("http_headers") {
		t, err := expandSystemAutomationActionHttpHeaders(d, v, "http_headers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["http-headers"] = t
		}
	}

	if v, ok := d.GetOk("form_data"); ok || d.HasChange("form_data") {
		t, err := expandSystemAutomationActionFormData(d, v, "form_data", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["form-data"] = t
		}
	}

	if v, ok := d.GetOk("headers"); ok || d.HasChange("headers") {
		t, err := expandSystemAutomationActionHeaders(d, v, "headers", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["headers"] = t
		}
	}

	if v, ok := d.GetOk("verify_host_cert"); ok {
		t, err := expandSystemAutomationActionVerifyHostCert(d, v, "verify_host_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-host-cert"] = t
		}
	}

	if v, ok := d.GetOk("script"); ok {
		t, err := expandSystemAutomationActionScript(d, v, "script", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["script"] = t
		}
	} else if d.HasChange("script") {
		obj["script"] = nil
	}

	if v, ok := d.GetOk("output_size"); ok {
		t, err := expandSystemAutomationActionOutputSize(d, v, "output_size", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-size"] = t
		}
	}

	if v, ok := d.GetOkExists("timeout"); ok {
		t, err := expandSystemAutomationActionTimeout(d, v, "timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["timeout"] = t
		}
	} else if d.HasChange("timeout") {
		obj["timeout"] = nil
	}

	if v, ok := d.GetOk("duration"); ok {
		t, err := expandSystemAutomationActionDuration(d, v, "duration", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["duration"] = t
		}
	}

	if v, ok := d.GetOkExists("output_interval"); ok {
		t, err := expandSystemAutomationActionOutputInterval(d, v, "output_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["output-interval"] = t
		}
	} else if d.HasChange("output_interval") {
		obj["output-interval"] = nil
	}

	if v, ok := d.GetOk("file_only"); ok {
		t, err := expandSystemAutomationActionFileOnly(d, v, "file_only", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["file-only"] = t
		}
	}

	if v, ok := d.GetOk("execute_security_fabric"); ok {
		t, err := expandSystemAutomationActionExecuteSecurityFabric(d, v, "execute_security_fabric", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["execute-security-fabric"] = t
		}
	}

	if v, ok := d.GetOk("accprofile"); ok {
		t, err := expandSystemAutomationActionAccprofile(d, v, "accprofile", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["accprofile"] = t
		}
	}

	if v, ok := d.GetOk("security_tag"); ok {
		t, err := expandSystemAutomationActionSecurityTag(d, v, "security_tag", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["security-tag"] = t
		}
	} else if d.HasChange("security_tag") {
		obj["security-tag"] = nil
	}

	if v, ok := d.GetOk("sdn_connector"); ok || d.HasChange("sdn_connector") {
		t, err := expandSystemAutomationActionSdnConnector(d, v, "sdn_connector", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["sdn-connector"] = t
		}
	}

	if v, ok := d.GetOk("regular_expression"); ok {
		t, err := expandSystemAutomationActionRegularExpression(d, v, "regular_expression", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["regular-expression"] = t
		}
	} else if d.HasChange("regular_expression") {
		obj["regular-expression"] = nil
	}

	if v, ok := d.GetOk("log_debug_print"); ok {
		t, err := expandSystemAutomationActionLogDebugPrint(d, v, "log_debug_print", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-debug-print"] = t
		}
	}

	return &obj, nil
}
