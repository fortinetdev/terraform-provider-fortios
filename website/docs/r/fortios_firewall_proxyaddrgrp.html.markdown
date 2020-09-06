---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_proxyaddrgrp"
description: |-
  Web proxy address group configuration.
---

# fortios_firewall_proxyaddrgrp
Web proxy address group configuration.

## Argument Reference

The following arguments are supported:

* `name` - Address group name.
* `type` - Source or destination address group type.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `member` - (Required) Members of address group.
* `color` - Integer value to determine the color of the icon in the GUI (1 - 32, default = 0, which sets value to 1).
* `tagging` - Config object tagging.
* `comment` - Optional comments.
* `visibility` - Enable/disable visibility of the object in the GUI.

The `member` block supports:

* `name` - Address name.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags.

The `tags` block supports:

* `name` - Tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall ProxyAddrgrp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_proxyaddrgrp.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
