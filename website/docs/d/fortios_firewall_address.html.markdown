---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address"
description: |-
  Get information on an fortios firewall address.
---

# Data Source: fortios_firewall_address
Use this data source to get information on an fortios firewall address

## Example Usage

```hcl
data "fortios_firewall_address" sample1 {
  name = "google-play"
}

output output1 {
  value = join(": ", [data.fortios_firewall_address.sample1.type, data.fortios_firewall_address.sample1.fqdn])
}
```

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall address.

## Attribute Reference

The following attributes are exported:

* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `subnet` - IP address and subnet mask of address.
* `type` - Type of address.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `fqdn` - Fully Qualified Domain Name address.
* `country` - IP addresses associated to a specific country.
* `wildcard_fqdn` - Fully Qualified Domain Name with wildcard characters.
* `cache_ttl` - Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
* `wildcard` - IP address and wildcard netmask.
* `sdn` - SDN.
* `tenant` - Tenant.
* `organization` - Organization domain name (Syntax: organization/domain).
* `epg_name` - Endpoint group name.
* `subnet_name` - Subnet name.
* `sdn_tag` - SDN Tag.
* `policy_group` - Policy group name.
* `comment` - Comment.
* `visibility` - Enable/disable address visibility in the GUI.
* `associated_interface` - Network interface associated with address.
* `color` - Color of icon on the GUI.
* `filter` - Match criteria filter.
* `obj_id` - Object ID for NSX.
* `list` - IP address list. The structure of `list` block is documented below.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `allow_routing` - Enable/disable use of this address in the static route configuration.

The `list` block contains:

* `ip` - IP.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.
