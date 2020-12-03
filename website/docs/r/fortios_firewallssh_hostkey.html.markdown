---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallssh_hostkey"
description: |-
  SSH proxy host public keys.
---

# fortios_firewallssh_hostkey
SSH proxy host public keys.

## Example Usage

```hcl
resource "fortios_firewallssh_hostkey" "trname" {
  hostname = "testmachine"
  ip       = "1.1.1.1"
  name     = "hostkeys1"
  port     = 22
  status   = "trusted"
  type     = "RSA"
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) SSH public key name.
* `status` - Set the trust status of the public key.
* `type` - Set the type of the public key.
* `nid` - Set the nid of the ECDSA key.
* `ip` - IP address of the SSH server.
* `port` - Port of the SSH server.
* `hostname` - Hostname of the SSH server.
* `public_key` - SSH public key.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallSsh HostKey can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallssh_hostkey.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
