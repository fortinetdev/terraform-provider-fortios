---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_log_guidisplay"
description: |-
  Configure how log messages are displayed on the GUI.
---

# fortios_log_guidisplay
Configure how log messages are displayed on the GUI.

## Example Usage

```hcl
resource "fortios_log_guidisplay" "trname" {
  fortiview_unscanned_apps = "disable"
  resolve_apps             = "enable"
  resolve_hosts            = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `resolve_hosts` - Enable/disable resolving IP addresses to hostname in log messages on the GUI using reverse DNS lookup Valid values: `enable`, `disable`.
* `resolve_apps` - Resolve unknown applications on the GUI using Fortinet's remote application database. Valid values: `enable`, `disable`.
* `fortiview_unscanned_apps` - Enable/disable showing unscanned traffic in FortiView application charts. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Log GuiDisplay can be imported using any of these accepted formats:
```
$ terraform import fortios_log_guidisplay.labelname LogGuiDisplay

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_log_guidisplay.labelname LogGuiDisplay
$ unset "FORTIOS_IMPORT_TABLE"
```
