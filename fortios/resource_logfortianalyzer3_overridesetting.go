// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Override FortiAnalyzer settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogFortianalyzer3OverrideSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzer3OverrideSettingUpdate,
		Read:   resourceLogFortianalyzer3OverrideSettingRead,
		Update: resourceLogFortianalyzer3OverrideSettingUpdate,
		Delete: resourceLogFortianalyzer3OverrideSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"override": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"use_management_vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ips_archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"server": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 127),
				Optional:     true,
				Computed:     true,
			},
			"certificate_verification": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"serial": &schema.Schema{
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
			"preshared_key": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"access_config": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"hmac_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"enc_algorithm": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"ssl_min_proto_version": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"conn_timeout": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 3600),
				Optional:     true,
				Computed:     true,
			},
			"monitor_keepalive_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 120),
				Optional:     true,
				Computed:     true,
			},
			"monitor_failure_retry_period": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 86400),
				Optional:     true,
				Computed:     true,
			},
			"mgmt_name": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"faz_type": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"certificate": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"source_ip": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"__change_ip": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 255),
				Optional:     true,
				Computed:     true,
			},
			"upload_option": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_interval": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"reliable": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"priority": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_log_rate": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 100000),
				Optional:     true,
				Computed:     true,
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

func resourceLogFortianalyzer3OverrideSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogFortianalyzer3OverrideSetting(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3OverrideSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzer3OverrideSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3OverrideSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzer3OverrideSetting")
	}

	return resourceLogFortianalyzer3OverrideSettingRead(d, m)
}

