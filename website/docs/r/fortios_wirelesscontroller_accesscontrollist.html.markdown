---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_accesscontrollist"
description: |-
  Configure WiFi bridge access control list.
---

# fortios_wirelesscontroller_accesscontrollist
Configure WiFi bridge access control list. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - AP access control list name.
* `comment` - Description.
* `layer3_ipv4_rules` - AP ACL layer3 ipv4 rule list. The structure of `layer3_ipv4_rules` block is documented below.
* `layer3_ipv6_rules` - AP ACL layer3 ipv6 rule list. The structure of `layer3_ipv6_rules` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `layer3_ipv4_rules` block supports:

* `rule_id` - Rule ID (1 - 65535).
* `comment` - Description.
* `srcaddr` - Source IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).
* `srcport` - Source port (0 - 65535, default = 0, meaning any).
* `dstaddr` - Destination IP address (any | local-LAN | IPv4 address[/<network mask | mask length>], default = any).
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `action` - Policy action (allow | deny).

The `layer3_ipv6_rules` block supports:

* `rule_id` - Rule ID (1 - 65535).
* `comment` - Description.
* `srcaddr` - Source IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `srcport` - Source port (0 - 65535, default = 0, meaning any).
* `dstaddr` - Destination IPv6 address (any | local-LAN | IPv6 address[/prefix length]), default = any.
* `dstport` - Destination port (0 - 65535, default = 0, meaning any).
* `protocol` - Protocol type as defined by IANA (0 - 255, default = 255, meaning any).
* `action` - Policy action (allow | deny).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController AccessControlList can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_accesscontrollist.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
