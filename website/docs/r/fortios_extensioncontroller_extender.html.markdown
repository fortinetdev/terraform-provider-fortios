---
subcategory: "FortiGate Extension-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extensioncontroller_extender"
description: |-
  Extender controller configuration.
---

# fortios_extensioncontroller_extender
Extender controller configuration. Applies to FortiOS Version `>= 7.2.1`.

## Argument Reference

The following arguments are supported:

* `name` - FortiExtender entry name.
* `fosid` - FortiExtender serial number.
* `authorized` - FortiExtender Administration (enable or disable).
* `ext_name` - FortiExtender name.
* `description` - Description.
* `vdom` - VDOM.
* `device_id` - Device ID.
* `extension_type` - Extension type for this FortiExtender. Valid values: `wan-extension`, `lan-extension`.
* `profile` - FortiExtender profile configuration.
* `override_allowaccess` - Enable to override the extender profile management access configuration. Valid values: `enable`, `disable`.
* `allowaccess` - Control management access to the managed extender. Separate entries with a space. Valid values: `ping`, `telnet`, `http`, `https`, `ssh`, `snmp`.
* `override_login_password_change` - Enable to override the extender profile login-password (administrator password) setting. Valid values: `enable`, `disable`.
* `login_password_change` - Change or reset the administrator password of a managed extender (yes, default, or no, default = no). Valid values: `yes`, `default`, `no`.
* `login_password` - Set the managed extender's administrator password.
* `override_enforce_bandwidth` - Enable to override the extender profile enforce-bandwidth setting. Valid values: `enable`, `disable`.
* `enforce_bandwidth` - Enable/disable enforcement of bandwidth on LAN extension interface. Valid values: `enable`, `disable`.
* `bandwidth_limit` - FortiExtender LAN extension bandwidth limit (Mbps).
* `wan_extension` - FortiExtender wan extension configuration. The structure of `wan_extension` block is documented below.
* `firmware_provision_latest` - Enable/disable one-time automatic provisioning of the latest firmware version. Valid values: `disable`, `once`.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `wan_extension` block supports:

* `modem1_extension` - FortiExtender interface name.
* `modem2_extension` - FortiExtender interface name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController Extender can be imported using any of these accepted formats:
```
$ terraform import fortios_extensioncontroller_extender.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_extensioncontroller_extender.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
