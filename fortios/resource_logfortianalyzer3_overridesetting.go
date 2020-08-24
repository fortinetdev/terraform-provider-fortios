// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Override FortiAnalyzer settings.

package fortios

import (
	"fmt"
	"log"
	"strconv"

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
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
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
		},
	}
}

func resourceLogFortianalyzer3OverrideSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortianalyzer3OverrideSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzer3OverrideSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzer3OverrideSetting(obj, mkey)
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

	err := c.DeleteLogFortianalyzer3OverrideSetting(mkey)
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

	o, err := c.ReadLogFortianalyzer3OverrideSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3OverrideSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzer3OverrideSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzer3OverrideSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzer3OverrideSettingOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUseManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingHmacAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingConnTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingMonitorKeepalivePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingMonitorFailureRetryPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingMgmtName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingFazType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSetting__Change_Ip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUploadOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUploadDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzer3OverrideSettingReliable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogFortianalyzer3OverrideSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("override", flattenLogFortianalyzer3OverrideSettingOverride(o["override"], d, "override")); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("use_management_vdom", flattenLogFortianalyzer3OverrideSettingUseManagementVdom(o["use-management-vdom"], d, "use_management_vdom")); err != nil {
		if !fortiAPIPatch(o["use-management-vdom"]) {
			return fmt.Errorf("Error reading use_management_vdom: %v", err)
		}
	}

	if err = d.Set("status", flattenLogFortianalyzer3OverrideSettingStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogFortianalyzer3OverrideSettingIpsArchive(o["ips-archive"], d, "ips_archive")); err != nil {
		if !fortiAPIPatch(o["ips-archive"]) {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("server", flattenLogFortianalyzer3OverrideSettingServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("hmac_algorithm", flattenLogFortianalyzer3OverrideSettingHmacAlgorithm(o["hmac-algorithm"], d, "hmac_algorithm")); err != nil {
		if !fortiAPIPatch(o["hmac-algorithm"]) {
			return fmt.Errorf("Error reading hmac_algorithm: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortianalyzer3OverrideSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogFortianalyzer3OverrideSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenLogFortianalyzer3OverrideSettingConnTimeout(o["conn-timeout"], d, "conn_timeout")); err != nil {
		if !fortiAPIPatch(o["conn-timeout"]) {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("monitor_keepalive_period", flattenLogFortianalyzer3OverrideSettingMonitorKeepalivePeriod(o["monitor-keepalive-period"], d, "monitor_keepalive_period")); err != nil {
		if !fortiAPIPatch(o["monitor-keepalive-period"]) {
			return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
		}
	}

	if err = d.Set("monitor_failure_retry_period", flattenLogFortianalyzer3OverrideSettingMonitorFailureRetryPeriod(o["monitor-failure-retry-period"], d, "monitor_failure_retry_period")); err != nil {
		if !fortiAPIPatch(o["monitor-failure-retry-period"]) {
			return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
		}
	}

	if err = d.Set("mgmt_name", flattenLogFortianalyzer3OverrideSettingMgmtName(o["mgmt-name"], d, "mgmt_name")); err != nil {
		if !fortiAPIPatch(o["mgmt-name"]) {
			return fmt.Errorf("Error reading mgmt_name: %v", err)
		}
	}

	if err = d.Set("faz_type", flattenLogFortianalyzer3OverrideSettingFazType(o["faz-type"], d, "faz_type")); err != nil {
		if !fortiAPIPatch(o["faz-type"]) {
			return fmt.Errorf("Error reading faz_type: %v", err)
		}
	}

	if err = d.Set("certificate", flattenLogFortianalyzer3OverrideSettingCertificate(o["certificate"], d, "certificate")); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogFortianalyzer3OverrideSettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("__change_ip", flattenLogFortianalyzer3OverrideSetting__Change_Ip(o["__change_ip"], d, "__change_ip")); err != nil {
		if !fortiAPIPatch(o["__change_ip"]) {
			return fmt.Errorf("Error reading __change_ip: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenLogFortianalyzer3OverrideSettingUploadOption(o["upload-option"], d, "upload_option")); err != nil {
		if !fortiAPIPatch(o["upload-option"]) {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenLogFortianalyzer3OverrideSettingUploadInterval(o["upload-interval"], d, "upload_interval")); err != nil {
		if !fortiAPIPatch(o["upload-interval"]) {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortianalyzer3OverrideSettingUploadDay(o["upload-day"], d, "upload_day")); err != nil {
		if !fortiAPIPatch(o["upload-day"]) {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortianalyzer3OverrideSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if !fortiAPIPatch(o["upload-time"]) {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("reliable", flattenLogFortianalyzer3OverrideSettingReliable(o["reliable"], d, "reliable")); err != nil {
		if !fortiAPIPatch(o["reliable"]) {
			return fmt.Errorf("Error reading reliable: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzer3OverrideSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogFortianalyzer3OverrideSettingOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUseManagementVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingHmacAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingConnTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingMonitorKeepalivePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingMonitorFailureRetryPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingMgmtName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingFazType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSetting__Change_Ip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUploadOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUploadDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzer3OverrideSettingReliable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzer3OverrideSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("override"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("use_management_vdom"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingUseManagementVdom(d, v, "use_management_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-management-vdom"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingIpsArchive(d, v, "ips_archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-archive"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("hmac_algorithm"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingHmacAlgorithm(d, v, "hmac_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hmac-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingConnTimeout(d, v, "conn_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-timeout"] = t
		}
	}

	if v, ok := d.GetOk("monitor_keepalive_period"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingMonitorKeepalivePeriod(d, v, "monitor_keepalive_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-keepalive-period"] = t
		}
	}

	if v, ok := d.GetOk("monitor_failure_retry_period"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingMonitorFailureRetryPeriod(d, v, "monitor_failure_retry_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-failure-retry-period"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_name"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingMgmtName(d, v, "mgmt_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt-name"] = t
		}
	}

	if v, ok := d.GetOk("faz_type"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingFazType(d, v, "faz_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-type"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("__change_ip"); ok {
		t, err := expandLogFortianalyzer3OverrideSetting__Change_Ip(d, v, "__change_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["__change_ip"] = t
		}
	}

	if v, ok := d.GetOk("upload_option"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingUploadOption(d, v, "upload_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-option"] = t
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingUploadInterval(d, v, "upload_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-interval"] = t
		}
	}

	if v, ok := d.GetOk("upload_day"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingUploadDay(d, v, "upload_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-day"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	if v, ok := d.GetOk("reliable"); ok {
		t, err := expandLogFortianalyzer3OverrideSettingReliable(d, v, "reliable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reliable"] = t
		}
	}

	return &obj, nil
}
