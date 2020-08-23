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

* `status` - Enable/disable logging to FortiCloud.
* `upload_option` - Configure how log messages are sent to FortiCloud.
* `upload_interval` - Frequency of uploading log files to FortiCloud.
* `upload_day` - Day of week to roll logs.
* `upload_time` - Time of day to roll logs (hh:mm).
* `enc_algorithm` - Enable and set the SSL security level for for sending encrypted logs to FortiCloud.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `source_ip` - Source IP address used to connect FortiCloud.


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
