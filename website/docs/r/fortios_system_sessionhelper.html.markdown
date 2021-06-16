---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sessionhelper"
description: |-
  Configure session helper.
---

# fortios_system_sessionhelper
Configure session helper.

## Example Usage

```hcl
resource "fortios_system_sessionhelper" "trname" {
  fosid    = 33
  name     = "ftp"
  port     = 3234
  protocol = 17
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Session helper ID.
* `name` - (Required) Helper name.
* `protocol` - (Required) Protocol number.
* `port` - (Required) Protocol port.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System SessionHelper can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_sessionhelper.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
