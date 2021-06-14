---
subcategory: "FortiGate DPDK"
layout: "fortios"
page_title: "FortiOS: fortios_dpdk_cpus"
description: |-
  Configure CPUs enabled to run engines in each DPDK stage.
---

# fortios_dpdk_cpus
Configure CPUs enabled to run engines in each DPDK stage.

## Argument Reference

The following arguments are supported:

* `rx_cpus` - CPUs enabled to run DPDK RX engines.
* `vnp_cpus` - CPUs enabled to run DPDK VNP engines.
* `ips_cpus` - CPUs enabled to run DPDK IPS engines.
* `tx_cpus` - CPUs enabled to run DPDK TX engines.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Dpdk Cpus can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_dpdk_cpus.labelname DpdkCpus
$ unset "FORTIOS_IMPORT_TABLE"
```
