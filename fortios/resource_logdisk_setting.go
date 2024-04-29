// Copyright 2020 Fortinet, Inc. All rights reserved.
// Author: Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu)
// Documentation:
// Frank Shen (@frankshen01), Hongbin Lu (@fgtdev-hblu),
// Xing Li (@lix-fortinet), Yue Wang (@yuew-ftnt), Yuffie Zhu (@yuffiezhu)

// Description: Settings for local disk logging.

package fortios

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
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
			"vdomparam": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
				Computed: true,
			},
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
				Sensitive:    true,
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
		},
	}
}

func resourceLogDiskSettingUpdate(d *schema.ResourceData, m interface{}) error {
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

	obj, err := getObjectLogDiskSetting(d, false, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating LogDiskSetting resource while getting object: %v", err)
	}

	o, err := c.UpdateLogDiskSetting(obj, mkey, vdomparam)
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

	vdomparam := ""

	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	obj, err := getObjectLogDiskSetting(d, true, c.Fv)

	if err != nil {
		return fmt.Errorf("Error updating LogDiskSetting resource while getting object: %v", err)
	}

	_, err = c.UpdateLogDiskSetting(obj, mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error clearing LogDiskSetting resource: %v", err)
	}

	d.SetId("")

	return nil
}

func resourceLogDiskSettingRead(d *schema.ResourceData, m interface{}) error {
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

	o, err := c.ReadLogDiskSetting(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskSetting resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectLogDiskSetting(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading LogDiskSetting resource from API: %v", err)
	}
	return nil
}

func flattenLogDiskSettingStatus(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingIpsArchive(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingMaxLogFileSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingMaxPolicyPacketCaptureSize(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingRollSchedule(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingRollDay(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingRollTime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingDiskfull(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingLogQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingDlpArchiveQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingReportQuota(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingMaximumLogAge(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUpload(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadDestination(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadip(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadport(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingSourceIp(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploaduser(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadpass(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploaddir(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadtype(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadsched(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadtime(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadDeleteFiles(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingUploadSslConn(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingFullFirstWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingFullSecondWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingFullFinalWarningThreshold(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingInterfaceSelectMethod(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenLogDiskSettingInterface(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectLogDiskSetting(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("status", flattenLogDiskSettingStatus(o["status"], d, "status", sv)); err != nil {
		if !fortiAPIPatch(o["status"]) {
			return fmt.Errorf("Error reading status: %v", err)
		}
	}

	if err = d.Set("ips_archive", flattenLogDiskSettingIpsArchive(o["ips-archive"], d, "ips_archive", sv)); err != nil {
		if !fortiAPIPatch(o["ips-archive"]) {
			return fmt.Errorf("Error reading ips_archive: %v", err)
		}
	}

	if err = d.Set("max_log_file_size", flattenLogDiskSettingMaxLogFileSize(o["max-log-file-size"], d, "max_log_file_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-log-file-size"]) {
			return fmt.Errorf("Error reading max_log_file_size: %v", err)
		}
	}

	if err = d.Set("max_policy_packet_capture_size", flattenLogDiskSettingMaxPolicyPacketCaptureSize(o["max-policy-packet-capture-size"], d, "max_policy_packet_capture_size", sv)); err != nil {
		if !fortiAPIPatch(o["max-policy-packet-capture-size"]) {
			return fmt.Errorf("Error reading max_policy_packet_capture_size: %v", err)
		}
	}

	if err = d.Set("roll_schedule", flattenLogDiskSettingRollSchedule(o["roll-schedule"], d, "roll_schedule", sv)); err != nil {
		if !fortiAPIPatch(o["roll-schedule"]) {
			return fmt.Errorf("Error reading roll_schedule: %v", err)
		}
	}

	if err = d.Set("roll_day", flattenLogDiskSettingRollDay(o["roll-day"], d, "roll_day", sv)); err != nil {
		if !fortiAPIPatch(o["roll-day"]) {
			return fmt.Errorf("Error reading roll_day: %v", err)
		}
	}

	if err = d.Set("roll_time", flattenLogDiskSettingRollTime(o["roll-time"], d, "roll_time", sv)); err != nil {
		if !fortiAPIPatch(o["roll-time"]) {
			return fmt.Errorf("Error reading roll_time: %v", err)
		}
	}

	if err = d.Set("diskfull", flattenLogDiskSettingDiskfull(o["diskfull"], d, "diskfull", sv)); err != nil {
		if !fortiAPIPatch(o["diskfull"]) {
			return fmt.Errorf("Error reading diskfull: %v", err)
		}
	}

	if err = d.Set("log_quota", flattenLogDiskSettingLogQuota(o["log-quota"], d, "log_quota", sv)); err != nil {
		if !fortiAPIPatch(o["log-quota"]) {
			return fmt.Errorf("Error reading log_quota: %v", err)
		}
	}

	if err = d.Set("dlp_archive_quota", flattenLogDiskSettingDlpArchiveQuota(o["dlp-archive-quota"], d, "dlp_archive_quota", sv)); err != nil {
		if !fortiAPIPatch(o["dlp-archive-quota"]) {
			return fmt.Errorf("Error reading dlp_archive_quota: %v", err)
		}
	}

	if err = d.Set("report_quota", flattenLogDiskSettingReportQuota(o["report-quota"], d, "report_quota", sv)); err != nil {
		if !fortiAPIPatch(o["report-quota"]) {
			return fmt.Errorf("Error reading report_quota: %v", err)
		}
	}

	if err = d.Set("maximum_log_age", flattenLogDiskSettingMaximumLogAge(o["maximum-log-age"], d, "maximum_log_age", sv)); err != nil {
		if !fortiAPIPatch(o["maximum-log-age"]) {
			return fmt.Errorf("Error reading maximum_log_age: %v", err)
		}
	}

	if err = d.Set("upload", flattenLogDiskSettingUpload(o["upload"], d, "upload", sv)); err != nil {
		if !fortiAPIPatch(o["upload"]) {
			return fmt.Errorf("Error reading upload: %v", err)
		}
	}

	if err = d.Set("upload_destination", flattenLogDiskSettingUploadDestination(o["upload-destination"], d, "upload_destination", sv)); err != nil {
		if !fortiAPIPatch(o["upload-destination"]) {
			return fmt.Errorf("Error reading upload_destination: %v", err)
		}
	}

	if err = d.Set("uploadip", flattenLogDiskSettingUploadip(o["uploadip"], d, "uploadip", sv)); err != nil {
		if !fortiAPIPatch(o["uploadip"]) {
			return fmt.Errorf("Error reading uploadip: %v", err)
		}
	}

	if err = d.Set("uploadport", flattenLogDiskSettingUploadport(o["uploadport"], d, "uploadport", sv)); err != nil {
		if !fortiAPIPatch(o["uploadport"]) {
			return fmt.Errorf("Error reading uploadport: %v", err)
		}
	}

	if err = d.Set("source_ip", flattenLogDiskSettingSourceIp(o["source-ip"], d, "source_ip", sv)); err != nil {
		if !fortiAPIPatch(o["source-ip"]) {
			return fmt.Errorf("Error reading source_ip: %v", err)
		}
	}

	if err = d.Set("uploaduser", flattenLogDiskSettingUploaduser(o["uploaduser"], d, "uploaduser", sv)); err != nil {
		if !fortiAPIPatch(o["uploaduser"]) {
			return fmt.Errorf("Error reading uploaduser: %v", err)
		}
	}

	if err = d.Set("uploaddir", flattenLogDiskSettingUploaddir(o["uploaddir"], d, "uploaddir", sv)); err != nil {
		if !fortiAPIPatch(o["uploaddir"]) {
			return fmt.Errorf("Error reading uploaddir: %v", err)
		}
	}

	if err = d.Set("uploadtype", flattenLogDiskSettingUploadtype(o["uploadtype"], d, "uploadtype", sv)); err != nil {
		if !fortiAPIPatch(o["uploadtype"]) {
			return fmt.Errorf("Error reading uploadtype: %v", err)
		}
	}

	if err = d.Set("uploadsched", flattenLogDiskSettingUploadsched(o["uploadsched"], d, "uploadsched", sv)); err != nil {
		if !fortiAPIPatch(o["uploadsched"]) {
			return fmt.Errorf("Error reading uploadsched: %v", err)
		}
	}

	if err = d.Set("uploadtime", flattenLogDiskSettingUploadtime(o["uploadtime"], d, "uploadtime", sv)); err != nil {
		if !fortiAPIPatch(o["uploadtime"]) {
			return fmt.Errorf("Error reading uploadtime: %v", err)
		}
	}

	if err = d.Set("upload_delete_files", flattenLogDiskSettingUploadDeleteFiles(o["upload-delete-files"], d, "upload_delete_files", sv)); err != nil {
		if !fortiAPIPatch(o["upload-delete-files"]) {
			return fmt.Errorf("Error reading upload_delete_files: %v", err)
		}
	}

	if err = d.Set("upload_ssl_conn", flattenLogDiskSettingUploadSslConn(o["upload-ssl-conn"], d, "upload_ssl_conn", sv)); err != nil {
		if !fortiAPIPatch(o["upload-ssl-conn"]) {
			return fmt.Errorf("Error reading upload_ssl_conn: %v", err)
		}
	}

	if err = d.Set("full_first_warning_threshold", flattenLogDiskSettingFullFirstWarningThreshold(o["full-first-warning-threshold"], d, "full_first_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["full-first-warning-threshold"]) {
			return fmt.Errorf("Error reading full_first_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_second_warning_threshold", flattenLogDiskSettingFullSecondWarningThreshold(o["full-second-warning-threshold"], d, "full_second_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["full-second-warning-threshold"]) {
			return fmt.Errorf("Error reading full_second_warning_threshold: %v", err)
		}
	}

	if err = d.Set("full_final_warning_threshold", flattenLogDiskSettingFullFinalWarningThreshold(o["full-final-warning-threshold"], d, "full_final_warning_threshold", sv)); err != nil {
		if !fortiAPIPatch(o["full-final-warning-threshold"]) {
			return fmt.Errorf("Error reading full_final_warning_threshold: %v", err)
		}
	}

	if err = d.Set("interface_select_method", flattenLogDiskSettingInterfaceSelectMethod(o["interface-select-method"], d, "interface_select_method", sv)); err != nil {
		if !fortiAPIPatch(o["interface-select-method"]) {
			return fmt.Errorf("Error reading interface_select_method: %v", err)
		}
	}

	if err = d.Set("interface", flattenLogDiskSettingInterface(o["interface"], d, "interface", sv)); err != nil {
		if !fortiAPIPatch(o["interface"]) {
			return fmt.Errorf("Error reading interface: %v", err)
		}
	}

	return nil
}

func flattenLogDiskSettingFortiTestDebug(d *schema.ResourceData, fosdebugsn int, fosdebugbeg int, fosdebugend int) {
	log.Printf(strconv.Itoa(fosdebugsn))
	e := validation.IntBetween(fosdebugbeg, fosdebugend)
	log.Printf("ER List: %v, %v", strings.Split("FortiOS Ver", " "), e)
}

func expandLogDiskSettingStatus(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingIpsArchive(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaxLogFileSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaxPolicyPacketCaptureSize(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollSchedule(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollDay(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingRollTime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingDiskfull(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingLogQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingDlpArchiveQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingReportQuota(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingMaximumLogAge(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUpload(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadDestination(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadip(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadport(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingSourceIp(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploaduser(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadpass(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploaddir(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadtype(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadsched(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadtime(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadDeleteFiles(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingUploadSslConn(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullFirstWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullSecondWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingFullFinalWarningThreshold(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingInterfaceSelectMethod(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandLogDiskSettingInterface(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectLogDiskSetting(d *schema.ResourceData, setArgNil bool, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("status"); ok {
		if setArgNil {
			obj["status"] = nil
		} else {
			t, err := expandLogDiskSettingStatus(d, v, "status", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["status"] = t
			}
		}
	}

	if v, ok := d.GetOk("ips_archive"); ok {
		if setArgNil {
			obj["ips-archive"] = nil
		} else {
			t, err := expandLogDiskSettingIpsArchive(d, v, "ips_archive", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["ips-archive"] = t
			}
		}
	}

	if v, ok := d.GetOk("max_log_file_size"); ok {
		if setArgNil {
			obj["max-log-file-size"] = nil
		} else {
			t, err := expandLogDiskSettingMaxLogFileSize(d, v, "max_log_file_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-log-file-size"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("max_policy_packet_capture_size"); ok {
		if setArgNil {
			obj["max-policy-packet-capture-size"] = nil
		} else {
			t, err := expandLogDiskSettingMaxPolicyPacketCaptureSize(d, v, "max_policy_packet_capture_size", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["max-policy-packet-capture-size"] = t
			}
		}
	}

	if v, ok := d.GetOk("roll_schedule"); ok {
		if setArgNil {
			obj["roll-schedule"] = nil
		} else {
			t, err := expandLogDiskSettingRollSchedule(d, v, "roll_schedule", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["roll-schedule"] = t
			}
		}
	}

	if v, ok := d.GetOk("roll_day"); ok {
		if setArgNil {
			obj["roll-day"] = nil
		} else {
			t, err := expandLogDiskSettingRollDay(d, v, "roll_day", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["roll-day"] = t
			}
		}
	}

	if v, ok := d.GetOk("roll_time"); ok {
		if setArgNil {
			obj["roll-time"] = nil
		} else {
			t, err := expandLogDiskSettingRollTime(d, v, "roll_time", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["roll-time"] = t
			}
		}
	}

	if v, ok := d.GetOk("diskfull"); ok {
		if setArgNil {
			obj["diskfull"] = nil
		} else {
			t, err := expandLogDiskSettingDiskfull(d, v, "diskfull", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["diskfull"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("log_quota"); ok {
		if setArgNil {
			obj["log-quota"] = nil
		} else {
			t, err := expandLogDiskSettingLogQuota(d, v, "log_quota", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["log-quota"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("dlp_archive_quota"); ok {
		if setArgNil {
			obj["dlp-archive-quota"] = nil
		} else {
			t, err := expandLogDiskSettingDlpArchiveQuota(d, v, "dlp_archive_quota", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["dlp-archive-quota"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("report_quota"); ok {
		if setArgNil {
			obj["report-quota"] = nil
		} else {
			t, err := expandLogDiskSettingReportQuota(d, v, "report_quota", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["report-quota"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("maximum_log_age"); ok {
		if setArgNil {
			obj["maximum-log-age"] = nil
		} else {
			t, err := expandLogDiskSettingMaximumLogAge(d, v, "maximum_log_age", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["maximum-log-age"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload"); ok {
		if setArgNil {
			obj["upload"] = nil
		} else {
			t, err := expandLogDiskSettingUpload(d, v, "upload", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_destination"); ok {
		if setArgNil {
			obj["upload-destination"] = nil
		} else {
			t, err := expandLogDiskSettingUploadDestination(d, v, "upload_destination", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-destination"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadip"); ok {
		if setArgNil {
			obj["uploadip"] = nil
		} else {
			t, err := expandLogDiskSettingUploadip(d, v, "uploadip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadip"] = t
			}
		}
	}

	if v, ok := d.GetOkExists("uploadport"); ok {
		if setArgNil {
			obj["uploadport"] = nil
		} else {
			t, err := expandLogDiskSettingUploadport(d, v, "uploadport", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadport"] = t
			}
		}
	}

	if v, ok := d.GetOk("source_ip"); ok {
		if setArgNil {
			obj["source-ip"] = nil
		} else {
			t, err := expandLogDiskSettingSourceIp(d, v, "source_ip", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["source-ip"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploaduser"); ok {
		if setArgNil {
			obj["uploaduser"] = nil
		} else {
			t, err := expandLogDiskSettingUploaduser(d, v, "uploaduser", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploaduser"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadpass"); ok {
		if setArgNil {
			obj["uploadpass"] = nil
		} else {
			t, err := expandLogDiskSettingUploadpass(d, v, "uploadpass", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadpass"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploaddir"); ok {
		if setArgNil {
			obj["uploaddir"] = nil
		} else {
			t, err := expandLogDiskSettingUploaddir(d, v, "uploaddir", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploaddir"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadtype"); ok {
		if setArgNil {
			obj["uploadtype"] = nil
		} else {
			t, err := expandLogDiskSettingUploadtype(d, v, "uploadtype", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadtype"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadsched"); ok {
		if setArgNil {
			obj["uploadsched"] = nil
		} else {
			t, err := expandLogDiskSettingUploadsched(d, v, "uploadsched", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadsched"] = t
			}
		}
	}

	if v, ok := d.GetOk("uploadtime"); ok {
		if setArgNil {
			obj["uploadtime"] = nil
		} else {
			t, err := expandLogDiskSettingUploadtime(d, v, "uploadtime", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["uploadtime"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_delete_files"); ok {
		if setArgNil {
			obj["upload-delete-files"] = nil
		} else {
			t, err := expandLogDiskSettingUploadDeleteFiles(d, v, "upload_delete_files", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-delete-files"] = t
			}
		}
	}

	if v, ok := d.GetOk("upload_ssl_conn"); ok {
		if setArgNil {
			obj["upload-ssl-conn"] = nil
		} else {
			t, err := expandLogDiskSettingUploadSslConn(d, v, "upload_ssl_conn", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["upload-ssl-conn"] = t
			}
		}
	}

	if v, ok := d.GetOk("full_first_warning_threshold"); ok {
		if setArgNil {
			obj["full-first-warning-threshold"] = nil
		} else {
			t, err := expandLogDiskSettingFullFirstWarningThreshold(d, v, "full_first_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["full-first-warning-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("full_second_warning_threshold"); ok {
		if setArgNil {
			obj["full-second-warning-threshold"] = nil
		} else {
			t, err := expandLogDiskSettingFullSecondWarningThreshold(d, v, "full_second_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["full-second-warning-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("full_final_warning_threshold"); ok {
		if setArgNil {
			obj["full-final-warning-threshold"] = nil
		} else {
			t, err := expandLogDiskSettingFullFinalWarningThreshold(d, v, "full_final_warning_threshold", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["full-final-warning-threshold"] = t
			}
		}
	}

	if v, ok := d.GetOk("interface_select_method"); ok {
		if setArgNil {
			obj["interface-select-method"] = nil
		} else {
			t, err := expandLogDiskSettingInterfaceSelectMethod(d, v, "interface_select_method", sv)
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
			t, err := expandLogDiskSettingInterface(d, v, "interface", sv)
			if err != nil {
				return &obj, err
			} else if t != nil {
				obj["interface"] = t
			}
		}
	}

	return &obj, nil
}
