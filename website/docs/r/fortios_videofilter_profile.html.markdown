---
subcategory: "FortiGate Videofilter"
layout: "fortios"
page_title: "FortiOS: fortios_videofilter_profile"
description: |-
  Configure VideoFilter profile.
---

# fortios_videofilter_profile
Configure VideoFilter profile. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `comment` - Comment.
* `youtube_channel_filter` - Set YouTube channel filter.
* `fortiguard_category` - Configure FortiGuard categories. The structure of `fortiguard_category` block is documented below.
* `youtube` - Enable/disable YouTube video source. Valid values: `enable`, `disable`.
* `vimeo` - Enable/disable Vimeo video source. Valid values: `enable`, `disable`.
* `dailymotion` - Enable/disable Dailymotion video source. Valid values: `enable`, `disable`.
* `replacemsg_group` - Replacement message group.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `fortiguard_category` block supports:

* `filters` - Configure VideoFilter FortiGuard category. The structure of `filters` block is documented below.

The `filters` block supports:

* `id` - ID.
* `action` - VideoFilter action. Valid values: `allow`, `monitor`, `block`.
* `category_id` - Category ID.
* `log` - Enable/disable logging. Valid values: `enable`, `disable`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Videofilter Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_videofilter_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_videofilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
