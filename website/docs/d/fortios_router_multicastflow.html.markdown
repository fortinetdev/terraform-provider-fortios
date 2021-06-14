---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_multicastflow"
description: |-
  Get information on an fortios router multicastflow.
---

# Data Source: fortios_router_multicastflow
Use this data source to get information on an fortios router multicastflow

## Argument Reference

* `name` - (Required) Specify the name of the desired router multicastflow.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `comments` - Comment.
* `flows` - Multicast-flow entries. The structure of `flows` block is documented below.

The `flows` block contains:

* `id` - Flow ID.
* `group_addr` - Multicast group IP address.
* `source_addr` - Multicast source IP address.

