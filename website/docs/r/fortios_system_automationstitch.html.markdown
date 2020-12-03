---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_automationstitch"
description: |-
  Automation stitches.
---

# fortios_system_automationstitch
Automation stitches.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `status` - (Required) Enable/disable this stitch.
* `trigger` - (Required) Trigger name.
* `action` - Action names. The structure of `action` block is documented below.
* `destination` - Serial number/HA group-name of destination devices. The structure of `destination` block is documented below.

The `action` block supports:

* `name` - Action name.

The `destination` block supports:

* `name` - Destination name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System AutomationStitch can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_automationstitch.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
