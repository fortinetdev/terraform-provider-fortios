---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_internetservicename"
description: |-
  Define internet service names.
---

# fortios_firewall_internetservicename
Define internet service names. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - Internet Service name.
* `type` - Internet Service name type. Valid values: `default`, `location`.
* `internet_service_id` - Internet Service ID.
* `country_id` - Country or Area ID.
* `region_id` - Region ID.
* `city_id` - City ID.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall InternetServiceName can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_internetservicename.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
