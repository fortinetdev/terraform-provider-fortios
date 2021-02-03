---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ips"
description: |-
  Configure IPS system settings.
---

# fortios_system_ips
Configure IPS system settings. Applies to FortiOS Version `>= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `signature_hold_time` - Time to hold and monitor IPS signatures. Format <#d##h> (day range: 0 - 7, hour range: 0 - 23, max hold time: 7d0h, default hold time: 0d0h).

* `override_signature_hold_by_id` - Enable/disable override of hold of triggering signatures that are specified by IDs regardless of hold.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ips can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ips.labelname SystemIps
$ unset "FORTIOS_IMPORT_TABLE"
```
