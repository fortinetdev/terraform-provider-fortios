---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_prefixlist"
description: |-
  Get information on an fortios router prefixlist.
---

# Data Source: fortios_router_prefixlist
Use this data source to get information on an fortios router prefixlist

## Argument Reference

* `name` - (Required) Specify the name of the desired router prefixlist.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comments` - Comment.
* `rule` - IPv4 prefix list rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix.
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.
* `ge` - Minimum prefix length to be matched (0 - 32).
* `le` - Maximum prefix length to be matched (0 - 32).
* `flags` - Flags.

