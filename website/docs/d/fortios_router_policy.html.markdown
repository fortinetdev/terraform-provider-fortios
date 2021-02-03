---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_policy"
description: |-
  Get information on an fortios router policy.
---

# Data Source: fortios_router_policy
Use this data source to get information on an fortios router policy

## Argument Reference

* `seq_num` - (Required) Specify the seq_num of the desired router policy.

## Attribute Reference

The following attributes are exported:

* `seq_num` - Sequence number.
* `input_device` - Incoming interface name. The structure of `input_device` block is documented below.
* `input_device_negate` - Enable/disable negation of input device match.
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
* `internet_service_id` - Destination Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_custom` - Custom Destination Internet Service name. The structure of `internet_service_custom` block is documented below.

The `input_device` block contains:

* `name` - Interface name.

The `src` block contains:

* `subnet` - IP and mask.

The `srcaddr` block contains:

* `name` - Address/group name.

The `dst` block contains:

* `subnet` - IP and mask.

The `dstaddr` block contains:

* `name` - Address/group name.

The `internet_service_id` block contains:

* `id` - Destination Internet Service ID.

The `internet_service_custom` block contains:

* `name` - Custom Destination Internet Service name.

