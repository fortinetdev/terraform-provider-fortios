// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Yuffie Zhu (@yuffiezhu), Yue Wang (@yuew-ftnt)

// Description: Settings for local disk logging.

package fortios

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
)

func resourceLogDiskSetting() *schema.Resource {
	return &schema.Resource{
		Create: resourceLogDiskSettingUpdate,
		Read:   resourceLogDiskSettingRead,
		Update: resourceLogDiskSettingUpdate,
		Delete: resourceLogDiskSettingDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"status": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"ips_archive": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"max_log_file_size": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 100),
				Optional:     true,
				Computed:     true,
			},
			"max_policy_packet_capture_size": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"roll_schedule": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roll_day": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"roll_time": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"diskfull": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"log_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"dlp_archive_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"report_quota": &schema.Schema{
				Type:     schema.TypeInt,
				Optional: true,
				Computed: true,
			},
			"maximum_log_age": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 3650),
				Optional:     true,
				Computed:     true,
			},
			"upload": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_destination": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploadip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploadport": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(0, 65535),
				Optional:     true,
				Computed:     true,
			},
			"source_ip": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploaduser": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 35),
				Optional:     true,
				Computed:     true,
			},
			"uploadpass": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 128),
				Optional:     true,
			},
			"uploaddir": &schema.Schema{
				Type:         schema.TypeString,
				ValidateFunc: validation.StringLenBetween(0, 63),
				Optional:     true,
				Computed:     true,
			},
			"uploadtype": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploadsched": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"uploadtime": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_delete_files": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"upload_ssl_conn": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
			},
			"full_first_warning_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(1, 98),
				Optional:     true,
				Computed:     true,
			},
			"full_second_warning_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(2, 99),
				Optional:     true,
				Computed:     true,
			},
			"full_final_warning_threshold": &schema.Schema{
				Type:         schema.TypeInt,
				ValidateFunc: validation.IntBetween(3, 100),
				Optional:     true,
				Computed:     true,
			},
		},
	}
}

func resourceLogDiskSettingUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	obj, err := getObjectLogDiskSetting(d)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogDiskSetting(obj, mkey)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskSetting resource: %v", err)
	}

	log.Printf(strconv.Itoa(c.Retries))
	if o["mkey"] != nil && o["mkey"] != "" {
		d.SetId(o["mkey"].(string))
	} else {
		d.SetId("LogDiskSetting")
	}

	return resourceLogDiskSettingRead(d, m)
}

func resourceLogDiskSettingDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	err := c.DeleteLogDiskSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error deleting LogDiskSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogDiskSettingRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	o, err := c.ReadLogDiskSetting(mkey)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogDiskSetting(d, o)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogDiskSettingStatus(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingMaxLogFileSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingMaxPolicyPacketCaptureSize(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingRollSchedule(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingRollDay(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingRollTime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingDiskfull(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingLogQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingDlpArchiveQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingReportQuota(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingMaximumLogAge(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUpload(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadDestination(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadip(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadport(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingSourceIp(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploaduser(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadpass(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploaddir(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadtype(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadsched(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadtime(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadDeleteFiles(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingUploadSslConn(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingFullFirstWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingFullSecondWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func flattenLogDiskSettingFullFinalWarningThreshold(v interface{}, d *schema.ResourceData, pre string) interface{} {
	return v
}

func refreshObjectLogDiskSetting(d *schema.ResourceData, o map[string]interface{}) error {
	var err error

	if err = d.Set("status", flattenLogDiskSettingStatus(o["status"], d, "status")); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogDiskSettingIpsArchive(o["ips-archive"], d, "ips_archive")); err != nil {
		if !fortiAPIPatch(o["ips-archive"]) {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("max_log_file_size", flattenLogDiskSettingMaxLogFileSize(o["max-log-file-size"], d, "max_log_file_size")); err != nil {
		if !fortiAPIPatch(o["max-log-file-size"]) {
			return fmt.Errorf("Error reading max_log_file_size: %v", err)
		}
	}

	if err = d.Set("max_policy_packet_capture_size", flattenLogDiskSettingMaxPolicyPacketCaptureSize(o["max-policy-packet-capture-size"], d, "max_policy_packet_capture_size")); err != nil {
		if !fortiAPIPatch(o["max-policy-packet-capture-size"]) {
			return fmt.Errorf("Error reading max_policy_packet_capture_size: %v", err)
		}
	}

	if err = d.Set("roll_schedule", flattenLogDiskSettingRollSchedule(o["roll-schedule"], d, "roll_schedule")); err != nil {
		if !fortiAPIPatch(o["roll-schedule"]) {
			return fmt.Errorf("Error reading roll_schedule: %v", err)
		}
	}

	if err = d.Set("roll_day", flattenLogDiskSettingRollDay(o["roll-day"], d, "roll_day")); err != nil {
		if !fortiAPIPatch(o["roll-day"]) {
			return fmt.Errorf("Error reading roll_day: %v", err)
		}
	}

	if err = d.Set("roll_time", flattenLogDiskSettingRollTime(o["roll-time"], d, "roll_time")); err != nil {
		if !fortiAPIPatch(o["roll-time"]) {
			return fmt.Errorf("Error reading roll_time: %v", err)
		}
	}

	if err = d.Set("diskfull", flattenLogDiskSettingDiskfull(o["diskfull"], d, "diskfull")); err != nil {
		if !fortiAPIPatch(o["diskfull"]) {
			return fmt.Errorf("Error reading diskfull: %v", err)
		}
	}

	if err = d.Set("log_quota", flattenLogDiskSettingLogQuota(o["log-quota"], d, "log_quota")); err != nil {
		if !fortiAPIPatch(o["log-quota"]) {
			return fmt.Errorf("Error reading log_quota: %v", err)
		}
	}

	if err = d.Set("dlp_archive_quota", flattenLogDiskSettingDlpArchiveQuota(o["dlp-archive-quota"], d, "dlp_archive_quota")); err != nil {
		if !fortiAPIPatch(o["dlp-archive-quota"]) {
			return fmt.Errorf("Error reading dlp_archive_quota: %v", err)
		}
	}

	if err = d.Set("report_quota", flattenLogDiskSettingReportQuota(o["report-quota"], d, "report_quota")); err != nil {
		if !fortiAPIPatch(o["report-quota"]) {
			return fmt.Errorf("Error reading report_quota: %v", err)
		}
	}

	if err = d.Set("maximum_log_age", flattenLogDiskSettingMaximumLogAge(o["maximum-log-age"], d, "maximum_log_age")); err != nil {
		if !fortiAPIPatch(o["maximum-log-age"]) {
			return fmt.Errorf("Error reading maximum_log_age: %v", err)
		}
	}

	if err = d.Set("upload", flattenLogDiskSettingUpload(o["upload"], d, "upload")); err != nil {
		if !fortiAPIPatch(o["upload"]) {
			return fmt.Errorf("Error reading upload: %v", err)
		}
	}

	if err = d.Set("upload_destination", flattenLogDiskSettingUploadDestination(o["upload-destination"], d, "upload_destination")); err != nil {
		if !fortiAPIPatch(o["upload-destination"]) {
			return fmt.Errorf("Error reading upload_destination: %v", err)
		}
	}

	if err = d.Set("uploadip", flattenLogDiskSettingUploadip(o["uploadip"], d, "uploadip")); err != nil {
		if !fortiAPIPatch(o["uploadip"]) {
			return fmt.Errorf("Error reading uploadip: %v", err)
		}
	}

	if err = d.Set("uploadport", flattenLogDiskSettingUploadport(o["uploadport"], d, "uploadport")); err != nil {
		if !fortiAPIPatch(o["uploadport"]) {
			return fmt.Errorf("Error reading uploadport: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogDiskSettingSourceIp(o["source-ip"], d, "source_ip")); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("uploaduser", flattenLogDiskSettingUploaduser(o["uploaduser"], d, "uploaduser")); err != nil {
		if !fortiAPIPatch(o["uploaduser"]) {
			return fmt.Errorf("Error reading uploaduser: %v", err)
		}
	}

	if err = d.Set("uploadpass", flattenLogDiskSettingUploadpass(o["uploadpass"], d, "uploadpass")); err != nil {
		if !fortiAPIPatch(o["uploadpass"]) {
			return fmt.Errorf("Error reading uploadpass: %v", err)
		}
	}

	if err = d.Set("uploaddir", flattenLogDiskSettingUploaddir(o["uploaddir"], d, "uploaddir")); err != nil {
		if !fortiAPIPatch(o["uploaddir"]) {
			return fmt.Errorf("Error reading uploaddir: %v", err)
		}
	}

	if err = d.Set("uploadtype", flattenLogDiskSettingUploadtype(o["uploadtype"], d, "uploadtype")); err != nil {
		if !fortiAPIPatch(o["uploadtype"]) {
			return fmt.Errorf("Error reading uploadtype: %v", err)
		}
	}

	if err = d.Set("uploadsched", flattenLogDiskSettingUploadsched(o["uploadsched"], d, "uploadsched")); err != nil {
		if !fortiAPIPatch(o["uploadsched"]) {
			return fmt.Errorf("Error reading uploadsched: %v", err)
		}
	}

	if err = d.Set("uploadtime", flattenLogDiskSettingUploadtime(o["uploadtime"], d, "uploadtime")); err != nil {
		if !fortiAPIPatch(o["uploadtime"]) {
			return fmt.Errorf("Error reading uploadtime: %v", err)
		}
	}

	if err = d.Set("upload_delete_files", flattenLogDiskSettingUploadDeleteFiles(o["upload-delete-files"], d, "upload_delete_files")); err != nil {
		if !fortiAPIPatch(o["upload-delete-files"]) {
			return fmt.Errorf("Error reading upload_delete_files: %v", err)
		}
	}

	if err = d.Set("upload_ssl_conn", flattenLogDiskSettingUploadSslConn(o["upload-ssl-conn"], d, "upload_ssl_conn")); err != nil {
		if !fortiAPIPatch(o["upload-ssl-conn"]) {
			return fmt.Errorf("Error reading upload_ssl_conn: %v", err)
		}
	}

	if err = d.Set("full_first_warning_threshold", flattenLogDiskSettingFullFirstWarningThreshold(o["full-first-warning-threshold"], d, "full_first_warning_threshold")); err != nil {
		if !fortiAPIPatch(o["full-first-warning-threshold"]) {
			return fmt.Errorf("Error reading full_first_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_second_warning_threshold", flattenLogDiskSettingFullSecondWarningThreshold(o["full-second-warning-threshold"], d, "full_second_warning_threshold")); err != nil {
		if !fortiAPIPatch(o["full-second-warning-threshold"]) {
			return fmt.Errorf("Error reading full_second_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_final_warning_threshold", flattenLogDiskSettingFullFinalWarningThreshold(o["full-final-warning-threshold"], d, "full_final_warning_threshold")); err != nil {
		if !fortiAPIPatch(o["full-final-warning-threshold"]) {
			return fmt.Errorf("Error reading full_final_warning_threshold: %v", err)
		}
	}

	return nil
}

func flattenLogDiskSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v", e)
}

func expandLogDiskSettingStatus(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaxLogFileSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaxPolicyPacketCaptureSize(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollSchedule(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollDay(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollTime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingDiskfull(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingLogQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingDlpArchiveQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingReportQuota(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaximumLogAge(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUpload(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadDestination(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadip(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadport(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingSourceIp(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploaduser(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadpass(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploaddir(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadtype(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadsched(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadtime(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadDeleteFiles(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadSslConn(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullFirstWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullSecondWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullFinalWarningThreshold(d *schema.ResourceData, v interface{}, pre string) (interface{}, error) {
	return v, nil
}

func getObjectLogDiskSetting(d *schema.ResourceData) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		t, err := expandLogDiskSettingStatus(d, v, "status")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["status"] = t
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok {
		t, err := expandLogDiskSettingIpsArchive(d, v, "ips_archive")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["ips-archive"] = t
		}
	}

	if v, ok := d.GetOk("max_log_file_size"); ok {
		t, err := expandLogDiskSettingMaxLogFileSize(d, v, "max_log_file_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-log-file-size"] = t
		}
	}

	if v, ok := d.GetOk("max_policy_packet_capture_size"); ok {
		t, err := expandLogDiskSettingMaxPolicyPacketCaptureSize(d, v, "max_policy_packet_capture_size")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["max-policy-packet-capture-size"] = t
		}
	}

	if v, ok := d.GetOk("roll_schedule"); ok {
		t, err := expandLogDiskSettingRollSchedule(d, v, "roll_schedule")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roll-schedule"] = t
		}
	}

	if v, ok := d.GetOk("roll_day"); ok {
		t, err := expandLogDiskSettingRollDay(d, v, "roll_day")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roll-day"] = t
		}
	}

	if v, ok := d.GetOk("roll_time"); ok {
		t, err := expandLogDiskSettingRollTime(d, v, "roll_time")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["roll-time"] = t
		}
	}

	if v, ok := d.GetOk("diskfull"); ok {
		t, err := expandLogDiskSettingDiskfull(d, v, "diskfull")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["diskfull"] = t
		}
	}

	if v, ok := d.GetOk("log_quota"); ok {
		t, err := expandLogDiskSettingLogQuota(d, v, "log_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["log-quota"] = t
		}
	}

	if v, ok := d.GetOk("dlp_archive_quota"); ok {
		t, err := expandLogDiskSettingDlpArchiveQuota(d, v, "dlp_archive_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["dlp-archive-quota"] = t
		}
	}

	if v, ok := d.GetOk("report_quota"); ok {
		t, err := expandLogDiskSettingReportQuota(d, v, "report_quota")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["report-quota"] = t
		}
	}

	if v, ok := d.GetOk("maximum_log_age"); ok {
		t, err := expandLogDiskSettingMaximumLogAge(d, v, "maximum_log_age")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["maximum-log-age"] = t
		}
	}

	if v, ok := d.GetOk("upload"); ok {
		t, err := expandLogDiskSettingUpload(d, v, "upload")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload"] = t
		}
	}

	if v, ok := d.GetOk("upload_destination"); ok {
		t, err := expandLogDiskSettingUploadDestination(d, v, "upload_destination")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-destination"] = t
		}
	}

	if v, ok := d.GetOk("uploadip"); ok {
		t, err := expandLogDiskSettingUploadip(d, v, "uploadip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadip"] = t
		}
	}

	if v, ok := d.GetOk("uploadport"); ok {
		t, err := expandLogDiskSettingUploadport(d, v, "uploadport")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadport"] = t
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		t, err := expandLogDiskSettingSourceIp(d, v, "source_ip")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["source-ip"] = t
		}
	}

	if v, ok := d.GetOk("uploaduser"); ok {
		t, err := expandLogDiskSettingUploaduser(d, v, "uploaduser")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploaduser"] = t
		}
	}

	if v, ok := d.GetOk("uploadpass"); ok {
		t, err := expandLogDiskSettingUploadpass(d, v, "uploadpass")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadpass"] = t
		}
	}

	if v, ok := d.GetOk("uploaddir"); ok {
		t, err := expandLogDiskSettingUploaddir(d, v, "uploaddir")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploaddir"] = t
		}
	}

	if v, ok := d.GetOk("uploadtype"); ok {
		t, err := expandLogDiskSettingUploadtype(d, v, "uploadtype")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadtype"] = t
		}
	}

	if v, ok := d.GetOk("uploadsched"); ok {
		t, err := expandLogDiskSettingUploadsched(d, v, "uploadsched")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadsched"] = t
		}
	}

	if v, ok := d.GetOk("uploadtime"); ok {
		t, err := expandLogDiskSettingUploadtime(d, v, "uploadtime")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["uploadtime"] = t
		}
	}

	if v, ok := d.GetOk("upload_delete_files"); ok {
		t, err := expandLogDiskSettingUploadDeleteFiles(d, v, "upload_delete_files")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-delete-files"] = t
		}
	}

	if v, ok := d.GetOk("upload_ssl_conn"); ok {
		t, err := expandLogDiskSettingUploadSslConn(d, v, "upload_ssl_conn")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["upload-ssl-conn"] = t
		}
	}

	if v, ok := d.GetOk("full_first_warning_threshold"); ok {
		t, err := expandLogDiskSettingFullFirstWarningThreshold(d, v, "full_first_warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-first-warning-threshold"] = t
		}
	}

	if v, ok := d.GetOk("full_second_warning_threshold"); ok {
		t, err := expandLogDiskSettingFullSecondWarningThreshold(d, v, "full_second_warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-second-warning-threshold"] = t
		}
	}

	if v, ok := d.GetOk("full_final_warning_threshold"); ok {
		t, err := expandLogDiskSettingFullFinalWarningThreshold(d, v, "full_final_warning_threshold")
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["full-final-warning-threshold"] = t
		}
	}

	return &obj, nil
}
