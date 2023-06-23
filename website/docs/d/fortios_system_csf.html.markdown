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
* `upstream` - IP/FQDN of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_ip` - IP address of the FortiGate upstream from this FortiGate in the Security Fabric.
* `upstream_port` - The port number to use to communicate with the FortiGate upstream from this FortiGate in the Security Fabric (default = 8013).
* `group_name` - Security Fabric group name. All FortiGates in a Security Fabric must have the same group name.
* `group_password` - Security Fabric group password. All FortiGates in a Security Fabric must have the same group password.
* `accept_auth_by_cert` - Accept connections with unknown certificates and ask admin for approval.
* `log_unification` - Enable/disable broadcast of discovery messages for log unification.
* `configuration_sync` - Configuration sync mode.
* `fabric_object_unification` - Fabric CMDB Object Unification
* `saml_configuration_sync` - SAML setting configuration synchronization.
* `management_ip` - Management IP address of this FortiGate. Used to log into this FortiGate from another FortiGate in the Security Fabric.
* `management_port` - Overriding port for management connection (Overrides admin port).
* `authorization_request_type` - Authorization request type.
* `certificate` - Certificate.
* `fabric_workers` - Number of worker processes for Security Fabric daemon.
* `downstream_access` - Enable/disable downstream device access to this device's configuration and data.
* `downstream_accprofile` - Default access profile for requests from downstream devices.
* `fixed_key` - Auto-generated fixed key used when this device is the root. (Will automatically be generated if not set.)
* `trusted_list` - Pre-authorized and blocked security fabric nodes. The structure of `trusted_list` block is documented below.
* `fabric_connector` - Fabric connector configuration. The structure of `fabric_connector` block is documented below.
* `forticloud_account_enforcement` - Fabric FortiCloud account unification.
* `file_mgmt` - Enable/disable Security Fabric daemon file management.
* `file_quota` - Maximum amount of memory that can be used by the daemon files (in bytes).
* `file_quota_warning` - Warn when the set percentage of quota has been used.
* `fabric_device` - Fabric device configuration. The structure of `fabric_device` block is documented below.

The `trusted_list` block contains:

* `name` - Name.
* `authorization_type` - Authorization type.
* `serial` - Serial.
* `certificate` - Certificate.
* `action` - Security fabric authorization action.
* `ha_members` - HA members.
* `downstream_authorization` - Trust authorizations by this node's administrator.
* `index` - Index of the downstream in tree.

The `fabric_connector` block contains:

* `serial` - Serial.
* `accprofile` - Override access profile.
* `configuration_write_access` - Enable/disable downstream device write access to configuration.
* `vdom` - Virtual domains that the connector has access to. If none are set, the connector will only have access to the VDOM that it joins the Security Fabric through. The structure of `vdom` block is documented below.

The `vdom` block contains:

* `name` - Virtual domain name.

The `fabric_device` block contains:

* `name` - Device name.
* `device_ip` - Device IP.
* `https_port` - HTTPS port for fabric device.
* `access_token` - Device access token.
* `device_type` - Device type.
* `login` - Device login name.
* `password` - Device login password.

