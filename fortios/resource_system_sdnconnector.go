// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure connection to SDN Connector.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

func resourceSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemSdnConnectorCreate,
		Read:   resourceSystemSdnConnectorRead,
		Update: resourceSystemSdnConnectorUpdate,
		Delete: resourceSystemSdnConnectorDelete,

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
				ValidateFunc: validation.StringLenBetween(0, 35),
				ForceNew:     true,
				Optional:     true,
				Computed:     true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"proxy": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
			},
			"ha_status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"verify_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 15),
							Optional:     true,
						},
					},
				},
			},
			"server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"message_server_port": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
			},
			"username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"vcenter_server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"vcenter_username": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 64),
				Optional:     true,
			},
			"vcenter_password": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"access_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Sensitive:    true,
			},
			"secret_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 59),
				Optional:     true,
				Sensitive:    true,
			},
			"region": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
			},
			"vpc_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
			},
			"alt_resource_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"external_account_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_arn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 2047),
							Optional:     true,
						},
						"external_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 1399),
							Optional:     true,
						},
						"region_list": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 31),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"tenant_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"subscription_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"login_endpoint": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"resource_url": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"client_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"client_secret": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 59),
				Optional:     true,
				Sensitive:    true,
			},
			"resource_group": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"azure_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"nic": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"peer_nic": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
									},
									"private_ip": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 39),
										Optional:     true,
									},
									"public_ip": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
									},
									"resource_group": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"route_table": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"subscription_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"resource_group": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"route": &schema.Schema{
							Type:     schema.TypeList,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 63),
										Optional:     true,
									},
									"next_hop": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"user_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"compartment_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compartment_id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
						},
					},
				},
			},
			"oci_region_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 31),
							Optional:     true,
						},
					},
				},
			},
			"compartment_id": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"oci_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"oci_region_type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"oci_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"oci_fingerprint": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
			},
			"external_ip": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
					},
				},
			},
			"route": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
					},
				},
			},
			"forwarding_rule": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_name": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
						"target": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 63),
							Optional:     true,
						},
					},
				},
			},
			"gcp_project_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 127),
							Optional:     true,
						},
						"gcp_zone_list": &schema.Schema{
							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:         schema.TypeString,
										ValidateFunc: validation.StringLenBetween(0, 127),
										Optional:     true,
									},
								},
							},
						},
					},
				},
			},
			"use_metadata_iam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"gcp_project": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"service_account": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"key_passwd": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
				Sensitive:    true,
			},
			"private_key": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"secret_token": &schema.Schema{
				Type:      schema.TypeString,
				Optional:  true,
				Sensitive: true,
			},
			"domain": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"group_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"server_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"server_ca_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
			},
			"api_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 59),
				Optional:     true,
				Sensitive:    true,
			},
			"compute_generation": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 2),
				Optional:     true,
				Computed:     true,
			},
			"ibm_region_gen1": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ibm_region_gen2": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"ibm_region": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3600),
				Optional:     true,
				Computed:     true,
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

func resourceSystemSdnConnectorCreate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdnConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemSdnConnector resource while getting object: %v", err)
	}

	o, err := c.CreateSystemSdnConnector(obj, vdomparam)

	if err != nil {
		return fmt.Errorf("Error creating SystemSdnConnector resource: %v", err)
	}

	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSdnConnector")
	}

	return resourceSystemSdnConnectorRead(d, m)
}

func resourceSystemSdnConnectorUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectSystemSdnConnector(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnector resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemSdnConnector(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating SystemSdnConnector resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemSdnConnector")
	}

	return resourceSystemSdnConnectorRead(d, m)
}

func resourceSystemSdnConnectorDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteSystemSdnConnector(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting SystemSdnConnector resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemSdnConnectorRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadSystemSdnConnector(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnConnector resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemSdnConnector(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemSdnConnector resource from API: %v", err)
	}
	return nil
}

func flattenSystemSdnConnectorName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorProxy(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorHaStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorVerifyCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServerList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenSystemSdnConnectorServerListIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "ip", d)
	return result
}

func flattenSystemSdnConnectorServerListIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSdnConnectorMessageServerPort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSdnConnectorUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorVcenterServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorVcenterUsername(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorVpcId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorAltResourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorExternalAccountList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role_arn"
		if cur_v, ok := i["role-arn"]; ok {
			tmp["role_arn"] = flattenSystemSdnConnectorExternalAccountListRoleArn(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_id"
		if cur_v, ok := i["external-id"]; ok {
			tmp["external_id"] = flattenSystemSdnConnectorExternalAccountListExternalId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region_list"
		if cur_v, ok := i["region-list"]; ok {
			tmp["region_list"] = flattenSystemSdnConnectorExternalAccountListRegionList(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "role_arn", d)
	return result
}

func flattenSystemSdnConnectorExternalAccountListRoleArn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorExternalAccountListExternalId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorExternalAccountListRegionList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if cur_v, ok := i["region"]; ok {
			tmp["region"] = flattenSystemSdnConnectorExternalAccountListRegionListRegion(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "region", d)
	return result
}

func flattenSystemSdnConnectorExternalAccountListRegionListRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorTenantId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorSubscriptionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorLoginEndpoint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorResourceUrl(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorClientId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorResourceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorAzureRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNic(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSdnConnectorNicName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_nic"
		if cur_v, ok := i["peer-nic"]; ok {
			tmp["peer_nic"] = flattenSystemSdnConnectorNicPeerNic(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if cur_v, ok := i["ip"]; ok {
			tmp["ip"] = flattenSystemSdnConnectorNicIp(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorNicName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicPeerNic(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicIp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSdnConnectorNicIpName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "private_ip"
		if cur_v, ok := i["private-ip"]; ok {
			tmp["private_ip"] = flattenSystemSdnConnectorNicIpPrivateIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if cur_v, ok := i["public-ip"]; ok {
			tmp["public_ip"] = flattenSystemSdnConnectorNicIpPublicIp(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if cur_v, ok := i["resource-group"]; ok {
			tmp["resource_group"] = flattenSystemSdnConnectorNicIpResourceGroup(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorNicIpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicIpPrivateIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicIpPublicIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorNicIpResourceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTable(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSdnConnectorRouteTableName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if cur_v, ok := i["subscription-id"]; ok {
			tmp["subscription_id"] = flattenSystemSdnConnectorRouteTableSubscriptionId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if cur_v, ok := i["resource-group"]; ok {
			tmp["resource_group"] = flattenSystemSdnConnectorRouteTableResourceGroup(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if cur_v, ok := i["route"]; ok {
			tmp["route"] = flattenSystemSdnConnectorRouteTableRoute(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorRouteTableName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableSubscriptionId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableResourceGroup(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableRoute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSdnConnectorRouteTableRouteName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if cur_v, ok := i["next-hop"]; ok {
			tmp["next_hop"] = flattenSystemSdnConnectorRouteTableRouteNextHop(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorRouteTableRouteName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRouteTableRouteNextHop(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorUserId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorCompartmentList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "compartment_id"
		if cur_v, ok := i["compartment-id"]; ok {
			tmp["compartment_id"] = flattenSystemSdnConnectorCompartmentListCompartmentId(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "compartment_id", d)
	return result
}

func flattenSystemSdnConnectorCompartmentListCompartmentId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegionList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if cur_v, ok := i["region"]; ok {
			tmp["region"] = flattenSystemSdnConnectorOciRegionListRegion(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "region", d)
	return result
}

func flattenSystemSdnConnectorOciRegionListRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorCompartmentId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciRegionType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorOciFingerprint(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorExternalIp(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSdnConnectorExternalIpName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorExternalIpName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorRoute(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSdnConnectorRouteName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorRouteName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorForwardingRule(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_name"
		if cur_v, ok := i["rule-name"]; ok {
			tmp["rule_name"] = flattenSystemSdnConnectorForwardingRuleRuleName(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if cur_v, ok := i["target"]; ok {
			tmp["target"] = flattenSystemSdnConnectorForwardingRuleTarget(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "rule_name", d)
	return result
}

func flattenSystemSdnConnectorForwardingRuleRuleName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorForwardingRuleTarget(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorGcpProjectList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemSdnConnectorGcpProjectListId(cur_v, d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gcp_zone_list"
		if cur_v, ok := i["gcp-zone-list"]; ok {
			tmp["gcp_zone_list"] = flattenSystemSdnConnectorGcpProjectListGcpZoneList(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemSdnConnectorGcpProjectListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorGcpProjectListGcpZoneList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemSdnConnectorGcpProjectListGcpZoneListName(cur_v, d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenSystemSdnConnectorGcpProjectListGcpZoneListName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorUseMetadataIam(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorGcpProject(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServiceAccount(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorDomain(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorGroupName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServerCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorServerCaCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorComputeGeneration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func flattenSystemSdnConnectorIbmRegionGen1(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorIbmRegionGen2(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorIbmRegion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemSdnConnectorUpdateInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return convintf2i(v)
}

func refreshObjectSystemSdnConnector(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error
	var b_get_all_tables bool
	if get_all_tables, ok := d.GetOk("get_all_tables"); ok {
		b_get_all_tables = get_all_tables.(string) == "true"
	} else {
		b_get_all_tables = isImportTable()
	}

	if err = d.Set("name", flattenSystemSdnConnectorName(o["name"], d, "name", sv)); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", flattenSystemSdnConnectorStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemSdnConnectorType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("proxy", flattenSystemSdnConnectorProxy(o["proxy"], d, "proxy", sv)); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("ha_status", flattenSystemSdnConnectorHaStatus(o["ha-status"], d, "ha_status", sv)); err != nil {
		if !fortiAPIPatch(o["ha-status"]) {
			return fmt.Errorf("Error reading ha_status: %v", err)
		}
	}

	if err = d.Set("verify_certificate", flattenSystemSdnConnectorVerifyCertificate(o["verify-certificate"], d, "verify_certificate", sv)); err != nil {
		if !fortiAPIPatch(o["verify-certificate"]) {
			return fmt.Errorf("Error reading verify_certificate: %v", err)
		}
	}

	if err = d.Set("server", flattenSystemSdnConnectorServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("server_list", flattenSystemSdnConnectorServerList(o["server-list"], d, "server_list", sv)); err != nil {
			if !fortiAPIPatch(o["server-list"]) {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_list"); ok {
			if err = d.Set("server_list", flattenSystemSdnConnectorServerList(o["server-list"], d, "server_list", sv)); err != nil {
				if !fortiAPIPatch(o["server-list"]) {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("server_port", flattenSystemSdnConnectorServerPort(o["server-port"], d, "server_port", sv)); err != nil {
		if !fortiAPIPatch(o["server-port"]) {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("message_server_port", flattenSystemSdnConnectorMessageServerPort(o["message-server-port"], d, "message_server_port", sv)); err != nil {
		if !fortiAPIPatch(o["message-server-port"]) {
			return fmt.Errorf("Error reading message_server_port: %v", err)
		}
	}

	if err = d.Set("username", flattenSystemSdnConnectorUsername(o["username"], d, "username", sv)); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("vcenter_server", flattenSystemSdnConnectorVcenterServer(o["vcenter-server"], d, "vcenter_server", sv)); err != nil {
		if !fortiAPIPatch(o["vcenter-server"]) {
			return fmt.Errorf("Error reading vcenter_server: %v", err)
		}
	}

	if err = d.Set("vcenter_username", flattenSystemSdnConnectorVcenterUsername(o["vcenter-username"], d, "vcenter_username", sv)); err != nil {
		if !fortiAPIPatch(o["vcenter-username"]) {
			return fmt.Errorf("Error reading vcenter_username: %v", err)
		}
	}

	if err = d.Set("region", flattenSystemSdnConnectorRegion(o["region"], d, "region", sv)); err != nil {
		if !fortiAPIPatch(o["region"]) {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("vpc_id", flattenSystemSdnConnectorVpcId(o["vpc-id"], d, "vpc_id", sv)); err != nil {
		if !fortiAPIPatch(o["vpc-id"]) {
			return fmt.Errorf("Error reading vpc_id: %v", err)
		}
	}

	if err = d.Set("alt_resource_ip", flattenSystemSdnConnectorAltResourceIp(o["alt-resource-ip"], d, "alt_resource_ip", sv)); err != nil {
		if !fortiAPIPatch(o["alt-resource-ip"]) {
			return fmt.Errorf("Error reading alt_resource_ip: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("external_account_list", flattenSystemSdnConnectorExternalAccountList(o["external-account-list"], d, "external_account_list", sv)); err != nil {
			if !fortiAPIPatch(o["external-account-list"]) {
				return fmt.Errorf("Error reading external_account_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("external_account_list"); ok {
			if err = d.Set("external_account_list", flattenSystemSdnConnectorExternalAccountList(o["external-account-list"], d, "external_account_list", sv)); err != nil {
				if !fortiAPIPatch(o["external-account-list"]) {
					return fmt.Errorf("Error reading external_account_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("tenant_id", flattenSystemSdnConnectorTenantId(o["tenant-id"], d, "tenant_id", sv)); err != nil {
		if !fortiAPIPatch(o["tenant-id"]) {
			return fmt.Errorf("Error reading tenant_id: %v", err)
		}
	}

	if err = d.Set("subscription_id", flattenSystemSdnConnectorSubscriptionId(o["subscription-id"], d, "subscription_id", sv)); err != nil {
		if !fortiAPIPatch(o["subscription-id"]) {
			return fmt.Errorf("Error reading subscription_id: %v", err)
		}
	}

	if err = d.Set("login_endpoint", flattenSystemSdnConnectorLoginEndpoint(o["login-endpoint"], d, "login_endpoint", sv)); err != nil {
		if !fortiAPIPatch(o["login-endpoint"]) {
			return fmt.Errorf("Error reading login_endpoint: %v", err)
		}
	}

	if err = d.Set("resource_url", flattenSystemSdnConnectorResourceUrl(o["resource-url"], d, "resource_url", sv)); err != nil {
		if !fortiAPIPatch(o["resource-url"]) {
			return fmt.Errorf("Error reading resource_url: %v", err)
		}
	}

	if err = d.Set("client_id", flattenSystemSdnConnectorClientId(o["client-id"], d, "client_id", sv)); err != nil {
		if !fortiAPIPatch(o["client-id"]) {
			return fmt.Errorf("Error reading client_id: %v", err)
		}
	}

	if err = d.Set("resource_group", flattenSystemSdnConnectorResourceGroup(o["resource-group"], d, "resource_group", sv)); err != nil {
		if !fortiAPIPatch(o["resource-group"]) {
			return fmt.Errorf("Error reading resource_group: %v", err)
		}
	}

	if err = d.Set("azure_region", flattenSystemSdnConnectorAzureRegion(o["azure-region"], d, "azure_region", sv)); err != nil {
		if !fortiAPIPatch(o["azure-region"]) {
			return fmt.Errorf("Error reading azure_region: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("nic", flattenSystemSdnConnectorNic(o["nic"], d, "nic", sv)); err != nil {
			if !fortiAPIPatch(o["nic"]) {
				return fmt.Errorf("Error reading nic: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("nic"); ok {
			if err = d.Set("nic", flattenSystemSdnConnectorNic(o["nic"], d, "nic", sv)); err != nil {
				if !fortiAPIPatch(o["nic"]) {
					return fmt.Errorf("Error reading nic: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("route_table", flattenSystemSdnConnectorRouteTable(o["route-table"], d, "route_table", sv)); err != nil {
			if !fortiAPIPatch(o["route-table"]) {
				return fmt.Errorf("Error reading route_table: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route_table"); ok {
			if err = d.Set("route_table", flattenSystemSdnConnectorRouteTable(o["route-table"], d, "route_table", sv)); err != nil {
				if !fortiAPIPatch(o["route-table"]) {
					return fmt.Errorf("Error reading route_table: %v", err)
				}
			}
		}
	}

	if err = d.Set("user_id", flattenSystemSdnConnectorUserId(o["user-id"], d, "user_id", sv)); err != nil {
		if !fortiAPIPatch(o["user-id"]) {
			return fmt.Errorf("Error reading user_id: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("compartment_list", flattenSystemSdnConnectorCompartmentList(o["compartment-list"], d, "compartment_list", sv)); err != nil {
			if !fortiAPIPatch(o["compartment-list"]) {
				return fmt.Errorf("Error reading compartment_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("compartment_list"); ok {
			if err = d.Set("compartment_list", flattenSystemSdnConnectorCompartmentList(o["compartment-list"], d, "compartment_list", sv)); err != nil {
				if !fortiAPIPatch(o["compartment-list"]) {
					return fmt.Errorf("Error reading compartment_list: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("oci_region_list", flattenSystemSdnConnectorOciRegionList(o["oci-region-list"], d, "oci_region_list", sv)); err != nil {
			if !fortiAPIPatch(o["oci-region-list"]) {
				return fmt.Errorf("Error reading oci_region_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("oci_region_list"); ok {
			if err = d.Set("oci_region_list", flattenSystemSdnConnectorOciRegionList(o["oci-region-list"], d, "oci_region_list", sv)); err != nil {
				if !fortiAPIPatch(o["oci-region-list"]) {
					return fmt.Errorf("Error reading oci_region_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("compartment_id", flattenSystemSdnConnectorCompartmentId(o["compartment-id"], d, "compartment_id", sv)); err != nil {
		if !fortiAPIPatch(o["compartment-id"]) {
			return fmt.Errorf("Error reading compartment_id: %v", err)
		}
	}

	if err = d.Set("oci_region", flattenSystemSdnConnectorOciRegion(o["oci-region"], d, "oci_region", sv)); err != nil {
		if !fortiAPIPatch(o["oci-region"]) {
			return fmt.Errorf("Error reading oci_region: %v", err)
		}
	}

	if err = d.Set("oci_region_type", flattenSystemSdnConnectorOciRegionType(o["oci-region-type"], d, "oci_region_type", sv)); err != nil {
		if !fortiAPIPatch(o["oci-region-type"]) {
			return fmt.Errorf("Error reading oci_region_type: %v", err)
		}
	}

	if err = d.Set("oci_cert", flattenSystemSdnConnectorOciCert(o["oci-cert"], d, "oci_cert", sv)); err != nil {
		if !fortiAPIPatch(o["oci-cert"]) {
			return fmt.Errorf("Error reading oci_cert: %v", err)
		}
	}

	if err = d.Set("oci_fingerprint", flattenSystemSdnConnectorOciFingerprint(o["oci-fingerprint"], d, "oci_fingerprint", sv)); err != nil {
		if !fortiAPIPatch(o["oci-fingerprint"]) {
			return fmt.Errorf("Error reading oci_fingerprint: %v", err)
		}
	}

	if b_get_all_tables {
		if err = d.Set("external_ip", flattenSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip", sv)); err != nil {
			if !fortiAPIPatch(o["external-ip"]) {
				return fmt.Errorf("Error reading external_ip: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("external_ip"); ok {
			if err = d.Set("external_ip", flattenSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip", sv)); err != nil {
				if !fortiAPIPatch(o["external-ip"]) {
					return fmt.Errorf("Error reading external_ip: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("route", flattenSystemSdnConnectorRoute(o["route"], d, "route", sv)); err != nil {
			if !fortiAPIPatch(o["route"]) {
				return fmt.Errorf("Error reading route: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("route"); ok {
			if err = d.Set("route", flattenSystemSdnConnectorRoute(o["route"], d, "route", sv)); err != nil {
				if !fortiAPIPatch(o["route"]) {
					return fmt.Errorf("Error reading route: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("forwarding_rule", flattenSystemSdnConnectorForwardingRule(o["forwarding-rule"], d, "forwarding_rule", sv)); err != nil {
			if !fortiAPIPatch(o["forwarding-rule"]) {
				return fmt.Errorf("Error reading forwarding_rule: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("forwarding_rule"); ok {
			if err = d.Set("forwarding_rule", flattenSystemSdnConnectorForwardingRule(o["forwarding-rule"], d, "forwarding_rule", sv)); err != nil {
				if !fortiAPIPatch(o["forwarding-rule"]) {
					return fmt.Errorf("Error reading forwarding_rule: %v", err)
				}
			}
		}
	}

	if b_get_all_tables {
		if err = d.Set("gcp_project_list", flattenSystemSdnConnectorGcpProjectList(o["gcp-project-list"], d, "gcp_project_list", sv)); err != nil {
			if !fortiAPIPatch(o["gcp-project-list"]) {
				return fmt.Errorf("Error reading gcp_project_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("gcp_project_list"); ok {
			if err = d.Set("gcp_project_list", flattenSystemSdnConnectorGcpProjectList(o["gcp-project-list"], d, "gcp_project_list", sv)); err != nil {
				if !fortiAPIPatch(o["gcp-project-list"]) {
					return fmt.Errorf("Error reading gcp_project_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("use_metadata_iam", flattenSystemSdnConnectorUseMetadataIam(o["use-metadata-iam"], d, "use_metadata_iam", sv)); err != nil {
		if !fortiAPIPatch(o["use-metadata-iam"]) {
			return fmt.Errorf("Error reading use_metadata_iam: %v", err)
		}
	}

	if err = d.Set("gcp_project", flattenSystemSdnConnectorGcpProject(o["gcp-project"], d, "gcp_project", sv)); err != nil {
		if !fortiAPIPatch(o["gcp-project"]) {
			return fmt.Errorf("Error reading gcp_project: %v", err)
		}
	}

	if err = d.Set("service_account", flattenSystemSdnConnectorServiceAccount(o["service-account"], d, "service_account", sv)); err != nil {
		if !fortiAPIPatch(o["service-account"]) {
			return fmt.Errorf("Error reading service_account: %v", err)
		}
	}

	if err = d.Set("domain", flattenSystemSdnConnectorDomain(o["domain"], d, "domain", sv)); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemSdnConnectorGroupName(o["group-name"], d, "group_name", sv)); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("server_cert", flattenSystemSdnConnectorServerCert(o["server-cert"], d, "server_cert", sv)); err != nil {
		if !fortiAPIPatch(o["server-cert"]) {
			return fmt.Errorf("Error reading server_cert: %v", err)
		}
	}

	if err = d.Set("server_ca_cert", flattenSystemSdnConnectorServerCaCert(o["server-ca-cert"], d, "server_ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["server-ca-cert"]) {
			return fmt.Errorf("Error reading server_ca_cert: %v", err)
		}
	}

	if err = d.Set("compute_generation", flattenSystemSdnConnectorComputeGeneration(o["compute-generation"], d, "compute_generation", sv)); err != nil {
		if !fortiAPIPatch(o["compute-generation"]) {
			return fmt.Errorf("Error reading compute_generation: %v", err)
		}
	}

	if err = d.Set("ibm_region_gen1", flattenSystemSdnConnectorIbmRegionGen1(o["ibm-region-gen1"], d, "ibm_region_gen1", sv)); err != nil {
		if !fortiAPIPatch(o["ibm-region-gen1"]) {
			return fmt.Errorf("Error reading ibm_region_gen1: %v", err)
		}
	}

	if err = d.Set("ibm_region_gen2", flattenSystemSdnConnectorIbmRegionGen2(o["ibm-region-gen2"], d, "ibm_region_gen2", sv)); err != nil {
		if !fortiAPIPatch(o["ibm-region-gen2"]) {
			return fmt.Errorf("Error reading ibm_region_gen2: %v", err)
		}
	}

	if err = d.Set("ibm_region", flattenSystemSdnConnectorIbmRegion(o["ibm-region"], d, "ibm_region", sv)); err != nil {
		if !fortiAPIPatch(o["ibm-region"]) {
			return fmt.Errorf("Error reading ibm_region: %v", err)
		}
	}

	if err = d.Set("update_interval", flattenSystemSdnConnectorUpdateInterval(o["update-interval"], d, "update_interval", sv)); err != nil {
		if !fortiAPIPatch(o["update-interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	return nil
}

func flattenSystemSdnConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemSdnConnectorName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorProxy(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorHaStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVerifyCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["ip"], _ = expandSystemSdnConnectorServerListIp(d, i["ip"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorServerListIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorMessageServerPort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVcenterServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVcenterUsername(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVcenterPassword(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorAccessKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSecretKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorVpcId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorAltResourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorExternalAccountList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role_arn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["role-arn"], _ = expandSystemSdnConnectorExternalAccountListRoleArn(d, i["role_arn"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["role-arn"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["external-id"], _ = expandSystemSdnConnectorExternalAccountListExternalId(d, i["external_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["external-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region_list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["region-list"], _ = expandSystemSdnConnectorExternalAccountListRegionList(d, i["region_list"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["region-list"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorExternalAccountListRoleArn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorExternalAccountListExternalId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorExternalAccountListRegionList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["region"], _ = expandSystemSdnConnectorExternalAccountListRegionListRegion(d, i["region"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorExternalAccountListRegionListRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorTenantId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSubscriptionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorLoginEndpoint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorResourceUrl(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorClientId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorClientSecret(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorResourceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorAzureRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemSdnConnectorNicName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "peer_nic"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["peer-nic"], _ = expandSystemSdnConnectorNicPeerNic(d, i["peer_nic"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["peer-nic"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ip"], _ = expandSystemSdnConnectorNicIp(d, i["ip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["ip"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorNicName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicPeerNic(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemSdnConnectorNicIpName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "private_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["private-ip"], _ = expandSystemSdnConnectorNicIpPrivateIp(d, i["private_ip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["private-ip"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["public-ip"], _ = expandSystemSdnConnectorNicIpPublicIp(d, i["public_ip"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["public-ip"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["resource-group"], _ = expandSystemSdnConnectorNicIpResourceGroup(d, i["resource_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["resource-group"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorNicIpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicIpPrivateIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicIpPublicIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorNicIpResourceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemSdnConnectorRouteTableName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["subscription-id"], _ = expandSystemSdnConnectorRouteTableSubscriptionId(d, i["subscription_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["subscription-id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["resource-group"], _ = expandSystemSdnConnectorRouteTableResourceGroup(d, i["resource_group"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["resource-group"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["route"], _ = expandSystemSdnConnectorRouteTableRoute(d, i["route"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["route"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteTableName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableSubscriptionId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableResourceGroup(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemSdnConnectorRouteTableRouteName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["next-hop"], _ = expandSystemSdnConnectorRouteTableRouteNextHop(d, i["next_hop"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["next-hop"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteTableRouteName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRouteTableRouteNextHop(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUserId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorCompartmentList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "compartment_id"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["compartment-id"], _ = expandSystemSdnConnectorCompartmentListCompartmentId(d, i["compartment_id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["compartment-id"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorCompartmentListCompartmentId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegionList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["region"], _ = expandSystemSdnConnectorOciRegionListRegion(d, i["region"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["region"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorOciRegionListRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorCompartmentId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciRegionType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorOciFingerprint(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorExternalIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemSdnConnectorExternalIpName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorExternalIpName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorRoute(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemSdnConnectorRouteName(d, i["name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["name"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorRouteName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorForwardingRule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_name"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["rule-name"], _ = expandSystemSdnConnectorForwardingRuleRuleName(d, i["rule_name"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["rule-name"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["target"], _ = expandSystemSdnConnectorForwardingRuleTarget(d, i["target"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["target"] = nil
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorForwardingRuleRuleName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorForwardingRuleTarget(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorGcpProjectList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemSdnConnectorGcpProjectListId(d, i["id"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["id"] = nil
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gcp_zone_list"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["gcp-zone-list"], _ = expandSystemSdnConnectorGcpProjectListGcpZoneList(d, i["gcp_zone_list"], pre_append, sv)
		} else if d.HasChange(pre_append) {
			tmp["gcp-zone-list"] = make([]string, 0)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorGcpProjectListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorGcpProjectListGcpZoneList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		tmp["name"], _ = expandSystemSdnConnectorGcpProjectListGcpZoneListName(d, i["name"], pre_append, sv)

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemSdnConnectorGcpProjectListGcpZoneListName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUseMetadataIam(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorGcpProject(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServiceAccount(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorKeyPasswd(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorPrivateKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorSecretToken(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorDomain(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorGroupName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServerCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorServerCaCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorApiKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorComputeGeneration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorIbmRegionGen1(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorIbmRegionGen2(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorIbmRegion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemSdnConnectorUpdateInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemSdnConnector(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("name"); ok {
		t, err := expandSystemSdnConnectorName(d, v, "name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["name"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemSdnConnectorStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemSdnConnectorType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("proxy"); ok {
		t, err := expandSystemSdnConnectorProxy(d, v, "proxy", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["proxy"] = t
		}
	} else if d.HasChange("proxy") {
		obj["proxy"] = nil
	}

	if v, ok := d.GetOk("ha_status"); ok {
		t, err := expandSystemSdnConnectorHaStatus(d, v, "ha_status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ha-status"] = t
		}
	}

	if v, ok := d.GetOk("verify_certificate"); ok {
		t, err := expandSystemSdnConnectorVerifyCertificate(d, v, "verify_certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["verify-certificate"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandSystemSdnConnectorServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	} else if d.HasChange("server") {
		obj["server"] = nil
	}

	if v, ok := d.GetOk("server_list"); ok || d.HasChange("server_list") {
		t, err := expandSystemSdnConnectorServerList(d, v, "server_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	if v, ok := d.GetOkExists("server_port"); ok {
		t, err := expandSystemSdnConnectorServerPort(d, v, "server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-port"] = t
		}
	}

	if v, ok := d.GetOkExists("message_server_port"); ok {
		t, err := expandSystemSdnConnectorMessageServerPort(d, v, "message_server_port", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["message-server-port"] = t
		}
	} else if d.HasChange("message_server_port") {
		obj["message-server-port"] = nil
	}

	if v, ok := d.GetOk("username"); ok {
		t, err := expandSystemSdnConnectorUsername(d, v, "username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["username"] = t
		}
	} else if d.HasChange("username") {
		obj["username"] = nil
	}

	if v, ok := d.GetOk("password"); ok {
		t, err := expandSystemSdnConnectorPassword(d, v, "password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["password"] = t
		}
	} else if d.HasChange("password") {
		obj["password"] = nil
	}

	if v, ok := d.GetOk("vcenter_server"); ok {
		t, err := expandSystemSdnConnectorVcenterServer(d, v, "vcenter_server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-server"] = t
		}
	} else if d.HasChange("vcenter_server") {
		obj["vcenter-server"] = nil
	}

	if v, ok := d.GetOk("vcenter_username"); ok {
		t, err := expandSystemSdnConnectorVcenterUsername(d, v, "vcenter_username", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-username"] = t
		}
	} else if d.HasChange("vcenter_username") {
		obj["vcenter-username"] = nil
	}

	if v, ok := d.GetOk("vcenter_password"); ok {
		t, err := expandSystemSdnConnectorVcenterPassword(d, v, "vcenter_password", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vcenter-password"] = t
		}
	} else if d.HasChange("vcenter_password") {
		obj["vcenter-password"] = nil
	}

	if v, ok := d.GetOk("access_key"); ok {
		t, err := expandSystemSdnConnectorAccessKey(d, v, "access_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-key"] = t
		}
	} else if d.HasChange("access_key") {
		obj["access-key"] = nil
	}

	if v, ok := d.GetOk("secret_key"); ok {
		t, err := expandSystemSdnConnectorSecretKey(d, v, "secret_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-key"] = t
		}
	} else if d.HasChange("secret_key") {
		obj["secret-key"] = nil
	}

	if v, ok := d.GetOk("region"); ok {
		t, err := expandSystemSdnConnectorRegion(d, v, "region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["region"] = t
		}
	} else if d.HasChange("region") {
		obj["region"] = nil
	}

	if v, ok := d.GetOk("vpc_id"); ok {
		t, err := expandSystemSdnConnectorVpcId(d, v, "vpc_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vpc-id"] = t
		}
	} else if d.HasChange("vpc_id") {
		obj["vpc-id"] = nil
	}

	if v, ok := d.GetOk("alt_resource_ip"); ok {
		t, err := expandSystemSdnConnectorAltResourceIp(d, v, "alt_resource_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["alt-resource-ip"] = t
		}
	}

	if v, ok := d.GetOk("external_account_list"); ok || d.HasChange("external_account_list") {
		t, err := expandSystemSdnConnectorExternalAccountList(d, v, "external_account_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-account-list"] = t
		}
	}

	if v, ok := d.GetOk("tenant_id"); ok {
		t, err := expandSystemSdnConnectorTenantId(d, v, "tenant_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["tenant-id"] = t
		}
	} else if d.HasChange("tenant_id") {
		obj["tenant-id"] = nil
	}

	if v, ok := d.GetOk("subscription_id"); ok {
		t, err := expandSystemSdnConnectorSubscriptionId(d, v, "subscription_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["subscription-id"] = t
		}
	} else if d.HasChange("subscription_id") {
		obj["subscription-id"] = nil
	}

	if v, ok := d.GetOk("login_endpoint"); ok {
		t, err := expandSystemSdnConnectorLoginEndpoint(d, v, "login_endpoint", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["login-endpoint"] = t
		}
	} else if d.HasChange("login_endpoint") {
		obj["login-endpoint"] = nil
	}

	if v, ok := d.GetOk("resource_url"); ok {
		t, err := expandSystemSdnConnectorResourceUrl(d, v, "resource_url", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-url"] = t
		}
	} else if d.HasChange("resource_url") {
		obj["resource-url"] = nil
	}

	if v, ok := d.GetOk("client_id"); ok {
		t, err := expandSystemSdnConnectorClientId(d, v, "client_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-id"] = t
		}
	} else if d.HasChange("client_id") {
		obj["client-id"] = nil
	}

	if v, ok := d.GetOk("client_secret"); ok {
		t, err := expandSystemSdnConnectorClientSecret(d, v, "client_secret", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["client-secret"] = t
		}
	} else if d.HasChange("client_secret") {
		obj["client-secret"] = nil
	}

	if v, ok := d.GetOk("resource_group"); ok {
		t, err := expandSystemSdnConnectorResourceGroup(d, v, "resource_group", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["resource-group"] = t
		}
	} else if d.HasChange("resource_group") {
		obj["resource-group"] = nil
	}

	if v, ok := d.GetOk("azure_region"); ok {
		t, err := expandSystemSdnConnectorAzureRegion(d, v, "azure_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["azure-region"] = t
		}
	}

	if v, ok := d.GetOk("nic"); ok || d.HasChange("nic") {
		t, err := expandSystemSdnConnectorNic(d, v, "nic", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["nic"] = t
		}
	}

	if v, ok := d.GetOk("route_table"); ok || d.HasChange("route_table") {
		t, err := expandSystemSdnConnectorRouteTable(d, v, "route_table", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route-table"] = t
		}
	}

	if v, ok := d.GetOk("user_id"); ok {
		t, err := expandSystemSdnConnectorUserId(d, v, "user_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["user-id"] = t
		}
	} else if d.HasChange("user_id") {
		obj["user-id"] = nil
	}

	if v, ok := d.GetOk("compartment_list"); ok || d.HasChange("compartment_list") {
		t, err := expandSystemSdnConnectorCompartmentList(d, v, "compartment_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compartment-list"] = t
		}
	}

	if v, ok := d.GetOk("oci_region_list"); ok || d.HasChange("oci_region_list") {
		t, err := expandSystemSdnConnectorOciRegionList(d, v, "oci_region_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region-list"] = t
		}
	}

	if v, ok := d.GetOk("compartment_id"); ok {
		t, err := expandSystemSdnConnectorCompartmentId(d, v, "compartment_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compartment-id"] = t
		}
	} else if d.HasChange("compartment_id") {
		obj["compartment-id"] = nil
	}

	if v, ok := d.GetOk("oci_region"); ok {
		t, err := expandSystemSdnConnectorOciRegion(d, v, "oci_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region"] = t
		}
	} else if d.HasChange("oci_region") {
		obj["oci-region"] = nil
	}

	if v, ok := d.GetOk("oci_region_type"); ok {
		t, err := expandSystemSdnConnectorOciRegionType(d, v, "oci_region_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-region-type"] = t
		}
	}

	if v, ok := d.GetOk("oci_cert"); ok {
		t, err := expandSystemSdnConnectorOciCert(d, v, "oci_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-cert"] = t
		}
	} else if d.HasChange("oci_cert") {
		obj["oci-cert"] = nil
	}

	if v, ok := d.GetOk("oci_fingerprint"); ok {
		t, err := expandSystemSdnConnectorOciFingerprint(d, v, "oci_fingerprint", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["oci-fingerprint"] = t
		}
	} else if d.HasChange("oci_fingerprint") {
		obj["oci-fingerprint"] = nil
	}

	if v, ok := d.GetOk("external_ip"); ok || d.HasChange("external_ip") {
		t, err := expandSystemSdnConnectorExternalIp(d, v, "external_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["external-ip"] = t
		}
	}

	if v, ok := d.GetOk("route"); ok || d.HasChange("route") {
		t, err := expandSystemSdnConnectorRoute(d, v, "route", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["route"] = t
		}
	}

	if v, ok := d.GetOk("forwarding_rule"); ok || d.HasChange("forwarding_rule") {
		t, err := expandSystemSdnConnectorForwardingRule(d, v, "forwarding_rule", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["forwarding-rule"] = t
		}
	}

	if v, ok := d.GetOk("gcp_project_list"); ok || d.HasChange("gcp_project_list") {
		t, err := expandSystemSdnConnectorGcpProjectList(d, v, "gcp_project_list", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-project-list"] = t
		}
	}

	if v, ok := d.GetOk("use_metadata_iam"); ok {
		t, err := expandSystemSdnConnectorUseMetadataIam(d, v, "use_metadata_iam", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-metadata-iam"] = t
		}
	}

	if v, ok := d.GetOk("gcp_project"); ok {
		t, err := expandSystemSdnConnectorGcpProject(d, v, "gcp_project", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["gcp-project"] = t
		}
	} else if d.HasChange("gcp_project") {
		obj["gcp-project"] = nil
	}

	if v, ok := d.GetOk("service_account"); ok {
		t, err := expandSystemSdnConnectorServiceAccount(d, v, "service_account", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["service-account"] = t
		}
	} else if d.HasChange("service_account") {
		obj["service-account"] = nil
	}

	if v, ok := d.GetOk("key_passwd"); ok {
		t, err := expandSystemSdnConnectorKeyPasswd(d, v, "key_passwd", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key-passwd"] = t
		}
	} else if d.HasChange("key_passwd") {
		obj["key-passwd"] = nil
	}

	if v, ok := d.GetOk("private_key"); ok {
		t, err := expandSystemSdnConnectorPrivateKey(d, v, "private_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["private-key"] = t
		}
	} else if d.HasChange("private_key") {
		obj["private-key"] = nil
	}

	if v, ok := d.GetOk("secret_token"); ok {
		t, err := expandSystemSdnConnectorSecretToken(d, v, "secret_token", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["secret-token"] = t
		}
	} else if d.HasChange("secret_token") {
		obj["secret-token"] = nil
	}

	if v, ok := d.GetOk("domain"); ok {
		t, err := expandSystemSdnConnectorDomain(d, v, "domain", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["domain"] = t
		}
	} else if d.HasChange("domain") {
		obj["domain"] = nil
	}

	if v, ok := d.GetOk("group_name"); ok {
		t, err := expandSystemSdnConnectorGroupName(d, v, "group_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	} else if d.HasChange("group_name") {
		obj["group-name"] = nil
	}

	if v, ok := d.GetOk("server_cert"); ok {
		t, err := expandSystemSdnConnectorServerCert(d, v, "server_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-cert"] = t
		}
	} else if d.HasChange("server_cert") {
		obj["server-cert"] = nil
	}

	if v, ok := d.GetOk("server_ca_cert"); ok {
		t, err := expandSystemSdnConnectorServerCaCert(d, v, "server_ca_cert", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-ca-cert"] = t
		}
	} else if d.HasChange("server_ca_cert") {
		obj["server-ca-cert"] = nil
	}

	if v, ok := d.GetOk("api_key"); ok {
		t, err := expandSystemSdnConnectorApiKey(d, v, "api_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["api-key"] = t
		}
	} else if d.HasChange("api_key") {
		obj["api-key"] = nil
	}

	if v, ok := d.GetOk("compute_generation"); ok {
		t, err := expandSystemSdnConnectorComputeGeneration(d, v, "compute_generation", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["compute-generation"] = t
		}
	}

	if v, ok := d.GetOk("ibm_region_gen1"); ok {
		t, err := expandSystemSdnConnectorIbmRegionGen1(d, v, "ibm_region_gen1", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region-gen1"] = t
		}
	} else if d.HasChange("ibm_region_gen1") {
		obj["ibm-region-gen1"] = nil
	}

	if v, ok := d.GetOk("ibm_region_gen2"); ok {
		t, err := expandSystemSdnConnectorIbmRegionGen2(d, v, "ibm_region_gen2", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region-gen2"] = t
		}
	} else if d.HasChange("ibm_region_gen2") {
		obj["ibm-region-gen2"] = nil
	}

	if v, ok := d.GetOk("ibm_region"); ok {
		t, err := expandSystemSdnConnectorIbmRegion(d, v, "ibm_region", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ibm-region"] = t
		}
	}

	if v, ok := d.GetOkExists("update_interval"); ok {
		t, err := expandSystemSdnConnectorUpdateInterval(d, v, "update_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["update-interval"] = t
		}
	}

	return &obj, nil
}
