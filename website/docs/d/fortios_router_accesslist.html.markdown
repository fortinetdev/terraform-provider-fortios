---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_accesslist"
description: |-
  Get information on an fortios router accesslist.
---

# Data Source: fortios_router_accesslist
Use this data source to get information on an fortios router accesslist

## Argument Reference

* `name` - (Required) Specify the name of the desired router accesslist.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comments` - Comment.
* `rule` - Rule. The structure of `rule` block is documented below.

The `rule` block contains:

* `id` - Rule ID.
* `action` - Permit or deny this IP address and netmask prefix.
* `prefix` - IPv4 prefix to define regular filter criteria, such as "any" or subnets.
* `wildcard` - Wildcard to define Cisco-style wildcard filter criteria.
* `exact_match` - Enable/disable exact match.
* `flags` - Flags.

