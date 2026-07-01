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
* `max_clients` - Wi-Fi maximum clients (0 - 512, default = 0, which means no limit).
* `broadcast_ssid` - Wi-Fi broadcast SSID enable / disable. Valid values: `disable`, `enable`.
* `security` - Wi-Fi security. Valid values: `OPEN`, `WPA2-Personal`, `WPA-WPA2-Personal`, `WPA3-SAE`, `WPA3-SAE-Transition`, `WPA2-Enterprise`, `WPA3-Enterprise-only`, `WPA3-Enterprise-transition`, `WPA3-Enterprise-192-bit`.
* `dtim` - Wi-Fi DTIM (1 - 255, default = 1).
* `rts_threshold` - Wi-Fi RTS threshold (256 - 2347, default = 2347, which disables RTS/CTS).
* `pmf` - Enable/disable Wi-Fi PMF (default = disable). Valid values: `disabled`, `optional`, `required`.
* `target_wake_time` - Enable/disable Wi-Fi 802.11AX target wake time (default = enable). Valid values: `disable`, `enable`.
* `bss_color_partial` - Enable/disable Wi-Fi 802.11AX BSS color partial (default = enable). Valid values: `disable`, `enable`.
* `mu_mimo` - Enable/disable Wi-Fi multi-user MIMO (default = enable). Valid values: `disable`, `enable`.
* `passphrase` - Wi-Fi passphrase.
* `sae_password` - Wi-Fi SAE Password.
* `auth_server_address` - Wi-Fi Authentication Server Address (IPv4 format).
* `auth_server_port` - Wi-Fi Authentication Server Port.
* `auth_server_secret` - Wi-Fi Authentication Server Secret.
* `ip_address` - Extender ip address.
* `start_ip` - Start ip address.
* `end_ip` - End ip address.
* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
* `security_mode` - Turn on captive portal authentication for this Wi-Fi interface. Valid values: `none`, `captive-portal`.
* `security_external_web` - URL of external authentication web server.
* `security_redirect_url` - Optional URL for redirecting users after they pass captive portal authentication.
* `security_exempt_list` - Name of security exempt list.
* `security_groups` - User groups that can authenticate with the captive portal. The structure of `security_groups` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `security_groups` block supports:

* `name` - Names of user groups that can authenticate with the captive portal.


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
