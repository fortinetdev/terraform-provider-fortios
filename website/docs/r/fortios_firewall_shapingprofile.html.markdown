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
* `type` - Select shaping profile type: policing / queuing. Valid values: `policing`, `queuing`.
* `default_class_id` - (Required) Default class ID to handle unclassified packets (including all local traffic).
* `shaping_entries` - Define shaping entries of this shaping profile. The structure of `shaping_entries` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `shaping_entries` block supports:

* `id` - ID number.
* `class_id` - Class ID.
* `priority` - Priority.
* `guaranteed_bandwidth_percentage` - Guaranteed bandwith in percentage.
* `maximum_bandwidth_percentage` - Maximum bandwith in percentage.
* `limit` - Hard limit on the real queue size in packets.
* `burst_in_msec` - Number of bytes that can be burst at maximum-bandwidth speed. Formula: burst = maximum-bandwidth*burst-in-msec.
* `cburst_in_msec` - Number of bytes that can be burst as fast as the interface can transmit. Formula: cburst = maximum-bandwidth*cburst-in-msec.
* `red_probability` - Maximum probability (in percentage) for RED marking.
* `min` - Average queue size in packets at which RED drop becomes a possibility.
* `max` - Average queue size in packets at which RED drop probability is maximal.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{profile_name}}.

## Import

Firewall ShapingProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_shapingprofile.labelname {{profile_name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_shapingprofile.labelname {{profile_name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
