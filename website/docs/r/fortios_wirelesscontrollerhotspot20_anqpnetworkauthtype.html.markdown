---
subcategory: "FortiGate Wireless-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype"
description: |-
  Configure network authentication type.
---

# fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype
Configure network authentication type.

## Example Usage

```hcl
resource "fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype" "trname" {
  auth_type = "acceptance-of-terms"
  name      = "1"
  url       = "www.example.com"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Authentication type name.
* `auth_type` - Network authentication type. Valid values: `acceptance-of-terms`, `online-enrollment`, `http-redirection`, `dns-redirection`.
* `url` - Redirect URL.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpNetworkAuthType can be imported using any of these accepted formats:
```
$ terraform import fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
