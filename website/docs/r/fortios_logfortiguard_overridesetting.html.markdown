---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortiguard_overridesetting"
description: |-
  Override global FortiCloud logging settings for this VDOM.
---

# fortios_logfortiguard_overridesetting
Override global FortiCloud logging settings for this VDOM.

## Example Usage

```hcl
resource "fortios_logfortiguard_overridesetting" "trname" {
  override        = "disable"
  status          = "disable"
  upload_interval = "daily"
  upload_option   = "5-minute"
  upload_time     = "00:00"
}
```

## Argument Reference

The following arguments are supported:

* `override` - Overriding FortiCloud settings for this VDOM or use global settings. Valid values: `enable`, `disable`.
* `status` - Enable/disable logging to FortiCloud. Valid values: `enable`, `disable`.
* `upload_option` - Configure how log messages are sent to FortiCloud. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
* `upload_interval` - Frequency of uploading log files to FortiCloud. Valid values: `daily`, `weekly`, `monthly`.
* `upload_day` - Day of week to roll logs.
* `upload_time` - Time of day to roll logs (hh:mm).
* `priority` - Set log transmission priority. Valid values: `default`, `low`.
* `max_log_rate` - FortiCloud maximum log rate in MBps (0 = unlimited).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortiguard OverrideSetting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logfortiguard_overridesetting.labelname LogFortiguardOverrideSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
