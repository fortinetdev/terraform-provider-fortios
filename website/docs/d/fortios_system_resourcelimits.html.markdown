---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_resourcelimits"
description: |-
  Get information on fortios system resourcelimits.
---

# Data Source: fortios_system_resourcelimits
Use this data source to get information on fortios system resourcelimits

## Argument Reference

No arguments available for the data source.

## Attribute Reference

The following attributes are exported:

* `session` - Maximum number of sessions.
* `ipsec_phase1` - Maximum number of VPN IPsec phase1 tunnels.
* `ipsec_phase2` - Maximum number of VPN IPsec phase2 tunnels.
* `ipsec_phase1_interface` - Maximum number of VPN IPsec phase1 interface tunnels.
* `ipsec_phase2_interface` - Maximum number of VPN IPsec phase2 interface tunnels.
* `dialup_tunnel` - Maximum number of dial-up tunnels.
* `firewall_policy` - Maximum number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
* `firewall_address` - Maximum number of firewall addresses (IPv4, IPv6, multicast).
* `firewall_addrgrp` - Maximum number of firewall address groups (IPv4, IPv6).
* `custom_service` - Maximum number of firewall custom services.
* `service_group` - Maximum number of firewall service groups.
* `onetime_schedule` - Maximum number of firewall one-time schedules.
* `recurring_schedule` - Maximum number of firewall recurring schedules.
* `user` - Maximum number of local users.
* `user_group` - Maximum number of user groups.
* `sslvpn` - Maximum number of SSL-VPN.
* `proxy` - Maximum number of concurrent proxy users.
* `log_disk_quota` - Log disk quota in MB.

