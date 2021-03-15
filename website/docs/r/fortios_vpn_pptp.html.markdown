---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpn_pptp"
description: |-
  Configure PPTP.
---

# fortios_vpn_pptp
Configure PPTP.

## Example Usage

```hcl
resource "fortios_vpn_pptp" "trname" {
  eip      = "1.1.1.22"
  ip_mode  = "range"
  local_ip = "0.0.0.0"
  sip      = "1.1.1.1"
  status   = "enable"
  usrgrp   = "Guest-group"
}
```

## Argument Reference

The following arguments are supported:

* `status` - (Required) Enable/disable FortiGate as a PPTP gateway. Valid values: `enable`, `disable`.
* `ip_mode` - IP assignment mode for PPTP client. Valid values: `range`, `usrgrp`.
* `eip` - End IP.
* `sip` - Start IP.
* `local_ip` - Local IP to be used for peer's remote IP.
* `usrgrp` - User group.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Vpn Pptp can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpn_pptp.labelname VpnPptp
$ unset "FORTIOS_IMPORT_TABLE"
```
