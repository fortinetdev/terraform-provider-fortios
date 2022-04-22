---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logwebtrends_setting"
description: |-
  Settings for WebTrends.
---

# fortios_logwebtrends_setting
Settings for WebTrends.

## Example Usage

```hcl
resource "fortios_logwebtrends_setting" "trname" {
  status = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable logging to WebTrends. Valid values: `enable`, `disable`.
* `server` - Address of the remote WebTrends server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogWebtrends Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_logwebtrends_setting.labelname LogWebtrendsSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_logwebtrends_setting.labelname LogWebtrendsSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
