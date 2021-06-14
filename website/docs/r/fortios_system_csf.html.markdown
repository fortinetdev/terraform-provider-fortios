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

* `status` - (Required) Enable/disable Security Fabric. Valid values: `enable`, `disable`.
* `upstream_ip` - IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_port` - The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
* `group_name` - Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
* `group_password` - Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
* `accept_auth_by_cert` - Accept connections with unknown certificates and ask admin for approval. Valid values: `disable`, `enable`.
* `configuration_sync` - Configuration sync mode. Valid values: `default`, `local`.
* `fabric_object_unification` - Fabric CMDB Object Unification Valid values: `default`, `local`.
* `saml_configuration_sync` - SAML setting configuration synchronization. Valid values: `default`, `local`.
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `authorization_request_type` - Authorization request type. Valid values: `serial`, `certificate`.
* `certificate` - Certificate.
* `fixed_key` - Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
* `trusted_list` - Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
* `fabric_device` - Fabric device configuration. The structure of `fabric_device` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `trusted_list` block supports:

* `name` - Name.
* `authorization_type` - Authorization type. Valid values: `serial`, `certificate`.
* `serial` - Serial.
* `certificate` - Certificate.
* `action` - Security fabric authorization action. Valid values: `accept`, `deny`.
* `ha_members` - HA members.
* `downstream_authorization` - Trust authorizations by this node's administrator. Valid values: `enable`, `disable`.

The `fabric_device` block supports:

* `name` - Device name.
* `device_ip` - Device IP.
* `https_port` - HTTPS port for fabric device.
* `access_token` - Device access token.
* `device_type` - Device type. Valid values: `fortimail`.
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
