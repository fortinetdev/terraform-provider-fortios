---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ippool"
description: |-
  Configure IPv4 IP pools.
---

# fortios_firewall_ippool
Configure IPv4 IP pools.

## Example Usage

```hcl
resource "fortios_firewall_ippool" "trname" {
  arp_reply           = "enable"
  block_size          = 128
  endip               = "1.0.0.20"
  name                = "ippools1"
  num_blocks_per_user = 8
  pba_timeout         = 30
  permit_any_host     = "disable"
  source_endip        = "0.0.0.0"
  source_startip      = "0.0.0.0"
  startip             = "1.0.0.0"
  type                = "overload"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IP pool name.
* `type` - IP pool type. On FortiOS versions 6.2.0-7.4.1: overload, one-to-one, fixed port range, or port block allocation. On FortiOS versions >= 7.4.2: overload, one-to-one, fixed-port-range, port-block-allocation, cgn-resource-allocation (hyperscale vdom only). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
* `startip` - (Required) First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `endip` - (Required) Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startport` - First port number (inclusive) in the range for the address pool (1024 - 65535, Default: 5117).
* `endport` - Final port number (inclusive) in the range for the address pool (1024 - 65535, Default: 65533).
* `source_startip` -  First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `block_size` - Number of addresses in a block (64 - 4096, default = 128).
* `port_per_user` - Number of port for each user (32 - 60416, default = 0, which is auto).
* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_timeout` - Port block allocation timeout (seconds).
* `pba_interim_log` - Port block allocation interim logging interval (600 - 86400 seconds, default = 0 which disables interim logging).
* `permit_any_host` - Enable/disable full cone NAT. Valid values: `disable`, `enable`.
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
* `associated_interface` - Associated interface name.
* `comments` - Comment.
* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.
* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
* `source_prefix6` - Source IPv6 network to be translated (format = xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx/xxx, default = ::/0).
* `client_prefix_length` - Subnet length of a single deterministic NAT64 client (1 - 128, default = 64).
* `tcp_session_quota` - Maximum number of concurrent TCP sessions allowed per client (0 - 2097000, default = 0 which means no limit).
* `udp_session_quota` - Maximum number of concurrent UDP sessions allowed per client (0 - 2097000, default = 0 which means no limit).
* `icmp_session_quota` - Maximum number of concurrent ICMP sessions allowed per client (0 - 2097000, default = 0 which means no limit).
* `privileged_port_use_pba` - Enable/disable selection of the external port from the port block allocation for NAT'ing privileged ports (deafult = disable). Valid values: `disable`, `enable`.
* `subnet_broadcast_in_ippool` - Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Ippool can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_ippool.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_ippool.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
