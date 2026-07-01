---
subcategory: "FortiGate Switch-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_switchcontrollersecuritypolicy_admin"
description: |-
  Configure fortiswitch's admin security-policy.
---

# fortios_switchcontrollersecuritypolicy_admin
Configure fortiswitch's admin security-policy. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - Policy name.
* `auto` - Automatically set based on the host ip connected via the Fortilink interface. Valid values: `disable`, `enable`.
* `trusthost1` - Trusted IPv4 host.
* `trusthost2` - Trusted IPv4 host.
* `trusthost3` - Trusted IPv4 host.
* `trusthost4` - Trusted IPv4 host.
* `trusthost5` - Trusted IPv4 host.
* `trusthost6` - Trusted IPv4 host.
* `trusthost7` - Trusted IPv4 host.
* `trusthost8` - Trusted IPv4 host.
* `trusthost9` - Trusted IPv4 host.
* `trusthost10` - Trusted IPv4 host.
* `ip6_trusthost1` - Trusted IPv6 host.
* `ip6_trusthost2` - Trusted IPv6 host.
* `ip6_trusthost3` - Trusted IPv6 host.
* `ip6_trusthost4` - Trusted IPv6 host.
* `ip6_trusthost5` - Trusted IPv6 host.
* `ip6_trusthost6` - Trusted IPv6 host.
* `ip6_trusthost7` - Trusted IPv6 host.
* `ip6_trusthost8` - Trusted IPv6 host.
* `ip6_trusthost9` - Trusted IPv6 host.
* `ip6_trusthost10` - Trusted IPv6 host.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

SwitchControllerSecurityPolicy Admin can be imported using any of these accepted formats:
```
$ terraform import fortios_switchcontrollersecuritypolicy_admin.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_switchcontrollersecuritypolicy_admin.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
