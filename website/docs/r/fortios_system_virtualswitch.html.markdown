---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualswitch"
description: |-
  Configure virtual hardware switch interfaces.
---

# fortios_system_virtualswitch
Configure virtual hardware switch interfaces.

## Argument Reference

The following arguments are supported:

* `name` - Name of the virtual switch.
* `physical_switch` - Physical switch parent.
* `port` - Configure member ports. The structure of `port` block is documented below.
* `span` - Enable/disable SPAN. Valid values: `disable`, `enable`.
* `span_source_port` - SPAN source ports.
* `span_dest_port` - SPAN destination port.
* `span_direction` - SPAN direction. Valid values: `rx`, `tx`, `both`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `port` block supports:

* `name` - Physical interface name.
* `speed` - Interface speed. Valid values: `auto`, `10full`, `10half`, `100full`, `100half`, `1000full`, `1000half`, `1000auto`.
* `status` - Interface status. Valid values: `up`, `down`.
* `alias` - Alias.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VirtualSwitch can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_virtualswitch.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
