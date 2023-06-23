---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ndproxy"
description: |-
  Configure IPv6 neighbor discovery proxy (RFC4389).
---

# fortios_system_ndproxy
Configure IPv6 neighbor discovery proxy (RFC4389).

## Example Usage

```hcl
resource "fortios_system_ndproxy" "trname" {
  status = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable neighbor discovery proxy. Valid values: `enable`, `disable`.
* `member` - Interfaces using the neighbor discovery proxy. The structure of `member` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `member` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System NdProxy can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ndproxy.labelname SystemNdProxy

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ndproxy.labelname SystemNdProxy
$ unset "FORTIOS_IMPORT_TABLE"
```
