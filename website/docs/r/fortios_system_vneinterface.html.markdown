---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vneinterface"
description: |-
  Configure virtual network enabler tunnels.
---

# fortios_system_vneinterface
Configure virtual network enabler tunnels. Applies to FortiOS Version `>= 7.6.0`.

## Argument Reference

The following arguments are supported:

* `name` - VNE tunnel name.
* `interface` - Interface name.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `bmr_hostname` - BMR hostname.
* `auto_asic_offload` - Enable/disable tunnel ASIC offloading. Valid values: `enable`, `disable`.
* `ipv4_address` - Tunnel IPv4 address and netmask.
* `br` - IPv6 address or FQDN of the border relay.
* `update_url` - URL of provisioning server.
* `mode` - VNE tunnel mode. Valid values: `map-e`, `fixed-ip`, `ds-lite`.
* `http_username` - HTTP authentication user name.
* `http_password` - HTTP authentication password.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

System VneInterface can be imported using any of these accepted formats:
```
$ terraform import fortios_system_vneinterface.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_vneinterface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
