---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vnetunnel"
description: |-
  Configure virtual network enabler tunnel.
---

# fortios_system_vnetunnel
Configure virtual network enabler tunnel. Applies to FortiOS Version `>= 6.4.1`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable VNE tunnel. Valid values: `enable`, `disable`.
* `interface` - Interface name.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `bmr_hostname` - BMR hostname.
* `auto_asic_offload` - Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
* `ipv4_address` - Tunnel IPv4 address and netmask.
* `br` - Border relay IPv6 address.
* `update_url` - URL of provisioning server.
* `mode` - VNE tunnel mode.
* `http_username` - HTTP authentication user name.
* `http_password` - HTTP authentication password.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VneTunnel can be imported using any of these accepted formats:
```
$ terraform import fortios_system_vnetunnel.labelname SystemVneTunnel

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_vnetunnel.labelname SystemVneTunnel
$ unset "FORTIOS_IMPORT_TABLE"
```
