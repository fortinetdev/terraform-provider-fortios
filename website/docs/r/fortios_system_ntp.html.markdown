---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ntp"
description: |-
  Configure system NTP information.
---

# fortios_system_ntp
Configure system NTP information.

## Example Usage

```hcl
resource "fortios_system_ntp" "trname" {
  ntpsync      = "enable"
  server_mode  = "disable"
  source_ip    = "0.0.0.0"
  source_ip6   = "::"
  syncinterval = 1
  type         = "fortiguard"
}
```

## Argument Reference

The following arguments are supported:

* `ntpsync` - Enable/disable setting the FortiGate system time by synchronizing with an NTP Server. Valid values: `enable`, `disable`.
* `type` - Use the FortiGuard NTP server or any other available NTP Server. Valid values: `fortiguard`, `custom`.
* `syncinterval` - NTP synchronization interval (1 - 1440 min).
* `ntpserver` - Configure the FortiGate to connect to any available third-party NTP server. The structure of `ntpserver` block is documented below.
* `source_ip` - Source IP address for communication to the NTP server.
* `source_ip6` - Source IPv6 address for communication to the NTP server.
* `server_mode` - Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server. Valid values: `enable`, `disable`.
* `authentication` - Enable/disable authentication. Valid values: `enable`, `disable`.
* `key_type` - Key type for authentication (MD5, SHA1). Valid values: `MD5`, `SHA1`.
* `key` - Key for authentication.
* `key_id` - Key ID for authentication.
* `interface` - FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services. The structure of `interface` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ntpserver` block supports:

* `id` - NTP server ID.
* `server` - IP address or hostname of the NTP Server.
* `ntpv3` - Enable to use NTPv3 instead of NTPv4. Valid values: `enable`, `disable`.
* `authentication` - Enable/disable MD5/SHA1 authentication. Valid values: `enable`, `disable`.
* `key` - Key for MD5/SHA1 authentication.
* `key_id` - Key ID for authentication.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.

The `interface` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ntp can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ntp.labelname SystemNtp

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ntp.labelname SystemNtp
$ unset "FORTIOS_IMPORT_TABLE"
```
