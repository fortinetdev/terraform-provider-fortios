---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bfd6"
description: |-
  Get information on fortios router bfd6.
---

# Data Source: fortios_router_bfd6
Use this data source to get information on fortios router bfd6

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `neighbor` - Configure neighbor of IPv6 BFD. The structure of `neighbor` block is documented below.

The `neighbor` block contains:

* `ip6_address` - IPv6 address of the BFD neighbor.
* `interface` - Interface to the BFD neighbor.

