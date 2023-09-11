---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ftmpush"
description: |-
  Configure FortiToken Mobile push services.
---

# fortios_system_ftmpush
Configure FortiToken Mobile push services.

## Example Usage

```hcl
resource "fortios_system_ftmpush" "trname" {
  server_ip   = "0.0.0.0"
  server_port = 4433
  status      = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `proxy` - Enable/disable communication to the proxy server in FortiGuard configuration. Valid values: `enable`, `disable`.
* `server_port` - Port to communicate with FortiToken Mobile push services server (1 - 65535, default = 4433).
* `server_cert` - Name of the server certificate to be used for SSL (default = Fortinet_Factory).
* `server_ip` - IPv4 address of FortiToken Mobile push services server (format: xxx.xxx.xxx.xxx).
* `server` - IPv4 address or domain name of FortiToken Mobile push services server.
* `status` - Enable/disable the use of FortiToken Mobile push services. Valid values: `enable`, `disable`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FtmPush can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ftmpush.labelname SystemFtmPush

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ftmpush.labelname SystemFtmPush
$ unset "FORTIOS_IMPORT_TABLE"
```