func resourceLogFortianalyzer3OverrideSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteLogFortianalyzer3OverrideSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting LogFortianalyzer3OverrideSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzer3OverrideSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadLogFortianalyzer3OverrideSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3OverrideSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzer3OverrideSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3OverrideSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzer3OverrideSettingOverride(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUseManagementVdom(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingServer(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingCertificateVerification(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingSerial(v interface{}, d *schema.ResourceData, pre string, sv string) []map[string]interface{} {
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

			tmp["name"] = flattenLogFortianalyzer3OverrideSettingSerialName(i["name"], d, pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	dynamic_sort_subtable(result, "name", d)
	return result
}

func flattenLogFortianalyzer3OverrideSettingSerialName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingPresharedKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingAccessConfig(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingHmacAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingConnTimeout(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingMonitorKeepalivePeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingMonitorFailureRetryPeriod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingMgmtName(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingFazType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingCertificate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSetting__Change_Ip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUploadOption(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUploadDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUploadTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingReliable(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingPriority(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingMaxLogRate(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogFortianalyzer3OverrideSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("override", flattenLogFortianalyzer3OverrideSettingOverride(o["override"], d, "override", sv)); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("use_management_vdom", flattenLogFortianalyzer3OverrideSettingUseManagementVdom(o["use-management-vdom"], d, "use_management_vdom", sv)); err != nil {
		if !fortiAPIPatch(o["use-management-vdom"]) {
			return fmt.Errorf("Error reading use_management_vdom: %v", err)
		}
	}

	if err = d.Set("status", flattenLogFortianalyzer3OverrideSettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogFortianalyzer3OverrideSettingIpsArchive(o["ips-archive"], d, "ips_archive", sv)); err != nil {
		if !fortiAPIPatch(o["ips-archive"]) {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("server", flattenLogFortianalyzer3OverrideSettingServer(o["server"], d, "server", sv)); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("certificate_verification", flattenLogFortianalyzer3OverrideSettingCertificateVerification(o["certificate-verification"], d, "certificate_verification", sv)); err != nil {
		if !fortiAPIPatch(o["certificate-verification"]) {
			return fmt.Errorf("Error reading certificate_verification: %v", err)
		}
	}

	if isImportTable() {
		if err = d.Set("serial", flattenLogFortianalyzer3OverrideSettingSerial(o["serial"], d, "serial", sv)); err != nil {
			if !fortiAPIPatch(o["serial"]) {
				return fmt.Errorf("Error reading serial: %v", err)
			}
		}
	} else {
		if _, ok := d.GetOk("serial"); ok {
			if err = d.Set("serial", flattenLogFortianalyzer3OverrideSettingSerial(o["serial"], d, "serial", sv)); err != nil {
				if !fortiAPIPatch(o["serial"]) {
					return fmt.Errorf("Error reading serial: %v", err)
				}
			}
		}
	}

	if err = d.Set("preshared_key", flattenLogFortianalyzer3OverrideSettingPresharedKey(o["preshared-key"], d, "preshared_key", sv)); err != nil {
		if !fortiAPIPatch(o["preshared-key"]) {
			return fmt.Errorf("Error reading preshared_key: %v", err)
		}
	}

	if err = d.Set("access_config", flattenLogFortianalyzer3OverrideSettingAccessConfig(o["access-config"], d, "access_config", sv)); err != nil {
		if !fortiAPIPatch(o["access-config"]) {
			return fmt.Errorf("Error reading access_config: %v", err)
		}
	}

	if err = d.Set("hmac_algorithm", flattenLogFortianalyzer3OverrideSettingHmacAlgorithm(o["hmac-algorithm"], d, "hmac_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["hmac-algorithm"]) {
			return fmt.Errorf("Error reading hmac_algorithm: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortianalyzer3OverrideSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm", sv)); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogFortianalyzer3OverrideSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version", sv)); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenLogFortianalyzer3OverrideSettingConnTimeout(o["conn-timeout"], d, "conn_timeout", sv)); err != nil {
		if !fortiAPIPatch(o["conn-timeout"]) {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("monitor_keepalive_period", flattenLogFortianalyzer3OverrideSettingMonitorKeepalivePeriod(o["monitor-keepalive-period"], d, "monitor_keepalive_period", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-keepalive-period"]) {
			return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
		}
	}

	if err = d.Set("monitor_failure_retry_period", flattenLogFortianalyzer3OverrideSettingMonitorFailureRetryPeriod(o["monitor-failure-retry-period"], d, "monitor_failure_retry_period", sv)); err != nil {
		if !fortiAPIPatch(o["monitor-failure-retry-period"]) {
			return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
		}
	}

	if err = d.Set("mgmt_name", flattenLogFortianalyzer3OverrideSettingMgmtName(o["mgmt-name"], d, "mgmt_name", sv)); err != nil {
		if !fortiAPIPatch(o["mgmt-name"]) {
			return fmt.Errorf("Error reading mgmt_name: %v", err)
		}
	}

	if err = d.Set("faz_type", flattenLogFortianalyzer3OverrideSettingFazType(o["faz-type"], d, "faz_type", sv)); err != nil {
		if !fortiAPIPatch(o["faz-type"]) {
			return fmt.Errorf("Error reading faz_type: %v", err)
		}
	}

	if err = d.Set("certificate", flattenLogFortianalyzer3OverrideSettingCertificate(o["certificate"], d, "certificate", sv)); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogFortianalyzer3OverrideSettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("__change_ip", flattenLogFortianalyzer3OverrideSetting__Change_Ip(o["__change_ip"], d, "__change_ip", sv)); err != nil {
		if !fortiAPIPatch(o["__change_ip"]) {
			return fmt.Errorf("Error reading __change_ip: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenLogFortianalyzer3OverrideSettingUploadOption(o["upload-option"], d, "upload_option", sv)); err != nil {
		if !fortiAPIPatch(o["upload-option"]) {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenLogFortianalyzer3OverrideSettingUploadInterval(o["upload-interval"], d, "upload_interval", sv)); err != nil {
		if !fortiAPIPatch(o["upload-interval"]) {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortianalyzer3OverrideSettingUploadDay(o["upload-day"], d, "upload_day", sv)); err != nil {
		if !fortiAPIPatch(o["upload-day"]) {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortianalyzer3OverrideSettingUploadTime(o["upload-time"], d, "upload_time", sv)); err != nil {
		if !fortiAPIPatch(o["upload-time"]) {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("reliable", flattenLogFortianalyzer3OverrideSettingReliable(o["reliable"], d, "reliable", sv)); err != nil {
		if !fortiAPIPatch(o["reliable"]) {
			return fmt.Errorf("Error reading reliable: %v", err)
		}
	}

	if err = d.Set("priority", flattenLogFortianalyzer3OverrideSettingPriority(o["priority"], d, "priority", sv)); err != nil {
		if !fortiAPIPatch(o["priority"]) {
			return fmt.Errorf("Error reading priority: %v", err)
		}
	}

	if err = d.Set("max_log_rate", flattenLogFortianalyzer3OverrideSettingMaxLogRate(o["max-log-rate"], d, "max_log_rate", sv)); err != nil {
		if !fortiAPIPatch(o["max-log-rate"]) {
			return fmt.Errorf("Error reading max_log_rate: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenLogFortianalyzer3OverrideSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogFortianalyzer3OverrideSettingInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzer3OverrideSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogFortianalyzer3OverrideSettingOverride(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUseManagementVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingServer(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingCertificateVerification(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingSerial(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
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

			tmp["name"], _ = expandLogFortianalyzer3OverrideSettingSerialName(d, i["name"], pre_append, sv)
		}

		result = append(result, tmp)

		con += 1
	}

	return result, nil
}

func expandLogFortianalyzer3OverrideSettingSerialName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingPresharedKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingAccessConfig(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingHmacAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingConnTimeout(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingMonitorKeepalivePeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingMonitorFailureRetryPeriod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingMgmtName(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingFazType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingCertificate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSetting__Change_Ip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUploadOption(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUploadDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUploadTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingReliable(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingPriority(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingMaxLogRate(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzer3OverrideSetting(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("override"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingOverride(d, v, "override", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("use_management_vdom"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingUseManagementVdom(d, v, "use_management_vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-management-vdom"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingStatus(d, v, "status", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingIpsArchive(d, v, "ips_archive", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-archive"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingServer(d, v, "server", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("certificate_verification"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingCertificateVerification(d, v, "certificate_verification", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate-verification"] = t
		}
	}

	if v, ok := d.GetOk("serial"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingSerial(d, v, "serial", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["serial"] = t
		}
	}

	if v, ok := d.GetOk("preshared_key"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingPresharedKey(d, v, "preshared_key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["preshared-key"] = t
		}
	}

	if v, ok := d.GetOk("access_config"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingAccessConfig(d, v, "access_config", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["access-config"] = t
		}
	}

	if v, ok := d.GetOk("hmac_algorithm"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingHmacAlgorithm(d, v, "hmac_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hmac-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingEncAlgorithm(d, v, "enc_algorithm", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingSslMinProtoVersion(d, v, "ssl_min_proto_version", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingConnTimeout(d, v, "conn_timeout", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-timeout"] = t
		}
	}

	if v, ok := d.GetOk("monitor_keepalive_period"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingMonitorKeepalivePeriod(d, v, "monitor_keepalive_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-keepalive-period"] = t
		}
	}

	if v, ok := d.GetOk("monitor_failure_retry_period"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingMonitorFailureRetryPeriod(d, v, "monitor_failure_retry_period", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-failure-retry-period"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_name"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingMgmtName(d, v, "mgmt_name", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt-name"] = t
		}
	}

	if v, ok := d.GetOkExists("faz_type"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingFazType(d, v, "faz_type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-type"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingCertificate(d, v, "certificate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingSourceIp(d, v, "source_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOkExists("__change_ip"); ok {

		t, err := expandLogFortianalyzer3OverrideSetting__Change_Ip(d, v, "__change_ip", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["__change_ip"] = t
		}
	}

	if v, ok := d.GetOk("upload_option"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingUploadOption(d, v, "upload_option", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-option"] = t
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingUploadInterval(d, v, "upload_interval", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-interval"] = t
		}
	}

	if v, ok := d.GetOk("upload_day"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingUploadDay(d, v, "upload_day", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-day"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingUploadTime(d, v, "upload_time", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	if v, ok := d.GetOk("reliable"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingReliable(d, v, "reliable", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reliable"] = t
		}
	}

	if v, ok := d.GetOk("priority"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingPriority(d, v, "priority", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["priority"] = t
		}
	}

	if v, ok := d.GetOkExists("max_log_rate"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingMaxLogRate(d, v, "max_log_rate", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-rate"] = t
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingInterfaceSelectMethod(d, v, "interface_select_method", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface-select-method"] = t
		}
	}

	if v, ok := d.GetOk("interface"); ok {

		t, err := expandLogFortianalyzer3OverrideSettingInterface(d, v, "interface", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["interface"] = t
		}
	}

	return &obj, nil
}
