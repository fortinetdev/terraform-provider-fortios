---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sittunnel"
description: |-
  Configure IPv6 tunnel over IPv4.
---

# fortios_system_sittunnel
Configure IPv6 tunnel over IPv4.

## Example Usage

```hcl
resource "fortios_system_sittunnel" "trname" {
  destination = "1.1.1.1"
  interface   = "port2"
  ip6         = "::/0"
  name        = "sittunnel1"
  source      = "2.2.2.2"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Tunnel name.
* `source` - Source IP address of the tunnel.
* `destination` - (Required) Destination IP address of the tunnel.
* `ip6` - IPv6 address of the tunnel.
* `interface` - Interface name.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SitTunnel can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_sittunnel.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
