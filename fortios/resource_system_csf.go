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
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"uid": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"upstream": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 255),
				Optional:     true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upstream_interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upstream_interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
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
			},
			"management_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
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
			"legacy_authentication": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"downstream_accprofile": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
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
						},
						"accprofile": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 35),
							Optional:     true,
						},
						"configuration_write_access": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"vdom": &schema.Schema{
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
						},
						"login": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 64),
							Optional:     true,
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

func flattenSystemCsfUid(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfUpstream(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfUpstreamInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfUpstreamInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfUpstreamIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfUpstreamPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemCsfGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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
	return convintf2i(v)
}

func flattenSystemCsfAuthorizationRequestType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricWorkers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemCsfDownstreamAccess(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfLegacyAuthentication(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfDownstreamAccprofile(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "serial", "serial")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenSystemCsfTrustedListName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["authorization-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization_type"
			}
			tmp["authorization_type"] = flattenSystemCsfTrustedListAuthorizationType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["serial"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
			}
			tmp["serial"] = flattenSystemCsfTrustedListSerial(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["certificate"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
			}
			tmp["certificate"] = flattenSystemCsfTrustedListCertificate(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["action"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
			}
			tmp["action"] = flattenSystemCsfTrustedListAction(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["ha-members"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
			}
			tmp["ha_members"] = flattenSystemCsfTrustedListHaMembers(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["downstream-authorization"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
			}
			tmp["downstream_authorization"] = flattenSystemCsfTrustedListDownstreamAuthorization(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["index"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
			}
			tmp["index"] = flattenSystemCsfTrustedListIndex(cur_v, d, pre_append, sv)
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
	return convintf2i(v)
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "serial", "serial")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["serial"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
			}
			tmp["serial"] = flattenSystemCsfFabricConnectorSerial(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["accprofile"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "accprofile"
			}
			tmp["accprofile"] = flattenSystemCsfFabricConnectorAccprofile(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["configuration-write-access"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "configuration_write_access"
			}
			tmp["configuration_write_access"] = flattenSystemCsfFabricConnectorConfigurationWriteAccess(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["vdom"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
			}
			tmp["vdom"] = flattenSystemCsfFabricConnectorVdom(cur_v, d, pre_append, sv)
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenSystemCsfFabricConnectorVdomName(cur_v, d, pre_append, sv)
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
	return convintf2i(v)
}

func flattenSystemCsfFileQuotaWarning(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
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

	tf_list := []interface{}{}
	if tf_v, ok := d.GetOk(pre); ok {
		if tf_list, ok = tf_v.([]interface{}); !ok {
			log.Printf("[DEBUG] Argument %v could not convert to []interface{}.", pre)
		}
	}

	parsed_list := mergeBlock(tf_list, l, "name", "name")

	con := 0
	for _, r := range parsed_list {
		tmp := make(map[string]interface{})
		i := r.(map[string]interface{})
		tf_exist := i["tf_exist"].(bool)

		if cur_v, ok := i["name"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "name"
			}
			tmp["name"] = flattenSystemCsfFabricDeviceName(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["device-ip"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "device_ip"
			}
			tmp["device_ip"] = flattenSystemCsfFabricDeviceDeviceIp(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["https-port"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "https_port"
			}
			tmp["https_port"] = flattenSystemCsfFabricDeviceHttpsPort(cur_v, d, pre_append, sv)
		}

		if _, ok := i["access-token"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "access_token"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["access_token"] = c
				}
			}
		}

		if cur_v, ok := i["device-type"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
			}
			tmp["device_type"] = flattenSystemCsfFabricDeviceDeviceType(cur_v, d, pre_append, sv)
		}

		if cur_v, ok := i["login"]; ok {
			pre_append := ""
			if tf_exist {
				pre_append = pre + "." + strconv.Itoa(con) + "." + "login"
			}
			tmp["login"] = flattenSystemCsfFabricDeviceLogin(cur_v, d, pre_append, sv)
		}

		if _, ok := i["password"]; ok {
			if tf_exist {
				pre_append := pre + "." + strconv.Itoa(con) + "." + "password"
				c := d.Get(pre_append).(string)
				if c != "" {
					tmp["password"] = c
				}
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
	return convintf2i(v)
}

func flattenSystemCsfFabricDeviceDeviceType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceLogin(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
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

	if err = d.Set("uid", flattenSystemCsfUid(o["uid"], d, "uid", sv)); err != nil {
		if !fortiAPIPatch(o["uid"]) {
			return fmt.Errorf("Error reading uid: %v", err)
		}
	}

	if err = d.Set("upstream", flattenSystemCsfUpstream(o["upstream"], d, "upstream", sv)); err != nil {
		if !fortiAPIPatch(o["upstream"]) {
			return fmt.Errorf("Error reading upstream: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenSystemCsfSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("upstream_interface_select_method", flattenSystemCsfUpstreamInterfaceSelectMethod(o["upstream-interface-select-method"], d, "upstream_interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["upstream-interface-select-method"]) {
			return fmt.Errorf("Error reading upstream_interface_select_method: %v", err)
		}
	}

	if err = d.Set("upstream_interface", flattenSystemCsfUpstreamInterface(o["upstream-interface"], d, "upstream_interface", sv)); err != nil {
		if !fortiAPIPatch(o["upstream-interface"]) {
			return fmt.Errorf("Error reading upstream_interface: %v", err)
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

	if err = d.Set("legacy_authentication", flattenSystemCsfLegacyAuthentication(o["legacy-authentication"], d, "legacy_authentication", sv)); err != nil {
		if !fortiAPIPatch(o["legacy-authentication"]) {
			return fmt.Errorf("Error reading legacy_authentication: %v", err)
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

func expandSystemCsfUid(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstream(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

func expandSystemCsfLegacyAuthentication(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["authorization-type"], _ = expandSystemCsfTrustedListAuthorizationType(d, i["authorization_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["serial"], _ = expandSystemCsfTrustedListSerial(d, i["serial"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["serial"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["certificate"], _ = expandSystemCsfTrustedListCertificate(d, i["certificate"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["certificate"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandSystemCsfTrustedListAction(d, i["action"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ha-members"], _ = expandSystemCsfTrustedListHaMembers(d, i["ha_members"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ha-members"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["downstream-authorization"], _ = expandSystemCsfTrustedListDownstreamAuthorization(d, i["downstream_authorization"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["index"], _ = expandSystemCsfTrustedListIndex(d, i["index"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["index"] = nil
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
		} else if d.HasChange(pre_append) {
			tmp["serial"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accprofile"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["accprofile"], _ = expandSystemCsfFabricConnectorAccprofile(d, i["accprofile"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["accprofile"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "configuration_write_access"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["configuration-write-access"], _ = expandSystemCsfFabricConnectorConfigurationWriteAccess(d, i["configuration_write_access"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["vdom"], _ = expandSystemCsfFabricConnectorVdom(d, i["vdom"], pre_append, sv)
		} else if d.HasChange(pre_append) {
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

		tmp["name"], _ = expandSystemCsfFabricConnectorVdomName(d, i["name"], pre_append, sv)

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
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
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
		} else if d.HasChange(pre_append) {
			tmp["access-token"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device-type"], _ = expandSystemCsfFabricDeviceDeviceType(d, i["device_type"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["device-type"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "login"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["login"], _ = expandSystemCsfFabricDeviceLogin(d, i["login"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["login"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password"], _ = expandSystemCsfFabricDevicePassword(d, i["password"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["password"] = nil
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

	if v, ok := d.GetOk("uid"); ok {
		if setArgNil {
			obj["uid"] = nil
		} else {
			t, err := expandSystemCsfUid(d, v, "uid", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uid"] = t
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
	} else if d.HasChange("upstream") {
		obj["upstream"] = nil
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {
			t, err := expandSystemCsfSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("upstream_interface_select_method"); ok {
		if setArgNil {
			obj["upstream-interface-select-method"] = nil
		} else {
			t, err := expandSystemCsfUpstreamInterfaceSelectMethod(d, v, "upstream_interface_select_method", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upstream-interface-select-method"] = t
			}
		}
	}

	if v, ok := d.GetOk("upstream_interface"); ok {
		if setArgNil {
			obj["upstream-interface"] = nil
		} else {
			t, err := expandSystemCsfUpstreamInterface(d, v, "upstream_interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upstream-interface"] = t
			}
		}
	} else if d.HasChange("upstream_interface") {
		obj["upstream-interface"] = nil
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
	} else if d.HasChange("group_password") {
		obj["group-password"] = nil
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
	} else if d.HasChange("management_ip") {
		obj["management-ip"] = nil
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
	} else if d.HasChange("management_port") {
		obj["management-port"] = nil
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
	} else if d.HasChange("certificate") {
		obj["certificate"] = nil
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

	if v, ok := d.GetOk("legacy_authentication"); ok {
		if setArgNil {
			obj["legacy-authentication"] = nil
		} else {
			t, err := expandSystemCsfLegacyAuthentication(d, v, "legacy_authentication", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["legacy-authentication"] = t
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
	} else if d.HasChange("downstream_accprofile") {
		obj["downstream-accprofile"] = nil
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
	} else if d.HasChange("fixed_key") {
		obj["fixed-key"] = nil
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
