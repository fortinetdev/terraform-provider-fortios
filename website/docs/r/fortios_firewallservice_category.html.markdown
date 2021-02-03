---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_category"
description: |-
  Configure service categories.
---

# fortios_firewallservice_category
Configure service categories.

## Example Usage

```hcl
resource "fortios_firewallservice_category" "trname" {
  name = "s1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Service category name.
* `comment` - Comment.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallService Category can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallservice_category.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
