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

* `inter_controller_mode` - Configure inter-controller mode (disable, l2-roaming, 1+1, default = disable).
* `inter_controller_key` - Secret key for inter-controller communications.
* `inter_controller_pri` - Configure inter-controller's priority (primary or secondary, default = primary).
* `fast_failover_max` - Maximum number of retransmissions for fast failover HA messages between peer wireless controllers (3 - 64, default = 10).
* `fast_failover_wait` - Minimum wait time before an AP transitions from secondary controller to primary controller (10 - 86400 sec, default = 10).
* `inter_controller_peer` - Fast failover peer wireless controller list.

The `inter_controller_peer` block supports:

* `id` - ID.
* `peer_ip` - Peer wireless controller's IP address.
* `peer_port` - Port used by the wireless controller's for inter-controller communications (1024 - 49150, default = 5246).
* `peer_priority` - Peer wireless controller's priority (primary or secondary, default = primary).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

WirelessController InterController can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_intercontroller.labelname WirelessControllerInterController
$ unset "FORTIOS_IMPORT_TABLE"
```
