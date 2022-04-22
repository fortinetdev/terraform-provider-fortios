---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_l2tp"
description: |-
  Configure L2TP.
---

# fortios_vpn_l2tp
Configure L2TP.

## Argument Reference

The following arguments are supported:

* `eip` - End IP.
* `sip` - Start IP.
* `status` - (Required) Enable/disable FortiGate as a L2TP gateway. Valid values: `enable`, `disable`.
* `usrgrp` - User group.
* `enforce_ipsec` - Enable/disable IPsec enforcement. Valid values: `enable`, `disable`.
* `lcp_echo_interval` - Time in seconds between PPPoE Link Control Protocol (LCP) echo requests.
* `lcp_max_echo_fails` - Maximum number of missed LCP echo messages before disconnect.
* `hello_interval` - L2TP hello message interval in seconds (0 - 3600 sec, default = 60).
* `compress` - Enable/disable data compression. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn L2Tp can be imported using any of these accepted formats:
```
$ terraform import fortios_vpn_l2tp.labelname VpnL2Tp

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpn_l2tp.labelname VpnL2Tp
$ unset "FORTIOS_IMPORT_TABLE"
```
