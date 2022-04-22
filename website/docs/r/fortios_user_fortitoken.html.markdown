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
* `status` - Status Valid values: `active`, `lock`.
* `seed` - Token seed.
* `comments` - Comment.
* `license` - Mobile token license.
* `activation_code` - Mobile token user activation-code.
* `activation_expire` - Mobile token user activation-code expire time.
* `reg_id` - Device Reg ID.
* `os_ver` - Device Mobile Version.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{serial_number}}.

## Import

User Fortitoken can be imported using any of these accepted formats:
```
$ terraform import fortios_user_fortitoken.labelname {{serial_number}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_user_fortitoken.labelname {{serial_number}}
$ unset "FORTIOS_IMPORT_TABLE"
```
