---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_externalresource"
description: |-
  Get information on an fortios system externalresource.
---

# Data Source: fortios_system_externalresource
Use this data source to get information on an fortios system externalresource

## Argument Reference

* `name` - (Required) Specify the name of the desired system externalresource.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - External resource name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `status` - Enable/disable user resource.
* `type` - User resource type.
* `namespace` - Generic external connector address namespace.
* `object_array_path` - JSON Path to array of generic addresses in resource.
* `address_name_field` - JSON Path to address name in generic address entry.
* `address_data_field` - JSON Path to address data in generic address entry.
* `address_comment_field` - JSON Path to address description in generic address entry.
* `update_method` - External resource update method.
* `category` - User resource category.
* `username` - HTTP basic authentication user name.
* `password` - HTTP basic authentication password.
* `client_cert_auth` - Enable/disable using client certificate for TLS authentication.
* `client_cert` - Client certificate name.
* `comments` - Comment.
* `resource` - URI of external resource.
* `user_agent` - Override HTTP User-Agent header used when retrieving this external resource.
* `server_identity_check` - Certificate verification option.
* `refresh_rate` - Time interval to refresh external resource (1 - 43200 min, default = 5 min).
* `source_ip` - Source IPv4 address used to communicate with server.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.

