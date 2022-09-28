---
subcategory: "FortiGate Extension-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_extensioncontroller_fortigateprofile"
description: |-
  FortiGate connector profile configuration.
---

# fortios_extensioncontroller_fortigateprofile
FortiGate connector profile configuration. Applies to FortiOS Version `>= 7.2.1`.

## Argument Reference

The following arguments are supported:

* `name` - FortiGate connector profile name.
* `fosid` - ID.
* `extension` - Extension option. Valid values: `lan-extension`.
* `lan_extension` - FortiGate connector LAN extension configuration. The structure of `lan_extension` block is documented below.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `lan_extension` block supports:

* `ipsec_tunnel` - IPsec tunnel name.
* `backhaul_interface` - IPsec phase1 interface.
* `backhaul_ip` - IPsec phase1 IPv4/FQDN. Used to specify the external IP/FQDN when the FortiGate unit is behind a NAT device.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

ExtensionController FortigateProfile can be imported using any of these accepted formats:
```
$ terraform import fortios_extensioncontroller_fortigateprofile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_extensioncontroller_fortigateprofile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
