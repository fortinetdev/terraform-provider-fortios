---
subcategory: "FortiGate AntiVirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_heuristic"
description: |-
  Configure global heuristic options.
---

# fortios_antivirus_heuristic
Configure global heuristic options.

## Example Usage

```hcl
resource "fortios_antivirus_heuristic" "trname" {
  mode = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `mode` - Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Antivirus Heuristic can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_antivirus_heuristic.labelname AntivirusHeuristic
$ unset "FORTIOS_IMPORT_TABLE"
```
