---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_manualkey"
description: |-
  Configure IPsec manual keys.
---

# fortios_vpnipsec_manualkey
Configure IPsec manual keys.

## Example Usage

```hcl
resource "fortios_vpnipsec_manualkey" "trname" {
  authentication = "md5"
  authkey        = "EE32CA121ECD772A-ECACAABA212345EC"
  enckey         = "-"
  encryption     = "null"
  interface      = "port4"
  local_gw       = "0.0.0.0"
  localspi       = "0x100"
  name           = "mk1"
  remote_gw      = "1.1.1.1"
  remotespi      = "0x100"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) IPsec tunnel name.
* `interface` - (Required) Name of the physical, aggregate, or VLAN interface.
* `remote_gw` - (Required) Peer gateway.
* `local_gw` - Local gateway.
* `authentication` - (Required) Authentication algorithm. Must be the same for both ends of the tunnel.
* `encryption` - (Required) Encryption algorithm. Must be the same for both ends of the tunnel.
* `authkey` - Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
* `enckey` - Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
* `localspi` - Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
* `remotespi` - Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnIpsec Manualkey can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnipsec_manualkey.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
