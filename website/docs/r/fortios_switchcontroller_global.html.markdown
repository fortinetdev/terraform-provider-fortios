---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontroller_global"
description: |-
  Configure FortiSwitch global settings.
---

# fortios_switchcontroller_global
Configure FortiSwitch global settings.

## Example Usage

```hcl
resource "fortios_switchcontroller_global" "trname" {
  allow_multiple_interfaces = "disable"
  https_image_push          = "disable"
  log_mac_limit_violations  = "disable"
  mac_aging_interval        = 332
  mac_retention_period      = 24
  mac_violation_timer       = 0
}
```

## Argument Reference

The following arguments are supported:

* `mac_aging_interval` - Time after which an inactive MAC is aged out (10 - 1000000 sec, default = 300, 0 = disable).
* `allow_multiple_interfaces` - Enable/disable multiple FortiLink interfaces for redundant connections between a managed FortiSwitch and FortiGate. Valid values: `enable`, `disable`.
* `https_image_push` - Enable/disable image push to FortiSwitch using HTTPS. Valid values: `enable`, `disable`.
* `vlan_all_mode` - VLAN configuration mode, user-defined-vlans or all-possible-vlans. Valid values: `all`, `defined`.
* `vlan_optimization` - FortiLink VLAN optimization. Valid values: `enable`, `disable`.
* `disable_discovery` - Prevent this FortiSwitch from discovering. The structure of `disable_discovery` block is documented below.
* `mac_retention_period` - Time in hours after which an inactive MAC is removed from client DB.
* `default_virtual_switch_vlan` - Default VLAN for ports when added to the virtual-switch.
* `dhcp_server_access_list` - Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
* `log_mac_limit_violations` - Enable/disable logs for Learning Limit Violations. Valid values: `enable`, `disable`.
* `mac_violation_timer` - Set timeout for Learning Limit Violations (0 = disabled).
* `sn_dns_resolution` - Enable/disable DNS resolution of the FortiSwitch unit's IP address by use of its serial number. Valid values: `enable`, `disable`.
* `mac_event_logging` - Enable/disable MAC address event logging. Valid values: `enable`, `disable`.
* `bounce_quarantined_link` - Enable/disable bouncing (administratively bring the link down, up) of a switch port where a quarantined device was seen last. Helps to re-initiate the DHCP process for a device. Valid values: `disable`, `enable`.
* `quarantine_mode` - Quarantine mode. Valid values: `by-vlan`, `by-redirect`.
* `update_user_device` - Control which sources update the device user list. Valid values: `mac-cache`, `lldp`, `dhcp-snooping`, `l2-db`, `l3-db`.
* `custom_command` - List of custom commands to be pushed to all FortiSwitches in the VDOM. The structure of `custom_command` block is documented below.
* `fips_enforce` - Enable/disable enforcement of FIPS on managed FortiSwitch devices. Valid values: `disable`, `enable`.
* `firmware_provision_on_authorization` - Enable/disable automatic provisioning of latest firmware on authorization. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `disable_discovery` block supports:

* `name` - Managed device ID.

The `custom_command` block supports:

* `command_entry` - List of FortiSwitch commands.
* `command_name` - Name of custom command to push to all FortiSwitches in VDOM.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SwitchController Global can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_switchcontroller_global.labelname SwitchControllerGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
