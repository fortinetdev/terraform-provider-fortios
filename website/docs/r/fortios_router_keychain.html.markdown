---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_keychain"
description: |-
  Configure key-chain.
---

# fortios_router_keychain
Configure key-chain.

## Example Usage

```hcl
resource "fortios_router_keychain" "trname" {
  name = "1"

  key {
    accept_lifetime = "04:00:00 01 01 2008 04:00:00 01 01 2022"
    key_string      = "ewiwn3i23232s212"
    send_lifetime   = "04:00:00 01 01 2008 04:00:00 01 01 2022"
  }
}
```

## Argument Reference


The following arguments are supported:

* `name` - (Required) Key-chain name.
* `key` - Configuration method to edit key settings. The structure of `key` block is documented below.

The `key` block supports:

* `id` - Key ID (0 - 2147483647).
* `accept_lifetime` - Lifetime of received authentication key (format: hh:mm:ss day month year).
* `send_lifetime` - Lifetime of sent authentication key (format: hh:mm:ss day month year).
* `key_string` - Password for the key (max. = 35 characters).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router KeyChain can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_router_keychain.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
