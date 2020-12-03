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

* `name` - (Required) Decoder name.
* `parameter` - IPS group parameters. The structure of `parameter` block is documented below.

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
