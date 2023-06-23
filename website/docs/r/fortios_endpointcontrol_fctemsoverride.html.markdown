---
subcategory: "FortiGate Endpoint-Control"
layout: "fortios"
page_title: "FortiOS: fortios_endpointcontrol_fctemsoverride"
description: |-
  Configure FortiClient Enterprise Management Server (EMS) entries.
---

# fortios_endpointcontrol_fctemsoverride
Configure FortiClient Enterprise Management Server (EMS) entries. Applies to FortiOS Version `>= 7.4.0`.

## Argument Reference

The following arguments are supported:

* `ems_id` - EMS ID in order (1 - 7).
* `status` - Enable or disable this EMS configuration. Valid values: `enable`, `disable`.
* `name` - FortiClient Enterprise Management Server (EMS) name.
* `dirty_reason` - Dirty Reason for FortiClient EMS. Valid values: `none`, `mismatched-ems-sn`.
* `fortinetone_cloud_authentication` - Enable/disable authentication of FortiClient EMS Cloud through FortiCloud account. Valid values: `enable`, `disable`.
* `server` - FortiClient EMS FQDN or IPv4 address.
* `https_port` - FortiClient EMS HTTPS access port number. (1 - 65535, default: 443).
* `serial_number` - EMS Serial Number.
* `tenant_id` - EMS Tenant ID.
* `source_ip` - REST API call source IP.
* `pull_sysinfo` - Enable/disable pulling SysInfo from EMS. Valid values: `enable`, `disable`.
* `pull_vulnerabilities` - Enable/disable pulling vulnerabilities from EMS. Valid values: `enable`, `disable`.
* `pull_avatars` - Enable/disable pulling avatars from EMS. Valid values: `enable`, `disable`.
* `pull_tags` - Enable/disable pulling FortiClient user tags from EMS. Valid values: `enable`, `disable`.
* `pull_malware_hash` - Enable/disable pulling FortiClient malware hash from EMS. Valid values: `enable`, `disable`.
* `cloud_server_type` - Cloud server type. Valid values: `production`, `alpha`, `beta`.
* `capabilities` - List of EMS capabilities. Valid values: `fabric-auth`, `silent-approval`, `websocket`, `websocket-malware`, `push-ca-certs`, `common-tags-api`, `tenant-id`, `single-vdom-connector`.
* `call_timeout` - FortiClient EMS call timeout in seconds (1 - 180 seconds, default = 30).
* `out_of_sync_threshold` - Outdated resource threshold in seconds (10 - 3600, default = 180).
* `websocket_override` - Enable/disable override behavior for how this FortiGate unit connects to EMS using a WebSocket connection. Valid values: `disable`, `enable`.
* `preserve_ssl_session` - Enable/disable preservation of EMS SSL session connection. Warning, most users should not touch this setting. Valid values: `enable`, `disable`.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `trust_ca_cn` - Enable/disable trust of the EMS certificate issuer(CA) and common name(CN) for certificate auto-renewal. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{ems_id}}.

## Import

EndpointControl FctemsOverride can be imported using any of these accepted formats:
```
$ terraform import fortios_endpointcontrol_fctemsoverride.labelname {{ems_id}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_endpointcontrol_fctemsoverride.labelname {{ems_id}}
$ unset "FORTIOS_IMPORT_TABLE"
```
