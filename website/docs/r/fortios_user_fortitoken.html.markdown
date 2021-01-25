---
subcategory: "FortiGate User"
layout: "fortios"
page_title: "FortiOS: fortios_user_fortitoken"
description: |-
  Configure FortiToken.
---

# fortios_user_fortitoken
Configure FortiToken.

## Argument Reference


The following arguments are supported:

* `serial_number` - Serial number.
* `status` - Status
* `seed` - Token seed.
* `comments` - Comment.
* `license` - Mobile token license.
* `activation_code` - Mobile token user activation-code.
* `activation_expire` - Mobile token user activation-code expire time.
* `reg_id` - Device Reg ID.
* `os_ver` - Device Mobile Version.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial_number}}.

## Import

User Fortitoken can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_user_fortitoken.labelname {{serial_number}}
$ unset "FORTIOS_IMPORT_TABLE"
```
