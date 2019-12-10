---
layout: "fortios"
page_title: "FortiOS: fortios_networking_route_static"
sidebar_current: "docs-fortios-resource-networking-route-static"
description: |-
  Provides a resource to configure static route of FortiOS.
---

# fortios_networking_route_static
Provides a resource to configure static route of FortiOS.

## Example Usage
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_networking_route_static" "subnet" {
	dst = "110.2.2.122/32"
	gateway = "2.2.2.2"
	blackhole = "disable"
	distance = "22"
	weight = "3"
	priority = "3"
	device = "port2"
	comment = "Terraform test"
	status = "enable"
}

resource "fortios_networking_route_static" "internet_service" {
	internet_service = 5242881
	gateway = "2.2.2.2"
	blackhole = "disable"
	distance = "22"
	weight = "3"
	priority = "3"
	device = "port2"
	comment = "Terraform Test"
	status = "enable"
}
```

## Argument Reference
The following arguments are supported:

* `dst` - (Required) Destination IP and mask for this route.
* `gateway` - Gateway IP for this route.
* `blackhole` - Enable/disable black hole.
* `distance` - Administrative distance.
* `weight` - Administrative weight.
* `priority` - Administrative priority.
* `device` - (Required) Gateway out interface or tunnel.
* `comment` - Optional comments.
* `internet_service` - Application ID in the Internet service database.
* `status` - Enable/disable this static route. default is "enable".

## Attributes Reference
The following attributes are exported:

* `id` - The ID of the static route item.
* `dst` - Destination IP and mask for this route.
* `gateway` - Gateway IP for this route.
* `blackhole` - Enable/disable black hole.
* `distance` - Administrative distance.
* `weight` - Administrative weight.
* `priority` - Administrative priority.
* `device` - Gateway out interface or tunnel.
* `comment` - Optional comments.
* `internet_service` - Application ID in the Internet service database.
* `status` - Enable/disable this static route.
