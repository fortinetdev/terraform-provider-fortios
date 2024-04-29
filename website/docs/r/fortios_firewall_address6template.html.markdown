---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address6template"
description: |-
  Configure IPv6 address templates.
---

# fortios_firewall_address6template
Configure IPv6 address templates.

## Example Usage

```hcl
resource "fortios_firewall_address6template" "trname" {
  ip6                  = "2001:db8:0:b::/64"
  name                 = "1"
  subnet_segment_count = 2

  subnet_segment {
    bits      = 4
    exclusive = "disable"
    id        = 1
    name      = "country"
  }
  subnet_segment {
    bits      = 4
    exclusive = "disable"
    id        = 2
    name      = "state"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) IPv6 address template name.
* `ip6` - (Required) IPv6 address prefix.
* `subnet_segment_count` - (Required) Number of IPv6 subnet segments.
* `subnet_segment` - IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `subnet_segment` block supports:

* `id` - Subnet segment ID.
* `name` - Subnet segment name.
* `bits` - Number of bits.
* `exclusive` - Enable/disable exclusive value. Valid values: `enable`, `disable`.
* `values` - Subnet segment values. The structure of `values` block is documented below.

The `values` block supports:

* `name` - Subnet segment value name.
* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address6Template can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_address6template.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_address6template.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
