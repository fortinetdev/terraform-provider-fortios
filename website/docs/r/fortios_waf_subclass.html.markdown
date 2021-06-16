---
subcategory: "FortiGate WAF"
layout: "fortios"
page_title: "FortiOS: fortios_waf_subclass"
description: |-
  Hidden table for datasource.
---

# fortios_waf_subclass
Hidden table for datasource.

## Argument Reference

The following arguments are supported:

* `name` - Signature subclass name.
* `fosid` - Signature subclass ID.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Waf SubClass can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_waf_subclass.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
