---
subcategory: "FortiGate EmailFilter"
layout: "fortios"
page_title: "FortiOS: fortios_emailfilter_dnsbl"
description: |-
  Configure AntiSpam DNSBL/ORBL.
---

# fortios_emailfilter_dnsbl
Configure AntiSpam DNSBL/ORBL.

## Argument Reference

The following arguments are supported:

* `fosid` - ID.
* `name` - Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `entries` block supports:

* `status` - Enable/disable status. Valid values: `enable`, `disable`.
* `id` - DNSBL/ORBL entry ID.
* `server` - DNSBL or ORBL server name.
* `action` - Reject connection or mark as spam email. Valid values: `reject`, `spam`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Emailfilter Dnsbl can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_emailfilter_dnsbl.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
