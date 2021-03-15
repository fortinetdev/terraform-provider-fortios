---
subcategory: "FortiGate NSX-T"
layout: "fortios"
page_title: "FortiOS: fortios_nsxt_setting"
description: |-
  Configure NSX-T setting.
---

# fortios_nsxt_setting
Configure NSX-T setting. Applies to FortiOS Version `>= 7.0.0`.

## Argument Reference

The following arguments are supported:

* `liveness` - Enable/disable liveness detection packet forwarding. Valid values: `enable`, `disable`.
* `service` - Service name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Nsxt Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_nsxt_setting.labelname NsxtSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
