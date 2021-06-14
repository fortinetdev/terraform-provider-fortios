---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_keychain"
description: |-
  Get information on an fortios router keychain.
---

# Data Source: fortios_router_keychain
Use this data source to get information on an fortios router keychain

## Argument Reference

* `name` - (Required) Specify the name of the desired router keychain.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Key-chain name.
* `key` - Configuration method to edit key settings. The structure of `key` block is documented below.

The `key` block contains:

* `id` - Key ID (0 - 2147483647).
* `accept_lifetime` - Lifetime of received authentication key (format: hh:mm:ss day month year).
* `send_lifetime` - Lifetime of sent authentication key (format: hh:mm:ss day month year).
* `key_string` - Password for the key (max. = 35 characters).

