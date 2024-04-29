---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_prefixlist"
description: |-
  Configure IPv4 prefix lists.
---

# fortios_router_prefixlist
Configure IPv4 prefix lists.

## Example Usage

```hcl
resource "fortios_router_prefixlist" "trname" {
  name = "pr1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `comments` - Comment.
* `rule` - IPv4 prefix list rule. The structure of `rule` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `rule` block supports:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix. Valid values: `permit`, `deny`.
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.
* `ge` - Minimum prefix length to be matched (0 - 32).
* `le` - Maximum prefix length to be matched (0 - 32).
* `flags` - Flags.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router PrefixList can be imported using any of these accepted formats:
```
$ terraform import fortios_router_prefixlist.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_prefixlist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
