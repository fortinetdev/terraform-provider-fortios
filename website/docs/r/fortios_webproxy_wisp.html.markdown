---
subcategory: "FortiGate Web-Proxy"
layout: "fortios"
page_title: "FortiOS: fortios_webproxy_wisp"
description: |-
  Configure Wireless Internet service provider (WISP) servers.
---

# fortios_webproxy_wisp
Configure Wireless Internet service provider (WISP) servers.

## Example Usage

```hcl
resource "fortios_webproxy_wisp" "trname" {
  max_connections = 64
  name            = "1"
  outgoing_ip     = "0.0.0.0"
  server_ip       = "1.1.1.1"
  server_port     = 15868
  timeout         = 5
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Server name.
* `comment` - Comment.
* `outgoing_ip` - WISP outgoing IP address.
* `server_ip` - (Required) WISP server IP address.
* `server_port` - (Required) WISP server port (1 - 65535, default = 15868).
* `max_connections` - Maximum number of web proxy WISP connections (4 - 4096, default = 64).
* `timeout` - Period of time before WISP requests time out (1 - 15 sec, default = 5).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WebProxy Wisp can be imported using any of these accepted formats:
```
$ terraform import fortios_webproxy_wisp.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_webproxy_wisp.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
