---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dscpbasedpriority"
description: |-
  Configure DSCP based priority table.
---

# fortios_system_dscpbasedpriority
Configure DSCP based priority table.

## Example Usage

```hcl
resource "fortios_system_dscpbasedpriority" "trname" {
  ds       = 1
  fosid    = 1
  priority = "low"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Item ID.
* `ds` - DSCP(DiffServ) DS value (0 - 63).
* `priority` - DSCP based priority level. Valid values: `low`, `medium`, `high`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System DscpBasedPriority can be imported using any of these accepted formats:
```
$ terraform import fortios_system_dscpbasedpriority.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_dscpbasedpriority.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
