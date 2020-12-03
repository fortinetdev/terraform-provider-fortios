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

* `name` - (Required) SSH proxy local key name.
* `password` - Password for SSH private key.
* `private_key` - (Required) SSH proxy private key, encrypted with a password.
* `public_key` - (Required) SSH proxy public key.
* `source` - SSH proxy local key source type.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallSsh LocalKey can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallssh_localkey.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
