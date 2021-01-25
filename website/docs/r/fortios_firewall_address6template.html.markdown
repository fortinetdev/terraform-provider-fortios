---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address6template"
description: |-
  Configure IPv6 address templates.
---

# fortios_firewall_address6template
Configure IPv6 address templates.

## Example Usage

```hcl
resource "fortios_firewall_address6template" "trname" {
  ip6                  = "2001:db8:0:b::/64"
  name                 = "1"
  subnet_segment_count = 2

  subnet_segment {
    bits      = 4
    exclusive = "disable"
    id        = 1
    name      = "country"
  }
  subnet_segment {
    bits      = 4
    exclusive = "disable"
    id        = 2
    name      = "state"
  }
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) IPv6 address template name.
* `ip6` - (Required) IPv6 address prefix.
* `subnet_segment_count` - (Required) Number of IPv6 subnet segments.
* `subnet_segment` - IPv6 subnet segments. The structure of `subnet_segment` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `subnet_segment` block supports:

* `id` - Subnet segment ID.
* `name` - Subnet segment name.
* `bits` - Number of bits.
* `exclusive` - Enable/disable exclusive value.
* `values` - Subnet segment values. The structure of `values` block is documented below.

The `values` block supports:

* `name` - Subnet segment value name.
* `value` - Subnet segment value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Address6Template can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_address6template.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
