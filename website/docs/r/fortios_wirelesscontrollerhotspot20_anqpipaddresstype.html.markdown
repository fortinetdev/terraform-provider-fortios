---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpipaddresstype"
description: |-
  Configure IP address type availability.
---

# fortios_wirelesscontrollerhotspot20_anqpipaddresstype
Configure IP address type availability.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_anqpipaddresstype" "trname" {
  ipv4_address_type = "public"
  ipv6_address_type = "not-available"
  name              = "1"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IP type name.
* `ipv6_address_type` - IPv6 address type. Valid values: `not-available`, `available`, `not-known`.
* `ipv4_address_type` - IPv4 address type. Valid values: `not-available`, `public`, `port-restricted`, `single-NATed-private`, `double-NATed-private`, `port-restricted-and-single-NATed`, `port-restricted-and-double-NATed`, `not-known`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpIpAddressType can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_anqpipaddresstype.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_anqpipaddresstype.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
