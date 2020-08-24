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

func resourceLogFortianalyzerOverrideSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogFortianalyzerOverrideSettingUpdate,
		Read:   resourceLogFortianalyzerOverrideSettingRead,
		Update: resourceLogFortianalyzerOverrideSettingUpdate,
		Delete: resourceLogFortianalyzerOverrideSettingDelete,

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

func resourceLogFortianalyzerOverrideSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogFortianalyzerOverrideSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerOverrideSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogFortianalyzerOverrideSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogFortianalyzerOverrideSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogFortianalyzerOverrideSetting")
	}

	return resourceLogFortianalyzerOverrideSettingRead(d, m)
}

func resourceLogFortianalyzerOverrideSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteLogFortianalyzerOverrideSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting LogFortianalyzerOverrideSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogFortianalyzerOverrideSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogFortianalyzerOverrideSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerOverrideSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogFortianalyzerOverrideSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogFortianalyzerOverrideSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogFortianalyzerOverrideSettingOverride(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingUseManagementVdom(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingServer(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingHmacAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingEncAlgorithm(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingSslMinProtoVersion(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingConnTimeout(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingMonitorKeepalivePeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingMonitorFailureRetryPeriod(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingMgmtName(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingFazType(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingCertificate(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSetting__Change_Ip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingUploadOption(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingUploadInterval(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingUploadDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingUploadTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogFortianalyzerOverrideSettingReliable(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogFortianalyzerOverrideSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("override", flattenLogFortianalyzerOverrideSettingOverride(o["override"], d, "override")); err != nil {
		if !fortiAPIPatch(o["override"]) {
			return fmt.Errorf("Error reading override: %v", err)
		}
	}

	if err = d.Set("use_management_vdom", flattenLogFortianalyzerOverrideSettingUseManagementVdom(o["use-management-vdom"], d, "use_management_vdom")); err != nil {
		if !fortiAPIPatch(o["use-management-vdom"]) {
			return fmt.Errorf("Error reading use_management_vdom: %v", err)
		}
	}

	if err = d.Set("status", flattenLogFortianalyzerOverrideSettingStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogFortianalyzerOverrideSettingIpsArchive(o["ips-archive"], d, "ips_archive")); err != nil {
		if !fortiAPIPatch(o["ips-archive"]) {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("server", flattenLogFortianalyzerOverrideSettingServer(o["server"], d, "server")); err != nil {
		if !fortiAPIPatch(o["server"]) {
			return fmt.Errorf("Error reading server: %v", err)
		}
	}

	if err = d.Set("hmac_algorithm", flattenLogFortianalyzerOverrideSettingHmacAlgorithm(o["hmac-algorithm"], d, "hmac_algorithm")); err != nil {
		if !fortiAPIPatch(o["hmac-algorithm"]) {
			return fmt.Errorf("Error reading hmac_algorithm: %v", err)
		}
	}

	if err = d.Set("enc_algorithm", flattenLogFortianalyzerOverrideSettingEncAlgorithm(o["enc-algorithm"], d, "enc_algorithm")); err != nil {
		if !fortiAPIPatch(o["enc-algorithm"]) {
			return fmt.Errorf("Error reading enc_algorithm: %v", err)
		}
	}

	if err = d.Set("ssl_min_proto_version", flattenLogFortianalyzerOverrideSettingSslMinProtoVersion(o["ssl-min-proto-version"], d, "ssl_min_proto_version")); err != nil {
		if !fortiAPIPatch(o["ssl-min-proto-version"]) {
			return fmt.Errorf("Error reading ssl_min_proto_version: %v", err)
		}
	}

	if err = d.Set("conn_timeout", flattenLogFortianalyzerOverrideSettingConnTimeout(o["conn-timeout"], d, "conn_timeout")); err != nil {
		if !fortiAPIPatch(o["conn-timeout"]) {
			return fmt.Errorf("Error reading conn_timeout: %v", err)
		}
	}

	if err = d.Set("monitor_keepalive_period", flattenLogFortianalyzerOverrideSettingMonitorKeepalivePeriod(o["monitor-keepalive-period"], d, "monitor_keepalive_period")); err != nil {
		if !fortiAPIPatch(o["monitor-keepalive-period"]) {
			return fmt.Errorf("Error reading monitor_keepalive_period: %v", err)
		}
	}

	if err = d.Set("monitor_failure_retry_period", flattenLogFortianalyzerOverrideSettingMonitorFailureRetryPeriod(o["monitor-failure-retry-period"], d, "monitor_failure_retry_period")); err != nil {
		if !fortiAPIPatch(o["monitor-failure-retry-period"]) {
			return fmt.Errorf("Error reading monitor_failure_retry_period: %v", err)
		}
	}

	if err = d.Set("mgmt_name", flattenLogFortianalyzerOverrideSettingMgmtName(o["mgmt-name"], d, "mgmt_name")); err != nil {
		if !fortiAPIPatch(o["mgmt-name"]) {
			return fmt.Errorf("Error reading mgmt_name: %v", err)
		}
	}

	if err = d.Set("faz_type", flattenLogFortianalyzerOverrideSettingFazType(o["faz-type"], d, "faz_type")); err != nil {
		if !fortiAPIPatch(o["faz-type"]) {
			return fmt.Errorf("Error reading faz_type: %v", err)
		}
	}

	if err = d.Set("certificate", flattenLogFortianalyzerOverrideSettingCertificate(o["certificate"], d, "certificate")); err != nil {
		if !fortiAPIPatch(o["certificate"]) {
			return fmt.Errorf("Error reading certificate: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogFortianalyzerOverrideSettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("__change_ip", flattenLogFortianalyzerOverrideSetting__Change_Ip(o["__change_ip"], d, "__change_ip")); err != nil {
		if !fortiAPIPatch(o["__change_ip"]) {
			return fmt.Errorf("Error reading __change_ip: %v", err)
		}
	}

	if err = d.Set("upload_option", flattenLogFortianalyzerOverrideSettingUploadOption(o["upload-option"], d, "upload_option")); err != nil {
		if !fortiAPIPatch(o["upload-option"]) {
			return fmt.Errorf("Error reading upload_option: %v", err)
		}
	}

	if err = d.Set("upload_interval", flattenLogFortianalyzerOverrideSettingUploadInterval(o["upload-interval"], d, "upload_interval")); err != nil {
		if !fortiAPIPatch(o["upload-interval"]) {
			return fmt.Errorf("Error reading upload_interval: %v", err)
		}
	}

	if err = d.Set("upload_day", flattenLogFortianalyzerOverrideSettingUploadDay(o["upload-day"], d, "upload_day")); err != nil {
		if !fortiAPIPatch(o["upload-day"]) {
			return fmt.Errorf("Error reading upload_day: %v", err)
		}
	}

	if err = d.Set("upload_time", flattenLogFortianalyzerOverrideSettingUploadTime(o["upload-time"], d, "upload_time")); err != nil {
		if !fortiAPIPatch(o["upload-time"]) {
			return fmt.Errorf("Error reading upload_time: %v", err)
		}
	}

	if err = d.Set("reliable", flattenLogFortianalyzerOverrideSettingReliable(o["reliable"], d, "reliable")); err != nil {
		if !fortiAPIPatch(o["reliable"]) {
			return fmt.Errorf("Error reading reliable: %v", err)
		}
	}

	return nil
}

func flattenLogFortianalyzerOverrideSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogFortianalyzerOverrideSettingOverride(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingUseManagementVdom(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingServer(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingHmacAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingEncAlgorithm(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingSslMinProtoVersion(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingConnTimeout(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingMonitorKeepalivePeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingMonitorFailureRetryPeriod(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingMgmtName(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingFazType(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingCertificate(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSetting__Change_Ip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingUploadOption(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingUploadInterval(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingUploadDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingUploadTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogFortianalyzerOverrideSettingReliable(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogFortianalyzerOverrideSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("override"); ok {
		t, err := expandLogFortianalyzerOverrideSettingOverride(d, v, "override")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["override"] = t
		}
	}

	if v, ok := d.GetOk("use_management_vdom"); ok {
		t, err := expandLogFortianalyzerOverrideSettingUseManagementVdom(d, v, "use_management_vdom")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["use-management-vdom"] = t
		}
	}

	if v, ok := d.GetOk("status"); ok {
		t, err := expandLogFortianalyzerOverrideSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok {
		t, err := expandLogFortianalyzerOverrideSettingIpsArchive(d, v, "ips_archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-archive"] = t
		}
	}

	if v, ok := d.GetOk("server"); ok {
		t, err := expandLogFortianalyzerOverrideSettingServer(d, v, "server")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["server"] = t
		}
	}

	if v, ok := d.GetOk("hmac_algorithm"); ok {
		t, err := expandLogFortianalyzerOverrideSettingHmacAlgorithm(d, v, "hmac_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["hmac-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("enc_algorithm"); ok {
		t, err := expandLogFortianalyzerOverrideSettingEncAlgorithm(d, v, "enc_algorithm")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["enc-algorithm"] = t
		}
	}

	if v, ok := d.GetOk("ssl_min_proto_version"); ok {
		t, err := expandLogFortianalyzerOverrideSettingSslMinProtoVersion(d, v, "ssl_min_proto_version")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ssl-min-proto-version"] = t
		}
	}

	if v, ok := d.GetOk("conn_timeout"); ok {
		t, err := expandLogFortianalyzerOverrideSettingConnTimeout(d, v, "conn_timeout")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["conn-timeout"] = t
		}
	}

	if v, ok := d.GetOk("monitor_keepalive_period"); ok {
		t, err := expandLogFortianalyzerOverrideSettingMonitorKeepalivePeriod(d, v, "monitor_keepalive_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-keepalive-period"] = t
		}
	}

	if v, ok := d.GetOk("monitor_failure_retry_period"); ok {
		t, err := expandLogFortianalyzerOverrideSettingMonitorFailureRetryPeriod(d, v, "monitor_failure_retry_period")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["monitor-failure-retry-period"] = t
		}
	}

	if v, ok := d.GetOk("mgmt_name"); ok {
		t, err := expandLogFortianalyzerOverrideSettingMgmtName(d, v, "mgmt_name")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mgmt-name"] = t
		}
	}

	if v, ok := d.GetOk("faz_type"); ok {
		t, err := expandLogFortianalyzerOverrideSettingFazType(d, v, "faz_type")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["faz-type"] = t
		}
	}

	if v, ok := d.GetOk("certificate"); ok {
		t, err := expandLogFortianalyzerOverrideSettingCertificate(d, v, "certificate")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["certificate"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandLogFortianalyzerOverrideSettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("__change_ip"); ok {
		t, err := expandLogFortianalyzerOverrideSetting__Change_Ip(d, v, "__change_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["__change_ip"] = t
		}
	}

	if v, ok := d.GetOk("upload_option"); ok {
		t, err := expandLogFortianalyzerOverrideSettingUploadOption(d, v, "upload_option")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-option"] = t
		}
	}

	if v, ok := d.GetOk("upload_interval"); ok {
		t, err := expandLogFortianalyzerOverrideSettingUploadInterval(d, v, "upload_interval")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-interval"] = t
		}
	}

	if v, ok := d.GetOk("upload_day"); ok {
		t, err := expandLogFortianalyzerOverrideSettingUploadDay(d, v, "upload_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-day"] = t
		}
	}

	if v, ok := d.GetOk("upload_time"); ok {
		t, err := expandLogFortianalyzerOverrideSettingUploadTime(d, v, "upload_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-time"] = t
		}
	}

	if v, ok := d.GetOk("reliable"); ok {
		t, err := expandLogFortianalyzerOverrideSettingReliable(d, v, "reliable")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["reliable"] = t
		}
	}

	return &obj, nil
}
