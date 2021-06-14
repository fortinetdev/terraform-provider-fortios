---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicedefinition"
description: |-
  Get information on an fortios firewall internetservicedefinition.
---

# Data Source: fortios_firewall_internetservicedefinition
Use this data source to get information on an fortios firewall internetservicedefinition

## Argument Reference

* `fosid` - (Required) Specify the fosid of the desired firewall internetservicedefinition.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fosid` - Internet Service application list ID.
* `entry` - Protocol and port information in an Internet Service entry. The structure of `entry` block is documented below.

The `entry` block contains:

* `seq_num` - Entry sequence number.
* `category_id` - Internet Service category ID.
* `name` - Internet Service name.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the definition entry. The structure of `port_range` block is documented below.
* `port` - Integer value for ending TCP/UDP/SCTP destination port in range (0 to 65535). 0 means undefined.

The `port_range` block contains:

* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (1 to 65535).
* `end_port` - Ending TCP/UDP/SCTP destination port (1 to 65535).

