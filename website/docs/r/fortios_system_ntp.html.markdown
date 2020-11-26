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

* `ntpsync` - Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
* `type` - Use the FortiGuard NTP server or any other available NTP Server.
* `syncinterval` - NTP synchronization interval (1 - 1440 min).
* `ntpserver` - Configure the FortiGate to connect to any available third-party NTP server. The structure of `ntpserver` block is documented below.
* `source_ip` - Source IP address for communication to the NTP server.
* `source_ip6` - Source IPv6 address for communication to the NTP server.
* `server_mode` - Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server.
* `interface` - FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services. The structure of `interface` block is documented below.

The `ntpserver` block supports:

* `id` - NTP server ID.
* `server` - IP address or hostname of the NTP Server.
* `ntpv3` - Enable to use NTPv3 instead of NTPv4.
* `authentication` - Enable/disable MD5/SHA1 authentication.
* `key` - Key for MD5/SHA1 authentication.
* `key_id` - Key ID for authentication.

The `interface` block supports:

* `interface_name` - Interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ntp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_ntp.labelname SystemNtp
$ unset "FORTIOS_IMPORT_TABLE"
```
