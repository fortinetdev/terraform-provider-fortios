// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemCsf() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCsfUpdate,
		Read:   resourceSystemCsfRead,
		Update: resourceSystemCsfUpdate,
		Delete: resourceSystemCsfDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"upstream": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"upstream_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upstream_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 65535),
				Optional:     true,
				Computed:     true,
			},
			"group_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"group_password": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"accept_auth_by_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_unification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"configuration_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fabric_object_unification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"saml_configuration_sync": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"management_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"authorization_request_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"fabric_workers": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 4),
				Optional:     true,
				Computed:     true,
			},
			"downstream_access": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstream_accprofile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"fixed_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"trusted_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"authorization_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"serial": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 19),
							Optional:     true,
							Computed:     true,
						},
						"certificate": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 32767),
							Optional:     true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"ha_members": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 19),
							Optional:     true,
							Computed:     true,
						},
						"downstream_authorization": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"index": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(0, 1024),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"fabric_connector": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"serial": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 19),
							Optional:     true,
							Computed:     true,
						},
						"accprofile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"configuration_write_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom": &schema.Schema{
							Type:     schema.TypeList,
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
					},
				},
			},
			"forticloud_account_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_mgmt": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"file_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"file_quota_warning": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 99),
				Optional:     true,
				Computed:     true,
			},
			"fabric_device": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
							Computed:     true,
						},
						"device_ip": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"https_port": &schema.Schema{
							Type:         schema.TypeInt,
							ValidateFunc: validation.IntBetween(1, 65535),
							Optional:     true,
							Computed:     true,
						},
						"access_token": &schema.Schema{
							Type:      schema.TypeString,
							Optional:  true,
							Sensitive: true,
						},
						"device_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"login": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
							Computed:     true,
						},
						"password": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 128),
							Optional:     true,
							Sensitive:    true,
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

func resourceSystemCsfUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemCsf(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsf resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCsf(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsf resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCsf")
	}

	return resourceSystemCsfRead(d, m)
}

func resourceSystemCsfDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemCsf(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemCsf resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCsf(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemCsf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemCsf(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsf(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsf resource from API: %v", err)
	}
	return nil
}

func flattenSystemCsfStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfUpstream(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfUpstreamIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfUpstreamPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfGroupPassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfAcceptAuthByCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfLogUnification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfConfigurationSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricObjectUnification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfSamlConfigurationSync(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfManagementIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfManagementPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfAuthorizationRequestType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricWorkers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfDownstreamAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfDownstreamAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFixedKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfTrustedList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemCsfTrustedListName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization_type"
		if _, ok := i["authorization-type"]; ok {
			tmp["authorization_type"] = flattenSystemCsfTrustedListAuthorizationType(i["authorization-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			tmp["serial"] = flattenSystemCsfTrustedListSerial(i["serial"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := i["certificate"]; ok {
			tmp["certificate"] = flattenSystemCsfTrustedListCertificate(i["certificate"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenSystemCsfTrustedListAction(i["action"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := i["ha-members"]; ok {
			tmp["ha_members"] = flattenSystemCsfTrustedListHaMembers(i["ha-members"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := i["downstream-authorization"]; ok {
			tmp["downstream_authorization"] = flattenSystemCsfTrustedListDownstreamAuthorization(i["downstream-authorization"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			tmp["index"] = flattenSystemCsfTrustedListIndex(i["index"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "serial", d)
	return result
}

func flattenSystemCsfTrustedListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfTrustedListAuthorizationType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfTrustedListSerial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfTrustedListCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfTrustedListAction(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfTrustedListHaMembers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfTrustedListDownstreamAuthorization(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfTrustedListIndex(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricConnector(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			tmp["serial"] = flattenSystemCsfFabricConnectorSerial(i["serial"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accprofile"
		if _, ok := i["accprofile"]; ok {
			tmp["accprofile"] = flattenSystemCsfFabricConnectorAccprofile(i["accprofile"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "configuration_write_access"
		if _, ok := i["configuration-write-access"]; ok {
			tmp["configuration_write_access"] = flattenSystemCsfFabricConnectorConfigurationWriteAccess(i["configuration-write-access"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			tmp["vdom"] = flattenSystemCsfFabricConnectorVdom(i["vdom"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "serial", d)
	return result
}

func flattenSystemCsfFabricConnectorSerial(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricConnectorAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricConnectorConfigurationWriteAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricConnectorVdom(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemCsfFabricConnectorVdomName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemCsfFabricConnectorVdomName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfForticloudAccountEnforcement(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFileMgmt(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFileQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFileQuotaWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricDevice(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemCsfFabricDeviceName(i["name"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_ip"
		if _, ok := i["device-ip"]; ok {
			tmp["device_ip"] = flattenSystemCsfFabricDeviceDeviceIp(i["device-ip"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_port"
		if _, ok := i["https-port"]; ok {
			tmp["https_port"] = flattenSystemCsfFabricDeviceHttpsPort(i["https-port"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_token"
		if _, ok := i["access-token"]; ok {
			tmp["access_token"] = flattenSystemCsfFabricDeviceAccessToken(i["access-token"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["access_token"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := i["device-type"]; ok {
			tmp["device_type"] = flattenSystemCsfFabricDeviceDeviceType(i["device-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "login"
		if _, ok := i["login"]; ok {
			tmp["login"] = flattenSystemCsfFabricDeviceLogin(i["login"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			tmp["password"] = flattenSystemCsfFabricDevicePassword(i["password"], d, pre_append, sv)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["password"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemCsfFabricDeviceName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceDeviceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceHttpsPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceAccessToken(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceDeviceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceLogin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricDevicePassword(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCsf(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("status", flattenSystemCsfStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upstream", flattenSystemCsfUpstream(o["upstream"], d, "upstream", sv)); err != nil {
		if !fortiAPIPatch(o["upstream"]) {
			return fmt.Errorf("Error reading upstream: %v", err)
		}
	}

	if err = d.Set("upstream_ip", flattenSystemCsfUpstreamIp(o["upstream-ip"], d, "upstream_ip", sv)); err != nil {
		if !fortiAPIPatch(o["upstream-ip"]) {
			return fmt.Errorf("Error reading upstream_ip: %v", err)
		}
	}

	if err = d.Set("upstream_port", flattenSystemCsfUpstreamPort(o["upstream-port"], d, "upstream_port", sv)); err != nil {
		if !fortiAPIPatch(o["upstream-port"]) {
			return fmt.Errorf("Error reading upstream_port: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemCsfGroupName(o["group-name"], d, "group_name", sv)); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("accept_auth_by_cert", flattenSystemCsfAcceptAuthByCert(o["accept-auth-by-cert"], d, "accept_auth_by_cert", sv)); err != nil {
		if !fortiAPIPatch(o["accept-auth-by-cert"]) {
			return fmt.Errorf("Error reading accept_auth_by_cert: %v", err)
		}
	}

	if err = d.Set("log_unification", flattenSystemCsfLogUnification(o["log-unification"], d, "log_unification", sv)); err != nil {
		if !fortiAPIPatch(o["log-unification"]) {
			return fmt.Errorf("Error reading log_unification: %v", err)
		}
	}

	if err = d.Set("configuration_sync", flattenSystemCsfConfigurationSync(o["configuration-sync"], d, "configuration_sync", sv)); err != nil {
		if !fortiAPIPatch(o["configuration-sync"]) {
			return fmt.Errorf("Error reading configuration_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object_unification", flattenSystemCsfFabricObjectUnification(o["fabric-object-unification"], d, "fabric_object_unification", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-object-unification"]) {
			return fmt.Errorf("Error reading fabric_object_unification: %v", err)
		}
	}

	if err = d.Set("saml_configuration_sync", flattenSystemCsfSamlConfigurationSync(o["saml-configuration-sync"], d, "saml_configuration_sync", sv)); err != nil {
		if !fortiAPIPatch(o["saml-configuration-sync"]) {
			return fmt.Errorf("Error reading saml_configuration_sync: %v", err)
		}
	}

	if err = d.Set("management_ip", flattenSystemCsfManagementIp(o["management-ip"], d, "management_ip", sv)); err != nil {
		if !fortiAPIPatch(o["management-ip"]) {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("management_port", flattenSystemCsfManagementPort(o["management-port"], d, "management_port", sv)); err != nil {
		if !fortiAPIPatch(o["management-port"]) {
			return fmt.Errorf("Error reading management_port: %v", err)
		}
	}

	if err = d.Set("authorization_request_type", flattenSystemCsfAuthorizationRequestType(o["authorization-request-type"], d, "authorization_request_type", sv)); err != nil {
		if !fortiAPIPatch(o["authorization-request-type"]) {
			return fmt.Errorf("Error reading authorization_request_type: %v", err)
		}
	}

	if err = d.Set("certificate", flattenSystemCsfCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("fabric_workers", flattenSystemCsfFabricWorkers(o["fabric-workers"], d, "fabric_workers", sv)); err != nil {
		if !fortiAPIPatch(o["fabric-workers"]) {
			return fmt.Errorf("Error reading fabric_workers: %v", err)
		}
	}

	if err = d.Set("downstream_access", flattenSystemCsfDownstreamAccess(o["downstream-access"], d, "downstream_access", sv)); err != nil {
		if !fortiAPIPatch(o["downstream-access"]) {
			return fmt.Errorf("Error reading downstream_access: %v", err)
		}
	}

	if err = d.Set("downstream_accprofile", flattenSystemCsfDownstreamAccprofile(o["downstream-accprofile"], d, "downstream_accprofile", sv)); err != nil {
		if !fortiAPIPatch(o["downstream-accprofile"]) {
			return fmt.Errorf("Error reading downstream_accprofile: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("trusted_list", flattenSystemCsfTrustedList(o["trusted-list"], d, "trusted_list", sv)); err != nil {
			if !fortiAPIPatch(o["trusted-list"]) {
				return fmt.Errorf("Error reading trusted_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("trusted_list"); ok {
			if err = d.Set("trusted_list", flattenSystemCsfTrustedList(o["trusted-list"], d, "trusted_list", sv)); err != nil {
				if !fortiAPIPatch(o["trusted-list"]) {
					return fmt.Errorf("Error reading trusted_list: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("fabric_connector", flattenSystemCsfFabricConnector(o["fabric-connector"], d, "fabric_connector", sv)); err != nil {
			if !fortiAPIPatch(o["fabric-connector"]) {
				return fmt.Errorf("Error reading fabric_connector: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fabric_connector"); ok {
			if err = d.Set("fabric_connector", flattenSystemCsfFabricConnector(o["fabric-connector"], d, "fabric_connector", sv)); err != nil {
				if !fortiAPIPatch(o["fabric-connector"]) {
					return fmt.Errorf("Error reading fabric_connector: %v", err)
				}
			}
		}
	}

	if err = d.Set("forticloud_account_enforcement", flattenSystemCsfForticloudAccountEnforcement(o["forticloud-account-enforcement"], d, "forticloud_account_enforcement", sv)); err != nil {
		if !fortiAPIPatch(o["forticloud-account-enforcement"]) {
			return fmt.Errorf("Error reading forticloud_account_enforcement: %v", err)
		}
	}

	if err = d.Set("file_mgmt", flattenSystemCsfFileMgmt(o["file-mgmt"], d, "file_mgmt", sv)); err != nil {
		if !fortiAPIPatch(o["file-mgmt"]) {
			return fmt.Errorf("Error reading file_mgmt: %v", err)
		}
	}

	if err = d.Set("file_quota", flattenSystemCsfFileQuota(o["file-quota"], d, "file_quota", sv)); err != nil {
		if !fortiAPIPatch(o["file-quota"]) {
			return fmt.Errorf("Error reading file_quota: %v", err)
		}
	}

	if err = d.Set("file_quota_warning", flattenSystemCsfFileQuotaWarning(o["file-quota-warning"], d, "file_quota_warning", sv)); err != nil {
		if !fortiAPIPatch(o["file-quota-warning"]) {
			return fmt.Errorf("Error reading file_quota_warning: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("fabric_device", flattenSystemCsfFabricDevice(o["fabric-device"], d, "fabric_device", sv)); err != nil {
			if !fortiAPIPatch(o["fabric-device"]) {
				return fmt.Errorf("Error reading fabric_device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fabric_device"); ok {
			if err = d.Set("fabric_device", flattenSystemCsfFabricDevice(o["fabric-device"], d, "fabric_device", sv)); err != nil {
				if !fortiAPIPatch(o["fabric-device"]) {
					return fmt.Errorf("Error reading fabric_device: %v", err)
				}
			}
		}
	}

	return nil
}

func flattenSystemCsfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemCsfStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstream(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfGroupPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfAcceptAuthByCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfLogUnification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfConfigurationSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricObjectUnification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSamlConfigurationSync(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfManagementIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfManagementPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfAuthorizationRequestType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricWorkers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfDownstreamAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfDownstreamAccprofile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFixedKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemCsfTrustedListName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authorization-type"], _ = expandSystemCsfTrustedListAuthorizationType(d, i["authorization_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["serial"], _ = expandSystemCsfTrustedListSerial(d, i["serial"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["certificate"], _ = expandSystemCsfTrustedListCertificate(d, i["certificate"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandSystemCsfTrustedListAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ha-members"], _ = expandSystemCsfTrustedListHaMembers(d, i["ha_members"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["downstream-authorization"], _ = expandSystemCsfTrustedListDownstreamAuthorization(d, i["downstream_authorization"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["index"], _ = expandSystemCsfTrustedListIndex(d, i["index"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemCsfTrustedListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListAuthorizationType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListSerial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListAction(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListHaMembers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListDownstreamAuthorization(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListIndex(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnector(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["serial"], _ = expandSystemCsfFabricConnectorSerial(d, i["serial"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accprofile"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["accprofile"], _ = expandSystemCsfFabricConnectorAccprofile(d, i["accprofile"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "configuration_write_access"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["configuration-write-access"], _ = expandSystemCsfFabricConnectorConfigurationWriteAccess(d, i["configuration_write_access"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok || d.HasChange(pre_append) {
			tmp["vdom"], _ = expandSystemCsfFabricConnectorVdom(d, i["vdom"], pre_append, sv)
		} else {
			tmp["vdom"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemCsfFabricConnectorSerial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnectorAccprofile(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnectorConfigurationWriteAccess(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricConnectorVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemCsfFabricConnectorVdomName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemCsfFabricConnectorVdomName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfForticloudAccountEnforcement(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFileMgmt(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFileQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFileQuotaWarning(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDevice(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["name"], _ = expandSystemCsfFabricDeviceName(d, i["name"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device-ip"], _ = expandSystemCsfFabricDeviceDeviceIp(d, i["device_ip"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_port"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["https-port"], _ = expandSystemCsfFabricDeviceHttpsPort(d, i["https_port"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_token"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["access-token"], _ = expandSystemCsfFabricDeviceAccessToken(d, i["access_token"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device-type"], _ = expandSystemCsfFabricDeviceDeviceType(d, i["device_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "login"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["login"], _ = expandSystemCsfFabricDeviceLogin(d, i["login"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password"], _ = expandSystemCsfFabricDevicePassword(d, i["password"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemCsfFabricDeviceName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceDeviceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceHttpsPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceAccessToken(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceDeviceType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceLogin(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDevicePassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCsf(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandSystemCsfStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("upstream"); ok {
		if setArgNil {
			obj["upstream"] = nil
		} else {
			t, err := expandSystemCsfUpstream(d, v, "upstream", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upstream"] = t
			}
		}
	}

	if v, ok := d.GetOk("upstream_ip"); ok {
		if setArgNil {
			obj["upstream-ip"] = nil
		} else {
			t, err := expandSystemCsfUpstreamIp(d, v, "upstream_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upstream-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("upstream_port"); ok {
		if setArgNil {
			obj["upstream-port"] = nil
		} else {
			t, err := expandSystemCsfUpstreamPort(d, v, "upstream_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upstream-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("group_name"); ok {
		if setArgNil {
			obj["group-name"] = nil
		} else {
			t, err := expandSystemCsfGroupName(d, v, "group_name", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["group-name"] = t
			}
		}
	}

	if v, ok := d.GetOk("group_password"); ok {
		if setArgNil {
			obj["group-password"] = nil
		} else {
			t, err := expandSystemCsfGroupPassword(d, v, "group_password", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["group-password"] = t
			}
		}
	}

	if v, ok := d.GetOk("accept_auth_by_cert"); ok {
		if setArgNil {
			obj["accept-auth-by-cert"] = nil
		} else {
			t, err := expandSystemCsfAcceptAuthByCert(d, v, "accept_auth_by_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["accept-auth-by-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("log_unification"); ok {
		if setArgNil {
			obj["log-unification"] = nil
		} else {
			t, err := expandSystemCsfLogUnification(d, v, "log_unification", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-unification"] = t
			}
		}
	}

	if v, ok := d.GetOk("configuration_sync"); ok {
		if setArgNil {
			obj["configuration-sync"] = nil
		} else {
			t, err := expandSystemCsfConfigurationSync(d, v, "configuration_sync", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["configuration-sync"] = t
			}
		}
	}

	if v, ok := d.GetOk("fabric_object_unification"); ok {
		if setArgNil {
			obj["fabric-object-unification"] = nil
		} else {
			t, err := expandSystemCsfFabricObjectUnification(d, v, "fabric_object_unification", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fabric-object-unification"] = t
			}
		}
	}

	if v, ok := d.GetOk("saml_configuration_sync"); ok {
		if setArgNil {
			obj["saml-configuration-sync"] = nil
		} else {
			t, err := expandSystemCsfSamlConfigurationSync(d, v, "saml_configuration_sync", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["saml-configuration-sync"] = t
			}
		}
	}

	if v, ok := d.GetOk("management_ip"); ok {
		if setArgNil {
			obj["management-ip"] = nil
		} else {
			t, err := expandSystemCsfManagementIp(d, v, "management_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["management-ip"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("management_port"); ok {
		if setArgNil {
			obj["management-port"] = nil
		} else {
			t, err := expandSystemCsfManagementPort(d, v, "management_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["management-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("authorization_request_type"); ok {
		if setArgNil {
			obj["authorization-request-type"] = nil
		} else {
			t, err := expandSystemCsfAuthorizationRequestType(d, v, "authorization_request_type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["authorization-request-type"] = t
			}
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		if setArgNil {
			obj["certificate"] = nil
		} else {
			t, err := expandSystemCsfCertificate(d, v, "certificate", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["certificate"] = t
			}
		}
	}

	if v, ok := d.GetOk("fabric_workers"); ok {
		if setArgNil {
			obj["fabric-workers"] = nil
		} else {
			t, err := expandSystemCsfFabricWorkers(d, v, "fabric_workers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fabric-workers"] = t
			}
		}
	}

	if v, ok := d.GetOk("downstream_access"); ok {
		if setArgNil {
			obj["downstream-access"] = nil
		} else {
			t, err := expandSystemCsfDownstreamAccess(d, v, "downstream_access", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["downstream-access"] = t
			}
		}
	}

	if v, ok := d.GetOk("downstream_accprofile"); ok {
		if setArgNil {
			obj["downstream-accprofile"] = nil
		} else {
			t, err := expandSystemCsfDownstreamAccprofile(d, v, "downstream_accprofile", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["downstream-accprofile"] = t
			}
		}
	}

	if v, ok := d.GetOk("fixed_key"); ok {
		if setArgNil {
			obj["fixed-key"] = nil
		} else {
			t, err := expandSystemCsfFixedKey(d, v, "fixed_key", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fixed-key"] = t
			}
		}
	}

	if v, ok := d.GetOk("trusted_list"); ok || d.HasChange("trusted_list") {
		if setArgNil {
			obj["trusted-list"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemCsfTrustedList(d, v, "trusted_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["trusted-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("fabric_connector"); ok || d.HasChange("fabric_connector") {
		if setArgNil {
			obj["fabric-connector"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemCsfFabricConnector(d, v, "fabric_connector", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fabric-connector"] = t
			}
		}
	}

	if v, ok := d.GetOk("forticloud_account_enforcement"); ok {
		if setArgNil {
			obj["forticloud-account-enforcement"] = nil
		} else {
			t, err := expandSystemCsfForticloudAccountEnforcement(d, v, "forticloud_account_enforcement", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["forticloud-account-enforcement"] = t
			}
		}
	}

	if v, ok := d.GetOk("file_mgmt"); ok {
		if setArgNil {
			obj["file-mgmt"] = nil
		} else {
			t, err := expandSystemCsfFileMgmt(d, v, "file_mgmt", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["file-mgmt"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("file_quota"); ok {
		if setArgNil {
			obj["file-quota"] = nil
		} else {
			t, err := expandSystemCsfFileQuota(d, v, "file_quota", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["file-quota"] = t
			}
		}
	}

	if v, ok := d.GetOk("file_quota_warning"); ok {
		if setArgNil {
			obj["file-quota-warning"] = nil
		} else {
			t, err := expandSystemCsfFileQuotaWarning(d, v, "file_quota_warning", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["file-quota-warning"] = t
			}
		}
	}

	if v, ok := d.GetOk("fabric_device"); ok || d.HasChange("fabric_device") {
		if setArgNil {
			obj["fabric-device"] = make([]struct{}, 0)
		} else {
			t, err := expandSystemCsfFabricDevice(d, v, "fabric_device", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fabric-device"] = t
			}
		}
	}

	return &obj, nil
}
