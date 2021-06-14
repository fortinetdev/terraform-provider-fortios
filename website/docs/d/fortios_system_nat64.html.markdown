---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_nat64"
description: |-
  Get information on fortios system nat64.
---

# Data Source: fortios_system_nat64
Use this data source to get information on fortios system nat64

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable NAT64 (default = disable).
* `nat64_prefix` - NAT64 prefix must be ::/96 (default = 64:ff9b::/96).
* `secondary_prefix_status` - Enable/disable secondary NAT64 prefix.
* `secondary_prefix` - Secondary NAT64 prefix. The structure of `secondary_prefix` block is documented below.
* `always_synthesize_aaaa_record` - Enable/disable AAAA record synthesis (default = enable).
* `generate_ipv6_fragment_header` - Enable/disable IPv6 fragment header generation.
* `nat46_force_ipv4_packet_forwarding` - Enable/disable mandatory IPv4 packet forwarding in nat46.

The `secondary_prefix` block contains:

* `name` - NAT64 prefix name.
* `nat64_prefix` - NAT64 prefix.

