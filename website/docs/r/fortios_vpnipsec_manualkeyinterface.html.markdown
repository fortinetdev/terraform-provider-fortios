---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_manualkeyinterface"
description: |-
  Configure IPsec manual keys.
---

# fortios_vpnipsec_manualkeyinterface
Configure IPsec manual keys.

## Example Usage

```hcl
resource "fortios_vpnipsec_manualkeyinterface" "trname" {
  addr_type  = "4"
  auth_alg   = "null"
  auth_key   = "-"
  enc_alg    = "des"
  enc_key    = "CECA2184ACADAEEF"
  interface  = "port3"
  ip_version = "4"
  local_gw   = "0.0.0.0"
  local_gw6  = "::"
  local_spi  = "0x100"
  name       = "mi1"
  remote_gw  = "2.2.2.2"
  remote_gw6 = "::"
  remote_spi = "0x100"
}
```

## Argument Reference

The following arguments are supported:

* `name` - IPsec tunnel name.
* `interface` - (Required) Name of the physical, aggregate, or VLAN interface.
* `ip_version` - IP version to use for VPN interface.
* `addr_type` - IP version to use for IP packets.
* `remote_gw` - (Required) IPv4 address of the remote gateway's external interface.
* `remote_gw6` - (Required) Remote IPv6 address of VPN gateway.
* `local_gw` - IPv4 address of the local gateway's external interface.
* `local_gw6` - Local IPv6 address of VPN gateway.
* `auth_alg` - (Required) Authentication algorithm. Must be the same for both ends of the tunnel.
* `enc_alg` - (Required) Encryption algorithm. Must be the same for both ends of the tunnel.
* `auth_key` - (Required) Hexadecimal authentication key in 16-digit (8-byte) segments separated by hyphens.
* `enc_key` - (Required) Hexadecimal encryption key in 16-digit (8-byte) segments separated by hyphens.
* `local_spi` - (Required) Local SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.
* `remote_spi` - (Required) Remote SPI, a hexadecimal 8-digit (4-byte) tag. Discerns between two traffic streams with different encryption rules.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnIpsec ManualkeyInterface can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnipsec_manualkeyinterface.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
