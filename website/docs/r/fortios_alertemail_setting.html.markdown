---
subcategory: "FortiGate AlertEmail"
layout: "fortios"
page_title: "FortiOS: fortios_alertemail_setting"
description: |-
  Configure alert email settings.
---

# fortios_alertemail_setting
Configure alert email settings.

## Example Usage

```hcl
resource "fortios_alertemail_setting" "trname" {
  admin_login_logs           = "disable"
  alert_interval             = 2
  amc_interface_bypass_mode  = "disable"
  antivirus_logs             = "disable"
  configuration_changes_logs = "disable"
  critical_interval          = 3
  debug_interval             = 60
  email_interval             = 5
  emergency_interval         = 1
  error_interval             = 5
  fds_license_expiring_days  = 15
  information_interval       = 30
}
```

## Argument Reference

The following arguments are supported:

* `username` - Name that appears in the From: field of alert emails (max. 36 characters).
* `mailto1` - Email address to send alert email to (usually a system administrator) (max. 64 characters).
* `mailto2` - Optional second email address to send alert email to (max. 64 characters).
* `mailto3` - Optional third email address to send alert email to (max. 64 characters).
* `filter_mode` - How to filter log messages that are sent to alert emails.
* `email_interval` - Interval between sending alert emails (1 - 99999 min, default = 5).
* `IPS_logs` - Enable/disable IPS logs in alert email.
* `firewall_authentication_failure_logs` - Enable/disable firewall authentication failure logs in alert email.
* `HA_logs` - Enable/disable HA logs in alert email.
* `IPsec_errors_logs` - Enable/disable IPsec error logs in alert email.
* `FDS_update_logs` - Enable/disable FortiGuard update logs in alert email.
* `PPP_errors_logs` - Enable/disable PPP error logs in alert email.
* `sslvpn_authentication_errors_logs` - Enable/disable SSL-VPN authentication error logs in alert email.
* `antivirus_logs` - Enable/disable antivirus logs in alert email.
* `webfilter_logs` - Enable/disable web filter logs in alert email.
* `configuration_changes_logs` - Enable/disable configuration change logs in alert email.
* `violation_traffic_logs` - Enable/disable violation traffic logs in alert email.
* `admin_login_logs` - Enable/disable administrator login/logout logs in alert email.
* `FDS_license_expiring_warning` - Enable/disable FortiGuard license expiration warnings in alert email.
* `log_disk_usage_warning` - Enable/disable disk usage warnings in alert email.
* `fortiguard_log_quota_warning` - Enable/disable FortiCloud log quota warnings in alert email.
* `amc_interface_bypass_mode` - Enable/disable Fortinet Advanced Mezzanine Card (AMC) interface bypass mode logs in alert email.
* `FIPS_CC_errors` - Enable/disable FIPS and Common Criteria error logs in alert email.
* `FSSO_disconnect_logs` - Enable/disable logging of FSSO collector agent disconnect.
* `ssh_logs` - Enable/disable SSH logs in alert email.
* `FDS_license_expiring_days` - Number of days to send alert email prior to FortiGuard license expiration (1 - 100 days, default = 100).
* `local_disk_usage` - Disk usage percentage at which to send alert email (1 - 99 percent, default = 75).
* `emergency_interval` - Emergency alert interval in minutes.
* `alert_interval` - Alert alert interval in minutes.
* `critical_interval` - Critical alert interval in minutes.
* `error_interval` - Error alert interval in minutes.
* `warning_interval` - Warning alert interval in minutes.
* `notification_interval` - Notification alert interval in minutes.
* `information_interval` - Information alert interval in minutes.
* `debug_interval` - Debug alert interval in minutes.
* `severity` - Lowest severity level to log.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Alertemail Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_alertemail_setting.labelname AlertemailSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
