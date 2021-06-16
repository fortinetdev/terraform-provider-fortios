---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemdhcp6_server"
description: |-
  Configure DHCPv6 servers.
---

# fortios_systemdhcp6_server
Configure DHCPv6 servers.

## Example Usage

```hcl
resource "fortios_systemdhcp6_server" "trname" {
  fosid        = 1
  interface    = "port3"
  lease_time   = 604800
  rapid_commit = "disable"
  status       = "enable"
  subnet       = "2001:db8:1234:113::/64"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `status` - Enable/disable this DHCPv6 configuration. Valid values: `disable`, `enable`.
* `rapid_commit` - Enable/disable allow/disallow rapid commit. Valid values: `disable`, `enable`.
* `lease_time` - Lease time in seconds, 0 means unlimited.
* `dns_service` -  Options for assigning DNS servers to DHCPv6 clients. Valid values: `delegated`, `default`, `specify`.
* `dns_search_list` - DNS search list options. Valid values: `delegated`, `specify`.
* `dns_server1` - DNS server 1.
* `dns_server2` - DNS server 2.
* `dns_server3` - DNS server 3.
* `dns_server4` - DNS server 4.
* `domain` - Domain name suffix for the IP addresses that the DHCP server assigns to clients.
* `subnet` - (Required) Subnet or subnet-id if the IP mode is delegated.
* `interface` - (Required) DHCP server can assign IP configurations to clients connected to this interface.
* `option1` - Option 1.
* `option2` - Option 2.
* `option3` - Option 3.
* `upstream_interface` - Interface name from where delegated information is provided.
* `ip_mode` - Method used to assign client IP. Valid values: `range`, `delegated`.
* `prefix_mode` - Assigning a prefix from a DHCPv6 client or RA. Valid values: `dhcp6`, `ra`.
* `prefix_range` - DHCP prefix configuration. The structure of `prefix_range` block is documented below.
* `ip_range` - DHCP IP range configuration. The structure of `ip_range` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `prefix_range` block supports:

* `id` - ID.
* `start_prefix` - Start of prefix range.
* `end_prefix` - End of prefix range.
* `prefix_length` - Prefix length.

The `ip_range` block supports:

* `id` - ID.
* `start_ip` - Start of IP range.
* `end_ip` - End of IP range.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SystemDhcp6 Server can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemdhcp6_server.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
