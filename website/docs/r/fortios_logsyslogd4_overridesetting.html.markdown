---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logsyslogd4_overridesetting"
description: |-
  Override settings for remote syslog server.
---

# fortios_logsyslogd4_overridesetting
Override settings for remote syslog server.

## Argument Reference

The following arguments are supported:

* `override` - Enable/disable override syslog settings.
* `status` - Enable/disable remote syslog logging.
* `server` - Address of remote syslog server.
* `mode` - Remote syslog logging over UDP/Reliable TCP.
* `port` - Server listen port.
* `facility` - Remote syslog facility.
* `source_ip` - Source IP address of syslog.
* `format` - Log format.
* `priority` - Set log transmission priority.
* `max_log_rate` - Syslog maximum log rate in MBps (0 = unlimited).
* `enc_algorithm` - Enable/disable reliable syslogging with TLS encryption.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `certificate` - Certificate used to communicate with Syslog server.
* `custom_field_name` - Custom field name for CEF format logging. The structure of `custom_field_name` block is documented below.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.
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

LogSyslogd4 OverrideSetting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logsyslogd4_overridesetting.labelname LogSyslogd4OverrideSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
