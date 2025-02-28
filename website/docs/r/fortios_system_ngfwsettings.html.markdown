---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ngfwsettings"
description: |-
  Configure IPS NGFW policy-mode VDOM settings.
---

# fortios_system_ngfwsettings
Configure IPS NGFW policy-mode VDOM settings. Applies to FortiOS Version `>= 7.6.1`.

## Argument Reference

The following arguments are supported:

* `match_timeout` - Number of seconds to wait before a security policy match for an idle non-TCP session (0 - 1800, default = 300, 0 means unlimited).
* `tcp_match_timeout` - Number of seconds to wait before a security policy match for an idle TCP session (0 - 1800, default = 300, 0 means unlimited).
* `tcp_halfopen_match_timeout` - Number of seconds to wait before a security policy match for a session after one peer has sent an open session packet but the other has not responded (0 - 300, default = 8, 0 means unlimited).
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System NgfwSettings can be imported using any of these accepted formats:
```
$ terraform import fortios_system_ngfwsettings.labelname SystemNgfwSettings

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_ngfwsettings.labelname SystemNgfwSettings
$ unset "FORTIOS_IMPORT_TABLE"
```
