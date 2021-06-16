---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_trafficsniffer"
description: |-
  Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters.
---

# fortios_switchcontroller_trafficsniffer
Configure FortiSwitch RSPAN/ERSPAN traffic sniffing parameters.

## Argument Reference

The following arguments are supported:

* `mode` - Configure traffic sniffer mode. Valid values: `erspan-auto`, `rspan`, `none`.
* `erspan_ip` - Configure ERSPAN collector IP address.
* `target_mac` - Sniffer MACs to filter. The structure of `target_mac` block is documented below.
* `target_ip` - Sniffer IPs to filter. The structure of `target_ip` block is documented below.
* `target_port` - Sniffer ports to filter. The structure of `target_port` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_trafficsniffer.labelname SwitchControllerTrafficSniffer
$ unset "FORTIOS_IMPORT_TABLE"
```
