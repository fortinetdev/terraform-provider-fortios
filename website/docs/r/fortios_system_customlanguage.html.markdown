---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_customlanguage"
description: |-
  Configure custom languages.
---

# fortios_system_customlanguage
Configure custom languages.

## Example Usage

```hcl
resource "fortios_system_customlanguage" "trname" {
  filename = "en"
  name     = "my-en"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `filename` - (Required) Custom language file path.
* `comments` - Comment.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CustomLanguage can be imported using any of these accepted formats:
```
$ terraform import fortios_system_customlanguage.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_customlanguage.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
