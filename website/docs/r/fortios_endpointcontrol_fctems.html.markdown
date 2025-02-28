---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_fctems"
description: |-
  Configure FortiClient Enterprise Management Server (EMS) entries.
---

# fortios_endpointcontrol_fctems
Configure FortiClient Enterprise Management Server (EMS) entries. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `ems_id` - EMS ID in order. On FortiOS versions 7.0.8-7.0.17, 7.2.1-7.2.3: 1 - 5. On FortiOS versions >= 7.2.4: 1 - 7.
* `status` - Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
* `name` - FortiClient Enterprise Management Server (EMS) name.
* `dirty_reason` - Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
* `server` - FortiClient EMS FQDN or IPv4 address.
* `serial_number` - FortiClient EMS Serial Number.
* `tenant_id` - EMS Tenant ID.
* `fortinetone_cloud_authentication` - Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
* `cloud_authentication_access_key` - FortiClient EMS Cloud multitenancy access key
* `https_port` - FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
* `admin_username` - FortiClient EMS admin username.
* `admin_password` - FortiClient EMS admin password.
* `source_ip` - REST API call source IP.
* `pull_sysinfo` - Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
* `pull_vulnerabilities` - Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
* `pull_avatars` - Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
* `pull_tags` - Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
* `pull_malware_hash` - Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
* `cloud_server_type` - Cloud server type. Valid values: `production`, `alpha`, `beta`.
* `capabilities` - List of EMS capabilities.
* `call_timeout` - FortiClient EMS call timeout. On FortiOS versions 6.2.4-6.2.6: 500 - 30000 milliseconds, default = 5000. On FortiOS versions 6.4.0: 500 - 50000 milliseconds, default = 5000. On FortiOS versions >= 6.4.2: 1 - 180 seconds, default = 30. On FortiOS versions 6.4.1: 500 - 180000 milliseconds, default = 30000.
* `out_of_sync_threshold` - Outdated resource threshold in seconds (10 - 3600, default = 180).
* `send_tags_to_all_vdoms` - Relax restrictions on tags to send all EMS tags to all VDOMs Valid values: `enable`, `disable`.
* `websocket_override` - Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
* `preserve_ssl_session` - Enable/disable preservation of EMS SSL session connection. WARNING: Most users should not touch this setting! Valid values: `enable`, `disable`.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `trust_ca_cn` - Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
* `verifying_ca` - Lowest CA cert on Fortigate in verified EMS cert chain.
* `status_check_interval` - FortiClient EMS call timeout in seconds (1 - 120 seconds, default = 5).
* `certificate` - FortiClient EMS certificate.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

EndpointControl Fctems can be imported using any of these accepted formats:
```
$ terraform import fortios_endpointcontrol_fctems.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_endpointcontrol_fctems.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
