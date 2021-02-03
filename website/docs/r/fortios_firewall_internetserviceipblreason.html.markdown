---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetserviceipblreason"
description: |-
  IP blacklist reason.
---

# fortios_firewall_internetserviceipblreason
IP blacklist reason.

## Argument Reference

The following arguments are supported:

* `fosid` - IP blacklist reason ID.
* `name` - IP blacklist reason name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceIpblReason can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetserviceipblreason.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
