---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_centralsnatmap"
description: |-
  Configure central SNAT policies.
---

# fortios_firewall_centralsnatmap
Configure central SNAT policies.

## Example Usage

```hcl
resource "fortios_firewall_centralsnatmap" "trname" {
  nat       = "enable"
  nat_port  = "0"
  orig_port = "0"
  policyid  = 1
  protocol  = 33
  status    = "enable"

  dst_addr {
    name = "all"
  }

  dstintf {
    name = "port3"
  }

  orig_addr {
    name = "all"
  }

  srcintf {
    name = "port1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `policyid` - Policy ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `status` - Enable/disable the active status of this policy. Valid values: `enable`, `disable`.
* `type` - IPv4/IPv6 source NAT. Valid values: `ipv4`, `ipv6`.
* `orig_addr` - (Required) Original address. The structure of `orig_addr` block is documented below.
* `orig_addr6` - IPv6 Original address. The structure of `orig_addr6` block is documented below.
* `srcintf` - (Required) Source interface name from available interfaces. The structure of `srcintf` block is documented below.
* `dst_addr` - (Required) Destination address name from available addresses. The structure of `dst_addr` block is documented below.
* `dst_addr6` - IPv6 Destination address. The structure of `dst_addr6` block is documented below.
* `dstintf` - (Required) Destination interface name from available interfaces. The structure of `dstintf` block is documented below.
* `nat_ippool` - Name of the IP pools to be used to translate addresses from available IP Pools. The structure of `nat_ippool` block is documented below.
* `nat_ippool6` - IPv6 pools to be used for source NAT. The structure of `nat_ippool6` block is documented below.
* `protocol` - (Required) Integer value for the protocol type (0 - 255).
* `orig_port` - (Required) Original TCP port (0 to 65535).
* `nat_port` - Translated port or port range (0 to 65535).
* `nat` - (Required) Enable/disable source NAT. Valid values: `disable`, `enable`.
* `comments` - Comment.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `orig_addr` block supports:

* `name` - Address name.

The `orig_addr6` block supports:

* `name` - Address name.

The `srcintf` block supports:

* `name` - Interface name.

The `dst_addr` block supports:

* `name` - Address name.

The `dst_addr6` block supports:

* `name` - Address name.

The `dstintf` block supports:

* `name` - Interface name.

The `nat_ippool` block supports:

* `name` - IP pool name.

The `nat_ippool6` block supports:

* `name` - IPv6 pool name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall CentralSnatMap can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_centralsnatmap.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
