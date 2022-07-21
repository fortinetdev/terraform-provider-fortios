---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnssl_client"
description: |-
  client
---

# fortios_vpnssl_client
client Applies to FortiOS Version `>= 7.0.1`.

## Argument Reference

The following arguments are supported:

* `name` - SSL-VPN tunnel name.
* `comment` - Comment.
* `interface` - SSL interface to send/receive traffic over.
* `user` - Username to offer to the peer to authenticate the client.
* `psk` - Pre-shared secret to authenticate with the server (ASCII string or hexadecimal encoded with a leading 0x).
* `peer` - Authenticate peer's certificate with the peer/peergrp.
* `server` - IPv4, IPv6 or DNS address of the SSL-VPN server.
* `port` - SSL-VPN server port.
* `realm` - Realm name configured on SSL-VPN server.
* `status` - Enable/disable this SSL-VPN client configuration. Valid values: `enable`, `disable`.
* `certificate` - Certificate to offer to SSL-VPN server if it requests one.
* `source_ip` - IPv4 or IPv6 address to use as a source for the SSL-VPN connection to the server.
* `distance` - Distance for routes added by SSL-VPN (1 - 255).
* `priority` - Priority for routes added by SSL-VPN (0 - 4294967295).
* `class_id` - Traffic class ID.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

VpnSsl Client can be imported using any of these accepted formats:
```
$ terraform import fortios_vpnssl_client.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_vpnssl_client.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
