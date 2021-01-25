---
subcategory: "FortiGate Application"
layout: "fortios"
page_title: "FortiOS: fortios_application_custom"
description: |-
  Configure custom application signatures.
---

# fortios_application_custom
Configure custom application signatures.

## Argument Reference


The following arguments are supported:

* `tag` - Signature tag.
* `name` - Name of this custom application signature.
* `fosid` - Custom application category ID (use ? to view available options).
* `comment` - Comment.
* `signature` - The text that makes up the actual custom application signature.
* `category` - (Required) Custom application category ID (use ? to view available options).
* `protocol` - Custom application signature protocol.
* `technology` - Custom application signature technology.
* `behavior` - Custom application signature behavior.
* `vendor` - Custom application signature vendor.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{tag}}.

## Import

Application Custom can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_application_custom.labelname {{tag}}
$ unset "FORTIOS_IMPORT_TABLE"
```
