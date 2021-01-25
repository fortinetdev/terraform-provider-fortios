---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vxlan"
description: |-
  Configure VXLAN devices.
---

# fortios_system_vxlan
Configure VXLAN devices.

## Example Usage

```hcl
resource "fortios_system_vxlan" "trname" {
  dstport    = 4789
  interface  = "port3"
  ip_version = "ipv4-unicast"
  remote_ip {
    ip = "1.1.1.1"
  }
  name = "18"
  vni  = 3
}
```

## Argument Reference


The following arguments are supported:

* `name` - VXLAN device or interface name. Must be a unique interface name.
* `interface` - (Required) Outgoing interface for VXLAN encapsulated traffic.
* `vni` - (Required) VXLAN network ID.
* `ip_version` - (Required) IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast.
* `remote_ip` - IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip` block is documented below.
* `remote_ip6` - IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip6` block is documented below.
* `dstport` - VXLAN destination port (1 - 65535, default = 4789).
* `multicast_ttl` - VXLAN multicast TTL (1-255, default = 0).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `remote_ip` block supports:

* `ip` - IPv4 address.

The `remote_ip6` block supports:

* `ip6` - IPv6 address.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Vxlan can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_vxlan.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
