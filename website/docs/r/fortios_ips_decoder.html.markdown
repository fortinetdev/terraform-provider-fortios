---
subcategory: "FortiGate IPS"
layout: "fortios"
page_title: "FortiOS: fortios_ips_decoder"
description: |-
  Configure IPS decoder.
---

# fortios_ips_decoder
Configure IPS decoder.

## Argument Reference

The following arguments are supported:

* `name` - Decoder name.
* `parameter` - IPS group parameters. The structure of `parameter` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `parameter` block supports:

* `name` - Parameter name.
* `value` - Parameter value.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Ips Decoder can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_ips_decoder.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
