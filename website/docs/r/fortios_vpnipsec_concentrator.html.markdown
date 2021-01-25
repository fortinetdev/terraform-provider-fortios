---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnipsec_concentrator"
description: |-
  Concentrator configuration.
---

# fortios_vpnipsec_concentrator
Concentrator configuration.

## Example Usage

```hcl
resource "fortios_vpnipsec_concentrator" "trname" {
  name      = "1"
  src_check = "disable"
}
```

## Argument Reference


The following arguments are supported:

* `name` - Concentrator name.
* `src_check` - Enable to check source address of phase 2 selector. Disable to check only the destination selector.
* `member` - Names of up to 3 VPN tunnels to add to the concentrator. The structure of `member` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `member` block supports:

* `name` - Member name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnIpsec Concentrator can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_vpnipsec_concentrator.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
