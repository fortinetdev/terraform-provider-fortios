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

func dataSourceSystemSdnConnector() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemSdnConnectorRead,
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"proxy": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ha_status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"verify_certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"server_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"vcenter_server": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_username": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vcenter_password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"access_key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"secret_key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"vpc_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"alt_resource_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_account_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"role_arn": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"external_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"region_list": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"region": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"tenant_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"subscription_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"login_endpoint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"resource_url": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"client_secret": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"resource_group": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"azure_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"nic": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ip": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"public_ip": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"resource_group": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"route_table": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"subscription_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"resource_group": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"route": &schema.Schema{
							Type:     schema.TypeList,
							Computed: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{
									"name": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
									"next_hop": &schema.Schema{
										Type:     schema.TypeString,
										Computed: true,
									},
								},
							},
						},
					},
				},
			},
			"user_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"compartment_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"compartment_id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"oci_region_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"region": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"compartment_id": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_region_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"oci_fingerprint": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"external_ip": &schema.Schema{
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
			"route": &schema.Schema{
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
			"forwarding_rule": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"rule_name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"target": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
			"gcp_project_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"gcp_zone_list": &schema.Schema{
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
				},
			},
			"use_metadata_iam": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"gcp_project": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"service_account": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"key_passwd": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"private_key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"secret_token": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"domain": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"server_ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"api_key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"compute_generation": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"ibm_region_gen1": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ibm_region_gen2": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"ibm_region": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"update_interval": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
		},
	}
}

