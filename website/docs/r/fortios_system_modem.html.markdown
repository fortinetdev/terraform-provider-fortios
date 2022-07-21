---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_modem"
description: |-
  Configure MODEM.
---

# fortios_system_modem
Configure MODEM. Applies to FortiOS Version `7.0.4`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable Modem support (equivalent to bringing an interface up or down). Valid values: `enable`, `disable`.
* `pin_init` - AT command to set the PIN (AT+PIN=<pin>).
* `network_init` - AT command to set the Network name/type (AT+COPS=<mode>,[<format>,<oper>[,<AcT>]]).
* `lockdown_lac` - Allow connection only to the specified Location Area Code (LAC).
* `mode` - Set MODEM operation mode to redundant or standalone. Valid values: `standalone`, `redundant`.
* `auto_dial` - Enable/disable auto-dial after a reboot or disconnection. Valid values: `enable`, `disable`.
* `dial_on_demand` - Enable/disable to dial the modem when packets are routed to the modem interface. Valid values: `enable`, `disable`.
* `idle_timer` - MODEM connection idle time (1 - 9999 min, default = 5).
* `redial` - Redial limit (1 - 10 attempts, none = redial forever). Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
* `reset` - Number of dial attempts before resetting modem (0 = never reset).
* `holddown_timer` - Hold down timer in seconds (1 - 60 sec).
* `connect_timeout` - Connection completion timeout (30 - 255 sec, default = 90).
* `interface` - Name of redundant interface.
* `wireless_port` - Enter wireless port number, 0 for default, 1 for first port, ... (0 - 4294967295, default = 0)
* `dont_send_cr1` - Do not send CR when connected (ISP1). Valid values: `enable`, `disable`.
* `phone1` - Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
* `dial_cmd1` - Dial command (this is often an ATD or ATDT command).
* `username1` - User name to access the specified dialup account.
* `passwd1` - Password to access the specified dialup account.
* `extra_init1` - Extra initialization string to ISP 1.
* `peer_modem1` - Specify peer MODEM type for phone1. Valid values: `generic`, `actiontec`, `ascend_TNT`.
* `ppp_echo_request1` - Enable/disable PPP echo-request to ISP 1. Valid values: `enable`, `disable`.
* `authtype1` - Allowed authentication types for ISP 1. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
* `dont_send_cr2` - Do not send CR when connected (ISP2). Valid values: `enable`, `disable`.
* `phone2` - Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
* `dial_cmd2` - Dial command (this is often an ATD or ATDT command).
* `username2` - User name to access the specified dialup account.
* `passwd2` - Password to access the specified dialup account.
* `extra_init2` - Extra initialization string to ISP 2.
* `peer_modem2` - Specify peer MODEM type for phone2. Valid values: `generic`, `actiontec`, `ascend_TNT`.
* `ppp_echo_request2` - Enable/disable PPP echo-request to ISP 2. Valid values: `enable`, `disable`.
* `authtype2` - Allowed authentication types for ISP 2. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
* `dont_send_cr3` - Do not send CR when connected (ISP3). Valid values: `enable`, `disable`.
* `phone3` - Phone number to connect to the dialup account (must not contain spaces, and should include standard special characters).
* `dial_cmd3` - Dial command (this is often an ATD or ATDT command).
* `username3` - User name to access the specified dialup account.
* `passwd3` - Password to access the specified dialup account.
* `extra_init3` - Extra initialization string to ISP 3.
* `peer_modem3` - Specify peer MODEM type for phone3. Valid values: `generic`, `actiontec`, `ascend_TNT`.
* `ppp_echo_request3` - Enable/disable PPP echo-request to ISP 3. Valid values: `enable`, `disable`.
* `altmode` - Enable/disable altmode for installations using PPP in China. Valid values: `enable`, `disable`.
* `authtype3` - Allowed authentication types for ISP 3. Valid values: `pap`, `chap`, `mschap`, `mschapv2`.
* `traffic_check` - Enable/disable traffic-check. Valid values: `enable`, `disable`.
* `action` - Dial up/stop MODEM. Valid values: `dial`, `stop`, `none`.
* `distance` - Distance of learned routes (1 - 255, default = 1).
* `priority` - Priority of learned routes (0 - 4294967295, default = 0).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Modem can be imported using any of these accepted formats:
```
$ terraform import fortios_system_modem.labelname SystemModem

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_modem.labelname SystemModem
$ unset "FORTIOS_IMPORT_TABLE"
```
