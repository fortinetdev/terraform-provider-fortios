---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_lognulldevice_setting"
description: |-
  Settings for null device logging.
---

# fortios_lognulldevice_setting
Settings for null device logging.

## Example Usage

```hcl
resource "fortios_lognulldevice_setting" "trname" {
  status = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - (Required) Enable/disable statistics collection for when no external logging destination, such as FortiAnalyzer, is present (data is not saved). Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogNullDevice Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_lognulldevice_setting.labelname LogNullDeviceSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
