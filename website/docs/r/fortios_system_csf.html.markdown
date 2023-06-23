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
* `upstream` - IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_ip` - IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_port` - The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
* `group_name` - Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
* `group_password` - Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
* `accept_auth_by_cert` - Accept connections with unknown certificates and ask admin for approval. Valid values: `disable`, `enable`.
* `log_unification` - Enable/disable broadcast of discovery messages for log unification. Valid values: `disable`, `enable`.
* `configuration_sync` - Configuration sync mode. Valid values: `default`, `local`.
* `fabric_object_unification` - Fabric CMDB Object Unification Valid values: `default`, `local`.
* `saml_configuration_sync` - SAML setting configuration synchronization. Valid values: `default`, `local`.
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `authorization_request_type` - Authorization request type. Valid values: `serial`, `certificate`.
* `certificate` - Certificate.
* `fabric_workers` - Number of worker processes for Security Fabric daemon.
* `downstream_access` - Enable/disable downstream device access to this device's configuration and data. Valid values: `enable`, `disable`.
* `downstream_accprofile` - Default access profile for requests from downstream devices.
* `fixed_key` - Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
* `trusted_list` - Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
* `fabric_connector` - Fabric connector configuration. The structure of `fabric_connector` block is documented below.
* `forticloud_account_enforcement` - Fabric FortiCloud account unification. Valid values: `enable`, `disable`.
* `file_mgmt` - Enable/disable Security Fabric daemon file management. Valid values: `enable`, `disable`.
* `file_quota` - Maximum amount of memory that can be used by the daemon files (in bytes).
* `file_quota_warning` - Warn when the set percentage of quota has been used.
* `fabric_device` - Fabric device configuration. The structure of `fabric_device` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `trusted_list` block supports:

* `name` - Name.
* `authorization_type` - Authorization type. Valid values: `serial`, `certificate`.
* `serial` - Serial.
* `certificate` - Certificate.
* `action` - Security fabric authorization action. Valid values: `accept`, `deny`.
* `ha_members` - HA members.
* `downstream_authorization` - Trust authorizations by this node's administrator. Valid values: `enable`, `disable`.
* `index` - Index of the downstream in tree.

The `fabric_connector` block supports:

* `serial` - Serial.
* `accprofile` - Override access profile.
* `configuration_write_access` - Enable/disable downstream device write access to configuration. Valid values: `enable`, `disable`.
* `vdom` - Virtual domains that the connector has access to. If none are set, the connector will only have access to the VDOM that it joins the Security Fabric through. The structure of `vdom` block is documented below.

The `vdom` block supports:

* `name` - Virtual domain name.

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
$ terraform import fortios_system_csf.labelname SystemCsf

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_csf.labelname SystemCsf
$ unset "FORTIOS_IMPORT_TABLE"
```
