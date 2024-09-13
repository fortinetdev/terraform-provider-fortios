---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_rmonstat"
description: |-
  SNMP Remote Network Monitoring (RMON) Ethernet statistics configuration.
---

# fortios_systemsnmp_rmonstat
SNMP Remote Network Monitoring (RMON) Ethernet statistics configuration. Applies to FortiOS Version `>= 7.6.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - RMON Ethernet statistics entry ID.
* `source` - Data source of the Ethernet statistics entry.
* `owner` - Owner of the Ethernet statistics entry.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

SystemSnmp RmonStat can be imported using any of these accepted formats:
```
$ terraform import fortios_systemsnmp_rmonstat.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_systemsnmp_rmonstat.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
