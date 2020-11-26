---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_iptrust"
description: |-
  Configure AntiSpam IP trust.
---

# fortios_spamfilter_iptrust
Configure AntiSpam IP trust.

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter trusted IP addresses. The structure of `entries` block is documented below.

The `entries` block supports:

* `status` - Enable/disable status.
* `id` - Trusted IP entry ID.
* `addr_type` - Type of address.
* `ip4_subnet` - IPv4 network address or network address/subnet mask bits.
* `ip6_subnet` - IPv6 network address/subnet mask bits.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Spamfilter Iptrust can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_spamfilter_iptrust.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
