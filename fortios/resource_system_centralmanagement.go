// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Configure central management.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"mode": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_config_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"schedule_script_restore": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_push_configuration": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_push_firmware": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_remote_firmware_upgrade": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"allow_monitor": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial_number": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg_source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"fmg_source_ip6": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"local_cert": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"ca_cert": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"vdom": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 31),
				Optional:     true,
				Computed:     true,
			},
			"server_list": &schema.Schema{
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"id": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
						},
						"server_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"addr_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_address": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"server_address6": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
						},
						"fqdn": &schema.Schema{
							Type:         schema.TypeString,
							ValidateFunc: validation.StringLenBetween(0, 255),
							Optional:     true,
							Computed:     true,
						},
					},
				},
			},
			"fmg_update_port": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"include_default_servers": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface_select_method": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"interface": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 15),
				Optional:     true,
				Computed:     true,
			},
			"dynamic_sort_subtable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Default:  "false",
			},
		},
	}
}

func resourceSystemCentralManagementUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemCentralManagement(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource while getting object: %v", err)
	}

	o, err := c.UpdateSystemCentralManagement(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectSystemCentralManagement(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating SystemCentralManagement resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemCentralManagement(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing SystemCentralManagement resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceSystemCentralManagementRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadSystemCentralManagement(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemCentralManagement(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemCentralManagement resource from API: %v", err)
	}
	return nil
}

func flattenSystemCentralManagementMode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementScheduleConfigRestore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementScheduleScriptRestore(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowPushConfiguration(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowPushFirmware(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowRemoteFirmwareUpgrade(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementAllowMonitor(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementSerialNumber(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementFmg(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementFmgSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementFmgSourceIp6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementLocalCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementCaCert(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementServerList(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["id"] = flattenSystemCentralManagementServerListId(i["id"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := i["server-type"]; ok {

			tmp["server_type"] = flattenSystemCentralManagementServerListServerType(i["server-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := i["addr-type"]; ok {

			tmp["addr_type"] = flattenSystemCentralManagementServerListAddrType(i["addr-type"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address"
		if _, ok := i["server-address"]; ok {

			tmp["server_address"] = flattenSystemCentralManagementServerListServerAddress(i["server-address"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address6"
		if _, ok := i["server-address6"]; ok {

			tmp["server_address6"] = flattenSystemCentralManagementServerListServerAddress6(i["server-address6"], d, pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := i["fqdn"]; ok {

			tmp["fqdn"] = flattenSystemCentralManagementServerListFqdn(i["fqdn"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "id", d)
	return result
}

func flattenSystemCentralManagementServerListId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListAddrType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerAddress(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListServerAddress6(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementServerListFqdn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementFmgUpdatePort(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementIncludeDefaultServers(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemCentralManagementInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemCentralManagement(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mode", flattenSystemCentralManagementMode(o["mode"], d, "mode", sv)); err != nil {
		if !fortiAPIPatch(o["mode"]) {
			return fmt.Errorf("Error reading mode: %v", err)
		}
	}

	if err = d.Set("type", flattenSystemCentralManagementType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("schedule_config_restore", flattenSystemCentralManagementScheduleConfigRestore(o["schedule-config-restore"], d, "schedule_config_restore", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-config-restore"]) {
			return fmt.Errorf("Error reading schedule_config_restore: %v", err)
		}
	}

	if err = d.Set("schedule_script_restore", flattenSystemCentralManagementScheduleScriptRestore(o["schedule-script-restore"], d, "schedule_script_restore", sv)); err != nil {
		if !fortiAPIPatch(o["schedule-script-restore"]) {
			return fmt.Errorf("Error reading schedule_script_restore: %v", err)
		}
	}

	if err = d.Set("allow_push_configuration", flattenSystemCentralManagementAllowPushConfiguration(o["allow-push-configuration"], d, "allow_push_configuration", sv)); err != nil {
		if !fortiAPIPatch(o["allow-push-configuration"]) {
			return fmt.Errorf("Error reading allow_push_configuration: %v", err)
		}
	}

	if err = d.Set("allow_push_firmware", flattenSystemCentralManagementAllowPushFirmware(o["allow-push-firmware"], d, "allow_push_firmware", sv)); err != nil {
		if !fortiAPIPatch(o["allow-push-firmware"]) {
			return fmt.Errorf("Error reading allow_push_firmware: %v", err)
		}
	}

	if err = d.Set("allow_remote_firmware_upgrade", flattenSystemCentralManagementAllowRemoteFirmwareUpgrade(o["allow-remote-firmware-upgrade"], d, "allow_remote_firmware_upgrade", sv)); err != nil {
		if !fortiAPIPatch(o["allow-remote-firmware-upgrade"]) {
			return fmt.Errorf("Error reading allow_remote_firmware_upgrade: %v", err)
		}
	}

	if err = d.Set("allow_monitor", flattenSystemCentralManagementAllowMonitor(o["allow-monitor"], d, "allow_monitor", sv)); err != nil {
		if !fortiAPIPatch(o["allow-monitor"]) {
			return fmt.Errorf("Error reading allow_monitor: %v", err)
		}
	}

	if err = d.Set("serial_number", flattenSystemCentralManagementSerialNumber(o["serial-number"], d, "serial_number", sv)); err != nil {
		if !fortiAPIPatch(o["serial-number"]) {
			return fmt.Errorf("Error reading serial_number: %v", err)
		}
	}

	if err = d.Set("fmg", flattenSystemCentralManagementFmg(o["fmg"], d, "fmg", sv)); err != nil {
		if !fortiAPIPatch(o["fmg"]) {
			return fmt.Errorf("Error reading fmg: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip", flattenSystemCentralManagementFmgSourceIp(o["fmg-source-ip"], d, "fmg_source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["fmg-source-ip"]) {
			return fmt.Errorf("Error reading fmg_source_ip: %v", err)
		}
	}

	if err = d.Set("fmg_source_ip6", flattenSystemCentralManagementFmgSourceIp6(o["fmg-source-ip6"], d, "fmg_source_ip6", sv)); err != nil {
		if !fortiAPIPatch(o["fmg-source-ip6"]) {
			return fmt.Errorf("Error reading fmg_source_ip6: %v", err)
		}
	}

	if err = d.Set("local_cert", flattenSystemCentralManagementLocalCert(o["local-cert"], d, "local_cert", sv)); err != nil {
		if !fortiAPIPatch(o["local-cert"]) {
			return fmt.Errorf("Error reading local_cert: %v", err)
		}
	}

	if err = d.Set("ca_cert", flattenSystemCentralManagementCaCert(o["ca-cert"], d, "ca_cert", sv)); err != nil {
		if !fortiAPIPatch(o["ca-cert"]) {
			return fmt.Errorf("Error reading ca_cert: %v", err)
		}
	}

	if err = d.Set("vdom", flattenSystemCentralManagementVdom(o["vdom"], d, "vdom", sv)); err != nil {
		if !fortiAPIPatch(o["vdom"]) {
			return fmt.Errorf("Error reading vdom: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("server_list", flattenSystemCentralManagementServerList(o["server-list"], d, "server_list", sv)); err != nil {
			if !fortiAPIPatch(o["server-list"]) {
				return fmt.Errorf("Error reading server_list: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("server_list"); ok {
			if err = d.Set("server_list", flattenSystemCentralManagementServerList(o["server-list"], d, "server_list", sv)); err != nil {
				if !fortiAPIPatch(o["server-list"]) {
					return fmt.Errorf("Error reading server_list: %v", err)
				}
			}
		}
	}

	if err = d.Set("fmg_update_port", flattenSystemCentralManagementFmgUpdatePort(o["fmg-update-port"], d, "fmg_update_port", sv)); err != nil {
		if !fortiAPIPatch(o["fmg-update-port"]) {
			return fmt.Errorf("Error reading fmg_update_port: %v", err)
		}
	}

	if err = d.Set("include_default_servers", flattenSystemCentralManagementIncludeDefaultServers(o["include-default-servers"], d, "include_default_servers", sv)); err != nil {
		if !fortiAPIPatch(o["include-default-servers"]) {
			return fmt.Errorf("Error reading include_default_servers: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenSystemCentralManagementEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenSystemCentralManagementInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenSystemCentralManagementInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenSystemCentralManagementFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandSystemCentralManagementMode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementScheduleConfigRestore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementScheduleScriptRestore(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowPushConfiguration(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowPushFirmware(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowRemoteFirmwareUpgrade(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementAllowMonitor(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementSerialNumber(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmg(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmgSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmgSourceIp6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementLocalCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementCaCert(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerList(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

		pre_append = pre + "." + strconv.Itoa(con) + "." + "id"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["id"], _ = expandSystemCentralManagementServerListId(d, i["id"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server-type"], _ = expandSystemCentralManagementServerListServerType(d, i["server_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "addr_type"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["addr-type"], _ = expandSystemCentralManagementServerListAddrType(d, i["addr_type"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server-address"], _ = expandSystemCentralManagementServerListServerAddress(d, i["server_address"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "server_address6"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["server-address6"], _ = expandSystemCentralManagementServerListServerAddress6(d, i["server_address6"], pre_append, sv)
		}

		pre_append = pre + "." + strconv.Itoa(con) + "." + "fqdn"
		if _, ok := d.GetOk(pre_append); ok {

			tmp["fqdn"], _ = expandSystemCentralManagementServerListFqdn(d, i["fqdn"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandSystemCentralManagementServerListId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListAddrType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerAddress(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListServerAddress6(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementServerListFqdn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementFmgUpdatePort(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementIncludeDefaultServers(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemCentralManagementInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemCentralManagement(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mode"); ok {
		if setArgNil {
			obj["mode"] = nil
		} else {

			t, err := expandSystemCentralManagementMode(d, v, "mode", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["mode"] = t
			}
		}
	}

	if v, ok := d.GetOk("type"); ok {
		if setArgNil {
			obj["type"] = nil
		} else {

			t, err := expandSystemCentralManagementType(d, v, "type", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["type"] = t
			}
		}
	}

	if v, ok := d.GetOk("schedule_config_restore"); ok {
		if setArgNil {
			obj["schedule-config-restore"] = nil
		} else {

			t, err := expandSystemCentralManagementScheduleConfigRestore(d, v, "schedule_config_restore", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["schedule-config-restore"] = t
			}
		}
	}

	if v, ok := d.GetOk("schedule_script_restore"); ok {
		if setArgNil {
			obj["schedule-script-restore"] = nil
		} else {

			t, err := expandSystemCentralManagementScheduleScriptRestore(d, v, "schedule_script_restore", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["schedule-script-restore"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_push_configuration"); ok {
		if setArgNil {
			obj["allow-push-configuration"] = nil
		} else {

			t, err := expandSystemCentralManagementAllowPushConfiguration(d, v, "allow_push_configuration", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-push-configuration"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_push_firmware"); ok {
		if setArgNil {
			obj["allow-push-firmware"] = nil
		} else {

			t, err := expandSystemCentralManagementAllowPushFirmware(d, v, "allow_push_firmware", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-push-firmware"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_remote_firmware_upgrade"); ok {
		if setArgNil {
			obj["allow-remote-firmware-upgrade"] = nil
		} else {

			t, err := expandSystemCentralManagementAllowRemoteFirmwareUpgrade(d, v, "allow_remote_firmware_upgrade", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-remote-firmware-upgrade"] = t
			}
		}
	}

	if v, ok := d.GetOk("allow_monitor"); ok {
		if setArgNil {
			obj["allow-monitor"] = nil
		} else {

			t, err := expandSystemCentralManagementAllowMonitor(d, v, "allow_monitor", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["allow-monitor"] = t
			}
		}
	}

	if v, ok := d.GetOk("serial_number"); ok {
		if setArgNil {
			obj["serial-number"] = nil
		} else {

			t, err := expandSystemCentralManagementSerialNumber(d, v, "serial_number", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["serial-number"] = t
			}
		}
	}

	if v, ok := d.GetOk("fmg"); ok {
		if setArgNil {
			obj["fmg"] = nil
		} else {

			t, err := expandSystemCentralManagementFmg(d, v, "fmg", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fmg"] = t
			}
		}
	}

	if v, ok := d.GetOk("fmg_source_ip"); ok {
		if setArgNil {
			obj["fmg-source-ip"] = nil
		} else {

			t, err := expandSystemCentralManagementFmgSourceIp(d, v, "fmg_source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fmg-source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("fmg_source_ip6"); ok {
		if setArgNil {
			obj["fmg-source-ip6"] = nil
		} else {

			t, err := expandSystemCentralManagementFmgSourceIp6(d, v, "fmg_source_ip6", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fmg-source-ip6"] = t
			}
		}
	}

	if v, ok := d.GetOk("local_cert"); ok {
		if setArgNil {
			obj["local-cert"] = nil
		} else {

			t, err := expandSystemCentralManagementLocalCert(d, v, "local_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["local-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("ca_cert"); ok {
		if setArgNil {
			obj["ca-cert"] = nil
		} else {

			t, err := expandSystemCentralManagementCaCert(d, v, "ca_cert", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ca-cert"] = t
			}
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		if setArgNil {
			obj["vdom"] = nil
		} else {

			t, err := expandSystemCentralManagementVdom(d, v, "vdom", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["vdom"] = t
			}
		}
	}

	if v, ok := d.GetOk("server_list"); ok {
		if setArgNil {
			obj["server-list"] = make([]struct{}, 0)
		} else {

			t, err := expandSystemCentralManagementServerList(d, v, "server_list", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["server-list"] = t
			}
		}
	}

	if v, ok := d.GetOk("fmg_update_port"); ok {
		if setArgNil {
			obj["fmg-update-port"] = nil
		} else {

			t, err := expandSystemCentralManagementFmgUpdatePort(d, v, "fmg_update_port", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["fmg-update-port"] = t
			}
		}
	}

	if v, ok := d.GetOk("include_default_servers"); ok {
		if setArgNil {
			obj["include-default-servers"] = nil
		} else {

			t, err := expandSystemCentralManagementIncludeDefaultServers(d, v, "include_default_servers", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["include-default-servers"] = t
			}
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		if setArgNil {
			obj["enc-algorithm"] = nil
		} else {

			t, err := expandSystemCentralManagementEncAlgorithm(d, v, "enc_algorithm", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["enc-algorithm"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		if setArgNil {
			obj["interface-select-method"] = nil
		} else {

			t, err := expandSystemCentralManagementInterfaceSelectMethod(d, v, "interface_select_method", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface-select-method"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface"); ok {
		if setArgNil {
			obj["interface"] = nil
		} else {

			t, err := expandSystemCentralManagementInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	return &obj, nil
}
