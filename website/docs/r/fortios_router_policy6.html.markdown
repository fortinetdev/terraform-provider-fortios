---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_policy6"
description: |-
  Configure IPv6 routing policies.
---

# fortios_router_policy6
Configure IPv6 routing policies.

## Example Usage

```hcl
resource "fortios_router_policy6" "trname" {
  dst           = "::/0"
  end_port      = 65535
  gateway       = "::"
  input_device  = "port1"
  output_device = "port3"
  protocol      = 33
  seq_num       = 1
  src           = "2001:db8:85a3::8a2e:370:7334/128"
  start_port    = 1
  status        = "enable"
  tos           = "0x00"
  tos_mask      = "0x00"
}
```

## Argument Reference

The following arguments are supported:

* `seq_num` - Sequence number.
* `input_device` - (Required) Incoming interface name.
* `input_device_negate` - Enable/disable negation of input device match. Valid values: `enable`, `disable`.
* `src` - Source IPv6 prefix.
* `srcaddr` - Source address name. The structure of `srcaddr` block is documented below.
* `src_negate` - Enable/disable negating source address match. Valid values: `enable`, `disable`.
* `dst` - Destination IPv6 prefix.
* `dstaddr` - Destination address name. The structure of `dstaddr` block is documented below.
* `dst_negate` - Enable/disable negating destination address match. Valid values: `enable`, `disable`.
* `action` - Action of the policy route. Valid values: `deny`, `permit`.
* `protocol` - Protocol number (0 - 255).
* `start_port` - Start destination port number (1 - 65535).
* `end_port` - End destination port number (1 - 65535).
* `start_source_port` - Start source port number (1 - 65535).
* `end_source_port` - End source port number (1 - 65535).
* `gateway` - IPv6 address of the gateway.
* `output_device` - Outgoing interface name.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `status` - Enable/disable this policy route. Valid values: `enable`, `disable`.
* `comments` - Optional comments.
* `internet_service_id` - Destination Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_custom` - Custom Destination Internet Service name. The structure of `internet_service_custom` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcaddr` block supports:

* `name` - Address/group name.

The `dstaddr` block supports:

* `name` - Address/group name.

The `internet_service_id` block supports:

* `id` - Destination Internet Service ID.

The `internet_service_custom` block supports:

* `name` - Custom Destination Internet Service name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Policy6 can be imported using any of these accepted formats:
```
$ terraform import fortios_router_policy6.labelname {{seq_num}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_policy6.labelname {{seq_num}}
$ unset "FORTIOS_IMPORT_TABLE"
```
