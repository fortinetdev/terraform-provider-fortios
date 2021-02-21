---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_blockallowlist"
description: |-
  Configure anti-spam block/allow list.
---

# fortios_emailfilter_blockallowlist
Configure anti-spam block/allow list. Applies to FortiOS Version `>= 7.0.0`.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `name` - Name of table.
* `comment` - Optional comments.
* `entries` - Anti-spam block/allow entries. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `status` - Enable/disable status.
* `id` - Entry ID.
* `type` - Entry type.
* `action` - Reject, mark as spam or good email.
* `addr_type` - IP address type.
* `ip4_subnet` - IPv4 network address/subnet mask bits.
* `ip6_subnet` - IPv6 network address/subnet mask bits.
* `pattern_type` - Wildcard pattern or regular expression.
* `email_pattern` - Email address pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter BlockAllowList can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_emailfilter_blockallowlist.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
