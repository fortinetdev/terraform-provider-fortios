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
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetServiceBotnet can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_internetservicebotnet.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_internetservicebotnet.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
