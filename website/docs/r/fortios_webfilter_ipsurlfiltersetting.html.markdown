---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_ipsurlfiltersetting"
description: |-
  Configure IPS URL filter settings.
---

# fortios_webfilter_ipsurlfiltersetting
Configure IPS URL filter settings.

## Example Usage

```hcl
resource "fortios_webfilter_ipsurlfiltersetting" "trname" {
  distance = 1
  gateway  = "0.0.0.0"
}
```

## Argument Reference

The following arguments are supported:

* `device` - Interface for this route.
* `distance` - Administrative distance (1 - 255) for this route.
* `gateway` - Gateway IP address for this route.
* `geo_filter` - Filter based on geographical location. Route will NOT be installed if the resolved IP address belongs to the country in the filter.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Webfilter IpsUrlfilterSetting can be imported using any of these accepted formats:
```
$ terraform import fortios_webfilter_ipsurlfiltersetting.labelname WebfilterIpsUrlfilterSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webfilter_ipsurlfiltersetting.labelname WebfilterIpsUrlfilterSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