func dataSourceSystemSdnConnectorRead(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing SystemSdnConnector: type error")
	}

	o, err := c.ReadSystemSdnConnector(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemSdnConnector: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemSdnConnector(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemSdnConnector from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemSdnConnectorName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorProxy(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorHaStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorVerifyCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServerList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemSdnConnectorServerListIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorServerListIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServerPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorVcenterServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorVcenterUsername(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorVcenterPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorAccessKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSecretKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorVpcId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorAltResourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorExternalAccountList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "role_arn"
		if _, ok := i["role-arn"]; ok {
			tmp["role_arn"] = dataSourceFlattenSystemSdnConnectorExternalAccountListRoleArn(i["role-arn"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "external_id"
		if _, ok := i["external-id"]; ok {
			tmp["external_id"] = dataSourceFlattenSystemSdnConnectorExternalAccountListExternalId(i["external-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region_list"
		if _, ok := i["region-list"]; ok {
			tmp["region_list"] = dataSourceFlattenSystemSdnConnectorExternalAccountListRegionList(i["region-list"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorExternalAccountListRoleArn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorExternalAccountListExternalId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorExternalAccountListRegionList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if _, ok := i["region"]; ok {
			tmp["region"] = dataSourceFlattenSystemSdnConnectorExternalAccountListRegionListRegion(i["region"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorExternalAccountListRegionListRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorTenantId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSubscriptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorLoginEndpoint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorResourceUrl(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorClientId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorClientSecret(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorAzureRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorNic(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorNicName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ip"
		if _, ok := i["ip"]; ok {
			tmp["ip"] = dataSourceFlattenSystemSdnConnectorNicIp(i["ip"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorNicName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorNicIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorNicIpName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "public_ip"
		if _, ok := i["public-ip"]; ok {
			tmp["public_ip"] = dataSourceFlattenSystemSdnConnectorNicIpPublicIp(i["public-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {
			tmp["resource_group"] = dataSourceFlattenSystemSdnConnectorNicIpResourceGroup(i["resource-group"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorNicIpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorNicIpPublicIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorNicIpResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTable(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorRouteTableName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "subscription_id"
		if _, ok := i["subscription-id"]; ok {
			tmp["subscription_id"] = dataSourceFlattenSystemSdnConnectorRouteTableSubscriptionId(i["subscription-id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "resource_group"
		if _, ok := i["resource-group"]; ok {
			tmp["resource_group"] = dataSourceFlattenSystemSdnConnectorRouteTableResourceGroup(i["resource-group"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "route"
		if _, ok := i["route"]; ok {
			tmp["route"] = dataSourceFlattenSystemSdnConnectorRouteTableRoute(i["route"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorRouteTableName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTableSubscriptionId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTableResourceGroup(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTableRoute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorRouteTableRouteName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "next_hop"
		if _, ok := i["next-hop"]; ok {
			tmp["next_hop"] = dataSourceFlattenSystemSdnConnectorRouteTableRouteNextHop(i["next-hop"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorRouteTableRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRouteTableRouteNextHop(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUserId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorCompartmentList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "compartment_id"
		if _, ok := i["compartment-id"]; ok {
			tmp["compartment_id"] = dataSourceFlattenSystemSdnConnectorCompartmentListCompartmentId(i["compartment-id"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorCompartmentListCompartmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciRegionList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "region"
		if _, ok := i["region"]; ok {
			tmp["region"] = dataSourceFlattenSystemSdnConnectorOciRegionListRegion(i["region"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorOciRegionListRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorCompartmentId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciRegionType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorOciFingerprint(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorExternalIp(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorExternalIpName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorExternalIpName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorRoute(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorRouteName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorRouteName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorForwardingRule(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "rule_name"
		if _, ok := i["rule-name"]; ok {
			tmp["rule_name"] = dataSourceFlattenSystemSdnConnectorForwardingRuleRuleName(i["rule-name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "target"
		if _, ok := i["target"]; ok {
			tmp["target"] = dataSourceFlattenSystemSdnConnectorForwardingRuleTarget(i["target"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorForwardingRuleRuleName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorForwardingRuleTarget(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorGcpProjectList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = dataSourceFlattenSystemSdnConnectorGcpProjectListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "gcp_zone_list"
		if _, ok := i["gcp-zone-list"]; ok {
			tmp["gcp_zone_list"] = dataSourceFlattenSystemSdnConnectorGcpProjectListGcpZoneList(i["gcp-zone-list"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorGcpProjectListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorGcpProjectListGcpZoneList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemSdnConnectorGcpProjectListGcpZoneListName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemSdnConnectorGcpProjectListGcpZoneListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUseMetadataIam(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorGcpProject(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServiceAccount(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorKeyPasswd(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorPrivateKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorSecretToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorDomain(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServerCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorServerCaCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorApiKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorComputeGeneration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorIbmRegionGen1(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorIbmRegionGen2(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorIbmRegion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemSdnConnectorUpdateInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemSdnConnector(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("name", dataSourceFlattenSystemSdnConnectorName(o["name"], d, "name")); err != nil {
		if !fortiAPIPatch(o["name"]) {
			return fmt.Errorf("Error reading name: %v", err)
		}
	}

	if err = d.Set("status", dataSourceFlattenSystemSdnConnectorStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("type", dataSourceFlattenSystemSdnConnectorType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("proxy", dataSourceFlattenSystemSdnConnectorProxy(o["proxy"], d, "proxy")); err != nil {
		if !fortiAPIPatch(o["proxy"]) {
			return fmt.Errorf("Error reading proxy: %v", err)
		}
	}

	if err = d.Set("ha_status", dataSourceFlattenSystemSdnConnectorHaStatus(o["ha-status"], d, "ha_status")); err != nil {
		if !fortiAPIPatch(o["ha-status"]) {
			return fmt.Errorf("Error reading ha_status: %v", err)
		}
	}

	if err = d.Set("verify_certificate", dataSourceFlattenSystemSdnConnectorVerifyCertificate(o["verify-certificate"], d, "verify_certificate")); err != nil {
		if !fortiAPIPatch(o["verify-certificate"]) {
			return fmt.Errorf("Error reading verify_certificate: %v", err)
		}
	}

	if err = d.Set("server", dataSourceFlattenSystemSdnConnectorServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("server_list", dataSourceFlattenSystemSdnConnectorServerList(o["server-list"], d, "server_list")); err != nil {
		if !fortiAPIPatch(o["server-list"]) {
			return fmt.Errorf("Error reading server_list: %v", err)
		}
	}

	if err = d.Set("server_port", dataSourceFlattenSystemSdnConnectorServerPort(o["server-port"], d, "server_port")); err != nil {
		if !fortiAPIPatch(o["server-port"]) {
			return fmt.Errorf("Error reading server_port: %v", err)
		}
	}

	if err = d.Set("username", dataSourceFlattenSystemSdnConnectorUsername(o["username"], d, "username")); err != nil {
		if !fortiAPIPatch(o["username"]) {
			return fmt.Errorf("Error reading username: %v", err)
		}
	}

	if err = d.Set("vcenter_server", dataSourceFlattenSystemSdnConnectorVcenterServer(o["vcenter-server"], d, "vcenter_server")); err != nil {
		if !fortiAPIPatch(o["vcenter-server"]) {
			return fmt.Errorf("Error reading vcenter_server: %v", err)
		}
	}

	if err = d.Set("vcenter_username", dataSourceFlattenSystemSdnConnectorVcenterUsername(o["vcenter-username"], d, "vcenter_username")); err != nil {
		if !fortiAPIPatch(o["vcenter-username"]) {
			return fmt.Errorf("Error reading vcenter_username: %v", err)
		}
	}

	if err = d.Set("region", dataSourceFlattenSystemSdnConnectorRegion(o["region"], d, "region")); err != nil {
		if !fortiAPIPatch(o["region"]) {
			return fmt.Errorf("Error reading region: %v", err)
		}
	}

	if err = d.Set("vpc_id", dataSourceFlattenSystemSdnConnectorVpcId(o["vpc-id"], d, "vpc_id")); err != nil {
		if !fortiAPIPatch(o["vpc-id"]) {
			return fmt.Errorf("Error reading vpc_id: %v", err)
		}
	}

	if err = d.Set("alt_resource_ip", dataSourceFlattenSystemSdnConnectorAltResourceIp(o["alt-resource-ip"], d, "alt_resource_ip")); err != nil {
		if !fortiAPIPatch(o["alt-resource-ip"]) {
			return fmt.Errorf("Error reading alt_resource_ip: %v", err)
		}
	}

	if err = d.Set("external_account_list", dataSourceFlattenSystemSdnConnectorExternalAccountList(o["external-account-list"], d, "external_account_list")); err != nil {
		if !fortiAPIPatch(o["external-account-list"]) {
			return fmt.Errorf("Error reading external_account_list: %v", err)
		}
	}

	if err = d.Set("tenant_id", dataSourceFlattenSystemSdnConnectorTenantId(o["tenant-id"], d, "tenant_id")); err != nil {
		if !fortiAPIPatch(o["tenant-id"]) {
			return fmt.Errorf("Error reading tenant_id: %v", err)
		}
	}

	if err = d.Set("subscription_id", dataSourceFlattenSystemSdnConnectorSubscriptionId(o["subscription-id"], d, "subscription_id")); err != nil {
		if !fortiAPIPatch(o["subscription-id"]) {
			return fmt.Errorf("Error reading subscription_id: %v", err)
		}
	}

	if err = d.Set("login_endpoint", dataSourceFlattenSystemSdnConnectorLoginEndpoint(o["login-endpoint"], d, "login_endpoint")); err != nil {
		if !fortiAPIPatch(o["login-endpoint"]) {
			return fmt.Errorf("Error reading login_endpoint: %v", err)
		}
	}

	if err = d.Set("resource_url", dataSourceFlattenSystemSdnConnectorResourceUrl(o["resource-url"], d, "resource_url")); err != nil {
		if !fortiAPIPatch(o["resource-url"]) {
			return fmt.Errorf("Error reading resource_url: %v", err)
		}
	}

	if err = d.Set("client_id", dataSourceFlattenSystemSdnConnectorClientId(o["client-id"], d, "client_id")); err != nil {
		if !fortiAPIPatch(o["client-id"]) {
			return fmt.Errorf("Error reading client_id: %v", err)
		}
	}

	if err = d.Set("resource_group", dataSourceFlattenSystemSdnConnectorResourceGroup(o["resource-group"], d, "resource_group")); err != nil {
		if !fortiAPIPatch(o["resource-group"]) {
			return fmt.Errorf("Error reading resource_group: %v", err)
		}
	}

	if err = d.Set("azure_region", dataSourceFlattenSystemSdnConnectorAzureRegion(o["azure-region"], d, "azure_region")); err != nil {
		if !fortiAPIPatch(o["azure-region"]) {
			return fmt.Errorf("Error reading azure_region: %v", err)
		}
	}

	if err = d.Set("nic", dataSourceFlattenSystemSdnConnectorNic(o["nic"], d, "nic")); err != nil {
		if !fortiAPIPatch(o["nic"]) {
			return fmt.Errorf("Error reading nic: %v", err)
		}
	}

	if err = d.Set("route_table", dataSourceFlattenSystemSdnConnectorRouteTable(o["route-table"], d, "route_table")); err != nil {
		if !fortiAPIPatch(o["route-table"]) {
			return fmt.Errorf("Error reading route_table: %v", err)
		}
	}

	if err = d.Set("user_id", dataSourceFlattenSystemSdnConnectorUserId(o["user-id"], d, "user_id")); err != nil {
		if !fortiAPIPatch(o["user-id"]) {
			return fmt.Errorf("Error reading user_id: %v", err)
		}
	}

	if err = d.Set("compartment_list", dataSourceFlattenSystemSdnConnectorCompartmentList(o["compartment-list"], d, "compartment_list")); err != nil {
		if !fortiAPIPatch(o["compartment-list"]) {
			return fmt.Errorf("Error reading compartment_list: %v", err)
		}
	}

	if err = d.Set("oci_region_list", dataSourceFlattenSystemSdnConnectorOciRegionList(o["oci-region-list"], d, "oci_region_list")); err != nil {
		if !fortiAPIPatch(o["oci-region-list"]) {
			return fmt.Errorf("Error reading oci_region_list: %v", err)
		}
	}

	if err = d.Set("compartment_id", dataSourceFlattenSystemSdnConnectorCompartmentId(o["compartment-id"], d, "compartment_id")); err != nil {
		if !fortiAPIPatch(o["compartment-id"]) {
			return fmt.Errorf("Error reading compartment_id: %v", err)
		}
	}

	if err = d.Set("oci_region", dataSourceFlattenSystemSdnConnectorOciRegion(o["oci-region"], d, "oci_region")); err != nil {
		if !fortiAPIPatch(o["oci-region"]) {
			return fmt.Errorf("Error reading oci_region: %v", err)
		}
	}

	if err = d.Set("oci_region_type", dataSourceFlattenSystemSdnConnectorOciRegionType(o["oci-region-type"], d, "oci_region_type")); err != nil {
		if !fortiAPIPatch(o["oci-region-type"]) {
			return fmt.Errorf("Error reading oci_region_type: %v", err)
		}
	}

	if err = d.Set("oci_cert", dataSourceFlattenSystemSdnConnectorOciCert(o["oci-cert"], d, "oci_cert")); err != nil {
		if !fortiAPIPatch(o["oci-cert"]) {
			return fmt.Errorf("Error reading oci_cert: %v", err)
		}
	}

	if err = d.Set("oci_fingerprint", dataSourceFlattenSystemSdnConnectorOciFingerprint(o["oci-fingerprint"], d, "oci_fingerprint")); err != nil {
		if !fortiAPIPatch(o["oci-fingerprint"]) {
			return fmt.Errorf("Error reading oci_fingerprint: %v", err)
		}
	}

	if err = d.Set("external_ip", dataSourceFlattenSystemSdnConnectorExternalIp(o["external-ip"], d, "external_ip")); err != nil {
		if !fortiAPIPatch(o["external-ip"]) {
			return fmt.Errorf("Error reading external_ip: %v", err)
		}
	}

	if err = d.Set("route", dataSourceFlattenSystemSdnConnectorRoute(o["route"], d, "route")); err != nil {
		if !fortiAPIPatch(o["route"]) {
			return fmt.Errorf("Error reading route: %v", err)
		}
	}

	if err = d.Set("forwarding_rule", dataSourceFlattenSystemSdnConnectorForwardingRule(o["forwarding-rule"], d, "forwarding_rule")); err != nil {
		if !fortiAPIPatch(o["forwarding-rule"]) {
			return fmt.Errorf("Error reading forwarding_rule: %v", err)
		}
	}

	if err = d.Set("gcp_project_list", dataSourceFlattenSystemSdnConnectorGcpProjectList(o["gcp-project-list"], d, "gcp_project_list")); err != nil {
		if !fortiAPIPatch(o["gcp-project-list"]) {
			return fmt.Errorf("Error reading gcp_project_list: %v", err)
		}
	}

	if err = d.Set("use_metadata_iam", dataSourceFlattenSystemSdnConnectorUseMetadataIam(o["use-metadata-iam"], d, "use_metadata_iam")); err != nil {
		if !fortiAPIPatch(o["use-metadata-iam"]) {
			return fmt.Errorf("Error reading use_metadata_iam: %v", err)
		}
	}

	if err = d.Set("gcp_project", dataSourceFlattenSystemSdnConnectorGcpProject(o["gcp-project"], d, "gcp_project")); err != nil {
		if !fortiAPIPatch(o["gcp-project"]) {
			return fmt.Errorf("Error reading gcp_project: %v", err)
		}
	}

	if err = d.Set("service_account", dataSourceFlattenSystemSdnConnectorServiceAccount(o["service-account"], d, "service_account")); err != nil {
		if !fortiAPIPatch(o["service-account"]) {
			return fmt.Errorf("Error reading service_account: %v", err)
		}
	}

	if err = d.Set("domain", dataSourceFlattenSystemSdnConnectorDomain(o["domain"], d, "domain")); err != nil {
		if !fortiAPIPatch(o["domain"]) {
			return fmt.Errorf("Error reading domain: %v", err)
		}
	}

	if err = d.Set("group_name", dataSourceFlattenSystemSdnConnectorGroupName(o["group-name"], d, "group_name")); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("server_cert", dataSourceFlattenSystemSdnConnectorServerCert(o["server-cert"], d, "server_cert")); err != nil {
		if !fortiAPIPatch(o["server-cert"]) {
			return fmt.Errorf("Error reading server_cert: %v", err)
		}
	}

	if err = d.Set("server_ca_cert", dataSourceFlattenSystemSdnConnectorServerCaCert(o["server-ca-cert"], d, "server_ca_cert")); err != nil {
		if !fortiAPIPatch(o["server-ca-cert"]) {
			return fmt.Errorf("Error reading server_ca_cert: %v", err)
		}
	}

	if err = d.Set("compute_generation", dataSourceFlattenSystemSdnConnectorComputeGeneration(o["compute-generation"], d, "compute_generation")); err != nil {
		if !fortiAPIPatch(o["compute-generation"]) {
			return fmt.Errorf("Error reading compute_generation: %v", err)
		}
	}

	if err = d.Set("ibm_region_gen1", dataSourceFlattenSystemSdnConnectorIbmRegionGen1(o["ibm-region-gen1"], d, "ibm_region_gen1")); err != nil {
		if !fortiAPIPatch(o["ibm-region-gen1"]) {
			return fmt.Errorf("Error reading ibm_region_gen1: %v", err)
		}
	}

	if err = d.Set("ibm_region_gen2", dataSourceFlattenSystemSdnConnectorIbmRegionGen2(o["ibm-region-gen2"], d, "ibm_region_gen2")); err != nil {
		if !fortiAPIPatch(o["ibm-region-gen2"]) {
			return fmt.Errorf("Error reading ibm_region_gen2: %v", err)
		}
	}

	if err = d.Set("ibm_region", dataSourceFlattenSystemSdnConnectorIbmRegion(o["ibm-region"], d, "ibm_region")); err != nil {
		if !fortiAPIPatch(o["ibm-region"]) {
			return fmt.Errorf("Error reading ibm_region: %v", err)
		}
	}

	if err = d.Set("update_interval", dataSourceFlattenSystemSdnConnectorUpdateInterval(o["update-interval"], d, "update_interval")); err != nil {
		if !fortiAPIPatch(o["update-interval"]) {
			return fmt.Errorf("Error reading update_interval: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemSdnConnectorFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
