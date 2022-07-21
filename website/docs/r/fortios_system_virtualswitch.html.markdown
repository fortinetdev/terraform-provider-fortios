---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualswitch"
description: |-
  Configure virtual hardware switch interfaces.
---

# fortios_system_virtualswitch
Configure virtual hardware switch interfaces. Applies to FortiOS Version `7.0.4`.

## Argument Reference

The following arguments are supported:

* `name` - Name of the virtual switch.
* `physical_switch` - Physical switch parent.
* `vlan` - VLAN.
* `port` - Configure member ports. The structure of `port` block is documented below.
* `span` - Enable/disable SPAN. Valid values: `disable`, `enable`.
* `span_source_port` - SPAN source port.
* `span_dest_port` - SPAN destination port.
* `span_direction` - SPAN direction. Valid values: `rx`, `tx`, `both`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `port` block supports:

* `name` - Physical interface name.
* `alias` - Alias.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VirtualSwitch can be imported using any of these accepted formats:
```
$ terraform import fortios_system_virtualswitch.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_virtualswitch.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
