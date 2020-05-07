---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_system_network_route"
sidebar_current: "docs-fortios-fortimanager-resource-system-network-route"
description: |-
  Provides a resource to configure system network route for FortiManager.
---

# fortios_fmg_system_network_route
This resource supports updating system network route for FortiManager.

## Example Usage
```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	product = "fortimanager"
	insecure = false
	cabundlefile = "/path/yourCA.crt"
}

resource "fortios_fmg_system_network_route" "test1" {
	route_id = 5
	destination = "192.168.2.0 255.255.255.0"
	gateway = "192.168.2.1"
	device = "port4"
}
```

## Argument Reference
The following arguments are supported:

* `route_id` - (Required) Route id.
* `destination` - (Required) Destination Ip/Mask.
* `gateway` - (Required) Gateway Ip.
* `device` - (Required) Gateway out interface.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `route_id` - Route id.
* `destination` - Destination Ip/Mask.
* `gateway` - Gateway Ip.
* `device` - Gateway out interface.
