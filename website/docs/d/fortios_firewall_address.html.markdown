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
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Address name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `subnet` - IP address and subnet mask of address.
* `type` - Type of address.
* `sub_type` - Sub-type of address.
* `clearpass_spt` - SPT (System Posture Token) value.
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
* `obj_type` - Object type.
* `tag_detection_level` - Tag detection level of dynamic address object.
* `tag_type` - Tag type of dynamic address object.
* `comment` - Comment.
* `visibility` - Enable/disable address visibility in the GUI.
* `associated_interface` - Network interface associated with address.
* `color` - Color of icon on the GUI.
* `filter` - Match criteria filter.
* `sdn_addr_type` - Type of addresses to collect.
* `node_ip_only` - Enable/disable collection of node addresses only in Kubernetes.
* `obj_id` - Object ID for NSX.
* `list` - IP address list. The structure of `list` block is documented below.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `allow_routing` - Enable/disable use of this address in the static route configuration.
* `fabric_object` - Security Fabric global object setting.

The `macaddr` block contains:

* `macaddr` - MAC address ranges <start>[-<end>] separated by space.

The `fsso_group` block contains:

* `name` - FSSO group name.

The `list` block contains:

* `ip` - IP.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

