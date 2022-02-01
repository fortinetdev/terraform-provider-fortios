---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_acme"
description: |-
  Configure ACME client.
---

# fortios_system_acme
Configure ACME client. Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `interface` - Interface(s) on which the ACME client will listen for challenges. The structure of `interface` block is documented below.
* `accounts` - ACME accounts list. The structure of `accounts` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `interface` block supports:

* `interface_name` - Interface name.

The `accounts` block supports:

* `id` - Account id.
* `status` - Account status.
* `url` - Account url.
* `ca_url` - Account ca_url.
* `email` - Account email.
* `privatekey` - Account Private Key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Acme can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_acme.labelname SystemAcme
$ unset "FORTIOS_IMPORT_TABLE"
```
