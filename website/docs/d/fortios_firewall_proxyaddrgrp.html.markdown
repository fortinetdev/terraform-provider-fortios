---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxyaddrgrp"
description: |-
  Get information on an fortios firewall proxyaddrgrp.
---

# Data Source: fortios_firewall_proxyaddrgrp
Use this data source to get information on an fortios firewall proxyaddrgrp

## Argument Reference

* `name` - (Required) Specify the name of the desired firewall proxyaddrgrp.

## Attribute Reference

The following attributes are exported:

* `name` - Address group name.
* `type` - Source or destination address group type.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `member` - Members of address group. The structure of `member` block is documented below.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `comment` - Optional comments.
* `visibility` - Enable/disable visibility of the object in the GUI.

The `member` block contains:

* `name` - Address name.

The `tagging` block contains:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block contains:

* `name` - Tag name.

