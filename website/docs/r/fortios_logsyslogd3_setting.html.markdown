---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logsyslogd3_setting"
description: |-
  Global settings for remote syslog server.
---

# fortios_logsyslogd3_setting
Global settings for remote syslog server.

## Example Usage

```hcl
resource "fortios_logsyslogd3_setting" "trname" {
  enc_algorithm         = "disable"
  facility              = "local7"
  format                = "default"
  mode                  = "udp"
  port                  = 514
  ssl_min_proto_version = "default"
  status                = "disable"
  syslog_type           = 3
}
```

## Argument Reference


The following arguments are supported:

* `status` - Enable/disable remote syslog logging.
* `server` - Address of remote syslog server.
* `mode` - Remote syslog logging over UDP/Reliable TCP.
* `port` - Server listen port.
* `facility` - Remote syslog facility.
* `source_ip` - Source IP address of syslog.
* `format` - Log format.
* `enc_algorithm` - Enable/disable reliable syslogging with TLS encryption.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `certificate` - Certificate used to communicate with Syslog server.
* `custom_field_name` - Custom field name for CEF format logging. The structure of `custom_field_name` block is documented below.
* `syslog_type` - Hidden setting index of Syslog.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `custom_field_name` block supports:

* `id` - Entry ID.
* `name` - Field name.
* `custom` - Field custom name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogSyslogd3 Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logsyslogd3_setting.labelname LogSyslogd3Setting
$ unset "FORTIOS_IMPORT_TABLE"
```
