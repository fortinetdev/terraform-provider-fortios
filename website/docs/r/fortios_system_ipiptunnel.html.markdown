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
