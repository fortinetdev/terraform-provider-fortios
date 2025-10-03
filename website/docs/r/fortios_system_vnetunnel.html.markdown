---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vnetunnel"
description: |-
  Configure virtual network enabler tunnel.
---

# fortios_system_vnetunnel
Configure virtual network enabler tunnel. Applies to FortiOS Version `6.4.1,6.4.2,6.4.10,6.4.11,6.4.12,6.4.13,6.4.14,6.4.15,7.0.0,7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.0.14,7.0.15,7.0.16,7.0.17,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.2.7,7.2.8,7.2.9,7.2.10,7.2.11,7.2.12,7.4.0,7.4.1,7.4.2,7.4.3,7.4.4,7.4.5,7.4.6,7.4.7,7.4.8`.

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
