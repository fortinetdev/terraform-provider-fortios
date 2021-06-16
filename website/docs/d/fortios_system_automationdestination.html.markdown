---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationdestination"
description: |-
  Get information on an fortios system automationdestination.
---

# Data Source: fortios_system_automationdestination
Use this data source to get information on an fortios system automationdestination

## Argument Reference

* `name` - (Required) Specify the name of the desired system automationdestination.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `type` - Destination type.
* `destination` - Destinations. The structure of `destination` block is documented below.
* `ha_group_id` - Cluster group ID set for this destination (default = 0).

The `destination` block contains:

* `name` - Destination.

