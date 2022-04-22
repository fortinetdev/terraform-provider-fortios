---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_lldpsettings"
description: |-
  Configure FortiSwitch LLDP settings.
---

# fortios_switchcontroller_lldpsettings
Configure FortiSwitch LLDP settings.

## Example Usage

```hcl
resource "fortios_switchcontroller_lldpsettings" "trname" {
  fast_start_interval  = 2
  management_interface = "internal"
  status               = "enable"
  tx_hold              = 4
  tx_interval          = 30
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable LLDP global settings. Valid values: `enable`, `disable`.
* `tx_hold` - Number of tx-intervals before local LLDP data expires (1 - 16, default = 4). Packet TTL is tx-hold * tx-interval.
* `tx_interval` - Frequency of LLDP PDU transmission from FortiSwitch (5 - 4095 sec, default = 30). Packet TTL is tx-hold * tx-interval.
* `fast_start_interval` - Frequency of LLDP PDU transmission from FortiSwitch for the first 4 packets when the link is up (2 - 5 sec, default = 2, 0 = disable fast start).
* `management_interface` - Primary management interface to be advertised in LLDP and CDP PDUs. Valid values: `internal`, `mgmt`.
* `device_detection` - Enable/disable dynamic detection of LLDP neighbor devices for VLAN assignment. Valid values: `disable`, `enable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController LldpSettings can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_lldpsettings.labelname SwitchControllerLldpSettings

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_lldpsettings.labelname SwitchControllerLldpSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
