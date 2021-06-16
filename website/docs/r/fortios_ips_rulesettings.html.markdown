---
subcategory: "FortiGate IPS"
layout: "fortios"
page_title: "FortiOS: fortios_ips_rulesettings"
description: |-
  Configure IPS rule setting.
---

# fortios_ips_rulesettings
Configure IPS rule setting.

## Argument Reference

The following arguments are supported:

* `fosid` - Rule ID.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Ips RuleSettings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_ips_rulesettings.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
