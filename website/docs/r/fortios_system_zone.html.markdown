---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_zone"
description: |-
  Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.
---

# fortios_system_zone
Configure zones to group two or more interfaces. When a zone is created you can configure policies for the zone instead of individual interfaces in the zone.

## Example Usage

```hcl
resource "fortios_system_zone" "trname" {
  intrazone = "allow"
  name      = "zone1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Zone name.
* `tagging` - Config object tagging. The structure of `tagging` block is documented below.
* `intrazone` - Allow or deny traffic routing between different interfaces in the same zone (default = deny).
* `interface` - Add interfaces to this zone. Interfaces must not be assigned to another zone or have firewall policies defined. The structure of `interface` block is documented below.

The `tagging` block supports:

* `name` - Tagging entry name.
* `category` - Tag category.
* `tags` - Tags. The structure of `tags` block is documented below.

The `tags` block supports:

* `name` - Tag name.

The `interface` block supports:

* `interface_name` - Select interfaces to add to the zone.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System Zone can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_zone.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
