---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sdnproxy"
description: |-
  Configure SDN proxy.
---

# fortios_system_sdnproxy
Configure SDN proxy. Applies to FortiOS Version `>= 7.4.0`.

## Argument Reference

The following arguments are supported:

* `name` - SDN proxy name.
* `type` - Type of SDN proxy. Valid values: `general`, `fortimanager`.
* `server` - Server address of the SDN proxy.
* `server_port` - Port number of the SDN proxy.
* `username` - SDN proxy username.
* `password` - SDN proxy password.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System SdnProxy can be imported using any of these accepted formats:
```
$ terraform import fortios_system_sdnproxy.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_sdnproxy.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
