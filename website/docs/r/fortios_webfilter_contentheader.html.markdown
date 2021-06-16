---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_contentheader"
description: |-
  Configure content types used by Web filter.
---

# fortios_webfilter_contentheader
Configure content types used by Web filter.

## Example Usage

```hcl
resource "fortios_webfilter_contentheader" "trname" {
  fosid = 1
  name  = "earth"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - Configure content types used by web filter. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `pattern` - Content type (regular expression).
* `action` - Action to take for this content type. Valid values: `block`, `allow`, `exempt`.
* `category` - Categories that this content type applies to.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter ContentHeader can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webfilter_contentheader.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
