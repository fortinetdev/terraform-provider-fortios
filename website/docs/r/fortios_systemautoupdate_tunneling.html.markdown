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

* `status` - Enable/disable web proxy tunnelling.
* `address` - Web proxy IP address or FQDN.
* `port` - Web proxy port.
* `username` - Web proxy username.
* `password` - Web proxy password.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

SystemAutoupdate Tunneling can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_systemautoupdate_tunneling.labelname SystemAutoupdateTunneling
$ unset "FORTIOS_IMPORT_TABLE"
```
