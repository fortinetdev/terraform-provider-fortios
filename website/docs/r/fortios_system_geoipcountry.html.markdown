---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_geoipcountry"
description: |-
  Define geoip country name-ID table.
---

# fortios_system_geoipcountry
Define geoip country name-ID table.

## Argument Reference

The following arguments are supported:

* `fosid` - Country ID.
* `name` - Country name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System GeoipCountry can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_geoipcountry.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
