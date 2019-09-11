---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_system_network_interface"
sidebar_current: "docs-fortios-fortimanager-resource-system-network-interface"
description: |-
  Provides a resource to configure system network interface for FortiManager.
---

# fortios_fortimanager_system_network_interface
This resource supports updating system network interface for FortiManager.

## Example Usage
```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	provider = "fortimanager"
}

resource "fortios_fortimanager_system_network_interface" "test1" {
	name = "port3"
	ip = "1.1.1.3 255.255.255.0"
	description = ""
	status = "up"
	allow_access = ["ping","ssh","https", "http"]
	service_access = ["webfilter","fgtupdates"]
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Interface port.
* `ip` - Interface Ipaddress.
* `description` - Description.
* `status` - Interface status, Enum: ["down", "up"]
* `allow_access` - Allow managerment access to interface, Enum: ["ping", "https", "ssh", "snmp", "telnet", "http", "web"]
* `sevice_access` - Allow service access to interface, Enum: ["fgtupdates", "webfilter"]

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Interface port.
* `ip` - Interface Ipaddress.
* `description` - Description.
* `status` - Interface status.
* `allow_access` - Allow managerment access to interface.
* `sevice_access` - Allow service access to interface.
