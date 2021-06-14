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
* `src` - Source IPv6 prefix.
* `dst` - Destination IPv6 prefix.
* `protocol` - Protocol number (0 - 255).
* `start_port` - Start destination port number (1 - 65535).
* `end_port` - End destination port number (1 - 65535).
* `gateway` - IPv6 address of the gateway.
* `output_device` - Outgoing interface name.
* `tos` - Type of service bit pattern.
* `tos_mask` - Type of service evaluated bits.
* `status` - Enable/disable this policy route.
* `comments` - Optional comments.

