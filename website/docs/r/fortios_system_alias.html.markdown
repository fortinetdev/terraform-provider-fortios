---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_alias"
description: |-
  Configure alias command.
---

# fortios_system_alias
Configure alias command.

## Example Usage

```hcl
resource "fortios_system_alias" "trname" {
  name = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Alias command name.
* `command` - Command list to execute.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Alias can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_alias.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
