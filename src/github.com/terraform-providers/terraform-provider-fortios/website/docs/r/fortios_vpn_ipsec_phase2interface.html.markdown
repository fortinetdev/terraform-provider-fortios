---
layout: "fortios"
page_title: "FortiOS: fortios_vpn_ipsec_phase2interface"
sidebar_current: "docs-fortios-resource-vpn-ipsec-phase2interface"
description: |-
  Provides a resource to use phase2-interface to add or edit a phase 2 configuration on a route-based (interface mode) IPsec tunnel.
---

# fortios_vpn_ipsec_phase2interface
Provides a resource to use phase2-interface to add or edit a phase 2 configuration on a route-based (interface mode) IPsec tunnel.

## Example Usage for Site to Site/Pre-shared Key
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_vpn_ipsec_phase1interface" "test1" {
	name = "001Test"
	type = "static"
	interface = "port2"
	peertype = "any"
	proposal = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
	comments = "VPN 001Test P1"
	wizard_type = "static-fortigate"
	remote_gw = "1.2.2.2"
	psksecret = "testscecret112233445566778899"
}
```

```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_vpn_ipsec_phase2interface" "test2" {
	name = "001Test"
	phase1name = "001Test"
	proposal = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
	comments = "VPN 001Test P2"
	src_addr_type = "name"
	dst_addr_type = "name"
	src_name = "HQ-toBranch_local"
	dst_name = "HQ-toBranch_remote"
}
```

## Example Usage for Site to Site/Signature
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_vpn_ipsec_phase1interface" "test1" {
	name = "001Test"
	type = "static"
	interface = "port2"
	proposal = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
	comments = "VPN 001Test P1"
	wizard_type = "static-fortigate"
	remote_gw = "1.2.2.2"
	psksecret = "testscecret112233445566778899"
	certificate = ["Fortinet_SSL_ECDSA384"]
	peertype = "peer"
	peerid = ""
	peer = "2b_peer"
	peergrp = ""                                      
}
```

```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_vpn_ipsec_phase2interface" "test2" {
	name = "001Test"
	phase1name = "001Test"
	proposal = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
	comments = "VPN 001Test P2"
	src_addr_type = "name"
	dst_addr_type = "name"
	src_name = "HQ_toBranch_local"
	dst_name = "HQ_toBranch_remote"
}
```

## Example Usage for Remote Access/Pre-shared Key
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_vpn_ipsec_phase1interface" "test1" {
	name = "001Test"
	type = "dynamic"
	interface = "port2"
	peertype = "any"
	proposal = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
	comments = "VPN 001Test P1"
	wizard_type = "dialup-forticlient"
	remote_gw = "0.0.0.0"
	psksecret = "testscecret112233445566778899"
	ipv4_split_include = "d_split"
	split_include_service = ""
	ipv4_split_exclude = ""
}
```

```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_vpn_ipsec_phase2interface" "test2" {
	name = "001Test"
	phase1name = "001Test"
	proposal = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
	comments = "VPN 001Test P2"
	src_addr_type = "subnet"
	src_start_ip = "0.0.0.0"
	src_end_ip = "0.0.0.0"
	src_subnet = "0.0.0.0 0.0.0.0"
	dst_addr_type = "subnet"
	dst_start_ip = "0.0.0.0"
	dst_end_ip = "0.0.0.0"
	dst_subnet = "0.0.0.0 0.0.0.0"
}
```

## Example Usage for Remote Access/Signature
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_vpn_ipsec_phase1interface" "test1" {
	name = "001Test"
	type = "dynamic"
	interface = "port2"
	peertype = "any"
	proposal = "aes128-sha256 aes256-sha256 aes128-sha1 aes256-sha1"
	comments = "VPN 001Test P1"
	wizard_type = "dialup-forticlient"
	remote_gw = "1.2.2.2"
	psksecret = "testscecret112233445566778899"
	certificate = ["Fortinet_SSL_ECDSA384"]
	peertype = "peer"
	peerid = ""
	peer = "2b_peer"
	peergrp = "",        
	ipv4_split_include = "d_split"
	split_include_service = ""
	ipv4_split_exclude = ""
}
```

```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_vpn_ipsec_phase2interface" "test2" {
	name = "001Test"
	phase1name = "001Test"
	proposal = "aes128-sha1 aes256-sha1 aes128-sha256 aes256-sha256 aes128gcm aes256gcm chacha20poly1305"
	comments = "VPN 001Test P2"
	src_addr_type = "subnet"
	src_start_ip = "0.0.0.0"
	src_end_ip = "0.0.0.0"
	src_subnet = "0.0.0.0 0.0.0.0"
	dst_addr_type = "subnet"
	dst_start_ip = "0.0.0.0"
	dst_end_ip = "0.0.0.0"
	dst_subnet = "0.0.0.0 0.0.0.0"
}
```

## Argument Reference
The following arguments are supported:
* `name` - (Required) IPsec tunnel name.
* `phase1name` - (Required) Phase 1 determines the options required for phase 2.
* `proposal` - Phase2 proposal.
* `src_addr_type` - Local proxy ID type.
* `src_start_ip` - Local proxy ID start.
* `src_end_ip` - Local proxy ID end.
* `src_subnet` - Local proxy ID subnet.
* `dst_addr_type` - Local proxy ID type.
* `src_name` - Local proxy ID name. 
* `dst_name` - Remote proxy ID name.
* `dst_start_ip` - Remote proxy ID IPv4 start.
* `dst_end_ip` - Remote proxy ID IPv4 end.
* `dst_subnet` - Remote proxy ID IPv4 subnet.
* `comments` - Comment.

## Attributes Reference
The following attributes are exported:
* `id` - The ID of the phase2-interface.
* `name` - IPsec tunnel name.
* `phase1name` - Phase 1 determines the options required for phase 2.
* `proposal` - Phase2 proposal.
* `src_addr_type` - Local proxy ID type.
* `src_start_ip` - Local proxy ID start.
* `src_end_ip` - Local proxy ID end.
* `src_subnet` - Local proxy ID subnet.
* `dst_addr_type` - Local proxy ID type.
* `src_name` - Local proxy ID name. 
* `dst_name` - Remote proxy ID name.
* `dst_start_ip` - Remote proxy ID IPv4 start.
* `dst_end_ip` - Remote proxy ID IPv4 end.
* `dst_subnet` - Remote proxy ID IPv4 subnet.
* `comments` - Comment.

