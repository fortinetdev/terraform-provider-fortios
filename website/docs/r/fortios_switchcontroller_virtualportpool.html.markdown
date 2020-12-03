---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_virtualportpool"
description: |-
  Configure virtual pool.
---

# fortios_switchcontroller_virtualportpool
Configure virtual pool.

## Example Usage

```hcl
resource "fortios_switchcontroller_virtualportpool" "trname" {
  description = "virtualport"
  name        = "virtualportpool1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Virtual switch pool name.
* `description` - Virtual switch pool description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController VirtualPortPool can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_virtualportpool.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
