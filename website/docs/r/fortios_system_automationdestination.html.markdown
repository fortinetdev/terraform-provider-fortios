---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationdestination"
description: |-
  Automation destinations.
---

# fortios_system_automationdestination
Automation destinations.

## Example Usage

```hcl
resource "fortios_system_automationdestination" "trname" {
  ha_group_id = 0
  name        = "1"
  type        = "fortigate"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `type` - Destination type. Valid values: `fortigate`, `ha-cluster`.
* `destination` - Destinations. The structure of `destination` block is documented below.
* `ha_group_id` - Cluster group ID set for this destination (default = 0).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `destination` block supports:

* `name` - Destination.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationDestination can be imported using any of these accepted formats:
```
$ terraform import fortios_system_automationdestination.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_automationdestination.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
