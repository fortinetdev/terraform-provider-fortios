---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontroller_mpskprofile"
description: |-
  Configure MPSK profile.
---

# fortios_wirelesscontroller_mpskprofile
Configure MPSK profile. Applies to FortiOS Version `>= 6.4.2`.

## Argument Reference

The following arguments are supported:

* `name` - MPSK profile name.
* `mpsk_concurrent_clients` - Maximum number of concurrent clients that connect using the same passphrase in multiple PSK authentication (0 - 65535, default = 0, meaning no limitation).
* `mpsk_group` - List of multiple PSK groups. The structure of `mpsk_group` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `mpsk_group` block supports:

* `name` - MPSK group name.
* `vlan_type` - MPSK group VLAN options. Valid values: `no-vlan`, `fixed-vlan`.
* `vlan_id` - Optional VLAN ID.
* `mpsk_key` - List of multiple PSK entries. The structure of `mpsk_key` block is documented below.

The `mpsk_key` block supports:

* `name` - Pre-shared key name.
* `mac` - MAC address.
* `passphrase` - WPA Pre-shared key.
* `concurrent_client_limit_type` - MPSK client limit type options. Valid values: `default`, `unlimited`, `specified`.
* `concurrent_clients` - Number of clients that can connect using this pre-shared key (1 - 65535, default is 256).
* `comment` - Comment.
* `mpsk_schedules` - Firewall schedule for MPSK passphrase. The passphrase will be effective only when at least one schedule is valid. The structure of `mpsk_schedules` block is documented below.

The `mpsk_schedules` block supports:

* `name` - Schedule name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessController MpskProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontroller_mpskprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontroller_mpskprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
