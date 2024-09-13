---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomnetflow"
description: |-
  Configure NetFlow per VDOM.
---

# fortios_system_vdomnetflow
Configure NetFlow per VDOM.

## Example Usage

```hcl
resource "fortios_system_vdomnetflow" "trname" {
  collector_ip   = "0.0.0.0"
  collector_port = 2055
  source_ip      = "0.0.0.0"
  vdom_netflow   = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `vdom_netflow` - Enable/disable NetFlow per VDOM. Valid values: `enable`, `disable`.
* `collectors` - Netflow collectors. The structure of `collectors` block is documented below.
* `collector_ip` - NetFlow collector IP address.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `collectors` block supports:

* `id` - ID.
* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `source_ip_interface` - Name of the interface used to determine the source IP for exporting packets.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VdomNetflow can be imported using any of these accepted formats:
```
$ terraform import fortios_system_vdomnetflow.labelname SystemVdomNetflow

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_vdomnetflow.labelname SystemVdomNetflow
$ unset "FORTIOS_IMPORT_TABLE"
```
