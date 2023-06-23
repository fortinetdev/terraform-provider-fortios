---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_policy"
description: |-
  Configure IPv4 routing policies.
---

# fortios_router_policy
Configure IPv4 routing policies.

## Example Usage

```hcl
resource "fortios_router_policy" "trname" {
  action            = "permit"
  dst_negate        = "disable"
  end_port          = 25
  end_source_port   = 65535
  gateway           = "0.0.0.0"
  output_device     = "port2"
  protocol          = 6
  seq_num           = 1
  src_negate        = "disable"
  start_port        = 25
  start_source_port = 0
  status            = "enable"
  tos               = "0x00"
  tos_mask          = "0x00"

  input_device {
    name = "port1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `seq_num` - Sequence number.
* `input_device` - Incoming interface name. The structure of `input_device` block is documented below.
* `input_device_negate` - Enable/disable negation of input device match. Valid values: `enable`, `disable`.
* `src` - Source IP and mask (x.x.x.x/x). The structure of `src` block is documented below.
* `srcaddr` - Source address name. The structure of `srcaddr` block is documented below.
* `src_negate` - Enable/disable negating source address match. Valid values: `enable`, `disable`.
* `dst` - Destination IP and mask (x.x.x.x/x). The structure of `dst` block is documented below.
* `dstaddr` - Destination address name. The structure of `dstaddr` block is documented below.
* `dst_negate` - Enable/disable negating destination address match. Valid values: `enable`, `disable`.
* `action` - Action of the policy route. Valid values: `deny`, `permit`.
* `protocol` - Protocol number (0 - 255).
* `start_port` - Start destination port number (0 - 65535).
* `end_port` - End destination port number (0 - 65535).
* `start_source_port` - Start source port number (0 - 65535).
* `end_source_port` - End source port number (0 - 65535).
* `gateway` - IP address of the gateway.
* `output_device` - Outgoing interface name.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `status` - Enable/disable this policy route. Valid values: `enable`, `disable`.
* `comments` - Optional comments.
* `internet_service_id` - Destination Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_custom` - Custom Destination Internet Service name. The structure of `internet_service_custom` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `input_device` block supports:

* `name` - Interface name.

The `src` block supports:

* `subnet` - IP and mask.

The `srcaddr` block supports:

* `name` - Address/group name.

The `dst` block supports:

* `subnet` - IP and mask.

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

Router Policy can be imported using any of these accepted formats:
```
$ terraform import fortios_router_policy.labelname {{seq_num}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_policy.labelname {{seq_num}}
$ unset "FORTIOS_IMPORT_TABLE"
```
