---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_gretunnel"
description: |-
  Configure GRE tunnel.
---

# fortios_system_gretunnel
Configure GRE tunnel.

## Example Usage

```hcl
resource "fortios_system_gretunnel" "trname" {
  checksum_reception           = "disable"
  checksum_transmission        = "disable"
  dscp_copying                 = "disable"
  interface                    = "port3"
  ip_version                   = "4"
  keepalive_failtimes          = 10
  keepalive_interval           = 0
  key_inbound                  = 0
  key_outbound                 = 0
  local_gw                     = "3.3.3.3"
  local_gw6                    = "::"
  name                         = "gretunnel1"
  remote_gw                    = "1.1.1.1"
  remote_gw6                   = "::"
  sequence_number_reception    = "disable"
  sequence_number_transmission = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Tunnel name.
* `interface` - Interface name.
* `ip_version` - IP version to use for VPN interface. Valid values: `4`, `6`.
* `remote_gw6` - IPv6 address of the remote gateway.
* `local_gw6` - IPv6 address of the local gateway.
* `remote_gw` - (Required) IP address of the remote gateway.
* `local_gw` - (Required) IP address of the local gateway.
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway. Valid values: `disable`, `enable`.
* `sequence_number_transmission` - Enable/disable including of sequence numbers in transmitted GRE packets. Valid values: `disable`, `enable`.
* `sequence_number_reception` - Enable/disable validating sequence numbers in received GRE packets. Valid values: `disable`, `enable`.
* `checksum_transmission` - Enable/disable including checksums in transmitted GRE packets. Valid values: `disable`, `enable`.
* `checksum_reception` - Enable/disable validating checksums in received GRE packets. Valid values: `disable`, `enable`.
* `key_outbound` - Include this key in transmitted GRE packets (0 - 4294967295).
* `key_inbound` - Require received GRE packets contain this key (0 - 4294967295).
* `dscp_copying` - Enable/disable DSCP copying. Valid values: `disable`, `enable`.
* `diffservcode` - DiffServ setting to be applied to GRE tunnel outer IP header.
* `keepalive_interval` - Keepalive message interval (0 - 32767, 0 = disabled).
* `keepalive_failtimes` - Number of consecutive unreturned keepalive messages before a GRE connection is considered down (1 - 255).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System GreTunnel can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_gretunnel.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
