---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_zone"
description: |-
  Get information on an fortios system zone.
---

# Data Source: fortios_system_zone
Use this data source to get information on an fortios system zone

## Argument Reference

* `name` - (Required) Specify the name of the desired system zone.

## Attribute Reference

The following attributes are exported:

* `name` - Zone name.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `description` - Description.
* `intrazone` - Allow or deny traffic routing between different interfaces in the same zone (default = deny).
* `interface` - Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

The `interface` block contains:

* `interface_name` - Select interfaces to add to the zone.

