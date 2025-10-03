---
subcategory: "FortiGate VPN"
layout: "fortios"
page_title: "FortiOS: fortios_vpnssl_client"
description: |-
  client
---

# fortios_vpnssl_client
client Applies to FortiOS Version `7.0.1,7.0.2,7.0.3,7.0.4,7.0.5,7.0.6,7.0.7,7.0.8,7.0.9,7.0.10,7.0.11,7.0.12,7.0.13,7.0.14,7.0.15,7.0.16,7.0.17,7.2.0,7.2.1,7.2.2,7.2.3,7.2.4,7.2.6,7.2.7,7.2.8,7.2.9,7.2.10,7.2.11,7.2.12,7.4.0,7.4.1,7.4.2,7.4.3,7.4.4,7.4.5,7.4.6,7.4.7,7.4.8,7.6.0,7.6.1,7.6.2`.

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
* `priority` - Priority for routes added by SSL-VPN. On FortiOS versions 7.0.1-7.0.3: 0 - 4294967295. On FortiOS versions 7.0.4-7.6.2: 1 - 65535.
* `class_id` - Traffic class ID.
* `ipv4_subnets` - IPv4 subnets that the client is protecting.
* `ipv6_subnets` - IPv6 subnets that the client is protecting.
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
