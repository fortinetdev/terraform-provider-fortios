---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logdisk_setting"
description: |-
  Settings for local disk logging.
---

# fortios_logdisk_setting
Settings for local disk logging.

## Example Usage

```hcl
resource "fortios_logdisk_setting" "trname" {
  diskfull                       = "overwrite"
  dlp_archive_quota              = 0
  full_final_warning_threshold   = 95
  full_first_warning_threshold   = 75
  full_second_warning_threshold  = 90
  ips_archive                    = "enable"
  log_quota                      = 0
  max_log_file_size              = 20
  max_policy_packet_capture_size = 100
  maximum_log_age                = 7
  report_quota                   = 0
  roll_day                       = "sunday"
  roll_schedule                  = "daily"
  roll_time                      = "00:00"
  source_ip                      = "0.0.0.0"
  status                         = "enable"
  upload                         = "disable"
  upload_delete_files            = "enable"
  upload_destination             = "ftp-server"
  upload_ssl_conn                = "default"
  uploadip                       = "0.0.0.0"
  uploadport                     = 21
  uploadsched                    = "disable"
  uploadtime                     = "00:00"
  uploadtype                     = "traffic event virus webfilter IPS spamfilter dlp-archive anomaly voip dlp app-ctrl waf netscan gtp dns"
}
```

## Argument Reference

The following arguments are supported:

* `status` - (Required) Enable/disable local disk logging. Valid values: `enable`, `disable`.
* `ips_archive` - Enable/disable IPS packet archiving to the local disk. Valid values: `enable`, `disable`.
* `max_log_file_size` - Maximum log file size before rolling (1 - 100 Mbytes).
* `max_policy_packet_capture_size` - Maximum size of policy sniffer in MB (0 means unlimited).
* `roll_schedule` - Frequency to check log file for rolling. Valid values: `daily`, `weekly`.
* `roll_day` - Day of week on which to roll log file. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `roll_time` - Time of day to roll the log file (hh:mm).
* `diskfull` - Action to take when disk is full. The system can overwrite the oldest log messages or stop logging when the disk is full (default = overwrite). Valid values: `overwrite`, `nolog`.
* `log_quota` - Disk log quota (MB).
* `dlp_archive_quota` - DLP archive quota (MB).
* `report_quota` - Report quota (MB).
* `maximum_log_age` - Delete log files older than (days).
* `upload` - Enable/disable uploading log files when they are rolled. Valid values: `enable`, `disable`.
* `upload_destination` - The type of server to upload log files to. Only FTP is currently supported. Valid values: `ftp-server`.
* `uploadip` - IP address of the FTP server to upload log files to.
* `uploadport` - TCP port to use for communicating with the FTP server (default = 21).
* `source_ip` - Source IP address to use for uploading disk log files.
* `uploaduser` - Username required to log into the FTP server to upload disk log files.
* `uploadpass` - Password required to log into the FTP server to upload disk log files.
* `uploaddir` - The remote directory on the FTP server to upload log files to.
* `uploadtype` - Types of log files to upload. Separate multiple entries with a space.
* `uploadsched` - Set the schedule for uploading log files to the FTP server (default = disable = upload when rolling). Valid values: `disable`, `enable`.
* `uploadtime` - Time of day at which log files are uploaded if uploadsched is enabled (hh:mm or hh).
* `upload_delete_files` - Delete log files after uploading (default = enable). Valid values: `enable`, `disable`.
* `upload_ssl_conn` - Enable/disable encrypted FTPS communication to upload log files. Valid values: `default`, `high`, `low`, `disable`.
* `full_first_warning_threshold` - Log full first warning threshold as a percent (1 - 98, default = 75).
* `full_second_warning_threshold` - Log full second warning threshold as a percent (2 - 99, default = 90).
* `full_final_warning_threshold` - Log full final warning threshold as a percent (3 - 100, default = 95).
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogDisk Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logdisk_setting.labelname LogDiskSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
