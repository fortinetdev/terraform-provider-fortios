---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceextension"
description: |-
  Configure Internet Services Extension.
---

# fortios_firewall_internetserviceextension
Configure Internet Services Extension.

## Example Usage

```hcl
resource "fortios_firewall_internetserviceextension" "trname" {
  comment = "EIWE"
  fosid   = 65536
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Internet Service ID in the Internet Service database.
* `comment` - Comment.
* `entry` - Entries added to the Internet Service extension database. The structure of `entry` block is documented below.
* `disable_entry` - Disable entries in the Internet Service database. The structure of `disable_entry` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entry` block supports:

* `id` - Entry ID(1-255).
* `addr_mode` - Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the custom entry. The structure of `port_range` block is documented below.
* `dst` - Destination address or address group name. The structure of `dst` block is documented below.
* `dst6` - Destination address6 or address6 group name. The structure of `dst6` block is documented below.

The `port_range` block supports:

* `id` - Custom entry port range ID.
* `start_port` - Integer value for starting TCP/UDP/SCTP destination port in range (1 to 65535).
* `end_port` - Integer value for ending TCP/UDP/SCTP destination port in range (1 to 65535).

The `dst` block supports:

* `name` - Select the destination address or address group object from available options.

The `dst6` block supports:

* `name` - Select the destination address6 or address group object from available options.

The `disable_entry` block supports:

* `id` - Disable entry ID.
* `addr_mode` - Address mode (IPv4 or IPv6) Valid values: `ipv4`, `ipv6`.
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `port_range` - Port ranges in the disable entry. The structure of `port_range` block is documented below.
* `port` - Integer value for the TCP/IP port (0 - 65535).
* `ip_range` - IP ranges in the disable entry. The structure of `ip_range` block is documented below.
* `ip6_range` - IPv6 ranges in the disable entry. The structure of `ip6_range` block is documented below.

The `port_range` block supports:

* `id` - Custom entry port range ID.
* `start_port` - Starting TCP/UDP/SCTP destination port (1 to 65535).
* `end_port` - Ending TCP/UDP/SCTP destination port (1 to 65535).

The `ip_range` block supports:

* `id` - Disable entry range ID.
* `start_ip` - Start IP address.
* `end_ip` - End IP address.

The `ip6_range` block supports:

* `id` - Disable entry range ID.
* `start_ip6` - Start IPv6 address.
* `end_ip6` - End IPv6 address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceExtension can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_internetserviceextension.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_internetserviceextension.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
