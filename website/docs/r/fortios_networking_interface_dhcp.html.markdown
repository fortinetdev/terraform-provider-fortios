---
layout: "fortios"
page_title: "FortiOS: fortios_networking_interface_dhcp"
sidebar_current: "docs-fortios-resource-networking-interface-dhcp"
description: |-
  Provides a resource to configure interface DHCP settings of FortiOS.
---

# fortios_networking_interface_dhcp
Provides a resource to configure interface DHCP settings of FortiOS.

## Example Usage for vlan Interface
```hcl

provider "fortios" {
  hostname = "54.226.179.231"
  token    = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"
  vdom     = "root"
}


resource "fortios_networking_interface_dhcp" "dhcp_server" {
  interface    = "TestIntf"
  ip_range     = "10.10.10.10-10.10.10.100"
  default_gw   = "10.10.10.1"
  netmask      = "255.255.255.0"
  vdom         = "root"
  dns_service  = "default"
  dns_server_1 = "0.0.0.0"
  dns_server_2 = "0.0.0.0"
  dns_server_3 = "0.0.0.0"
  lease_time   = 500
}

```

## Argument Reference
The following arguments are supported:

* `interface` - (Required) Interface name to attache dhcp service.
* `ip_range` - (Required) IP Range to hand out from dhcp service. Syntax X.X.X.X-Y.Y.Y.Y where X is start.
* `netmask` - (Required) Netmask set on the IPs handed out from dhcp service.
* `default_gw` - Default gateway to hand out, default would be interface assigned gw.
* `dns_service` - DNS Service - Options default == Same as System DNS, local == Same as interface, specify == Assign DNS servers
* `dns_server_1` - DNS Server number 1 if dns_service = "specify"
* `dns_server_2` - DNS Server number 2 if dns_service = "specify"
* `dns_server_3` - DNS Server number 3 if dns_service = "specify"
* `lease_time` - IP Lease time, default value 604800


## Attributes Reference
The following attributes are exported:

* `id` - The Name of the interface.
* `status` - Service status, enable or disable
* `vdom` - Servic is in this virtual domain (VDOM).