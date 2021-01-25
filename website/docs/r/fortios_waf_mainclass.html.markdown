---
subcategory: "FortiGate WAF"
layout: "fortios"
page_title: "FortiOS: fortios_waf_mainclass"
description: |-
  Hidden table for datasource.
---

# fortios_waf_mainclass
Hidden table for datasource.

## Argument Reference


The following arguments are supported:

* `name` - Main signature class name.
* `fosid` - Main signature class ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Waf MainClass can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_waf_mainclass.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
