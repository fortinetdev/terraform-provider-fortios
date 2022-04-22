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

* `name` - SSH public key name.
* `status` - Set the trust status of the public key. Valid values: `trusted`, `revoked`.
* `type` - Set the type of the public key. Valid values: `RSA`, `DSA`, `ECDSA`, `ED25519`, `RSA-CA`, `DSA-CA`, `ECDSA-CA`, `ED25519-CA`.
* `nid` - Set the nid of the ECDSA key. Valid values: `256`, `384`, `521`.
* `usage` - Usage for this public key. Valid values: `transparent-proxy`, `access-proxy`.
* `ip` - IP address of the SSH server.
* `port` - Port of the SSH server.
* `hostname` - Hostname of the SSH server.
* `public_key` - SSH public key.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallSsh HostKey can be imported using any of these accepted formats:
```
$ terraform import fortios_firewallssh_hostkey.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallssh_hostkey.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
