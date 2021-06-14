---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollerqos_ipdscpmap"
description: |-
  Configure FortiSwitch QoS IP precedence/DSCP.
---

# fortios_switchcontrollerqos_ipdscpmap
Configure FortiSwitch QoS IP precedence/DSCP.

## Example Usage

```hcl
resource "fortios_switchcontrollerqos_ipdscpmap" "trname" {
  description = "DEIW"
  name        = "1"

  map {
    cos_queue = 3
    diffserv  = "CS0 CS1 AF11"
    name      = "1"
  }
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Dscp map name.
* `description` - Description of the ip-dscp map name.
* `map` - Maps between IP-DSCP value to COS queue. The structure of `map` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `map` block supports:

* `name` - Dscp mapping entry name.
* `cos_queue` - COS queue number.
* `diffserv` - Differentiated service. Valid values: `CS0`, `CS1`, `AF11`, `AF12`, `AF13`, `CS2`, `AF21`, `AF22`, `AF23`, `CS3`, `AF31`, `AF32`, `AF33`, `CS4`, `AF41`, `AF42`, `AF43`, `CS5`, `EF`, `CS6`, `CS7`.
* `ip_precedence` - IP Precedence. Valid values: `network-control`, `internetwork-control`, `critic-ecp`, `flashoverride`, `flash`, `immediate`, `priority`, `routine`.
* `value` - Raw values of DSCP (0 - 63).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerQos IpDscpMap can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontrollerqos_ipdscpmap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
