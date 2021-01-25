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
* `status` - (Required) Enable/disable FortiGate as a L2TP gateway.
* `usrgrp` - User group.
* `enforce_ipsec` - Enable/disable IPsec enforcement.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn L2Tp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpn_l2tp.labelname VpnL2Tp
$ unset "FORTIOS_IMPORT_TABLE"
```
