---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_externalresource"
description: |-
  Configure external resource.
---

# fortios_system_externalresource
Configure external resource.

## Example Usage

```hcl
resource "fortios_system_externalresource" "trname" {
  category     = 199
  name         = "externalresource1"
  refresh_rate = 5
  resource     = "https://tmpxxxxx.com"
  status       = "enable"
  type         = "category"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) External resource name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `status` - Enable/disable user resource. Valid values: `enable`, `disable`.
* `type` - User resource type.
* `namespace` - Generic external connector address namespace.
* `object_array_path` - JSON Path to array of generic addresses in resource.
* `address_name_field` - JSON Path to address name in generic address entry.
* `address_data_field` - JSON Path to address data in generic address entry.
* `address_comment_field` - JSON Path to address description in generic address entry.
* `update_method` - External resource update method. Valid values: `feed`, `push`.
* `category` - User resource category.
* `username` - HTTP basic authentication user name.
* `password` - HTTP basic authentication password.
* `client_cert_auth` - Enable/disable using client certificate for TLS authentication. Valid values: `enable`, `disable`.
* `client_cert` - Client certificate name.
* `comments` - Comment.
* `resource` - (Required) URI of external resource.
* `user_agent` - HTTP User-Agent header (default = 'curl/7.58.0').
* `server_identity_check` - Certificate verification option. Valid values: `none`, `basic`, `full`.
* `refresh_rate` - (Required) Time interval to refresh external resource (1 - 43200 min, default = 5 min).
* `source_ip` - Source IPv4 address used to communicate with server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ExternalResource can be imported using any of these accepted formats:
```
$ terraform import fortios_system_externalresource.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_externalresource.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
