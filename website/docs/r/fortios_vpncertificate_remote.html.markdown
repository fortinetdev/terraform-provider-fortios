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
* `range` - Either the global or VDOM IP address range for the remote certificate.
* `source` - Remote certificate source type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnCertificate Remote can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpncertificate_remote.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
