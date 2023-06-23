---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ike"
description: |-
  Configure IKE global attributes.
---

# fortios_system_ike
Configure IKE global attributes. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `embryonic_limit` - Maximum number of IPsec tunnels to negotiate simultaneously.
* `dh_multiprocess` - Enable/disable multiprocess Diffie-Hellman daemon for IKE. Valid values: `enable`, `disable`.
* `dh_worker_count` - Number of Diffie-Hellman workers to start.
* `dh_mode` - Use software (CPU) or hardware (CPX) to perform Diffie-Hellman calculations. Valid values: `software`, `hardware`.
* `dh_keypair_cache` - Enable/disable Diffie-Hellman key pair cache. Valid values: `enable`, `disable`.
* `dh_keypair_count` - Number of key pairs to pre-generate for each Diffie-Hellman group (per-worker).
* `dh_keypair_throttle` - Enable/disable Diffie-Hellman key pair cache CPU throttling. Valid values: `enable`, `disable`.
* `dh_group_1` - Diffie-Hellman group 1 (MODP-768). The structure of `dh_group_1` block is documented below.
* `dh_group_2` - Diffie-Hellman group 2 (MODP-1024). The structure of `dh_group_2` block is documented below.
* `dh_group_5` - Diffie-Hellman group 5 (MODP-1536). The structure of `dh_group_5` block is documented below.
* `dh_group_14` - Diffie-Hellman group 14 (MODP-2048). The structure of `dh_group_14` block is documented below.
* `dh_group_15` - Diffie-Hellman group 15 (MODP-3072). The structure of `dh_group_15` block is documented below.
* `dh_group_16` - Diffie-Hellman group 16 (MODP-4096). The structure of `dh_group_16` block is documented below.
* `dh_group_17` - Diffie-Hellman group 17 (MODP-6144). The structure of `dh_group_17` block is documented below.
* `dh_group_18` - Diffie-Hellman group 18 (MODP-8192). The structure of `dh_group_18` block is documented below.
* `dh_group_19` - Diffie-Hellman group 19 (EC-P256). The structure of `dh_group_19` block is documented below.
* `dh_group_20` - Diffie-Hellman group 20 (EC-P384). The structure of `dh_group_20` block is documented below.
* `dh_group_21` - Diffie-Hellman group 21 (EC-P521). The structure of `dh_group_21` block is documented below.
* `dh_group_27` - Diffie-Hellman group 27 (EC-P224BP). The structure of `dh_group_27` block is documented below.
* `dh_group_28` - Diffie-Hellman group 28 (EC-P256BP). The structure of `dh_group_28` block is documented below.
* `dh_group_29` - Diffie-Hellman group 29 (EC-P384BP). The structure of `dh_group_29` block is documented below.
* `dh_group_30` - Diffie-Hellman group 30 (EC-P512BP). The structure of `dh_group_30` block is documented below.
* `dh_group_31` - Diffie-Hellman group 31 (EC-X25519). The structure of `dh_group_31` block is documented below.
* `dh_group_32` - Diffie-Hellman group 32 (EC-X448). The structure of `dh_group_32` block is documented below.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `dh_group_1` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_2` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_5` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_14` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_15` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_16` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_17` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_18` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_19` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_20` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_21` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_27` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_28` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_29` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_30` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_31` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).

The `dh_group_32` block supports:

* `mode` - Use software (CPU) or hardware (CPX) to perform calculations for this Diffie-Hellman group. Valid values: `software`, `hardware`, `global`.
* `keypair_cache` - Configure custom key pair cache size for this Diffie-Hellman group. Valid values: `global`, `custom`.
* `keypair_count` - Number of key pairs to pre-generate for this Diffie-Hellman group (per-worker).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Ike can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ike.labelname SystemIke

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ike.labelname SystemIke
$ unset "FORTIOS_IMPORT_TABLE"
```
