---
subcategory: "FortiGate AntiVirus"
layout: "fortios"
page_title: "FortiOS: fortios_antivirus_heuristic"
description: |-
  Configure global heuristic options.
---

# fortios_antivirus_heuristic
Configure global heuristic options. Applies to FortiOS Version `<= 7.0.0`.

## Example Usage

```hcl
resource "fortios_antivirus_heuristic" "trname" {
  mode = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `mode` - Enable/disable heuristics and determine how the system behaves if heuristics detects a problem. Valid values: `pass`, `block`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Antivirus Heuristic can be imported using any of these accepted formats:
```
$ terraform import fortios_antivirus_heuristic.labelname AntivirusHeuristic

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_antivirus_heuristic.labelname AntivirusHeuristic
$ unset "FORTIOS_IMPORT_TABLE"
```
