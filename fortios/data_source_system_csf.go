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

func dataSourceSystemCsf() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSystemCsfRead,
		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},

			"status": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"upstream": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"upstream_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"upstream_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"group_name": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"group_password": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"accept_auth_by_cert": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"log_unification": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"configuration_sync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fabric_object_unification": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"saml_configuration_sync": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_ip": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"management_port": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"authorization_request_type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fabric_workers": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"downstream_access": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"downstream_accprofile": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"fixed_key": &schema.Schema{
				Type:      schema.TypeString,
				Sensitive: true,
				Computed:  true,
			},
			"trusted_list": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"authorization_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"certificate": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"action": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"ha_members": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"downstream_authorization": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"index": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
					},
				},
			},
			"fabric_connector": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"serial": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"accprofile": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"configuration_write_access": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"vdom": &schema.Schema{
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
			"forticloud_account_enforcement": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_mgmt": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
			},
			"file_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"file_quota_warning": &schema.Schema{
				Type:     schema.TypeInt,
				Computed: true,
			},
			"fabric_device": &schema.Schema{
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"device_ip": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"https_port": &schema.Schema{
							Type:     schema.TypeInt,
							Computed: true,
						},
						"access_token": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
						"device_type": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"login": &schema.Schema{
							Type:     schema.TypeString,
							Computed: true,
						},
						"password": &schema.Schema{
							Type:      schema.TypeString,
							Sensitive: true,
							Computed:  true,
						},
					},
				},
			},
		},
	}
}

func dataSourceSystemCsfRead(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	mkey := "SystemCsf"

	o, err := c.ReadSystemCsf(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error describing SystemCsf: %v", err)
	}

	if o == nil {
		d.SetId("")
		return nil
	}

	err = dataSourceRefreshObjectSystemCsf(d, o)
	if err != nil {
		return fmt.Errorf("Error describing SystemCsf from API: %v", err)
	}

	d.SetId(mkey)

	return nil
}

func dataSourceFlattenSystemCsfStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfUpstream(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfUpstreamIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfUpstreamPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfGroupPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfAcceptAuthByCert(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfLogUnification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfConfigurationSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricObjectUnification(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfSamlConfigurationSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfManagementIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfManagementPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfAuthorizationRequestType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricWorkers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfDownstreamAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfDownstreamAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFixedKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfTrustedList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemCsfTrustedListName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "authorization_type"
		if _, ok := i["authorization-type"]; ok {
			tmp["authorization_type"] = dataSourceFlattenSystemCsfTrustedListAuthorizationType(i["authorization-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			tmp["serial"] = dataSourceFlattenSystemCsfTrustedListSerial(i["serial"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "certificate"
		if _, ok := i["certificate"]; ok {
			tmp["certificate"] = dataSourceFlattenSystemCsfTrustedListCertificate(i["certificate"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = dataSourceFlattenSystemCsfTrustedListAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := i["ha-members"]; ok {
			tmp["ha_members"] = dataSourceFlattenSystemCsfTrustedListHaMembers(i["ha-members"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := i["downstream-authorization"]; ok {
			tmp["downstream_authorization"] = dataSourceFlattenSystemCsfTrustedListDownstreamAuthorization(i["downstream-authorization"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "index"
		if _, ok := i["index"]; ok {
			tmp["index"] = dataSourceFlattenSystemCsfTrustedListIndex(i["index"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemCsfTrustedListName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfTrustedListAuthorizationType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfTrustedListSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfTrustedListCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfTrustedListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfTrustedListHaMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfTrustedListDownstreamAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfTrustedListIndex(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricConnector(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := i["serial"]; ok {
			tmp["serial"] = dataSourceFlattenSystemCsfFabricConnectorSerial(i["serial"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "accprofile"
		if _, ok := i["accprofile"]; ok {
			tmp["accprofile"] = dataSourceFlattenSystemCsfFabricConnectorAccprofile(i["accprofile"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "configuration_write_access"
		if _, ok := i["configuration-write-access"]; ok {
			tmp["configuration_write_access"] = dataSourceFlattenSystemCsfFabricConnectorConfigurationWriteAccess(i["configuration-write-access"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "vdom"
		if _, ok := i["vdom"]; ok {
			tmp["vdom"] = dataSourceFlattenSystemCsfFabricConnectorVdom(i["vdom"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemCsfFabricConnectorSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricConnectorAccprofile(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricConnectorConfigurationWriteAccess(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricConnectorVdom(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemCsfFabricConnectorVdomName(i["name"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemCsfFabricConnectorVdomName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfForticloudAccountEnforcement(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFileMgmt(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFileQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFileQuotaWarning(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricDevice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = dataSourceFlattenSystemCsfFabricDeviceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_ip"
		if _, ok := i["device-ip"]; ok {
			tmp["device_ip"] = dataSourceFlattenSystemCsfFabricDeviceDeviceIp(i["device-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "https_port"
		if _, ok := i["https-port"]; ok {
			tmp["https_port"] = dataSourceFlattenSystemCsfFabricDeviceHttpsPort(i["https-port"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "access_token"
		if _, ok := i["access-token"]; ok {
			tmp["access_token"] = dataSourceFlattenSystemCsfFabricDeviceAccessToken(i["access-token"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["access_token"] = c
			}
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := i["device-type"]; ok {
			tmp["device_type"] = dataSourceFlattenSystemCsfFabricDeviceDeviceType(i["device-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "login"
		if _, ok := i["login"]; ok {
			tmp["login"] = dataSourceFlattenSystemCsfFabricDeviceLogin(i["login"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			tmp["password"] = dataSourceFlattenSystemCsfFabricDevicePassword(i["password"], d, pre_append)
			c := d.Get(pre_append).(string)
			if c != "" {
				tmp["password"] = c
			}
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func dataSourceFlattenSystemCsfFabricDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricDeviceDeviceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricDeviceHttpsPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricDeviceAccessToken(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricDeviceDeviceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricDeviceLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceFlattenSystemCsfFabricDevicePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func dataSourceRefreshObjectSystemCsf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", dataSourceFlattenSystemCsfStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upstream", dataSourceFlattenSystemCsfUpstream(o["upstream"], d, "upstream")); err != nil {
		if !fortiAPIPatch(o["upstream"]) {
			return fmt.Errorf("Error reading upstream: %v", err)
		}
	}

	if err = d.Set("upstream_ip", dataSourceFlattenSystemCsfUpstreamIp(o["upstream-ip"], d, "upstream_ip")); err != nil {
		if !fortiAPIPatch(o["upstream-ip"]) {
			return fmt.Errorf("Error reading upstream_ip: %v", err)
		}
	}

	if err = d.Set("upstream_port", dataSourceFlattenSystemCsfUpstreamPort(o["upstream-port"], d, "upstream_port")); err != nil {
		if !fortiAPIPatch(o["upstream-port"]) {
			return fmt.Errorf("Error reading upstream_port: %v", err)
		}
	}

	if err = d.Set("group_name", dataSourceFlattenSystemCsfGroupName(o["group-name"], d, "group_name")); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("accept_auth_by_cert", dataSourceFlattenSystemCsfAcceptAuthByCert(o["accept-auth-by-cert"], d, "accept_auth_by_cert")); err != nil {
		if !fortiAPIPatch(o["accept-auth-by-cert"]) {
			return fmt.Errorf("Error reading accept_auth_by_cert: %v", err)
		}
	}

	if err = d.Set("log_unification", dataSourceFlattenSystemCsfLogUnification(o["log-unification"], d, "log_unification")); err != nil {
		if !fortiAPIPatch(o["log-unification"]) {
			return fmt.Errorf("Error reading log_unification: %v", err)
		}
	}

	if err = d.Set("configuration_sync", dataSourceFlattenSystemCsfConfigurationSync(o["configuration-sync"], d, "configuration_sync")); err != nil {
		if !fortiAPIPatch(o["configuration-sync"]) {
			return fmt.Errorf("Error reading configuration_sync: %v", err)
		}
	}

	if err = d.Set("fabric_object_unification", dataSourceFlattenSystemCsfFabricObjectUnification(o["fabric-object-unification"], d, "fabric_object_unification")); err != nil {
		if !fortiAPIPatch(o["fabric-object-unification"]) {
			return fmt.Errorf("Error reading fabric_object_unification: %v", err)
		}
	}

	if err = d.Set("saml_configuration_sync", dataSourceFlattenSystemCsfSamlConfigurationSync(o["saml-configuration-sync"], d, "saml_configuration_sync")); err != nil {
		if !fortiAPIPatch(o["saml-configuration-sync"]) {
			return fmt.Errorf("Error reading saml_configuration_sync: %v", err)
		}
	}

	if err = d.Set("management_ip", dataSourceFlattenSystemCsfManagementIp(o["management-ip"], d, "management_ip")); err != nil {
		if !fortiAPIPatch(o["management-ip"]) {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("management_port", dataSourceFlattenSystemCsfManagementPort(o["management-port"], d, "management_port")); err != nil {
		if !fortiAPIPatch(o["management-port"]) {
			return fmt.Errorf("Error reading management_port: %v", err)
		}
	}

	if err = d.Set("authorization_request_type", dataSourceFlattenSystemCsfAuthorizationRequestType(o["authorization-request-type"], d, "authorization_request_type")); err != nil {
		if !fortiAPIPatch(o["authorization-request-type"]) {
			return fmt.Errorf("Error reading authorization_request_type: %v", err)
		}
	}

	if err = d.Set("certificate", dataSourceFlattenSystemCsfCertificate(o["certificate"], d, "certificate")); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("fabric_workers", dataSourceFlattenSystemCsfFabricWorkers(o["fabric-workers"], d, "fabric_workers")); err != nil {
		if !fortiAPIPatch(o["fabric-workers"]) {
			return fmt.Errorf("Error reading fabric_workers: %v", err)
		}
	}

	if err = d.Set("downstream_access", dataSourceFlattenSystemCsfDownstreamAccess(o["downstream-access"], d, "downstream_access")); err != nil {
		if !fortiAPIPatch(o["downstream-access"]) {
			return fmt.Errorf("Error reading downstream_access: %v", err)
		}
	}

	if err = d.Set("downstream_accprofile", dataSourceFlattenSystemCsfDownstreamAccprofile(o["downstream-accprofile"], d, "downstream_accprofile")); err != nil {
		if !fortiAPIPatch(o["downstream-accprofile"]) {
			return fmt.Errorf("Error reading downstream_accprofile: %v", err)
		}
	}

	if err = d.Set("trusted_list", dataSourceFlattenSystemCsfTrustedList(o["trusted-list"], d, "trusted_list")); err != nil {
		if !fortiAPIPatch(o["trusted-list"]) {
			return fmt.Errorf("Error reading trusted_list: %v", err)
		}
	}

	if err = d.Set("fabric_connector", dataSourceFlattenSystemCsfFabricConnector(o["fabric-connector"], d, "fabric_connector")); err != nil {
		if !fortiAPIPatch(o["fabric-connector"]) {
			return fmt.Errorf("Error reading fabric_connector: %v", err)
		}
	}

	if err = d.Set("forticloud_account_enforcement", dataSourceFlattenSystemCsfForticloudAccountEnforcement(o["forticloud-account-enforcement"], d, "forticloud_account_enforcement")); err != nil {
		if !fortiAPIPatch(o["forticloud-account-enforcement"]) {
			return fmt.Errorf("Error reading forticloud_account_enforcement: %v", err)
		}
	}

	if err = d.Set("file_mgmt", dataSourceFlattenSystemCsfFileMgmt(o["file-mgmt"], d, "file_mgmt")); err != nil {
		if !fortiAPIPatch(o["file-mgmt"]) {
			return fmt.Errorf("Error reading file_mgmt: %v", err)
		}
	}

	if err = d.Set("file_quota", dataSourceFlattenSystemCsfFileQuota(o["file-quota"], d, "file_quota")); err != nil {
		if !fortiAPIPatch(o["file-quota"]) {
			return fmt.Errorf("Error reading file_quota: %v", err)
		}
	}

	if err = d.Set("file_quota_warning", dataSourceFlattenSystemCsfFileQuotaWarning(o["file-quota-warning"], d, "file_quota_warning")); err != nil {
		if !fortiAPIPatch(o["file-quota-warning"]) {
			return fmt.Errorf("Error reading file_quota_warning: %v", err)
		}
	}

	if err = d.Set("fabric_device", dataSourceFlattenSystemCsfFabricDevice(o["fabric-device"], d, "fabric_device")); err != nil {
		if !fortiAPIPatch(o["fabric-device"]) {
			return fmt.Errorf("Error reading fabric_device: %v", err)
		}
	}

	return nil
}

func dataSourceFlattenSystemCsfFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}
