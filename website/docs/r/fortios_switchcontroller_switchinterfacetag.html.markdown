---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_switchinterfacetag"
description: |-
  Configure switch object tags.
---

# fortios_switchcontroller_switchinterfacetag
Configure switch object tags.

## Example Usage

```hcl
resource "fortios_switchcontroller_switchinterfacetag" "trname" {
  name = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Tag name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchController SwitchInterfaceTag can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_switchinterfacetag.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
