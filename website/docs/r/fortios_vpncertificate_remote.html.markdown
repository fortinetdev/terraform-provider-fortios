---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpncertificate_remote"
description: |-
  Remote certificate as a PEM file.
---

# fortios_vpncertificate_remote
Remote certificate as a PEM file.

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name.
* `remote` - Remote certificate.
* `range` - Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
* `source` - Remote certificate source type.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnCertificate Remote can be imported using any of these accepted formats:
```
$ terraform import fortios_vpncertificate_remote.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpncertificate_remote.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
