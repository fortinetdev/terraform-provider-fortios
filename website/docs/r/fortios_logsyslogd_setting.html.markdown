---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logsyslogd_setting"
description: |-
  Global settings for remote syslog server.
---

# fortios_logsyslogd_setting
Global settings for remote syslog server.

## Example Usage

```hcl
resource "fortios_logsyslogd_setting" "trname" {
  enc_algorithm         = "disable"
  facility              = "local7"
  format                = "default"
  mode                  = "udp"
  port                  = 514
  ssl_min_proto_version = "default"
  status                = "disable"
  syslog_type           = 1
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable remote syslog logging. Valid values: `enable`, `disable`.
* `server` - Address of remote syslog server.
* `mode` - Remote syslog logging over UDP/Reliable TCP. Valid values: `udp`, `legacy-reliable`, `reliable`.
* `port` - Server listen port.
* `facility` - Remote syslog facility. Valid values: `kernel`, `user`, `mail`, `daemon`, `auth`, `syslog`, `lpr`, `news`, `uucp`, `cron`, `authpriv`, `ftp`, `ntp`, `audit`, `alert`, `clock`, `local0`, `local1`, `local2`, `local3`, `local4`, `local5`, `local6`, `local7`.
* `source_ip` - Source IP address of syslog.
* `format` - Log format.
* `priority` - Set log transmission priority. Valid values: `default`, `low`.
* `max_log_rate` - Syslog maximum log rate in MBps (0 = unlimited).
* `enc_algorithm` - Enable/disable reliable syslogging with TLS encryption. Valid values: `high-medium`, `high`, `low`, `disable`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
* `certificate` - Certificate used to communicate with Syslog server.
* `custom_field_name` - Custom field name for CEF format logging. The structure of `custom_field_name` block is documented below.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `syslog_type` - Hidden setting index of Syslog.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `custom_field_name` block supports:

* `id` - Entry ID.
* `name` - Field name.
* `custom` - Field custom name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogSyslogd Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_logsyslogd_setting.labelname LogSyslogdSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_logsyslogd_setting.labelname LogSyslogdSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
