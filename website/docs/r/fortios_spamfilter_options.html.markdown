---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_options"
description: |-
  Configure AntiSpam options.
---

# fortios_spamfilter_options
Configure AntiSpam options. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_spamfilter_options" "trname" {
  dns_timeout = 7
}
```

## Argument Reference

The following arguments are supported:

* `dns_timeout` - DNS query time out (1 - 30 sec).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Spamfilter Options can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_spamfilter_options.labelname SpamfilterOptions
$ unset "FORTIOS_IMPORT_TABLE"
```
