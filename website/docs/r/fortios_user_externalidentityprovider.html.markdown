---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_externalidentityprovider"
description: |-
  Configure external identity provider.
---

# fortios_user_externalidentityprovider
Configure external identity provider. Applies to FortiOS Version `7.2.8,7.2.9,7.2.10,7.2.11,7.2.12,7.4.2,7.4.3,7.4.4,7.4.5,7.4.6,7.4.7,7.4.8,7.6.0,7.6.1,7.6.2,7.6.3,7.6.4`.

## Argument Reference

The following arguments are supported:

* `name` - External identity provider name.
* `type` - External identity provider type. Valid values: `ms-graph`.
* `version` - External identity API version. Valid values: `v1.0`, `beta`.
* `url` - External identity provider URL (e.g. "https://example.com:8080/api/v1").
* `user_attr_name` - User attribute name in authentication query.
* `group_attr_name` - Group attribute name in authentication query.
* `port` - External identity provider service port number (0 to use default).
* `source_ip` - Use this IPv4/v6 address to connect to the external identity provider.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `server_identity_check` - Enable/disable server's identity check against its certificate and subject alternative name(s). Valid values: `disable`, `enable`.
* `timeout` - Connection timeout value in seconds (default=5).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

User ExternalIdentityProvider can be imported using any of these accepted formats:
```
$ terraform import fortios_user_externalidentityprovider.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_externalidentityprovider.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
