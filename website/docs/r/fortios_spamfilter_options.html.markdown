---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_options"
description: |-
  Configure AntiSpam options.
---

# fortios_spamfilter_options
Configure AntiSpam options.

## Example Usage

```hcl
resource "fortios_spamfilter_options" "trname" {
  dns_timeout = 7
}
```

## Argument Reference


The following arguments are supported:

* `dns_timeout` - DNS query time out (1 - 30 sec).


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
