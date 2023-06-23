---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_extender"
description: |-
  Extender controller configuration.
---

# fortios_extendercontroller_extender
Extender controller configuration.

-> The resource applies to FortiOS Version < 6.4.2. For FortiOS Version >= 6.4.2, see `fortios_extendercontroller_extender1`.


## Example Usage

```hcl
resource "fortios_extendercontroller_extender" "trname" {
  admin               = "disable"
  billing_start_day   = 1
  conn_status         = 0
  dial_mode           = "always-connect"
  dial_status         = 0
  ext_name            = "332"
  fosid               = "1"
  initiated_update    = "disable"
  mode                = "standalone"
  modem_type          = "gsm/lte"
  multi_mode          = "auto"
  ppp_auth_protocol   = "auto"
  ppp_echo_request    = "disable"
  quota_limit_mb      = 0
  redial              = "none"
  roaming             = "disable"
  role                = "primary"
  vdom                = 0
  wimax_auth_protocol = "tls"
}
```

## Argument Reference

The following arguments are supported:

* `name` - FortiExtender entry name.
* `fosid` - (Required) FortiExtender serial number.
* `authorized` - FortiExtender Administration (enable or disable). Valid values: `disable`, `enable`.
* `admin` - (Required) FortiExtender Administration (enable or disable). Valid values: `disable`, `discovered`, `enable`.
* `ifname` - FortiExtender interface name.
* `vdom` - VDOM
* `device_id` - device-id
* `extension_type` - Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
* `override_allowaccess` - Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
* `override_login_password_change` - Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
* `login_password` - FortiExtender login password.
* `override_enforce_bandwidth` - Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `wan_extension` - FortiExtender wan extension configuration. The structure of `wan_extension` block is documented below.
* `profile` - FortiExtender profile configuration.
* `controller_report` - FortiExtender controller report configuration. The structure of `controller_report` block is documented below.
* `modem1` - Configuration options for modem 1. The structure of `modem1` block is documented below.
* `modem2` - Configuration options for modem 2. The structure of `modem2` block is documented below.
* `role` - (Required) FortiExtender work role(Primary, Secondary, None). Valid values: `none`, `primary`, `secondary`.
* `mode` - FortiExtender mode. Valid values: `standalone`, `redundant`.
* `dial_mode` - Dial mode (dial-on-demand or always-connect). Valid values: `dial-on-demand`, `always-connect`.
* `redial` - Number of redials allowed based on failed attempts. Valid values: `none`, `1`, `2`, `3`, `4`, `5`, `6`, `7`, `8`, `9`, `10`.
* `redundant_intf` - Redundant interface.
* `dial_status` - Dial status.
* `conn_status` - Connection status.
* `ext_name` - FortiExtender name.
* `description` - Description.
* `quota_limit_mb` - Monthly quota limit (MB).
* `billing_start_day` - Billing start day.
* `at_dial_script` - Initialization AT commands specific to the MODEM.
* `modem_passwd` - MODEM password.
* `initiated_update` - Allow/disallow network initiated updates to the MODEM. Valid values: `enable`, `disable`.
* `modem_type` - MODEM type (CDMA, GSM/LTE or WIMAX). Valid values: `cdma`, `gsm/lte`, `wimax`.
* `ppp_username` - PPP username.
* `ppp_password` - PPP password.
* `ppp_auth_protocol` - PPP authentication protocol (PAP,CHAP or auto). Valid values: `auto`, `pap`, `chap`.
* `ppp_echo_request` - Enable/disable PPP echo request. Valid values: `enable`, `disable`.
* `wimax_carrier` - WiMax carrier.
* `wimax_realm` - WiMax realm.
* `wimax_auth_protocol` - WiMax authentication protocol(TLS or TTLS). Valid values: `tls`, `ttls`.
* `sim_pin` - SIM PIN.
* `access_point_name` - Access point name(APN).
* `multi_mode` - MODEM mode of operation(3G,LTE,etc). Valid values: `auto`, `auto-3g`, `force-lte`, `force-3g`, `force-2g`.
* `roaming` - Enable/disable MODEM roaming. Valid values: `enable`, `disable`.
* `cdma_nai` - NAI for CDMA MODEMS.
* `aaa_shared_secret` - AAA shared secret.
* `ha_shared_secret` - HA shared secret.
* `primary_ha` - Primary HA.
* `secondary_ha` - Secondary HA.
* `cdma_aaa_spi` - CDMA AAA SPI.
* `cdma_ha_spi` - CDMA HA SPI.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `wan_extension` block supports:

* `modem1_extension` - FortiExtender interface name.
* `modem2_extension` - FortiExtender interface name.

The `controller_report` block supports:

* `status` - FortiExtender controller report status. Valid values: `disable`, `enable`.
* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.

The `modem1` block supports:

* `ifname` - FortiExtender interface name.
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

* `ifname` - FortiExtender interface name.
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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

ExtenderController Extender can be imported using any of these accepted formats:
```
$ terraform import fortios_extendercontroller_extender.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_extendercontroller_extender.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
