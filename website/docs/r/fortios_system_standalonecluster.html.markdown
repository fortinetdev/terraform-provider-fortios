---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_standalonecluster"
description: |-
  Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes.
---

# fortios_system_standalonecluster
Configure FortiGate Session Life Support Protocol (FGSP) cluster attributes. Applies to FortiOS Version `>= 6.4.0`.

## Argument Reference

The following arguments are supported:

* `standalone_group_id` - Cluster group ID (0 - 255). Must be the same for all members.
* `group_member_id` - Cluster member ID (0 - 3).
* `layer2_connection` - Indicate whether layer 2 connections are present among FGSP members. Valid values: `available`, `unavailable`.
* `session_sync_dev` - Offload session-sync process to kernel and sync sessions using connected interface(s) directly.
* `encryption` - Enable/disable encryption when synchronizing sessions. Valid values: `enable`, `disable`.
* `psksecret` - Pre-shared secret for session synchronization (ASCII string or hexadecimal encoded with a leading 0x).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System StandaloneCluster can be imported using any of these accepted formats:
```
$ terraform import fortios_system_standalonecluster.labelname SystemStandaloneCluster

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_standalonecluster.labelname SystemStandaloneCluster
$ unset "FORTIOS_IMPORT_TABLE"
```
