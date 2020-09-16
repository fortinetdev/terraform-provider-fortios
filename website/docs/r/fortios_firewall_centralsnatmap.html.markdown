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
* `status` - Enable/disable the active status of this policy.
* `orig_addr` - (Required) Original address.
* `srcintf` - (Required) Source interface name from available interfaces.
* `dst_addr` - (Required) Destination address name from available addresses.
* `dstintf` - (Required) Destination interface name from available interfaces.
* `nat_ippool` - Name of the IP pools to be used to translate addresses from available IP Pools.
* `protocol` - (Required) Integer value for the protocol type (0 - 255).
* `orig_port` - (Required) Original TCP port (0 to 65535).
* `nat_port` - Translated port or port range (0 to 65535).
* `nat` - (Required) Enable/disable source NAT.
* `comments` - Comment.

The `orig_addr` block supports:

* `name` - Address name.

The `srcintf` block supports:

* `name` - Interface name.

The `dst_addr` block supports:

* `name` - Address name.

The `dstintf` block supports:

* `name` - Interface name.

The `nat_ippool` block supports:

* `name` - IP pool name.


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
