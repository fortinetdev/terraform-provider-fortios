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
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `key` block supports:

* `id` - Key ID (0 - 2147483647).
* `accept_lifetime` - Lifetime of received authentication key (format: hh:mm:ss day month year).
* `send_lifetime` - Lifetime of sent authentication key (format: hh:mm:ss day month year).
* `key_string` - Password for the key (max. = 35 characters).
* `algorithm` - Cryptographic algorithm. Valid values: `md5`, `hmac-sha1`, `hmac-sha256`, `hmac-sha384`, `hmac-sha512`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Router KeyChain can be imported using any of these accepted formats:
```
$ terraform import fortios_router_keychain.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_router_keychain.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
