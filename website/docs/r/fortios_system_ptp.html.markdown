---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ptp"
description: |-
  Configure system PTP information.
---

# fortios_system_ptp
Configure system PTP information.

## Example Usage

```hcl
resource "fortios_system_ptp" "trname" {
  delay_mechanism  = "E2E"
  interface        = "port3"
  mode             = "multicast"
  request_interval = 1
  status           = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable setting the FortiGate system time by synchronizing with an PTP Server. Valid values: `enable`, `disable`.
* `mode` - Multicast transmission or hybrid transmission. Valid values: `multicast`, `hybrid`.
* `delay_mechanism` - End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.
* `request_interval` - The delay request value is the logarithmic mean interval in seconds between the delay request messages sent by the slave to the master.
* `interface` - (Required) PTP slave will reply through this interface.
* `server_mode` - Enable/disable FortiGate PTP server mode. Your FortiGate becomes an PTP server for other devices on your network. Valid values: `enable`, `disable`.
* `server_interface` - FortiGate interface(s) with PTP server mode enabled. Devices on your network can contact these interfaces for PTP services. The structure of `server_interface` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `server_interface` block supports:

* `id` - ID.
* `server_interface_name` - Interface name.
* `delay_mechanism` - End to end delay detection or peer to peer delay detection. Valid values: `E2E`, `P2P`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ptp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ptp.labelname SystemPtp
$ unset "FORTIOS_IMPORT_TABLE"
```
