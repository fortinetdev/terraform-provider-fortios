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
* `vlan_identity` - Identity of the VLAN. Commonly used for RADIUS Tunnel-Private-Group-Id. Valid values: `description`, `name`.
* `disable_discovery` - Prevent this FortiSwitch from discovering. The structure of `disable_discovery` block is documented below.
* `mac_retention_period` - Time in hours after which an inactive MAC is removed from client DB.
* `default_virtual_switch_vlan` - Default VLAN for ports when added to the virtual-switch.
* `dhcp_server_access_list` - Enable/disable DHCP snooping server access list. Valid values: `enable`, `disable`.
* `dhcp_option82_format` - DHCP option-82 format string. Valid values: `ascii`, `legacy`.
* `dhcp_option82_circuit_id` - List the parameters to be included to inform about client identification. Valid values: `intfname`, `vlan`, `hostname`, `mode`, `description`.
* `dhcp_option82_remote_id` - List the parameters to be included to inform about client identification. Valid values: `mac`, `hostname`, `ip`.
* `dhcp_snoop_client_req` - Client DHCP packet broadcast mode. Valid values: `drop-untrusted`, `forward-untrusted`.
* `dhcp_snoop_client_db_exp` - Expiry time for DHCP snooping server database entries (300 - 259200 sec, default = 86400 sec).
* `dhcp_snoop_db_per_port_learn_limit` - Per Interface dhcp-server entries learn limit (0 - 1024, default = 64).
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
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
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
$ terraform import fortios_switchcontroller_global.labelname SwitchControllerGlobal

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontroller_global.labelname SwitchControllerGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
