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
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


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

