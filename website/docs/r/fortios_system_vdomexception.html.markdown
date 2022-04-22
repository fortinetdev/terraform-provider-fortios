---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomexception"
description: |-
  Global configuration objects that can be configured independently for all VDOMs or for the defined VDOM scope.
---

# fortios_system_vdomexception
Global configuration objects that can be configured independently for all VDOMs or for the defined VDOM scope.

## Example Usage

```hcl
resource "fortios_system_vdomexception" "trname" {
  fosid  = 1
  object = "log.fortianalyzer.setting"
  oid    = 7150
  scope  = "all"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Index <1-4096>.
* `object` - (Required) Name of the configuration object that can be configured independently for all VDOMs.
* `oid` - Object ID.
* `scope` - Determine whether the configuration object can be configured separately for all VDOMs or if some VDOMs share the same configuration. Valid values: `all`, `inclusive`, `exclusive`.
* `vdom` - Names of the VDOMs. The structure of `vdom` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `vdom` block supports:

* `name` - VDOM name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System VdomException can be imported using any of these accepted formats:
```
$ terraform import fortios_system_vdomexception.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_vdomexception.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
