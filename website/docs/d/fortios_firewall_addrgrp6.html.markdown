---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_addrgrp6"
description: |-
  Get information on an fortios firewall addrgrp6.
---

# Data Source: fortios_firewall_addrgrp6
Use this data source to get information on an fortios firewall addrgrp6

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall addrgrp6.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - IPv6 address group name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `visibility` - Enable/disable address group6 visibility in the GUI.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets the value to 1).
* `comment` - Comment.
* `member` - Address objects contained within the group. The structure of `member` block is documented below.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.

The `member` block contains:

* `name` - Address6/addrgrp6 name.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

