---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_dnsbl"
description: |-
  Configure AntiSpam DNSBL/ORBL.
---

# fortios_spamfilter_dnsbl
Configure AntiSpam DNSBL/ORBL. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_spamfilter_dnsbl" "trname" {
  comment = "test"
  fosid   = 1
  name    = "s1"

  entries {
    action = "reject"
    server = "1.1.1.1"
    status = "enable"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - Spam filter DNSBL and ORBL server. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `status` - Enable/disable status. Valid values: `enable`, `disable`.
* `id` - DNSBL/ORBL entry ID.
* `server` - DNSBL or ORBL server name.
* `action` - Reject connection or mark as spam email. Valid values: `reject`, `spam`.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Spamfilter Dnsbl can be imported using any of these accepted formats:
```
$ terraform import fortios_spamfilter_dnsbl.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_spamfilter_dnsbl.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
