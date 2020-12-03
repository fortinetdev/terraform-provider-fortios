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

* `fosid` - (Required) Session helper ID.
* `name` - (Required) Helper name.
* `protocol` - (Required) Protocol number.
* `port` - (Required) Protocol port.


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
