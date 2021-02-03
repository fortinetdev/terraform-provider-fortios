---
subcategory: "FortiGate IPS"
layout: "fortios"
page_title: "FortiOS: fortios_ips_viewmap"
description: |-
  configure ips view-map
---

# fortios_ips_viewmap
configure ips view-map

## Argument Reference

The following arguments are supported:

* `fosid` - View ID.
* `vdom_id` - VDOM ID.
* `policy_id` - Policy ID.
* `id_policy_id` - ID-based policy ID.
* `which` - Policy.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Ips ViewMap can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_ips_viewmap.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
