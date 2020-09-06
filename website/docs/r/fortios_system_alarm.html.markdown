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

* `status` - Enable/disable alarm.
* `audible` - Enable/disable audible alarm.
* `groups` - Alarm groups.

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
* `fw_policy_violations` - Firewall policy violations.
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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_alarm.labelname SystemAlarm
$ unset "FORTIOS_IMPORT_TABLE"
```
