---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ipsurlfiltersetting6"
description: |-
  Configure IPS URL filter settings for IPv6.
---

# fortios_webfilter_ipsurlfiltersetting6
Configure IPS URL filter settings for IPv6.

## Example Usage

```hcl
resource "fortios_webfilter_ipsurlfiltersetting6" "trname" {
  distance = 1
  gateway6 = "::"
}
```

## Argument Reference

The following arguments are supported:

* `device` - Interface for this route.
* `distance` - Administrative distance (1 - 255) for this route.
* `gateway6` - Gateway IPv6 address for this route.
* `geo_filter` - Filter based on geographical location. Route will NOT be installed if the resolved IPv6 address belongs to the country in the filter.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter IpsUrlfilterSetting6 can be imported using any of these accepted formats:
```
$ terraform import fortios_webfilter_ipsurlfiltersetting6.labelname WebfilterIpsUrlfilterSetting6

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webfilter_ipsurlfiltersetting6.labelname WebfilterIpsUrlfilterSetting6
$ unset "FORTIOS_IMPORT_TABLE"
```
