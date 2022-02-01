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
* `sub_type` - Sub-type of address.
* `clearpass_spt` - SPT (System Posture Token) value. Valid values: `unknown`, `healthy`, `quarantine`, `checkup`, `transient`, `infected`.
* `macaddr` - Multiple MAC address ranges. The structure of `macaddr` block is documented below.
* `start_mac` - First MAC address in the range.
* `end_mac` - Last MAC address in the range.
* `start_ip` - First IP address (inclusive) in the range for the address.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `fqdn` - Fully Qualified Domain Name address.
* `country` - IP addresses associated to a specific country.
* `wildcard_fqdn` - Fully Qualified Domain Name with wildcard characters.
* `cache_ttl` - Defines the minimal TTL of individual IP addresses in FQDN cache measured in seconds.
* `wildcard` - IP address and wildcard netmask.
* `sdn` - SDN.
* `fsso_group` - FSSO group(s). The structure of `fsso_group` block is documented below.
* `interface` - Name of interface whose IP address is to be used.
* `tenant` - Tenant.
* `organization` - Organization domain name (Syntax: organization/domain).
* `epg_name` - Endpoint group name.
* `subnet_name` - Subnet name.
* `sdn_tag` - SDN Tag.
* `policy_group` - Policy group name.
* `obj_tag` - Tag of dynamic address object.
* `obj_type` - Object type. Valid values: `ip`, `mac`.
* `tag_detection_level` - Tag detection level of dynamic address object.
* `tag_type` - Tag type of dynamic address object.
* `comment` - Comment.
* `visibility` - Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
* `associated_interface` - Network interface associated with address.
* `color` - Color of icon on the GUI.
* `filter` - Match criteria filter.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `private`, `public`, `all`.
* `node_ip_only` - Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
* `obj_id` - Object ID for NSX.
* `list` - IP address list. The structure of `list` block is documented below.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `allow_routing` - Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `macaddr` block supports:

* `macaddr` - MAC address ranges <start>[-<end>] separated by space.

The `fsso_group` block supports:

* `name` - FSSO group name.

The `list` block supports:

* `ip` - IP.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

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
