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
* `type` - IP pool type (overload, one-to-one, fixed port range, or port block allocation). Valid values: `overload`, `one-to-one`, `fixed-port-range`, `port-block-allocation`.
* `startip` - (Required) First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `endip` - (Required) Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `startport` - First port number (inclusive) in the range for the address pool (Default: 5117).
* `endport` - Final port number (inclusive) in the range for the address pool (Default: 65533).
* `source_startip` -  First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `block_size` -  Number of addresses in a block (64 to 4096, default = 128).
* `port_per_user` -  Number of port for each user (32 to 60416, default = 0, auto).
* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_timeout` - Port block allocation timeout (seconds).
* `permit_any_host` - Enable/disable full cone NAT. Valid values: `disable`, `enable`.
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable). Valid values: `disable`, `enable`.
* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
* `associated_interface` - Associated interface name.
* `comments` - Comment.
* `nat64` - Enable/disable NAT64. Valid values: `disable`, `enable`.
* `add_nat64_route` - Enable/disable adding NAT64 route. Valid values: `disable`, `enable`.
* `subnet_broadcast_in_ippool` - Enable/disable inclusion of the subnetwork address and broadcast IP address in the NAT64 IP pool. Valid values: `disable`, `enable`.
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
