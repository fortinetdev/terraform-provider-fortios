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

* `name` - Name.
* `status` - (Required) Enable/disable this stitch.
* `trigger` - (Required) Trigger name.
* `action` - Action names.
* `destination` - Serial number/HA group-name of destination devices.

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
