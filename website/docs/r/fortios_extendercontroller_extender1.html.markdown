---
subcategory: "FortiGate Extender-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extendercontroller_extender1"
description: |-
  Extender controller configuration.(For FortiOS 6.4.2 and 6.4.2+)
---

# fortios_extendercontroller_extender1
Extender controller configuration.(For FortiOS 6.4.2 and 6.4.2+)

## Argument Reference

The following arguments are supported:

* `name` - (Required) FortiExtender entry name.
* `fosid` - FortiExtender serial number.
* `authorized` - (Required) FortiExtender Administration (enable or disable).
* `ext_name` - FortiExtender name.
* `description` - Description.
* `vdom` - VDOM
* `login_password` - FortiExtender login password.
* `controller_report` - FortiExtender controller report configuration.
* `modem1` - Configuration options for modem 1.
* `modem2` - Configuration options for modem 2.

The `controller_report` block supports:

* `status` - FortiExtender controller report status.
* `interval` - Controller report interval.
* `signal_threshold` - Controller report signal threshold.

The `modem1` block supports:

* `ifname` - FortiExtender interface name.
* `redundant_mode` - FortiExtender mode.
* `redundant_intf` - Redundant interface.
* `conn_status` - Connection status.
* `default_sim` - Default SIM selection.
* `gps` - FortiExtender GPS enable/disable.
* `sim1_pin` - SIM #1 PIN status.
* `sim2_pin` - SIM #2 PIN status.
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin_code` - SIM #2 PIN password.
* `preferred_carrier` - Preferred carrier.
* `auto_switch` - FortiExtender auto switch configuration.

The `auto_switch` block supports:

* `disconnect` - Auto switch by disconnect.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `disconnect_period` - Automatically switch based on disconnect period.
* `signal` - Automatically switch based on signal strength.
* `dataplan` - Automatically switch based on data usage.
* `switch_back` - Auto switch with switch back multi-options.
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).

The `modem2` block supports:

* `ifname` - FortiExtender interface name.
* `redundant_mode` - FortiExtender mode.
* `redundant_intf` - Redundant interface.
* `conn_status` - Connection status.
* `default_sim` - Default SIM selection.
* `gps` - FortiExtender GPS enable/disable.
* `sim1_pin` - SIM #1 PIN status.
* `sim2_pin` - SIM #2 PIN status.
* `sim1_pin_code` - SIM #1 PIN password.
* `sim2_pin_code` - SIM #2 PIN password.
* `preferred_carrier` - Preferred carrier.
* `auto_switch` - FortiExtender auto switch configuration.

The `auto_switch` block supports:

* `disconnect` - Auto switch by disconnect.
* `disconnect_threshold` - Automatically switch based on disconnect threshold.
* `disconnect_period` - Automatically switch based on disconnect period.
* `signal` - Automatically switch based on signal strength.
* `dataplan` - Automatically switch based on data usage.
* `switch_back` - Auto switch with switch back multi-options.
* `switch_back_time` - Automatically switch over to preferred SIM/carrier at a specified time in UTC (HH:MM).
* `switch_back_timer` - Automatically switch over to preferred SIM/carrier after the given time (3600 - 2147483647 sec).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtenderController Extender1 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_extendercontroller_extender1.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
