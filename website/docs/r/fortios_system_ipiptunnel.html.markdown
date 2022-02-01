---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipiptunnel"
description: |-
  Configure IP in IP Tunneling.
---

# fortios_system_ipiptunnel
Configure IP in IP Tunneling.

## Example Usage

```hcl
resource "fortios_system_ipiptunnel" "trname" {
  interface = "port3"
  local_gw  = "1.1.1.1"
  name      = "12"
  remote_gw = "2.2.2.2"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IPIP Tunnel name.
* `interface` - (Required) Interface name that is associated with the incoming traffic from available options.
* `remote_gw` - (Required) IPv4 address for the remote gateway.
* `local_gw` - (Required) IPv4 address for the local gateway.
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
* `auto_asic_offload` - Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System IpipTunnel can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ipiptunnel.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
