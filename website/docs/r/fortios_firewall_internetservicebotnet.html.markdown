---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicebotnet"
description: |-
  Show Internet Service botnet.
---

# fortios_firewall_internetservicebotnet
Show Internet Service botnet. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - Internet Service Botnet ID.
* `name` - Internet Service Botnet name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceBotnet can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetservicebotnet.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
