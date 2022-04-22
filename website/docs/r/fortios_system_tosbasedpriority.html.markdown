---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_tosbasedpriority"
description: |-
  Configure Type of Service (ToS) based priority table to set network traffic priorities.
---

# fortios_system_tosbasedpriority
Configure Type of Service (ToS) based priority table to set network traffic priorities.

## Example Usage

```hcl
resource "fortios_system_tosbasedpriority" "trname" {
  fosid    = 1
  priority = "low"
  tos      = 11
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Item ID.
* `tos` - Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
* `priority` - ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium). Valid values: `low`, `medium`, `high`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System TosBasedPriority can be imported using any of these accepted formats:
```
$ terraform import fortios_system_tosbasedpriority.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_tosbasedpriority.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
