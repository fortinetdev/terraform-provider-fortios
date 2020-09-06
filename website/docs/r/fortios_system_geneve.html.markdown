---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_geneve"
description: |-
  Configure GENEVE devices.
---

# fortios_system_geneve
Configure GENEVE devices.

## Example Usage

```hcl
resource "fortios_system_geneve" "trname" {
  dstport    = 22
  interface  = "port2"
  ip_version = "ipv4-unicast"
  name       = "1"
  remote_ip  = "1.1.1.1"
  remote_ip6 = "::"
  vni        = 0
}
```

## Argument Reference

The following arguments are supported:

* `name` - GENEVE device or interface name. Must be an unique interface name.
* `interface` - (Required) Outgoing interface for GENEVE encapsulated traffic.
* `vni` - (Required) GENEVE network ID.
* `ip_version` - (Required) IP version to use for the GENEVE interface and so for communication over the GENEVE. IPv4 or IPv6 unicast.
* `remote_ip` - (Required) IPv4 address of the GENEVE interface on the device at the remote end of the GENEVE.
* `remote_ip6` - (Required) IPv6 IP address of the GENEVE interface on the device at the remote end of the GENEVE.
* `dstport` - GENEVE destination port (1 - 65535, default = 6081).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Geneve can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_geneve.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
