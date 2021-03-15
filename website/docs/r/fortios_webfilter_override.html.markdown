---
subcategory: "FortiGate WebFilter"
layout: "fortios"
page_title: "FortiOS: fortios_webfilter_override"
description: |-
  Configure FortiGuard Web Filter administrative overrides.
---

# fortios_webfilter_override
Configure FortiGuard Web Filter administrative overrides.

## Example Usage

```hcl
resource "fortios_webfilter_override" "trname" {
  expires     = "2021/03/16 19:18:25"
  fosid       = 1
  ip          = "69.101.119.0"
  ip6         = "4565:7700::"
  new_profile = "monitor-all"
  old_profile = "default"
  scope       = "user"
  status      = "disable"
  user        = "Eew"

}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Override rule ID.
* `status` - Enable/disable override rule. Valid values: `enable`, `disable`.
* `scope` - Override either the specific user, user group, IPv4 address, or IPv6 address. Valid values: `user`, `user-group`, `ip`, `ip6`.
* `ip` - IPv4 address which the override applies.
* `user` - (Required) Name of the user which the override applies.
* `user_group` - Specify the user group for which the override applies.
* `old_profile` - (Required) Name of the web filter profile which the override applies.
* `new_profile` - (Required) Name of the new web filter profile used by the override.
* `ip6` - IPv6 address which the override applies.
* `expires` - (Required) Override expiration date and time, from 5 minutes to 365 from now (format: yyyy/mm/dd hh:mm:ss).
* `initiator` - Initiating user of override (read-only setting).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Webfilter Override can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_webfilter_override.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
