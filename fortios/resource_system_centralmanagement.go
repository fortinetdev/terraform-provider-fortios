// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Configure central management.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceSystemCentralManagement() *schema.Resource {
	return &schema.Resource{
		Create: resourceSystemCentralManagementUpdate,
		Read:   resourceSystemCentralManagementRead,
		Update: resourceSystemCentralManagementUpdate,
		Delete: resourceSystemCentralManagementDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mode": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_config_restore": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_script_restore": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_push_configuration": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_push_firmware": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_remote_firmware_upgrade": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_monitor": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg_source_ip": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg_source_ip6": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type: schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional: true,
				Computed: true,
			},
			"server_list": &schema.Schema{
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
						"server_type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addr_type": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_address": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_address6": &schema.Schema{
							Type: schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fqdn": &schema.Schema{
							Type: schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional: true,
							Computed: true,
						},
					},
				},
			},
			"include_default_servers": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type: schema.TypeString,
				Optional: true,
				Computed: true,
			},
		},
	}
}


func resourceSystemCentralManagementUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectSystemCentralManagement(d)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCentralManagement(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("SystemCentralManagement")
	}

	return resourceSystemCentralManagementRead(d, m)
}

func resourceSystemCentralManagementDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteSystemCentralManagement(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting SystemCentralManagement resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCentralManagementRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadSystemCentralManagement(mkey)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCentralManagement(d, o)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource from API: %v", err)
	}
	return nil
}


