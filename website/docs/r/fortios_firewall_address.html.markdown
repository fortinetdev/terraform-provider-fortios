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
* `route_tag` - route-tag address.
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
* `sso_attribute_value` - Name(s) of the RADIUS user groups that this address includes. The structure of `sso_attribute_value` block is documented below.
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
* `hw_vendor` - Dynamic address matching hardware vendor.
* `hw_version` - Dynamic address matching hardware version.
* `hw_model` - Dynamic address matching hardware model.
* `os` - Dynamic address matching operating system.
* `sw_version` - Dynamic address matching software version.
* `agent_id` - Telemetry agent id.
* `comment` - Comment.
* `visibility` - Enable/disable address visibility in the GUI. Valid values: `enable`, `disable`.
* `associated_interface` - Network interface associated with address.
* `display_with` - Display object with first tag, all tags, or just the icon & color. Valid values: `all-tags`, `first-tag-only`, `icon-and-color`.
* `custom_tags` - Custom tags. The structure of `custom_tags` block is documented below.
* `color` - Color of icon on the GUI.
* `filter` - Match criteria filter.
* `sdn_addr_type` - Type of addresses to collect. Valid values: `private`, `public`, `all`.
* `node_ip_only` - Enable/disable collection of node addresses only in Kubernetes. Valid values: `enable`, `disable`.
* `obj_id` - Object ID for NSX.
* `list` - IP address list. The structure of `list` block is documented below.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `allow_routing` - Enable/disable use of this address in the static route configuration. Valid values: `enable`, `disable`.
* `managed_subnetwork_size` - Number of IP addresses to be allocated by FortiIPAM for this address. Valid values: `4`, `8`, `16`, `32`, `64`, `128`, `256`, `512`, `1024`, `2048`, `4096`, `8192`, `16384`, `32768`, `65536`, `131072`, `262144`, `524288`, `1048576`, `2097152`, `4194304`, `8388608`, `16777216`.
* `ipam_allocate_unique` - Allocate unique subnet for FortiIPAM managed fabric-object address. Valid values: `enable`, `disable`.
* `passive_fqdn_learning` - Enable/disable passive learning of FQDNs.  When enabled, the FortiGate learns, trusts, and saves FQDNs from endpoint DNS queries (default = enable). Valid values: `disable`, `enable`.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `enable`, `disable`.
* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `macaddr` block supports:

* `macaddr` - MAC address ranges <start>[-<end>] separated by space.

The `fsso_group` block supports:

* `name` - FSSO group name.

The `sso_attribute_value` block supports:

* `name` - RADIUS user group name.

The `custom_tags` block supports:

* `name` - Names of custom tags used with this address.

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
* `addr_8021x` - 802.1X address. The structure of `addr_8021x` block is documented below.
* `obsolete` - Indicates whether the address can be used.

The `addr_8021x` block contains following attibutes:

* `interface` - Interface name.
* `mac` - MAC address.
* `acct_user` - Account user name.
* `ip` - IP address.
* `vlan_id` - VLAN ID.

## Import

Firewall Address can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_address.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_address.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
