---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipv6tunnel"
description: |-
  Configure IPv6/IPv4 in IPv6 tunnel.
---

# fortios_system_ipv6tunnel
Configure IPv6/IPv4 in IPv6 tunnel.

## Example Usage

```hcl
resource "fortios_system_ipv6tunnel" "trname" {
  destination = "2001:db8:85a3::8a2e:370:7324"
  interface   = "port3"
  name        = "11111"
  source      = "2001:db8:85a3::8a2e:370:7334"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IPv6 tunnel name.
* `source` - Local IPv6 address of the tunnel.
* `destination` - (Required) Remote IPv6 address of the tunnel.
* `interface` - Interface name.
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
* `auto_asic_offload` - Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Ipv6Tunnel can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ipv6tunnel.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
