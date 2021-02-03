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

## Attribute Reference

The following attributes are exported:

* `name` - Name.
* `type` - Destination type.
* `destination` - Destinations. The structure of `destination` block is documented below.
* `ha_group_id` - Cluster group ID set for this destination (default = 0).

The `destination` block contains:

* `name` - Destination.

