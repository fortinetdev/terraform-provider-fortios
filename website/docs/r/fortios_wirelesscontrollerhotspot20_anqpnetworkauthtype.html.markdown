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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

WirelessControllerHotspot20 AnqpNetworkAuthType can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_wirelesscontrollerhotspot20_anqpnetworkauthtype.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
