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
* `type` - User resource type. Valid values: `category`, `address`, `domain`, `malware`.
* `category` - User resource category.
* `username` - HTTP basic authentication user name.
* `password` - HTTP basic authentication password.
* `comments` - Comment.
* `resource` - (Required) URI of external resource.
* `user_agent` - Override HTTP User-Agent header used when retrieving this external resource.
* `refresh_rate` - (Required) Time interval to refresh external resource (1 - 43200 min, default = 5 min).
* `source_ip` - Source IPv4 address used to communicate with server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System ExternalResource can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_externalresource.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