func flattenSystemCentralManagementMode(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementScheduleConfigRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementScheduleScriptRestore(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowPushConfiguration(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowPushFirmware(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowRemoteFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowMonitor(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementSerialNumber(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementFmg(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementFmgSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementFmgSourceIp6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerList(v interface{}, d *schema.ResourceData, pre string) []map[string]interface{} {
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
			tmp["id"] = flattenSystemCentralManagementServerListId(i["id"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {
			tmp["server_type"] = flattenSystemCentralManagementServerListServerType(i["server-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {
			tmp["addr_type"] = flattenSystemCentralManagementServerListAddrType(i["addr-type"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address"
		if _, ok := i["server-address"]; ok {
			tmp["server_address"] = flattenSystemCentralManagementServerListServerAddress(i["server-address"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address6"
		if _, ok := i["server-address6"]; ok {
			tmp["server_address6"] = flattenSystemCentralManagementServerListServerAddress6(i["server-address6"], d, pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {
			tmp["fqdn"] = flattenSystemCentralManagementServerListFqdn(i["fqdn"], d, pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result
}

func flattenSystemCentralManagementServerListId(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListAddrType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerAddress(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerAddress6(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListFqdn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementIncludeDefaultServers(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenSystemCentralManagementEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}


func refreshObjectSystemCentralManagement(d *schema.ResourceData, o map[string]interface{}) error {
	var err error


	if err = d.Set("mode", flattenSystemCentralManagementMode(o["mode"], d, "mode")); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemCentralManagementType(o["type"], d, "type")); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("schedule_config_restore", flattenSystemCentralManagementScheduleConfigRestore(o["schedule-config-restore"], d, "schedule_config_restore")); err != nil {
		if !fortiAPIPatch(o["schedule-config-restore"]) {
			return fmt.Errorf("Error reading schedule_config_restore: %v", err)
		}
	}

	if err = d.Set("schedule_script_restore", flattenSystemCentralManagementScheduleScriptRestore(o["schedule-script-restore"], d, "schedule_script_restore")); err != nil {
		if !fortiAPIPatch(o["schedule-script-restore"]) {
			return fmt.Errorf("Error reading schedule_script_restore: %v", err)
		}
	}

	if err = d.Set("allow_push_configuration", flattenSystemCentralManagementAllowPushConfiguration(o["allow-push-configuration"], d, "allow_push_configuration")); err != nil {
		if !fortiAPIPatch(o["allow-push-configuration"]) {
			return fmt.Errorf("Error reading allow_push_configuration: %v", err)
		}
	}

	if err = d.Set("allow_push_firmware", flattenSystemCentralManagementAllowPushFirmware(o["allow-push-firmware"], d, "allow_push_firmware")); err != nil {
		if !fortiAPIPatch(o["allow-push-firmware"]) {
			return fmt.Errorf("Error reading allow_push_firmware: %v", err)
		}
	}

	if err = d.Set("allow_remote_firmware_upgrade", flattenSystemCentralManagementAllowRemoteFirmwareUpgrade(o["allow-remote-firmware-upgrade"], d, "allow_remote_firmware_upgrade")); err != nil {
		if !fortiAPIPatch(o["allow-remote-firmware-upgrade"]) {
			return fmt.Errorf("Error reading allow_remote_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("allow_monitor", flattenSystemCentralManagementAllowMonitor(o["allow-monitor"], d, "allow_monitor")); err != nil {
		if !fortiAPIPatch(o["allow-monitor"]) {
			return fmt.Errorf("Error reading allow_monitor: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenSystemCentralManagementSerialNumber(o["serial-number"], d, "serial_number")); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("fmg", flattenSystemCentralManagementFmg(o["fmg"], d, "fmg")); err != nil {
		if !fortiAPIPatch(o["fmg"]) {
			return fmt.Errorf("Error reading fmg: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip", flattenSystemCentralManagementFmgSourceIp(o["fmg-source-ip"], d, "fmg_source_ip")); err != nil {
		if !fortiAPIPatch(o["fmg-source-ip"]) {
			return fmt.Errorf("Error reading fmg_source_ip: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip6", flattenSystemCentralManagementFmgSourceIp6(o["fmg-source-ip6"], d, "fmg_source_ip6")); err != nil {
		if !fortiAPIPatch(o["fmg-source-ip6"]) {
			return fmt.Errorf("Error reading fmg_source_ip6: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemCentralManagementVdom(o["vdom"], d, "vdom")); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

    if isImportTable() {
        if err = d.Set("server_list", flattenSystemCentralManagementServerList(o["server-list"], d, "server_list")); err != nil {
            if !fortiAPIPatch(o["server-list"]) {
                return fmt.Errorf("Error reading server_list: %v", err)
            }
        }
    } else {
        if _, ok := d.GetOk("server_list"); ok {
            if err = d.Set("server_list", flattenSystemCentralManagementServerList(o["server-list"], d, "server_list")); err != nil {
                if !fortiAPIPatch(o["server-list"]) {
                    return fmt.Errorf("Error reading server_list: %v", err)
                }
            }
        }
    }

	if err = d.Set("include_default_servers", flattenSystemCentralManagementIncludeDefaultServers(o["include-default-servers"], d, "include_default_servers")); err != nil {
		if !fortiAPIPatch(o["include-default-servers"]) {
			return fmt.Errorf("Error reading include_default_servers: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystemCentralManagementEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}


	return nil
}

func flattenSystemCentralManagementFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}


func expandSystemCentralManagementMode(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementScheduleConfigRestore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementScheduleScriptRestore(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowPushConfiguration(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowPushFirmware(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowRemoteFirmwareUpgrade(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowMonitor(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementSerialNumber(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmg(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmgSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmgSourceIp6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerList(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
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
			tmp["id"], _ = expandSystemCentralManagementServerListId(d, i["id"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server-type"], _ = expandSystemCentralManagementServerListServerType(d, i["server_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["addr-type"], _ = expandSystemCentralManagementServerListAddrType(d, i["addr_type"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server-address"], _ = expandSystemCentralManagementServerListServerAddress(d, i["server_address"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address6"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["server-address6"], _ = expandSystemCentralManagementServerListServerAddress6(d, i["server_address6"], pre_append)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := d.GetOk(pre_append); ok {
			tmp["fqdn"], _ = expandSystemCentralManagementServerListFqdn(d, i["fqdn"], pre_append)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemCentralManagementServerListId(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListAddrType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerAddress(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerAddress6(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListFqdn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementIncludeDefaultServers(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}


func getObjectSystemCentralManagement(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})


	if v, ok := d.GetOk("mode"); ok {
		t, err := expandSystemCentralManagementMode(d, v, "mode")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mode"] = t
		}
	}

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemCentralManagementType(d, v, "type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("schedule_config_restore"); ok {
		t, err := expandSystemCentralManagementScheduleConfigRestore(d, v, "schedule_config_restore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-config-restore"] = t
		}
	}

	if v, ok := d.GetOk("schedule_script_restore"); ok {
		t, err := expandSystemCentralManagementScheduleScriptRestore(d, v, "schedule_script_restore")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["schedule-script-restore"] = t
		}
	}

	if v, ok := d.GetOk("allow_push_configuration"); ok {
		t, err := expandSystemCentralManagementAllowPushConfiguration(d, v, "allow_push_configuration")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-push-configuration"] = t
		}
	}

	if v, ok := d.GetOk("allow_push_firmware"); ok {
		t, err := expandSystemCentralManagementAllowPushFirmware(d, v, "allow_push_firmware")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-push-firmware"] = t
		}
	}

	if v, ok := d.GetOk("allow_remote_firmware_upgrade"); ok {
		t, err := expandSystemCentralManagementAllowRemoteFirmwareUpgrade(d, v, "allow_remote_firmware_upgrade")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-remote-firmware-upgrade"] = t
		}
	}

	if v, ok := d.GetOk("allow_monitor"); ok {
		t, err := expandSystemCentralManagementAllowMonitor(d, v, "allow_monitor")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["allow-monitor"] = t
		}
	}

	if v, ok := d.GetOk("serial_number"); ok {
		t, err := expandSystemCentralManagementSerialNumber(d, v, "serial_number")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial-number"] = t
		}
	}

	if v, ok := d.GetOk("fmg"); ok {
		t, err := expandSystemCentralManagementFmg(d, v, "fmg")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmg"] = t
		}
	}

	if v, ok := d.GetOk("fmg_source_ip"); ok {
		t, err := expandSystemCentralManagementFmgSourceIp(d, v, "fmg_source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmg-source-ip"] = t
		}
	}

	if v, ok := d.GetOk("fmg_source_ip6"); ok {
		t, err := expandSystemCentralManagementFmgSourceIp6(d, v, "fmg_source_ip6")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["fmg-source-ip6"] = t
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandSystemCentralManagementVdom(d, v, "vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["vdom"] = t
		}
	}

	if v, ok := d.GetOk("server_list"); ok {
		t, err := expandSystemCentralManagementServerList(d, v, "server_list")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server-list"] = t
		}
	}

	if v, ok := d.GetOk("include_default_servers"); ok {
		t, err := expandSystemCentralManagementIncludeDefaultServers(d, v, "include_default_servers")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["include-default-servers"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		t, err := expandSystemCentralManagementEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}


	return &obj, nil
}

