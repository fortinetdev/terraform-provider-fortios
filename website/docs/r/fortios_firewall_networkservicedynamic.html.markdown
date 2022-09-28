---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_networkservicedynamic"
description: |-
  Configure Dynamic Network Services.
---

# fortios_firewall_networkservicedynamic
Configure Dynamic Network Services. Applies to FortiOS Version `>= 7.2.1`.

## Argument Reference

The following arguments are supported:

* `name` - Dynamic Network Service name.
* `sdn` - SDN connector name.
* `comment` - Comment.
* `filter` - Match criteria filter.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall NetworkServiceDynamic can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_networkservicedynamic.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_networkservicedynamic.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
