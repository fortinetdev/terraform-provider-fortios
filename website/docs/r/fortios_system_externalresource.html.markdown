---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_externalresource"
description: |-
  Configure external resource.
---

# fortios_system_externalresource
Configure external resource.

## Example Usage

```hcl
resource "fortios_system_externalresource" "trname" {
  category     = 199
  name         = "externalresource1"
  refresh_rate = 5
  resource     = "https://tmpxxxxx.com"
  status       = "enable"
  type         = "category"
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) External resource name.
* `status` - Enable/disable user resource.
* `type` - User resource type.
* `category` - User resource category.
* `username` - HTTP basic authentication user name.
* `password` - HTTP basic authentication password.
* `comments` - Comment.
* `resource` - (Required) URI of external resource.
* `refresh_rate` - (Required) Time interval to refresh external resource (1 - 43200 min, default = 5 min).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ExternalResource can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_externalresource.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
