---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_csf"
description: |-
  Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.
---

# fortios_system_csf
Add this FortiGate to a Security Fabric or set up a new Security Fabric on this FortiGate.

## Example Usage

```hcl
resource "fortios_system_csf" "trname" {
  configuration_sync = "default"
  management_ip      = "0.0.0.0"
  management_port    = 33
  status             = "disable"
  upstream_ip        = "0.0.0.0"
  upstream_port      = 8013
  group_password     = "tmp"
}
```

## Argument Reference

The following arguments are supported:

* `status` - (Required) Enable/disable Security Fabric.
* `upstream_ip` - IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_port` - The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
* `group_name` - Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
* `group_password` - Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
* `configuration_sync` - Configuration sync mode.
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `fixed_key` - Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
* `trusted_list` - Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
* `fabric_device` - Fabric device configuration. The structure of `fabric_device` block is documented below.

The `trusted_list` block supports:

* `serial` - Serial.
* `action` - Security fabric authorization action.
* `ha_members` - HA members.
* `downstream_authorization` - Trust authorizations by this node's administrator.

The `fabric_device` block supports:

* `name` - Device name.
* `device_ip` - Device IP.
* `device_type` - Device type.
* `login` - Device login name.
* `password` - Device login password.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Csf can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_csf.labelname SystemCsf
$ unset "FORTIOS_IMPORT_TABLE"
```
