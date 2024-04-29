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
* `ip_version` - (Required) IP version to use for the VXLAN interface and so for communication over the VXLAN. IPv4 or IPv6 unicast or multicast. Valid values: `ipv4-unicast`, `ipv6-unicast`, `ipv4-multicast`, `ipv6-multicast`.
* `remote_ip` - IPv4 address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip` block is documented below.
* `remote_ip6` - IPv6 IP address of the VXLAN interface on the device at the remote end of the VXLAN. The structure of `remote_ip6` block is documented below.
* `dstport` - VXLAN destination port (1 - 65535, default = 4789).
* `multicast_ttl` - VXLAN multicast TTL (1-255, default = 0).
* `evpn_id` - EVPN instance.
* `learn_from_traffic` - Enable/disable VXLAN MAC learning from traffic. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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
$ terraform import fortios_system_vxlan.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_vxlan.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
