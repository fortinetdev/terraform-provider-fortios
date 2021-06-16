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

* `name` - Virtual switch pool name.
* `description` - Virtual switch pool description.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


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
