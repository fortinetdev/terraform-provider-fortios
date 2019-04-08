---
layout: "fortios"
page_title: "FortiOS: fortios_networking_interface_port"
sidebar_current: "docs-fortios-resource-networking-interface-port"
description: |-
  Provides a resource to configure interface settings of FortiOS.
---

# fortios_networking_interface_port
Provides a resource to configure interface settings of FortiOS.

## Example Usage for Loopback Interface
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_networking_interface_port" "loopback1" {
	ip = "23.123.33.10/24"
	allowaccess = "ping http"
	alias = "cc1"
	description = "description"
	status = "up"
	role = "lan"
	name = "myinterface1"
	vdom = "root"
	type = "loopback"
	mode = "static"
}
```

## Example Usage for VLAN Interface
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_networking_interface_port" "vlan1" {
	role = "lan"
	mode = "static"
	defaultgw = "enable"
	distance = "33"
	type = "vlan"
	vlanid = "3"
	name = "myinterface2"
	vdom = "root"
	ip = "3.123.33.10/24"
	interface = "port2"
	allowaccess = "ping"
}
```

## Example Usage for Physical Interface
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_networking_interface_port" "test1" {
	name = "port2"
	ip = "93.133.133.110/24"
	alias = "dkeeew"
	description = "description"
	status = "up"
	device_identification = "enable"
	tcp_mss = "3232"
	speed = "auto"
	mtu_override = "enable"
	mtu = "2933"
	role = "lan"
	allowaccess = "ping https"
	mode = "static"
	dns_server_override = "enable"
	defaultgw = "enable"
	distance = "33"
	type = "physical"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) If the interface is physical, the argument is the name of the interface.
* `type` - (Required) Interface type (support physical, vlan, loopback).
* `ip` - Interface IPv4 address and subnet mask, syntax` - X.X.X.X/24.
* `alias` - Alias will be displayed with the interface name to make it easier to distinguish.
* `status` - Bring the interface up or shut the interface down.
* `device_identification` - Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
* `tcp_mss` - TCP maximum segment size. 0 means do not change segment size.
* `speed` - Interface speed. The default setting and the options available depend on the interface hardware.
* `mtu_override` - Enable to set a custom MTU for this interface.
* `mtu` - MTU value for this interface.
* `role` - Interface role.
* `allowaccess` - Permitted types of management access to this interface.
* `mode` - (Required) Addressing mode.
* `dns_server_override` - Enable/disable use DNS acquired by DHCP or PPPoE.
* `defaultgw` - Enable to get the gateway IP from the DHCP or PPPoE server.
* `distance` - Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
* `description` - Description.
* `interface` - Interface name.
* `name` - Name.
* `vdom` - Interface is in this virtual domain (VDOM).
* `vlanid` - VLAN ID.

## Attributes Reference
The following attributes are exported:

* `id` - The Name of the interface.
* `ip` - Interface IPv4 address and subnet mask, syntax` - X.X.X.X/24.
* `alias` - Alias will be displayed with the interface name to make it easier to distinguish.
* `status` - Bring the interface up or shut the interface down.
* `device_identification` - Enable/disable passively gathering of device identity information about the devices on the network connected to this interface.
* `tcp_mss` - TCP maximum segment size. 0 means do not change segment size.
* `speed` - Interface speed. The default setting and the options available depend on the interface hardware.
* `mtu_override` - Enable to set a custom MTU for this interface.
* `mtu` - MTU value for this interface.
* `role` - Interface role.
* `allowaccess` - Permitted types of management access to this interface.
* `mode` - Addressing mode.
* `dns_server_override` - Enable/disable use DNS acquired by DHCP or PPPoE.
* `defaultgw` - Enable to get the gateway IP from the DHCP or PPPoE server.
* `distance` - Distance for routes learned through PPPoE or DHCP, lower distance indicates preferred route.
* `description` - Description.
* `type` - Interface type (support physical, vlan, loopback).
* `interface` - Interface name.
* `name` - Name.
* `vdom` - Interface is in this virtual domain (VDOM).
* `vlanid` - VLAN ID.

