---
subcategory: "FortiGate SpamFilter"
layout: "fortios"
page_title: "FortiOS: fortios_spamfilter_bwl"
description: |-
  Configure anti-spam black/white list.
---

# fortios_spamfilter_bwl
Configure anti-spam black/white list.

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
* `entries` - Anti-spam black/white list entries.

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

Spamfilter Bwl can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_spamfilter_bwl.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
