---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceipblvendor"
description: |-
  IP blacklist vendor.
---

# fortios_firewall_internetserviceipblvendor
IP blacklist vendor.

## Argument Reference

The following arguments are supported:

* `fosid` - IP blacklist vendor ID.
* `name` - IP blacklist vendor name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceIpblVendor can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetserviceipblvendor.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
