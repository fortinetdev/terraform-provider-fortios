---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address6"
description: |-
  Get information on an fortios firewall address6.
---

# Data Source: fortios_firewall_address6
Use this data source to get information on an fortios firewall address6

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall address6.

## Attribute Reference

The following attributes are exported:

* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `type` - Type of IPv6 address object (default = ipprefix).
* `start_mac` - First MAC address in the range.
* `end_mac` - Last MAC address in the range.
* `sdn` - SDN.
* `ip6` - IPv6 address prefix (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx).
* `start_ip` - First IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `end_ip` - Final IP address (inclusive) in the range for the address (format: xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx).
* `fqdn` - Fully qualified domain name.
* `country` - IPv6 addresses associated to a specific country.
* `cache_ttl` - Minimal TTL of individual IPv6 addresses in FQDN cache.
* `visibility` - Enable/disable the visibility of the object in the GUI.
* `color` - Integer value to determine the color of the icon in the GUI (range 1 to 32, default = 0, which sets the value to 1).
* `obj_id` - Object ID for NSX.
* `list` - IP address list. The structure of `list` block is documented below.
* `tagging` - Config object tagging The structure of `tagging` block is documented below.
* `comment` - Comment.
* `template` - IPv6 address template.
* `subnet_segment` - IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
* `host_type` - Host type.
* `host` - Host Address.

The `list` block contains:

* `ip` - IP.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

The `subnet_segment` block contains:

* `name` - Name.
* `type` - Subnet segment type.
* `value` - Subnet segment value.

