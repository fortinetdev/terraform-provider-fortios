---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_nat64"
description: |-
  Configure NAT64.
---

# fortios_system_nat64
Configure NAT64.

## Example Usage

```hcl
resource "fortios_system_nat64" "trname" {
  always_synthesize_aaaa_record      = "enable"
  generate_ipv6_fragment_header      = "disable"
  nat46_force_ipv4_packet_forwarding = "disable"
  nat64_prefix                       = "2001:1:2:3::/96"
  secondary_prefix_status            = "disable"
  status                             = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `status` - Enable/disable NAT64 (default = disable).
* `nat64_prefix` - (Required) NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
* `secondary_prefix_status` - Enable/disable secondary NAT64 prefix.
* `secondary_prefix` - Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
* `always_synthesize_aaaa_record` - Enable/disable AAAA record synthesis (default = enable).
* `generate_ipv6_fragment_header` - Enable/disable IPv6 fragment header generation.
* `nat46_force_ipv4_packet_forwarding` - Enable/disable mandatory IPv4 packet forwarding in nat46.

The `secondary_prefix` block supports:

* `name` - NAT64 prefix name.
* `nat64_prefix` - NAT64 prefix.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Nat64 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_nat64.labelname SystemNat64
$ unset "FORTIOS_IMPORT_TABLE"
```
