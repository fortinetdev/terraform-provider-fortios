---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_virtualwirepair"
description: |-
  Configure virtual wire pairs.
---

# fortios_system_virtualwirepair
Configure virtual wire pairs.

## Argument Reference

The following arguments are supported:

* `name` - Virtual-wire-pair name. Must be a unique interface name.
* `member` - (Required) Interfaces belong to the virtual-wire-pair. The structure of `member` block is documented below.
* `wildcard_vlan` - Enable/disable wildcard VLAN.
* `vlan_filter` - Set VLAN filters.

The `member` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VirtualWirePair can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_virtualwirepair.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
