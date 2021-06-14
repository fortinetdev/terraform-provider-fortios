---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fssopolling"
description: |-
  Configure Fortinet Single Sign On (FSSO) server.
---

# fortios_system_fssopolling
Configure Fortinet Single Sign On (FSSO) server.

## Example Usage

```hcl
resource "fortios_system_fssopolling" "trname" {
  authentication = "disable"
  listening_port = 8000
  status         = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable FSSO Polling Mode. Valid values: `enable`, `disable`.
* `listening_port` - Listening port to accept clients (1 - 65535).
* `authentication` - Enable/disable FSSO Agent Authentication. Valid values: `enable`, `disable`.
* `auth_password` - Password to connect to FSSO Agent.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System FssoPolling can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_fssopolling.labelname SystemFssoPolling
$ unset "FORTIOS_IMPORT_TABLE"
```
