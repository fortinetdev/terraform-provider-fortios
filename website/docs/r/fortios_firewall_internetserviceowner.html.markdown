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
