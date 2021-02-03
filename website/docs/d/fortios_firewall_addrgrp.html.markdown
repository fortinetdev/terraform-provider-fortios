---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_addrgrp"
description: |-
  Get information on an fortios firewall addrgrp.
---

# Data Source: fortios_firewall_addrgrp
Use this data source to get information on an fortios firewall addrgrp

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall addrgrp.

## Attribute Reference

The following attributes are exported:

* `name` - Address group name.
* `type` - Address group type.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `member` - Address objects contained within the group. The structure of `member` block is documented below.
* `comment` - Comment.
* `exclude` - Enable/disable address exclusion.
* `exclude_member` - Address exclusion member. The structure of `exclude_member` block is documented below.
* `visibility` - Enable/disable address visibility in the GUI.
* `color` - Color of icon on the GUI.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `allow_routing` - Enable/disable use of this group in the static route configuration.

The `member` block contains:

* `name` - Address name.

The `exclude_member` block contains:

* `name` - Address name.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

