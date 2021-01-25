---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_resourcelimits"
description: |-
  Configure resource limits.
---

# fortios_system_resourcelimits
Configure resource limits.

## Example Usage

```hcl
resource "fortios_system_resourcelimits" "trname" {
  custom_service         = 0
  dialup_tunnel          = 0
  firewall_address       = 41024
  firewall_addrgrp       = 10692
  firewall_policy        = 41024
  ipsec_phase1           = 2000
  ipsec_phase1_interface = 0
  ipsec_phase2           = 2000
  ipsec_phase2_interface = 0
  log_disk_quota         = 30235
  onetime_schedule       = 0
  proxy                  = 64000
  recurring_schedule     = 0
  service_group          = 0
  session                = 0
  sslvpn                 = 0
  user                   = 0
  user_group             = 0
}
```

## Argument Reference


The following arguments are supported:

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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System ResourceLimits can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_resourcelimits.labelname SystemResourceLimits
$ unset "FORTIOS_IMPORT_TABLE"
```
