---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_systemautoupdate_tunneling"
description: |-
  Configure web proxy tunnelling for the FDN.
---

# fortios_systemautoupdate_tunneling
Configure web proxy tunnelling for the FDN.

## Example Usage

```hcl
resource "fortios_systemautoupdate_tunneling" "trname" {
  port   = 0
  status = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable web proxy tunnelling. Valid values: `enable`, `disable`.
* `address` - Web proxy IP address or FQDN.
* `port` - Web proxy port.
* `username` - Web proxy username.
* `password` - Web proxy password.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemAutoupdate Tunneling can be imported using any of these accepted formats:
```
$ terraform import fortios_systemautoupdate_tunneling.labelname SystemAutoupdateTunneling

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_systemautoupdate_tunneling.labelname SystemAutoupdateTunneling
$ unset "FORTIOS_IMPORT_TABLE"
```
