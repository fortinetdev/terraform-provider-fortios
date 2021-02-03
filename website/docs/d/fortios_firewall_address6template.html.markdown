---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_address6template"
description: |-
  Get information on an fortios firewall address6template.
---

# Data Source: fortios_firewall_address6template
Use this data source to get information on an fortios firewall address6template

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall address6template.

## Attribute Reference

The following attributes are exported:

* `name` - IPv6 address template name.
* `ip6` - IPv6 address prefix.
* `subnet_segment_count` - Number of IPv6 subnet segments.
* `subnet_segment` - IPv6 subnet segments. The structure of `subnet_segment` block is documented below.

The `subnet_segment` block contains:

* `id` - Subnet segment ID.
* `name` - Subnet segment name.
* `bits` - Number of bits.
* `exclusive` - Enable/disable exclusive value.
* `values` - Subnet segment values. The structure of `values` block is documented below.

The `values` block contains:

* `name` - Subnet segment value name.
* `value` - Subnet segment value.

