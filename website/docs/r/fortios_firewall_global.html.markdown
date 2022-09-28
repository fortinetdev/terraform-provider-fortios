---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_global"
description: |-
  Global firewall settings.
---

# fortios_firewall_global
Global firewall settings. Applies to FortiOS Version `>= 7.2.1`.

## Argument Reference

The following arguments are supported:

* `banned_ip_persistency` - Persistency of banned IPs across power cycling. Valid values: `disabled`, `permanent-only`, `all`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall Global can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_global.labelname FirewallGlobal

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_global.labelname FirewallGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
