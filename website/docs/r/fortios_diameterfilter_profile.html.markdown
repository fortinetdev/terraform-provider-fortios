---
subcategory: "FortiGate Diameter-Filter"
layout: "fortios"
page_title: "FortiOS: fortios_diameterfilter_profile"
description: |-
  Configure Diameter filter profiles.
---

# fortios_diameterfilter_profile
Configure Diameter filter profiles. Applies to FortiOS Version `>= 7.4.2`.

## Argument Reference

The following arguments are supported:

* `name` - Profile name.
* `comment` - Comment.
* `monitor_all_messages` - Enable/disable logging for all User Name and Result Code AVP messages. Valid values: `disable`, `enable`.
* `log_packet` - Enable/disable packet log for triggered diameter settings. Valid values: `disable`, `enable`.
* `track_requests_answers` - Enable/disable validation that each answer has a corresponding request. Valid values: `disable`, `enable`.
* `missing_request_action` - Action to be taken for answers without corresponding request. Valid values: `allow`, `block`, `reset`, `monitor`.
* `protocol_version_invalid` - Action to be taken for invalid protocol version. Valid values: `allow`, `block`, `reset`, `monitor`.
* `message_length_invalid` - Action to be taken for invalid message length. Valid values: `allow`, `block`, `reset`, `monitor`.
* `request_error_flag_set` - Action to be taken for request messages with error flag set. Valid values: `allow`, `block`, `reset`, `monitor`.
* `cmd_flags_reserve_set` - Action to be taken for messages with cmd flag reserve bits set. Valid values: `allow`, `block`, `reset`, `monitor`.
* `command_code_invalid` - Action to be taken for messages with invalid command code. Valid values: `allow`, `block`, `reset`, `monitor`.
* `command_code_range` - Valid range for command codes (0-16777215).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

DiameterFilter Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_diameterfilter_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_diameterfilter_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
