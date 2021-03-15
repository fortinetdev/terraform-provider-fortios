---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ipv6ehfilter"
description: |-
  Configure IPv6 extension header filter.
---

# fortios_firewall_ipv6ehfilter
Configure IPv6 extension header filter.

## Example Usage

```hcl
resource "fortios_firewall_ipv6ehfilter" "trname" {
  auth     = "disable"
  dest_opt = "disable"
  fragment = "disable"
  hop_opt  = "disable"
  no_next  = "disable"
  routing  = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `hop_opt` - Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable). Valid values: `enable`, `disable`.
* `dest_opt` - Enable/disable blocking packets with Destination Options headers (default = disable). Valid values: `enable`, `disable`.
* `hdopt_type` - Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
* `routing` - Enable/disable blocking packets with Routing headers (default = enable). Valid values: `enable`, `disable`.
* `routing_type` - Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
* `fragment` - Enable/disable blocking packets with the Fragment header (default = disable). Valid values: `enable`, `disable`.
* `auth` - Enable/disable blocking packets with the Authentication header (default = disable). Valid values: `enable`, `disable`.
* `no_next` - Enable/disable blocking packets with the No Next header (default = disable) Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Firewall Ipv6EhFilter can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_ipv6ehfilter.labelname FirewallIpv6EhFilter
$ unset "FORTIOS_IMPORT_TABLE"
```
