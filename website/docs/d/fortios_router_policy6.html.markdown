---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_policy6"
description: |-
  Get information on an fortios router policy6.
---

# Data Source: fortios_router_policy6
Use this data source to get information on an fortios router policy6

## Argument Reference

* `seq_num` - (Required) Specify the seq_num of the desired router policy6.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `seq_num` - Sequence number.
* `input_device` - Incoming interface name.
* `input_device_negate` - Enable/disable negation of input device match.
* `src` - Source IPv6 prefix.
* `srcaddr` - Source address name. The structure of `srcaddr` block is documented below.
* `src_negate` - Enable/disable negating source address match.
* `dst` - Destination IPv6 prefix.
* `dstaddr` - Destination address name. The structure of `dstaddr` block is documented below.
* `dst_negate` - Enable/disable negating destination address match.
* `action` - Action of the policy route.
* `protocol` - Protocol number (0 - 255).
* `start_port` - Start destination port number (1 - 65535).
* `end_port` - End destination port number (1 - 65535).
* `start_source_port` - Start source port number (1 - 65535).
* `end_source_port` - End source port number (1 - 65535).
* `gateway` - IPv6 address of the gateway.
* `output_device` - Outgoing interface name.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `status` - Enable/disable this policy route.
* `comments` - Optional comments.
* `internet_service_id` - Destination Internet Service ID. The structure of `internet_service_id` block is documented below.
* `internet_service_custom` - Custom Destination Internet Service name. The structure of `internet_service_custom` block is documented below.
* `internet_service_fortiguard` - FortiGuard Destination Internet Service name. The structure of `internet_service_fortiguard` block is documented below.
* `users` - List of users. The structure of `users` block is documented below.
* `groups` - List of user groups. The structure of `groups` block is documented below.

The `srcaddr` block contains:

* `name` - Address/group name.

The `dstaddr` block contains:

* `name` - Address/group name.

The `internet_service_id` block contains:

* `id` - Destination Internet Service ID.

The `internet_service_custom` block contains:

* `name` - Custom Destination Internet Service name.

The `internet_service_fortiguard` block contains:

* `name` - FortiGuard Destination Internet Service name.

The `users` block contains:

* `name` - User name.

The `groups` block contains:

* `name` - Group name.

