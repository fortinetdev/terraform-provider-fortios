---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_bfd"
description: |-
  Get information on fortios router bfd.
---

# Data Source: fortios_router_bfd
Use this data source to get information on fortios router bfd

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `neighbor` - neighbor The structure of `neighbor` block is documented below.
* `multihop_template` - BFD multi-hop template table. The structure of `multihop_template` block is documented below.

The `neighbor` block contains:

* `ip` - IPv4 address of the BFD neighbor.
* `interface` - Interface name.

The `multihop_template` block contains:

* `id` - ID.
* `src` - Source prefix.
* `dst` - Destination prefix.
* `bfd_desired_min_tx` - BFD desired minimal transmit interval (milliseconds).
* `bfd_required_min_rx` - BFD required minimal receive interval (milliseconds).
* `bfd_detect_mult` - BFD detection multiplier.
* `auth_mode` - Authentication mode.
* `md5_key` - MD5 key of key ID 1.

