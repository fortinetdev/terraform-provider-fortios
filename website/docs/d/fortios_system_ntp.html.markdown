---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ntp"
description: |-
  Get information on fortios system ntp.
---

# Data Source: fortios_system_ntp
Use this data source to get information on fortios system ntp

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `ntpsync` - Enable/disable setting the FortiGate system time by synchronizing with an NTP Server.
* `type` - Use the FortiGuard NTP server or any other available NTP Server.
* `syncinterval` - NTP synchronization interval (1 - 1440 min).
* `ntpserver` - Configure the FortiGate to connect to any available third-party NTP server. The structure of `ntpserver` block is documented below.
* `source_ip` - Source IP address for communication to the NTP server.
* `source_ip6` - Source IPv6 address for communication to the NTP server.
* `server_mode` - Enable/disable FortiGate NTP Server Mode. Your FortiGate becomes an NTP server for other devices on your network. The FortiGate relays NTP requests to its configured NTP server.
* `authentication` - Enable/disable authentication.
* `key_type` - Key type for authentication (MD5, SHA1).
* `key` - Key for authentication.
* `key_id` - Key ID for authentication.
* `interface` - FortiGate interface(s) with NTP server mode enabled. Devices on your network can contact these interfaces for NTP services. The structure of `interface` block is documented below.

The `ntpserver` block contains:

* `id` - NTP server ID.
* `server` - IP address or hostname of the NTP Server.
* `ntpv3` - Enable to use NTPv3 instead of NTPv4.
* `authentication` - Enable/disable MD5/SHA1 authentication.
* `key_type` - Select NTP authentication type.
* `key` - Key for MD5/SHA1 authentication.
* `key_id` - Key ID for authentication.
* `ip_type` - Choose to connect to IPv4 or/and IPv6 NTP server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.

The `interface` block contains:

* `interface_name` - Interface name.

