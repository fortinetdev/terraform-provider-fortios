---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceextension"
description: |-
  Get information on an fortios firewall internetserviceextension.
---

# Data Source: fortios_firewall_internetserviceextension
Use this data source to get information on an fortios firewall internetserviceextension

## Argument Reference

* `fosid` - (Required) Specify the fosid of the desired firewall internetserviceextension.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fosid` - Internet Service ID in the Internet Service database.
* `comment` - Comment.
* `entry` - Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
* `disable_entry` - Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.

The `entry` block contains:

* `id` - Entry ID(1-255).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the custom entry. The structure of `port_range` block is documented below.
* `dst` - Destination address or address group name. The structure of `dst` block is documented below.

The `port_range` block contains:

* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).

The `dst` block contains:

* `name` - Select the destination address or address group object from available options.

The `disable_entry` block contains:

* `id` - Disable entry ID.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the disable entry. The structure of `port_range` block is documented below.
* `port` - Integer value for the TCP/IP port (0 - 65535).
* `ip_range` - IP ranges in the disable entry. The structure of `ip_range` block is documented below.

The `port_range` block contains:

* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (1 to 65535).
* `end_port` - Ending TCP/UDP/SCTP destination port (1 to 65535).

The `ip_range` block contains:

* `id` - Disable entry range ID.
* `start_ip` - Start IP address.
* `end_ip` - End IP address.

