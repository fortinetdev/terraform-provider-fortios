---
subcategory: "FortiGate DPDK"
layout: "fortios"
page_title: "FortiOS: fortios_dpdk_global"
description: |-
  Configure global DPDK options.
---

# fortios_dpdk_global
Configure global DPDK options. Applies to FortiOS Version `6.2.4,6.2.6,6.4.0,6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.0.14,7.0.15,7.0.16,7.0.17,7.0.18,7.0.19,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.2.7,7.2.8,7.2.9,7.2.10,7.2.11,7.2.12,7.4.0,7.4.1,7.4.2,7.4.3,7.4.4,7.4.5,7.4.6,7.4.7,7.4.8,7.4.9,7.4.11,7.4.12,7.6.0,7.6.1,7.6.2,7.6.3,7.6.4,7.6.5,7.6.6,7.6.7`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
* `interface` - Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
* `multiqueue` - Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
* `sleep_on_idle` - Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
* `elasticbuffer` - Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
* `protects` - Special arguments for device
* `per_session_accounting` - Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
* `ipsec_offload` - Enable/disable DPDK IPsec phase 2 offloading. Valid values: `disable`, `enable`.
* `frag_offload` - Enable/disable DPDK fragmentation/defragmentation offloading (default = enable). Valid values: `disable`, `enable`.
* `hugepage_percentage` - Percentage of main memory allocated to hugepages, which are available for DPDK operation.
* `mbufpool_percentage` - Percentage of main memory allocated to DPDK packet buffer.
* `session_table_percentage` - Percentage of main memory allocated to DPDK session table.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `interface` block supports:

* `interface_name` - Physical interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dpdk Global can be imported using any of these accepted formats:
```
$ terraform import fortios_dpdk_global.labelname DpdkGlobal

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_dpdk_global.labelname DpdkGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
