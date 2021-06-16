---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ippool6"
description: |-
  Configure IPv6 IP pools.
---

# fortios_firewall_ippool6
Configure IPv6 IP pools.

## Example Usage

```hcl
resource "fortios_firewall_ippool6" "trname" {
  endip   = "2001:3ca1:10f:1a:121b::19"
  name    = "ippool6s1"
  startip = "2001:3ca1:10f:1a:121b::10"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IPv6 IP pool name.
* `startip` - (Required) First IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `endip` - (Required) Final IPv6 address (inclusive) in the range for the address pool (format xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx:xxxx, Default: ::).
* `comments` - Comment.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall Ippool6 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_ippool6.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
