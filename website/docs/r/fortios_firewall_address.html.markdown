---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address"
description: |-
  Configure IPv4 addresses.
---

# fortios_firewall_address
Configure IPv4 addresses.

## Example Usage

```hcl
resource "fortios_firewall_address" "trname" {
  allow_routing        = "disable"
  associated_interface = "port2"
  color                = 3
  end_ip               = "255.255.255.0"
  name                 = "testaddress"
  start_ip             = "22.1.1.0"
  subnet               = "22.1.1.0 255.255.255.0"
  type                 = "ipmask"
  visibility           = "enable"
}
```

## Argument Reference

The following arguments are supported:

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
* `list` - IP address list.
* `tagging` - Config object tagging.
* `allow_routing` - Enable/disable use of this address in the static route configuration.

The `list` block supports:

* `ip` - IP.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags.

The `tags` block supports:

* `name` - Tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_address.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
