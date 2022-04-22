---
subcategory: "FortiGate IPS"
layout: "fortios"
page_title: "FortiOS: fortios_ips_viewmap"
description: |-
  configure ips view-map
---

# fortios_ips_viewmap
configure ips view-map Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `fosid` - View ID.
* `vdom_id` - VDOM ID.
* `policy_id` - Policy ID.
* `id_policy_id` - ID-based policy ID.
* `which` - Policy.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Ips ViewMap can be imported using any of these accepted formats:
```
$ terraform import fortios_ips_viewmap.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ips_viewmap.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
