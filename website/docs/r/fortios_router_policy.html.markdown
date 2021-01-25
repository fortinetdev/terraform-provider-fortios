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
* `src` - Source IP and mask (x.x.x.x/x). The structure of `src` block is documented below.
* `srcaddr` - Source address name. The structure of `srcaddr` block is documented below.
* `src_negate` - Enable/disable negating source address match.
* `dst` - Destination IP and mask (x.x.x.x/x). The structure of `dst` block is documented below.
* `dstaddr` - Destination address name. The structure of `dstaddr` block is documented below.
* `dst_negate` - Enable/disable negating destination address match.
* `action` - Action of the policy route.
* `protocol` - Protocol number (0 - 255).
* `start_port` - Start destination port number (0 - 65535).
* `end_port` - End destination port number (0 - 65535).
* `start_source_port` - Start source port number (0 - 65535).
* `end_source_port` - End source port number (0 - 65535).
* `gateway` - IP address of the gateway.
* `output_device` - Outgoing interface name.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `status` - Enable/disable this policy route.
* `comments` - Optional comments.

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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{seq_num}}.

## Import

Router Policy can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_policy.labelname {{seq_num}}
$ unset "FORTIOS_IMPORT_TABLE"
```
