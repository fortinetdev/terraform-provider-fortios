---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssh_localkey"
description: |-
  SSH proxy local keys.
---

# fortios_firewallssh_localkey
SSH proxy local keys.

## Argument Reference

The following arguments are supported:

* `name` - SSH proxy local key name.
* `password` - Password for SSH private key.
* `private_key` - (Required) SSH proxy private key, encrypted with a password.
* `public_key` - (Required) SSH proxy public key.
* `source` - SSH proxy local key source type. Valid values: `built-in`, `user`.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallSsh LocalKey can be imported using any of these accepted formats:
```
$ terraform import fortios_firewallssh_localkey.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallssh_localkey.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
