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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter IpsUrlfilterSetting6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webfilter_ipsurlfiltersetting6.labelname WebfilterIpsUrlfilterSetting6
$ unset "FORTIOS_IMPORT_TABLE"
```
