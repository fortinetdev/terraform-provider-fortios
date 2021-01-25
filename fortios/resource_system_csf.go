// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
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
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
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
			"configuration_sync": &schema.Schema{
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
				ValidateFunc: validation.IntBetween(1, 65535),
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
						"serial": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 19),
							Optional:     true,
							Computed:     true,
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
					},
				},
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
		},
	}
}

func resourceSystemCsfUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCsf(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCsf resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCsf(obj, mkey)
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

	err := c.DeleteSystemCsf(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCsf resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCsfRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemCsf(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsf resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCsf(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCsf resource from API: %v", err)
	}
	return nil
}

func flattenSystemCsfStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUpstreamIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfUpstreamPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfGroupName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfGroupPassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfConfigurationSync(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfManagementIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfManagementPort(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFixedKey(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["serial"] = flattenSystemCsfTrustedListSerial(i["serial"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := i["action"]; ok {
			tmp["action"] = flattenSystemCsfTrustedListAction(i["action"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := i["ha-members"]; ok {
			tmp["ha_members"] = flattenSystemCsfTrustedListHaMembers(i["ha-members"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := i["downstream-authorization"]; ok {
			tmp["downstream_authorization"] = flattenSystemCsfTrustedListDownstreamAuthorization(i["downstream-authorization"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "serial", d)
	return result
}

func flattenSystemCsfTrustedListSerial(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListAction(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListHaMembers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfTrustedListDownstreamAuthorization(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricDevice(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["name"] = flattenSystemCsfFabricDeviceName(i["name"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_ip"
		if _, ok := i["device-ip"]; ok {
			tmp["device_ip"] = flattenSystemCsfFabricDeviceDeviceIp(i["device-ip"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := i["device-type"]; ok {
			tmp["device_type"] = flattenSystemCsfFabricDeviceDeviceType(i["device-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "login"
		if _, ok := i["login"]; ok {
			tmp["login"] = flattenSystemCsfFabricDeviceLogin(i["login"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := i["password"]; ok {
			tmp["password"] = flattenSystemCsfFabricDevicePassword(i["password"], d, pre_append)
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

func flattenSystemCsfFabricDeviceName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceDeviceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceDeviceType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricDeviceLogin(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCsfFabricDevicePassword(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectSystemCsf(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenSystemCsfStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("upstream_ip", flattenSystemCsfUpstreamIp(o["upstream-ip"], d, "upstream_ip")); err != nil {
		if !fortiAPIPatch(o["upstream-ip"]) {
			return fmt.Errorf("Error reading upstream_ip: %v", err)
		}
	}

	if err = d.Set("upstream_port", flattenSystemCsfUpstreamPort(o["upstream-port"], d, "upstream_port")); err != nil {
		if !fortiAPIPatch(o["upstream-port"]) {
			return fmt.Errorf("Error reading upstream_port: %v", err)
		}
	}

	if err = d.Set("group_name", flattenSystemCsfGroupName(o["group-name"], d, "group_name")); err != nil {
		if !fortiAPIPatch(o["group-name"]) {
			return fmt.Errorf("Error reading group_name: %v", err)
		}
	}

	if err = d.Set("configuration_sync", flattenSystemCsfConfigurationSync(o["configuration-sync"], d, "configuration_sync")); err != nil {
		if !fortiAPIPatch(o["configuration-sync"]) {
			return fmt.Errorf("Error reading configuration_sync: %v", err)
		}
	}

	if err = d.Set("management_ip", flattenSystemCsfManagementIp(o["management-ip"], d, "management_ip")); err != nil {
		if !fortiAPIPatch(o["management-ip"]) {
			return fmt.Errorf("Error reading management_ip: %v", err)
		}
	}

	if err = d.Set("management_port", flattenSystemCsfManagementPort(o["management-port"], d, "management_port")); err != nil {
		if !fortiAPIPatch(o["management-port"]) {
			return fmt.Errorf("Error reading management_port: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("trusted_list", flattenSystemCsfTrustedList(o["trusted-list"], d, "trusted_list")); err != nil {
			if !fortiAPIPatch(o["trusted-list"]) {
				return fmt.Errorf("Error reading trusted_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("trusted_list"); ok {
			if err = d.Set("trusted_list", flattenSystemCsfTrustedList(o["trusted-list"], d, "trusted_list")); err != nil {
				if !fortiAPIPatch(o["trusted-list"]) {
					return fmt.Errorf("Error reading trusted_list: %v", err)
				}
			}
		}
	}

	if isImportTable() {
		if err = d.Set("fabric_device", flattenSystemCsfFabricDevice(o["fabric-device"], d, "fabric_device")); err != nil {
			if !fortiAPIPatch(o["fabric-device"]) {
				return fmt.Errorf("Error reading fabric_device: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("fabric_device"); ok {
			if err = d.Set("fabric_device", flattenSystemCsfFabricDevice(o["fabric-device"], d, "fabric_device")); err != nil {
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
	log.Printf("ER List: %v", e)
}

func expandSystemCsfStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfUpstreamPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfGroupName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfGroupPassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfConfigurationSync(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfManagementIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfManagementPort(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFixedKey(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "serial"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["serial"], _ = expandSystemCsfTrustedListSerial(d, i["serial"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "action"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["action"], _ = expandSystemCsfTrustedListAction(d, i["action"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "ha_members"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["ha-members"], _ = expandSystemCsfTrustedListHaMembers(d, i["ha_members"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "downstream_authorization"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["downstream-authorization"], _ = expandSystemCsfTrustedListDownstreamAuthorization(d, i["downstream_authorization"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemCsfTrustedListSerial(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListAction(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListHaMembers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfTrustedListDownstreamAuthorization(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDevice(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["name"], _ = expandSystemCsfFabricDeviceName(d, i["name"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_ip"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device-ip"], _ = expandSystemCsfFabricDeviceDeviceIp(d, i["device_ip"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "device_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["device-type"], _ = expandSystemCsfFabricDeviceDeviceType(d, i["device_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "login"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["login"], _ = expandSystemCsfFabricDeviceLogin(d, i["login"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "password"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["password"], _ = expandSystemCsfFabricDevicePassword(d, i["password"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemCsfFabricDeviceName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceDeviceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceDeviceType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDeviceLogin(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCsfFabricDevicePassword(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCsf(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandSystemCsfStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("upstream_ip"); ok {
		t, err := expandSystemCsfUpstreamIp(d, v, "upstream_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-ip"] = t
		}
	}

	if v, ok := d.GetOk("upstream_port"); ok {
		t, err := expandSystemCsfUpstreamPort(d, v, "upstream_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upstream-port"] = t
		}
	}

	if v, ok := d.GetOk("group_name"); ok {
		t, err := expandSystemCsfGroupName(d, v, "group_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-name"] = t
		}
	}

	if v, ok := d.GetOk("group_password"); ok {
		t, err := expandSystemCsfGroupPassword(d, v, "group_password")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["group-password"] = t
		}
	}

	if v, ok := d.GetOk("configuration_sync"); ok {
		t, err := expandSystemCsfConfigurationSync(d, v, "configuration_sync")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["configuration-sync"] = t
		}
	}

	if v, ok := d.GetOk("management_ip"); ok {
		t, err := expandSystemCsfManagementIp(d, v, "management_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-ip"] = t
		}
	}

	if v, ok := d.GetOk("management_port"); ok {
		t, err := expandSystemCsfManagementPort(d, v, "management_port")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["management-port"] = t
		}
	}

	if v, ok := d.GetOk("fixed_key"); ok {
		t, err := expandSystemCsfFixedKey(d, v, "fixed_key")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fixed-key"] = t
		}
	}

	if v, ok := d.GetOk("trusted_list"); ok {
		t, err := expandSystemCsfTrustedList(d, v, "trusted_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["trusted-list"] = t
		}
	}

	if v, ok := d.GetOk("fabric_device"); ok {
		t, err := expandSystemCsfFabricDevice(d, v, "fabric_device")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fabric-device"] = t
		}
	}

	return &obj, nil
}
