---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomproperty"
description: |-
  Configure VDOM property.
---

# fortios_system_vdomproperty
Configure VDOM property.

## Argument Reference

The following arguments are supported:

* `name` - (Required) VDOM name.
* `description` - Description.
* `snmp_index` - Permanent SNMP Index of the virtual domain (0 - 4294967295).
* `session` - Maximum guaranteed number of sessions.
* `ipsec_phase1` - Maximum guaranteed number of VPN IPsec phase 1 tunnels.
* `ipsec_phase2` - Maximum guaranteed number of VPN IPsec phase 2 tunnels.
* `ipsec_phase1_interface` - Maximum guaranteed number of VPN IPsec phase1 interface tunnels.
* `ipsec_phase2_interface` - Maximum guaranteed number of VPN IPsec phase2 interface tunnels.
* `dialup_tunnel` - Maximum guaranteed number of dial-up tunnels.
* `firewall_policy` - Maximum guaranteed number of firewall policies (IPv4, IPv6, policy46, policy64, DoS-policy4, DoS-policy6, multicast).
* `firewall_address` - Maximum guaranteed number of firewall addresses (IPv4, IPv6, multicast).
* `firewall_addrgrp` - Maximum guaranteed number of firewall address groups (IPv4, IPv6).
* `custom_service` - Maximum guaranteed number of firewall custom services.
* `service_group` - Maximum guaranteed number of firewall service groups.
* `onetime_schedule` - Maximum guaranteed number of firewall one-time schedules.
* `recurring_schedule` - Maximum guaranteed number of firewall recurring schedules.
* `user` - Maximum guaranteed number of local users.
* `user_group` - Maximum guaranteed number of user groups.
* `sslvpn` - Maximum guaranteed number of SSL-VPNs.
* `proxy` - Maximum guaranteed number of concurrent proxy users.
* `log_disk_quota` - Log disk quota in MB (range depends on how much disk space is available).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VdomProperty can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_vdomproperty.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
