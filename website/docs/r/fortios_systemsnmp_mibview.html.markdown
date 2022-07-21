---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemsnmp_mibview"
description: |-
  SNMP Access Control MIB View configuration.
---

# fortios_systemsnmp_mibview
SNMP Access Control MIB View configuration. Applies to FortiOS Version `>= 7.2.0`.

## Argument Reference

The following arguments are supported:

* `name` - MIB view name.
* `include` - The OID subtrees to be included in the view. Maximum 16 allowed.
* `exclude` - The OID subtrees to be excluded in the view. Maximum 64 allowed.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SystemSnmp MibView can be imported using any of these accepted formats:
```
$ terraform import fortios_systemsnmp_mibview.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_systemsnmp_mibview.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
