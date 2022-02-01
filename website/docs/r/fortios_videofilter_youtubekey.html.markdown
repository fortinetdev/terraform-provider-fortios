---
subcategory: "FortiGate Videofilter"
layout: "fortios"
page_title: "FortiOS: fortios_videofilter_youtubekey"
description: |-
  Configure YouTube API keys.
---

# fortios_videofilter_youtubekey
Configure YouTube API keys. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `key` - Key.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Videofilter YoutubeKey can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_videofilter_youtubekey.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
