---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_extenderprofile"
description: |-
  FortiExtender extender profile configuration.
---

# fortios_extendercontroller_extenderprofile
FortiExtender extender profile configuration. Applies to FortiOS Version `>= 7.0.2`.

## Argument Reference

The following arguments are supported:

* `name` - FortiExtender profile name
* `fosid` - id
* `model` - Model. Valid values: `FX201E`, `FX211E`, `FX200F`, `FXA11F`, `FXE11F`, `FXA21F`, `FXE21F`, `FXA22F`, `FXE22F`, `FX212F`, `FX311F`, `FX312F`, `FX511F`, `FVG21F`, `FVA21F`, `FVG22F`, `FVA22F`, `FX04DA`.
* `extension` - Extension option. Valid values: `wan-extension`, `lan-extension`.
* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
* `login_password` - Set the managed extender's administrator password.
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `cellular` - FortiExtender cellular configuration. The structure of `cellular` block is documented below.
* `lan_extension` - FortiExtender lan extension configuration. The structure of `lan_extension` block is documented below.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `cellular` block supports:

* `dataplan` - Dataplan names. The structure of `dataplan` block is documented below.
* `controller_report` - FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
* `sms_notification` - FortiExtender cellular SMS notification configuration. The structure of `sms_notification` block is documented below.
* `modem1` - Configuration options for modem 1. The structure of `modem1` block is documented below.
* `modem2` - Configuration options for modem 2. The structure of `modem2` block is documented below.

The `dataplan` block supports:

* `name` - Dataplan name.

The `controller_report` block supports:

* `status` - FortiExtender controller report status. Valid values: `disable`, `enable`.
* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.

The `sms_notification` block supports:

* `status` - FortiExtender SMS notification status. Valid values: `disable`, `enable`.
* `alert` - SMS alert list. The structure of `alert` block is documented below.
* `receiver` - SMS notification receiver list. The structure of `receiver` block is documented below.

The `alert` block supports:

* `system_reboot` - Display string when system rebooted.
* `data_exhausted` - Display string when data exhausted.
* `session_disconnect` - Display string when session disconnected.
* `low_signal_strength` - Display string when signal strength is low.
* `os_image_fallback` - Display string when falling back to a previous OS image.
* `mode_switch` - Display string when mode is switched.
* `fgt_backup_mode_switch` - Display string when FortiGate backup mode switched.

The `receiver` block supports:

* `name` - FortiExtender SMS notification receiver name.
* `status` - SMS notification receiver status. Valid values: `disable`, `enable`.
* `phone_number` - Receiver phone number.  Format: [+][country code][area code][local phone number].  For example: +16501234567.
* `alert` - Alert multi-options. Valid values: `system-reboot`, `data-exhausted`, `session-disconnect`, `low-signal-strength`, `mode-switch`, `os-image-fallback`, `fgt-backup-mode-switch`.

The `modem1` block supports:

* `redundant_mode` - FortiExtender mode. Valid values: `disable`, `enable`.
* `redundant_intf` - Redundant interface.
* `conn_status` - Connection status.
* `default_sim` - Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.
* `gps` - FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.
* `sim1_pin` - SIM #1 PIN status. Valid values: `disable`, `enable`.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable`, `enable`.
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin_code` - SIM #2 PIN password.
* `preferred_carrier` - Preferred carrier.
* `auto_switch` - FortiExtender auto switch configuration. The structure of `auto_switch` block is documented below.

The `auto_switch` block supports:

* `disconnect` - Auto switch by disconnect. Valid values: `disable`, `enable`.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `disconnect_period` - Automatically switch based on disconnect period.
* `signal` - Automatically switch based on signal strength. Valid values: `disable`, `enable`.
* `dataplan` - Automatically switch based on data usage. Valid values: `disable`, `enable`.
* `switch_back` - Auto switch with switch back multi-options. Valid values: `time`, `timer`.
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).

The `modem2` block supports:

* `redundant_mode` - FortiExtender mode. Valid values: `disable`, `enable`.
* `redundant_intf` - Redundant interface.
* `conn_status` - Connection status.
* `default_sim` - Default SIM selection. Valid values: `sim1`, `sim2`, `carrier`, `cost`.
* `gps` - FortiExtender GPS enable/disable. Valid values: `disable`, `enable`.
* `sim1_pin` - SIM #1 PIN status. Valid values: `disable`, `enable`.
* `sim2_pin` - SIM #2 PIN status. Valid values: `disable`, `enable`.
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin_code` - SIM #2 PIN password.
* `preferred_carrier` - Preferred carrier.
* `auto_switch` - FortiExtender auto switch configuration. The structure of `auto_switch` block is documented below.

The `auto_switch` block supports:

* `disconnect` - Auto switch by disconnect. Valid values: `disable`, `enable`.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `disconnect_period` - Automatically switch based on disconnect period.
* `signal` - Automatically switch based on signal strength. Valid values: `disable`, `enable`.
* `dataplan` - Automatically switch based on data usage. Valid values: `disable`, `enable`.
* `switch_back` - Auto switch with switch back multi-options. Valid values: `time`, `timer`.
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).

The `lan_extension` block supports:

* `link_loadbalance` - LAN extension link load balance strategy. Valid values: `activebackup`, `loadbalance`.
* `ipsec_tunnel` - IPsec tunnel name.
* `backhaul_interface` - IPsec phase1 interface.
* `backhaul_ip` - IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.
* `backhaul` - LAN extension backhaul tunnel configuration. The structure of `backhaul` block is documented below.

The `backhaul` block supports:

* `name` - FortiExtender LAN extension backhaul name
* `port` - FortiExtender uplink port. Valid values: `wan`, `lte1`, `lte2`, `port1`, `port2`, `port3`, `port4`, `port5`, `sfp`.
* `role` - FortiExtender uplink port. Valid values: `primary`, `secondary`.
* `weight` - WRR weight parameter


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtenderController ExtenderProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_extendercontroller_extenderprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_extendercontroller_extenderprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
