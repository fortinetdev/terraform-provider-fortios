---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_geoipcountry"
description: |-
  Define geoip country name-ID table.
---

# fortios_system_geoipcountry
Define geoip country name-ID table. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `fosid` - Country ID.
* `name` - Country name.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

System GeoipCountry can be imported using any of these accepted formats:
```
$ terraform import fortios_system_geoipcountry.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_geoipcountry.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
