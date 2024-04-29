---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_sysinfo"
description: |-
  SNMP system info configuration.
---

# fortios_systemsnmp_sysinfo
SNMP system info configuration.

## Example Usage

```hcl
resource "fortios_systemsnmp_sysinfo" "trname" {
  status                    = "disable"
  trap_high_cpu_threshold   = 80
  trap_log_full_threshold   = 90
  trap_low_memory_threshold = 80
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable SNMP. Valid values: `enable`, `disable`.
* `engine_id_type` - Local SNMP engineID type (text/hex/mac). Valid values: `text`, `hex`, `mac`.
* `engine_id` - Local SNMP engineID string. On FortiOS versions 6.2.0-7.0.0: maximum 24 characters. On FortiOS versions >= 7.0.1: maximum 27 characters.
* `description` - System description.
* `contact_info` - Contact information.
* `location` - System location.
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `append_index` - Enable/disable allowance of appending VDOM or interface index in some RFC tables. Valid values: `enable`, `disable`.
* `trap_free_memory_threshold` - Free memory usage when trap is sent.
* `trap_freeable_memory_threshold` - Freeable memory usage when trap is sent.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemSnmp Sysinfo can be imported using any of these accepted formats:
```
$ terraform import fortios_systemsnmp_sysinfo.labelname SystemSnmpSysinfo

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_systemsnmp_sysinfo.labelname SystemSnmpSysinfo
$ unset "FORTIOS_IMPORT_TABLE"
```
