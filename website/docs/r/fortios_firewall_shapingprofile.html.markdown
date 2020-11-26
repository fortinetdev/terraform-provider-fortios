---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_shapingprofile"
description: |-
  Configure shaping profiles.
---

# fortios_firewall_shapingprofile
Configure shaping profiles.

## Example Usage

```hcl
resource "fortios_firewall_shapingprofile" "trname" {
  default_class_id = 2
  profile_name     = "shapingprofiles1"

  shaping_entries {
    class_id                        = 2
    guaranteed_bandwidth_percentage = 33
    id                              = 1
    maximum_bandwidth_percentage    = 88
    priority                        = "high"
  }
}
```

## Argument Reference

The following arguments are supported:

* `profile_name` - (Required) Shaping profile name.
* `comment` - Comment.
* `default_class_id` - (Required) Default class ID to handle unclassified packets (including all local traffic).
* `shaping_entries` - Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.

The `shaping_entries` block supports:

* `id` - ID number.
* `class_id` - Class ID.
* `priority` - Priority.
* `guaranteed_bandwidth_percentage` - Guaranteed bandwith in percentage.
* `maximum_bandwidth_percentage` - Maximum bandwith in percentage.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profile_name}}.

## Import

Firewall ShapingProfile can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_shapingprofile.labelname {{profile_name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
