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
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
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
$ terraform import fortios_switchcontrollerqos_ipdscpmap.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollerqos_ipdscpmap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
