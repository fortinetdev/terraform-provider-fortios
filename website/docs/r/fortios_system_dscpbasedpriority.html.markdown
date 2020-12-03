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

* `fosid` - (Required) Item ID.
* `ds` - DSCP(DiffServ) DS value (0 - 63).
* `priority` - DSCP based priority level.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System DscpBasedPriority can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_dscpbasedpriority.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
