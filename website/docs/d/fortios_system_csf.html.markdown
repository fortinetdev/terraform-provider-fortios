---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_csf"
description: |-
  Get information on fortios system csf.
---

# Data Source: fortios_system_csf
Use this data source to get information on fortios system csf

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable Security Fabric.
* `upstream_ip` - IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_port` - The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
* `group_name` - Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
* `group_password` - Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
* `accept_auth_by_cert` - Accept connections with unknown certificates and ask admin for approval.
* `configuration_sync` - Configuration sync mode.
* `fabric_object_unification` - Fabric CMDB Object Unification
* `saml_configuration_sync` - SAML setting configuration synchronization.
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `authorization_request_type` - Authorization request type.
* `certificate` - Certificate.
* `fixed_key` - Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
* `trusted_list` - Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
* `fabric_device` - Fabric device configuration. The structure of `fabric_device` block is documented below.

The `trusted_list` block contains:

* `name` - Name.
* `authorization_type` - Authorization type.
* `serial` - Serial.
* `certificate` - Certificate.
* `action` - Security fabric authorization action.
* `ha_members` - HA members.
* `downstream_authorization` - Trust authorizations by this node's administrator.

The `fabric_device` block contains:

* `name` - Device name.
* `device_ip` - Device IP.
* `https_port` - HTTPS port for fabric device.
* `access_token` - Device access token.
* `device_type` - Device type.
* `login` - Device login name.
* `password` - Device login password.

