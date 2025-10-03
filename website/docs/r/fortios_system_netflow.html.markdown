---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_netflow"
description: |-
  Configure NetFlow.
---

# fortios_system_netflow
Configure NetFlow.

## Example Usage

```hcl
resource "fortios_system_netflow" "trname" {
  active_flow_timeout   = 30
  collector_ip          = "0.0.0.0"
  collector_port        = 2055
  inactive_flow_timeout = 15
  source_ip             = "0.0.0.0"
  template_tx_counter   = 20
  template_tx_timeout   = 30
}
```

## Argument Reference

The following arguments are supported:

* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `active_flow_timeout` - Timeout to report active flows. On FortiOS versions 6.2.0-7.0.0: 1 - 60 min, default = 30. On FortiOS versions >= 7.0.1: 60 - 3600 sec, default = 1800.
* `inactive_flow_timeout` - Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
* `template_tx_timeout` - Timeout for periodic template flowset transmission. On FortiOS versions 6.2.0-7.0.0: 1 - 1440 min, default = 30. On FortiOS versions >= 7.0.1: 60 - 86400 sec, default = 1800.
* `template_tx_counter` - Counter of flowset records before resending a template flowset record.
* `session_cache_size` - Maximum RAM usage allowed for Netflow session cache. Valid values: `min`, `default`, `max`.
* `exclusion_filters` - Exclusion filters The structure of `exclusion_filters` block is documented below.
* `collectors` - Netflow collectors. The structure of `collectors` block is documented below.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `exclusion_filters` block supports:

* `id` - Filter ID.
* `source_ip` - Session source address.
* `destination_ip` - Session destination address.
* `source_port` - Session source port number or range.
* `destination_port` - Session destination port number or range.
* `protocol` - Session IP protocol (0 - 255, default = 255, meaning any).

The `collectors` block supports:

* `id` - ID.
* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `source_ip_interface` - Name of the interface used to determine the source IP for exporting packets.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Netflow can be imported using any of these accepted formats:
```
$ terraform import fortios_system_netflow.labelname SystemNetflow

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_netflow.labelname SystemNetflow
$ unset "FORTIOS_IMPORT_TABLE"
```
