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

The `map` block supports:

* `name` - Dscp mapping entry name.
* `cos_queue` - COS queue number.
* `diffserv` - Differentiated service.
* `ip_precedence` - IP Precedence.
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
