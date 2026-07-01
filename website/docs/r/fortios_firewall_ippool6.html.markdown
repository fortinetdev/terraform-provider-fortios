---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ippool6"
description: |-
  Configure IPv6 IP pools.
---

# fortios_firewall_ippool6
Configure IPv6 IP pools.

## Example Usage

```hcl
resource "fortios_firewall_ippool6" "trname" {
  endip   = "2001:3ca1:10f:1a:121b::19"
  name    = "ippool6s1"
  startip = "2001:3ca1:10f:1a:121b::10"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IPv6 IP pool name.
* `type` - Configure IPv6 pool type (overload or NPTv6). Valid values: `overload`, `nptv6`.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `fabric_force_sync` - Enable/disable forced synchronization of configuration objects from the root FortiGate unit to the downstream devices.  Configuration conflict check is skipped. Valid values: `enable`, `disable`.
* `fabric_object_source` - Source of truth for fabric object. Valid values: `member`, `local`, `root`.
* `startip` - (Required) First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `endip` - (Required) Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `internal_prefix` - Internal NPTv6 prefix length (32 - 64).
* `external_prefix` - External NPTv6 prefix length (32 - 64).
* `comments` - Comment.
* `nat46` - Enable/disable NAT46. Valid values: `disable`, `enable`.
* `add_nat46_route` - Enable/disable adding NAT46 route. Valid values: `disable`, `enable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Ippool6 can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_ippool6.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_ippool6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
