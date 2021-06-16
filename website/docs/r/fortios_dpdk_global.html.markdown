---
subcategory: "FortiGate DPDK"
layout: "fortios"
page_title: "FortiOS: fortios_dpdk_global"
description: |-
  Configure global DPDK options.
---

# fortios_dpdk_global
Configure global DPDK options.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable DPDK operation for the entire system. Valid values: `disable`, `enable`.
* `interface` - Physical interfaces that enable DPDK. The structure of `interface` block is documented below.
* `multiqueue` - Enable/disable multi-queue RX/TX support for all DPDK ports. Valid values: `disable`, `enable`.
* `sleep_on_idle` - Enable/disable sleep-on-idle support for all FDH engines. Valid values: `disable`, `enable`.
* `elasticbuffer` - Enable/disable elasticbuffer support for all DPDK ports. Valid values: `disable`, `enable`.
* `per_session_accounting` - Enable/disable per-session accounting. Valid values: `disable`, `traffic-log-only`, `enable`.
* `hugepage_percentage` - Percentage of main memory allocated to hugepages, which are available for DPDK operation.
* `mbufpool_percentage` - Percentage of main memory allocated to DPDK packet buffer.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `interface` block supports:

* `interface_name` - Physical interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dpdk Global can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_dpdk_global.labelname DpdkGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
