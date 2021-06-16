---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_switchinterface"
description: |-
  Configure software switch interfaces by grouping physical and WiFi interfaces.
---

# fortios_system_switchinterface
Configure software switch interfaces by grouping physical and WiFi interfaces.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Interface name (name cannot be in use by any other interfaces, VLANs, or inter-VDOM links).
* `vdom` - VDOM that the software switch belongs to.
* `span_dest_port` - SPAN destination port name. All traffic on the SPAN source ports is echoed to the SPAN destination port.
* `span_source_port` - Physical interface name. Port spanning echoes all traffic on the SPAN source ports to the SPAN destination port. The structure of `span_source_port` block is documented below.
* `member` - Names of the interfaces that belong to the virtual switch. The structure of `member` block is documented below.
* `type` - Type of switch based on functionality: switch for normal functionality, or hub to duplicate packets to all port members. Valid values: `switch`, `hub`.
* `intra_switch_policy` - Allow any traffic between switch interfaces or require firewall policies to allow traffic between switch interfaces. Valid values: `implicit`, `explicit`.
* `mac_ttl` - Duration for which MAC addresses are held in the ARP table (300 - 8640000 sec, default = 300).
* `span` - Enable/disable port spanning. Port spanning echoes traffic received by the software switch to the span destination port. Valid values: `disable`, `enable`.
* `span_direction` - The direction in which the SPAN port operates, either: rx, tx, or both. Valid values: `rx`, `tx`, `both`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `span_source_port` block supports:

* `interface_name` - Physical interface name.

The `member` block supports:

* `interface_name` - Physical interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SwitchInterface can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_switchinterface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
