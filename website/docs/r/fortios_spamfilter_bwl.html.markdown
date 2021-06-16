---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_bwl"
description: |-
  Configure anti-spam black/white list.
---

# fortios_spamfilter_bwl
Configure anti-spam black/white list. Applies to FortiOS Version `<= 6.2.0`.

## Example Usage

```hcl
resource "fortios_spamfilter_bwl" "trname" {
  comment = "test"
  fosid   = 1
  name    = "s2"

  entries {
    action       = "reject"
    addr_type    = "ipv4"
    ip4_subnet   = "1.1.1.0 255.255.255.0"
    ip6_subnet   = "::/128"
    pattern_type = "wildcard"
    status       = "enable"
    type         = "ip"
  }
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - (Required) ID.
* `name` - (Required) Name of table.
* `comment` - Optional comments.
* `entries` - Anti-spam black/white list entries. The structure of `entries` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `entries` block supports:

* `status` - Enable/disable status. Valid values: `enable`, `disable`.
* `id` - Entry ID.
* `type` - Entry type. Valid values: `ip`, `email`.
* `action` - Reject, mark as spam or good email. Valid values: `reject`, `spam`, `clear`.
* `addr_type` - IP address type. Valid values: `ipv4`, `ipv6`.
* `ip4_subnet` - IPv4 network address/subnet mask bits.
* `ip6_subnet` - IPv6 network address/subnet mask bits.
* `pattern_type` - Wildcard pattern or regular expression. Valid values: `wildcard`, `regexp`.
* `email_pattern` - Email address pattern.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Spamfilter Bwl can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_spamfilter_bwl.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
