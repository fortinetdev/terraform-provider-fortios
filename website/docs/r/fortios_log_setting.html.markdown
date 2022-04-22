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

* `resolve_ip` - Enable/disable adding resolved domain names to traffic logs if possible. Valid values: `enable`, `disable`.
* `resolve_port` - Enable/disable adding resolved service names to traffic logs. Valid values: `enable`, `disable`.
* `log_user_in_upper` - Enable/disable logs with user-in-upper. Valid values: `enable`, `disable`.
* `fwpolicy_implicit_log` - Enable/disable implicit firewall policy logging. Valid values: `enable`, `disable`.
* `fwpolicy6_implicit_log` - Enable/disable implicit firewall policy6 logging. Valid values: `enable`, `disable`.
* `log_invalid_packet` - Enable/disable invalid packet traffic logging. Valid values: `enable`, `disable`.
* `local_in_allow` - Enable/disable local-in-allow logging. Valid values: `enable`, `disable`.
* `local_in_deny_unicast` - Enable/disable local-in-deny-unicast logging. Valid values: `enable`, `disable`.
* `local_in_deny_broadcast` - Enable/disable local-in-deny-broadcast logging. Valid values: `enable`, `disable`.
* `local_out` - Enable/disable local-out logging. Valid values: `enable`, `disable`.
* `daemon_log` - Enable/disable daemon logging. Valid values: `enable`, `disable`.
* `neighbor_event` - Enable/disable neighbor event logging. Valid values: `enable`, `disable`.
* `brief_traffic_format` - Enable/disable brief format traffic logging. Valid values: `enable`, `disable`.
* `user_anonymize` - Enable/disable anonymizing user names in log messages. Valid values: `enable`, `disable`.
* `expolicy_implicit_log` - Enable/disable explicit proxy firewall implicit policy logging. Valid values: `enable`, `disable`.
* `log_policy_comment` - Enable/disable inserting policy comments into traffic logs. Valid values: `enable`, `disable`.
* `log_policy_name` - Enable/disable inserting policy name into traffic logs. Valid values: `enable`, `disable`.
* `faz_override` - Enable/disable override FortiAnalyzer settings. Valid values: `enable`, `disable`.
* `syslog_override` - Enable/disable override Syslog settings. Valid values: `enable`, `disable`.
* `rest_api_set` - Enable/disable REST API POST/PUT/DELETE request logging. Valid values: `enable`, `disable`.
* `rest_api_get` - Enable/disable REST API GET request logging. Valid values: `enable`, `disable`.
* `custom_log_fields` - Custom fields to append to all log messages. The structure of `custom_log_fields` block is documented below.
* `anonymization_hash` - User name anonymization hash salt.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `custom_log_fields` block supports:

* `field_id` - Custom log field.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_log_setting.labelname LogSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_log_setting.labelname LogSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
