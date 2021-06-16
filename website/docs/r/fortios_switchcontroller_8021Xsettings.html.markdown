---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_8021Xsettings"
description: |-
  Configure global 802.1X settings.
---

# fortios_switchcontroller_8021Xsettings
Configure global 802.1X settings.

## Example Usage

```hcl
resource "fortios_switchcontroller_8021Xsettings" "trname" {
  link_down_auth     = "set-unauth"
  max_reauth_attempt = 3
  reauth_period      = 12
}
```

## Argument Reference

The following arguments are supported:

* `link_down_auth` - Interface-reauthentication state to set if a link is down. Valid values: `set-unauth`, `no-action`.
* `reauth_period` - Period of time to allow for reauthentication (1 - 1440 sec, default = 60, 0 = disable reauthentication).
* `max_reauth_attempt` - Maximum number of authentication attempts (0 - 15, default = 3).
* `tx_period` - 802.1X Tx period (seconds, default=30).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController 8021XSettings can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_8021Xsettings.labelname SwitchController8021XSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
