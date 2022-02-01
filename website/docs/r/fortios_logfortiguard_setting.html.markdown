---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortiguard_setting"
description: |-
  Configure logging to FortiCloud.
---

# fortios_logfortiguard_setting
Configure logging to FortiCloud.

## Example Usage

```hcl
resource "fortios_logfortiguard_setting" "trname" {
  enc_algorithm         = "high"
  source_ip             = "0.0.0.0"
  ssl_min_proto_version = "default"
  status                = "disable"
  upload_interval       = "daily"
  upload_option         = "5-minute"
  upload_time           = "00:00"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
* `upload_option` - Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
* `upload_interval` - Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
* `upload_day` - Day of week to roll logs.
* `upload_time` - Time of day to roll logs (hh:mm).
* `priority` - Set log transmission priority. Valid values: `default`, `low`.
* `max_log_rate` - FortiCloud maximum log rate in MBps (0 = unlimited).
* `access_config` - Enable/disable FortiCloud access to configuration and data. Valid values: `enable`, `disable`.
* `enc_algorithm` - Enable and set the SSL security level for for sending encrypted logs to FortiCloud. Valid values: `high-medium`, `high`, `low`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting). Valid values: `default`, `SSLv3`, `TLSv1`, `TLSv1-1`, `TLSv1-2`.
* `conn_timeout` - FortiGate Cloud connection timeout in seconds.
* `source_ip` - Source IP address used to connect FortiCloud.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortiguard Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logfortiguard_setting.labelname LogFortiguardSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
