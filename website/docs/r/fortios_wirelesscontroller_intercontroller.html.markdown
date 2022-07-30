---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_intercontroller"
description: |-
  Configure inter wireless controller operation.
---

# fortios_wirelesscontroller_intercontroller
Configure inter wireless controller operation.

## Example Usage

```hcl
resource "fortios_wirelesscontroller_intercontroller" "trname" {
  fast_failover_max     = 10
  fast_failover_wait    = 10
  inter_controller_key  = "ENC XXXX"
  inter_controller_mode = "disable"
  inter_controller_pri  = "primary"
}
```

## Argument Reference

The following arguments are supported:

* `inter_controller_mode` - Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable). Valid values: `disable`, `l2-roaming`, `1+1`.
* `l3_roaming` - Enable/disable layer 3 roaming (default = disable). Valid values: `enable`, `disable`.
* `inter_controller_key` - Secret key for inter-controller communications.
* `inter_controller_pri` - Configure inter-controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.
* `fast_failover_max` - Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
* `fast_failover_wait` - Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
* `inter_controller_peer` - Fast failover peer wireless controller list. The structure of `inter_controller_peer` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `inter_controller_peer` block supports:

* `id` - ID.
* `peer_ip` - Peer wireless controller's IP address.
* `peer_port` - Port used by the wireless controller's for inter-controller communications (1024 - 49150, default = 5246).
* `peer_priority` - Peer wireless controller's priority (primary or secondary, default = primary). Valid values: `primary`, `secondary`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController InterController can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_intercontroller.labelname WirelessControllerInterController

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_intercontroller.labelname WirelessControllerInterController
$ unset "FORTIOS_IMPORT_TABLE"
```
