---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_options"
description: |-
  Configure AntiSpam options.
---

# fortios_emailfilter_options
Configure AntiSpam options.

## Argument Reference

The following arguments are supported:

* `dns_timeout` - DNS query time out (1 - 30 sec).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Emailfilter Options can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_emailfilter_options.labelname EmailfilterOptions
$ unset "FORTIOS_IMPORT_TABLE"
```
