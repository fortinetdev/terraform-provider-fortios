---
subcategory: "FortiGate Certificate"
layout: "fortios"
page_title: "FortiOS: fortios_certificate_remote"
description: |-
  Remote certificate as a PEM file.
---

# fortios_certificate_remote
Remote certificate as a PEM file.

## Argument Reference

The following arguments are supported:

* `name` - Name.
* `remote` - Remote certificate.
* `range` - Either the global or VDOM IP address range for the remote certificate. Valid values: `global`, `vdom`.
* `source` - Remote certificate source type. Valid values: `factory`, `user`, `bundle`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Certificate Remote can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_certificate_remote.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
