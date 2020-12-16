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
* `type` - IP pool type (overload, one-to-one, fixed port range, or port block allocation).
* `startip` - (Required) First IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `endip` - (Required) Final IPv4 address (inclusive) in the range for the address pool (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_startip` -  First IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `source_endip` - Final IPv4 address (inclusive) in the range of the source addresses to be translated (format xxx.xxx.xxx.xxx, Default: 0.0.0.0).
* `block_size` -  Number of addresses in a block (64 to 4096, default = 128).
* `num_blocks_per_user` - Number of addresses blocks that can be used by a user (1 to 128, default = 8).
* `pba_timeout` - Port block allocation timeout (seconds).
* `permit_any_host` - Enable/disable full cone NAT.
* `arp_reply` - Enable/disable replying to ARP requests when an IP Pool is added to a policy (default = enable).
* `arp_intf` - Select an interface from available options that will reply to ARP requests. (If blank, any is selected).
* `associated_interface` - Associated interface name.
* `comments` - Comment.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Ippool can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_ippool.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
