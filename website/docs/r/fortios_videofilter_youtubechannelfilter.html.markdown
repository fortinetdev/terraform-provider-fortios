---
subcategory: "FortiGate Videofilter"
layout: "fortios"
page_title: "FortiOS: fortios_videofilter_youtubechannelfilter"
description: |-
  Configure YouTube channel filter.
---

# fortios_videofilter_youtubechannelfilter
Configure YouTube channel filter. Applies to FortiOS Version `7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.4.0,7.4.1`.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `name` - Name.
* `comment` - Comment.
* `default_action` - YouTube channel filter default action. Valid values: `allow`, `monitor`, `block`.
* `entries` - YouTube filter entries. The structure of `entries` block is documented below.
* `override_category` - Enable/disable overriding category filtering result. Valid values: `enable`, `disable`.
* `log` - Eanble/disable logging. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `id` - ID.
* `comment` - Comment.
* `action` - YouTube channel filter action. Valid values: `allow`, `monitor`, `block`.
* `channel_id` - Channel ID.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Videofilter YoutubeChannelFilter can be imported using any of these accepted formats:
```
$ terraform import fortios_videofilter_youtubechannelfilter.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_videofilter_youtubechannelfilter.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
