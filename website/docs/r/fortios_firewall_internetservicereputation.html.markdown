---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicereputation"
description: |-
  Show Internet Service reputation.
---

# fortios_firewall_internetservicereputation
Show Internet Service reputation.

## Argument Reference

The following arguments are supported:

* `fosid` - Internet Service Reputation ID.
* `description` - Description.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceReputation can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetservicereputation.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
