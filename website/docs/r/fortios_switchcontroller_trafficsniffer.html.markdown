---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_trafficsniffer"
description: |-
  Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters.
---

# fortios_switchcontroller_trafficsniffer
Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `mode` - Configure traffic sniffer mode. Valid values: `erspan-auto`, `rspan`, `none`.
* `erspan_ip` - Configure ERSPAN collector IP address.
* `target_mac` - Sniffer MACs to filter. The structure of `target_mac` block is documented below.
* `target_ip` - Sniffer IPs to filter. The structure of `target_ip` block is documented below.
* `target_port` - Sniffer ports to filter. The structure of `target_port` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `target_mac` block supports:

* `mac` - Sniffer MAC.
* `description` - Description for the sniffer MAC.

The `target_ip` block supports:

* `ip` - Sniffer IP.
* `description` - Description for the sniffer IP.

The `target_port` block supports:

* `switch_id` - Managed-switch ID.
* `description` - Description for the sniffer port entry.
* `in_ports` - Configure source ingress port interfaces. The structure of `in_ports` block is documented below.
* `out_ports` - Configure source egress port interfaces. The structure of `out_ports` block is documented below.

The `in_ports` block supports:

* `name` - Interface name.

The `out_ports` block supports:

* `name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController TrafficSniffer can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontroller_trafficsniffer.labelname SwitchControllerTrafficSniffer

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_trafficsniffer.labelname SwitchControllerTrafficSniffer
$ unset "FORTIOS_IMPORT_TABLE"
```
