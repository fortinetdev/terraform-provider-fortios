---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservice"
description: |-
  Show Internet Service application.
---

# fortios_firewall_internetservice
Show Internet Service application.

## Argument Reference

The following arguments are supported:

* `fosid` - Internet Service ID.
* `name` - Internet Service name.
* `reputation` - Reputation level of the Internet Service.
* `icon_id` - Icon ID of Internet Service.
* `sld_id` - Second Level Domain.
* `direction` - How this service may be used in a firewall policy (source, destination or both). Valid values: `src`, `dst`, `both`.
* `database` - Database name this Internet Service belongs to. Valid values: `isdb`, `irdb`.
* `ip_range_number` - Total number of IP ranges.
* `extra_ip_range_number` - Extra number of IP ranges.
* `ip_number` - Total number of IP addresses.
* `singularity` - Singular level of the Internet Service.
* `obsolete` - Indicates whether the Internet Service can be used.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall InternetService can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetservice.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
