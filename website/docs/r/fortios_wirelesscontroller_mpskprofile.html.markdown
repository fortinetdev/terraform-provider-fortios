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
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `mpsk_group` block supports:

* `name` - MPSK group name.
* `vlan_type` - MPSK group VLAN options.
* `vlan_id` - Optional VLAN ID.
* `mpsk_key` - List of multiple PSK entries. The structure of `mpsk_key` block is documented below.

The `mpsk_key` block supports:

* `name` - Pre-shared key name.
* `mac` - MAC address.
* `passphrase` - WPA Pre-shared key.
* `concurrent_client_limit_type` - MPSK client limit type options.
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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontroller_mpskprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
