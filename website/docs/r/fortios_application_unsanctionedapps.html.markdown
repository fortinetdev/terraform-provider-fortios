---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_unsanctionedapps"
description: |-
  Configure unsanctioned applications.
---

# fortios_application_unsanctionedapps
Configure unsanctioned applications. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - Entry ID.
* `type` - Type. Valid values: `app`, `category`.
* `app` - Application ID.
* `category` - Application category.
* `status` - Status. Valid values: `unsanctioned`, `sanctioned`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Application UnsanctionedApps can be imported using any of these accepted formats:
```
$ terraform import fortios_application_unsanctionedapps.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_application_unsanctionedapps.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
