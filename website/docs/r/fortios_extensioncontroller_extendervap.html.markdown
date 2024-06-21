---
subcategory: "FortiGate Extension-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extensioncontroller_extendervap"
description: |-
  FortiExtender wifi vap configuration.
---

# fortios_extensioncontroller_extendervap
FortiExtender wifi vap configuration. Applies to FortiOS Version `>= 7.4.4`.

## Argument Reference

The following arguments are supported:

* `name` - Wi-Fi VAP name.
* `type` - Wi-Fi VAP type local-vap / lan-extension-vap. Valid values: `local-vap`, `lan-ext-vap`.
* `ssid` - Wi-Fi SSID.
* `max_clients` - Wi-Fi max clients (0 - 512), default = 0 (no limit) 
* `broadcast_ssid` - Wi-Fi broadcast SSID enable / disable. Valid values: `disable`, `enable`.
* `security` - Wi-Fi security. Valid values: `OPEN`, `WPA2-Personal`, `WPA-WPA2-Personal`, `WPA3-SAE`, `WPA3-SAE-Transition`, `WPA2-Enterprise`, `WPA3-Enterprise-only`, `WPA3-Enterprise-transition`, `WPA3-Enterprise-192-bit`.
* `dtim` - Wi-Fi DTIM (1 - 255) default = 1.
* `rts_threshold` - Wi-Fi RTS Threshold (256 - 2347), default = 2347 (RTS/CTS disabled).
* `pmf` - Wi-Fi pmf enable/disable, default = disable. Valid values: `disabled`, `optional`, `required`.
* `target_wake_time` - Wi-Fi 802.11AX target wake time enable / disable, default = enable. Valid values: `disable`, `enable`.
* `bss_color_partial` - Wi-Fi 802.11AX bss color partial enable / disable, default = enable. Valid values: `disable`, `enable`.
* `mu_mimo` - Wi-Fi multi-user MIMO enable / disable, default = enable. Valid values: `disable`, `enable`.
* `passphrase` - Wi-Fi passphrase.
* `sae_password` - Wi-Fi SAE Password.
* `auth_server_address` - Wi-Fi Authentication Server Address (IPv4 format).
* `auth_server_port` - Wi-Fi Authentication Server Port.
* `auth_server_secret` - Wi-Fi Authentication Server Secret.
* `ip_address` - Extender ip address.
* `start_ip` - Start ip address.
* `end_ip` - End ip address.
* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController ExtenderVap can be imported using any of these accepted formats:
```
$ terraform import fortios_extensioncontroller_extendervap.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_extensioncontroller_extendervap.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
