---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sessionttl"
description: |-
  Configure global session TTL timers for this FortiGate.
---

# fortios_system_sessionttl
Configure global session TTL timers for this FortiGate.

## Example Usage

```hcl
resource "fortios_system_sessionttl" "trname" {
  default = "3600"
}
```

## Argument Reference

The following arguments are supported:

* `default` - Default timeout.
* `port` - Session TTL port. The structure of `port` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `port` block supports:

* `id` - Table entry ID.
* `protocol` - Protocol (0 - 255).
* `start_port` - Start port number.
* `end_port` - End port number.
* `timeout` - Session timeout (TTL).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System SessionTtl can be imported using any of these accepted formats:
```
$ terraform import fortios_system_sessionttl.labelname SystemSessionTtl

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_sessionttl.labelname SystemSessionTtl
$ unset "FORTIOS_IMPORT_TABLE"
```
