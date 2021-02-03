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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System CustomLanguage can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_customlanguage.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
