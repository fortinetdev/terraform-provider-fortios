---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_alarm"
description: |-
  Configure alarm.
---

# fortios_system_alarm
Configure alarm.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable alarm. Valid values: `enable`, `disable`.
* `audible` - Enable/disable audible alarm. Valid values: `enable`, `disable`.
* `groups` - Alarm groups. The structure of `groups` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `groups` block supports:

* `id` - Group ID.
* `period` - Time period in seconds (0 = from start up).
* `admin_auth_failure_threshold` - Admin authentication failure threshold.
* `admin_auth_lockout_threshold` - Admin authentication lockout threshold.
* `user_auth_failure_threshold` - User authentication failure threshold.
* `user_auth_lockout_threshold` - User authentication lockout threshold.
* `replay_attempt_threshold` - Replay attempt threshold.
* `self_test_failure_threshold` - Self-test failure threshold.
* `log_full_warning_threshold` - Log full warning threshold.
* `encryption_failure_threshold` - Encryption failure threshold.
* `decryption_failure_threshold` - Decryption failure threshold.
* `fw_policy_violations` - Firewall policy violations. The structure of `fw_policy_violations` block is documented below.
* `fw_policy_id` - Firewall policy ID.
* `fw_policy_id_threshold` - Firewall policy ID threshold.

The `fw_policy_violations` block supports:

* `id` - Firewall policy violations ID.
* `threshold` - Firewall policy violation threshold.
* `src_ip` - Source IP (0=all).
* `dst_ip` - Destination IP (0=all).
* `src_port` - Source port (0=all).
* `dst_port` - Destination port (0=all).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Alarm can be imported using any of these accepted formats:
```
$ terraform import fortios_system_alarm.labelname SystemAlarm

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_alarm.labelname SystemAlarm
$ unset "FORTIOS_IMPORT_TABLE"
```
