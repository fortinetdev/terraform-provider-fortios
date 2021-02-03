---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_objecttagging"
description: |-
  Get information on an fortios system objecttagging.
---

# Data Source: fortios_system_objecttagging
Use this data source to get information on an fortios system objecttagging

## Argument Reference

* `category` - (Required) Specify the category of the desired system objecttagging.

## Attribute Reference

The following attributes are exported:

* `category` - Tag Category.
* `address` - Address.
* `device` - Device.
* `interface` - Interface.
* `multiple` - Allow multiple tag selection.
* `color` - Color of icon on the GUI.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

