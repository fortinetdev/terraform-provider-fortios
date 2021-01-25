---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_setting"
description: |-
  Configure general log settings.
---

# fortios_log_setting
Configure general log settings.

## Example Usage

```hcl
resource "fortios_log_setting" "trname" {
  brief_traffic_format    = "disable"
  daemon_log              = "disable"
  expolicy_implicit_log   = "disable"
  faz_override            = "disable"
  fwpolicy6_implicit_log  = "disable"
  fwpolicy_implicit_log   = "disable"
  local_in_allow          = "disable"
  local_in_deny_broadcast = "disable"
  local_in_deny_unicast   = "disable"
  local_out               = "disable"
  log_invalid_packet      = "disable"
  log_policy_comment      = "disable"
  log_policy_name         = "disable"
  log_user_in_upper       = "disable"
  neighbor_event          = "disable"
  resolve_ip              = "disable"
  resolve_port            = "enable"
  syslog_override         = "disable"
  user_anonymize          = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `resolve_ip` - Enable/disable adding resolved domain names to traffic logs if possible.
* `resolve_port` - Enable/disable adding resolved service names to traffic logs.
* `log_user_in_upper` - Enable/disable logs with user-in-upper.
* `fwpolicy_implicit_log` - Enable/disable implicit firewall policy logging.
* `fwpolicy6_implicit_log` - Enable/disable implicit firewall policy6 logging.
* `log_invalid_packet` - Enable/disable invalid packet traffic logging.
* `local_in_allow` - Enable/disable local-in-allow logging.
* `local_in_deny_unicast` - Enable/disable local-in-deny-unicast logging.
* `local_in_deny_broadcast` - Enable/disable local-in-deny-broadcast logging.
* `local_out` - Enable/disable local-out logging.
* `daemon_log` - Enable/disable daemon logging.
* `neighbor_event` - Enable/disable neighbor event logging.
* `brief_traffic_format` - Enable/disable brief format traffic logging.
* `user_anonymize` - Enable/disable anonymizing user names in log messages.
* `expolicy_implicit_log` - Enable/disable explicit proxy firewall implicit policy logging.
* `log_policy_comment` - Enable/disable inserting policy comments into traffic logs.
* `log_policy_name` - Enable/disable inserting policy name into traffic logs.
* `faz_override` - Enable/disable override FortiAnalyzer settings.
* `syslog_override` - Enable/disable override Syslog settings.
* `custom_log_fields` - Custom fields to append to all log messages. The structure of `custom_log_fields` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_log_fields` block supports:

* `field_id` - Custom log field.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_log_setting.labelname LogSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
