---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_sysinfo"
description: |-
  Get information on fortios systemsnmp sysinfo.
---

# Data Source: fortios_systemsnmp_sysinfo
Use this data source to get information on fortios systemsnmp sysinfo

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable SNMP.
* `engine_id_type` - Local SNMP engineID type (text/hex/mac).
* `engine_id` - Local SNMP engineID string (maximum 24 characters).
* `description` - System description.
* `contact_info` - Contact information.
* `location` - System location.
* `trap_high_cpu_threshold` - CPU usage when trap is sent.
* `trap_low_memory_threshold` - Memory usage when trap is sent.
* `trap_log_full_threshold` - Log disk usage when trap is sent.
* `trap_free_memory_threshold` - Free memory usage when trap is sent.
* `trap_freeable_memory_threshold` - Freeable memory usage when trap is sent.

