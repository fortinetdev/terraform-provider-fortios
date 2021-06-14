---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceowner"
description: |-
  Internet Service owner.
---

# fortios_firewall_internetserviceowner
Internet Service owner.

## Argument Reference

The following arguments are supported:

* `fosid` - Internet Service owner ID.
* `name` - Internet Service owner name.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceOwner can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetserviceowner.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
