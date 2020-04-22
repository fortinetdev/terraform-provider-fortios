---
layout: "fortios"
page_title: "FortiOS: fortios_networking_interface_dhcp_ipres"
sidebar_current: "docs-fortios-resource-networking-interface-dhcp-ipres"
description: |-
  Provides a resource to configure interface DHCP IP Reservation settings of FortiOS.
---

# fortios_networking_interface_dhcp_ipres
Provides a resource to configure interface DHCP IP Reservation settings of FortiOS.

## Example Usage
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

resource "fortios_networking_interface_dhcp_ipres" "reserv_1" {
  interface   = fortios_networking_interface_dhcp.dhcp_server.interface
  ip          = "10.10.10.12"
  mac         = "82:05:27:89:50:01"
  description = "Terraform Provisioned Reservation"
}

```

## Argument Reference
The following arguments are supported:

* `interface` - (Required) Interface name where dhcp service is running.
* `ip` - (Required) IP to hand out from dhcp service. 
* `mac` - (Required) MAC adrress for the reservation.
* `description` - Optional Description


## Attributes Reference
The following attributes are exported:

* `id` - The Name of the interface.
