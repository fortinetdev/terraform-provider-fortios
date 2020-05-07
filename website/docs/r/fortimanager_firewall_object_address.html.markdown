---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_firewall_object_address"
sidebar_current: "docs-fortios-fortimanager-resource-firewall-object-address"
description: |-
  Provides a resource to configure firewall object address for FortiManager.
---

# fortios_fmg_firewall_object_address
This resource supports Create/Read/Update/Delete firewall object address for FortiManager.

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

resource "fortios_fmg_firewall_object_address" "test1" {
	name = "foa_test"
	type = "fqdn"
	comment = "test obj address"
	fqdn = "fqdn.google.com"
	associated_intf = "any"
}

resource "fortios_fmg_firewall_object_address" "test2" {
	name = "foa_test2"
	type = "ipmask"
	comment = "test obj address"
	subnet = "2.2.2.0 255.255.255.0"
	associated_intf = "any"
	allow_routing = "disable"
}

resource "fortios_fmg_firewall_object_address" "test3" {
	name = "foa_test3"
	type = "iprange"
	comment = "test obj address"
	associated_intf = "any"
	start_ip = "2.2.2.1"
	end_ip = "2.2.2.100"
}

```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Address name.
* `type` - Type of address, Enum: ["ipmask", "iprange", "fqdn"].
* `comment` - Comments.
* `fqdn` - Fully Qualified Domain Name address.
* `associated_intf` - Network interface associated with address.
* `subnet` - IPv4 address/mask
* `start_ip` - First IP address (inclusive) in the range for the address.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `allow_routing` - Enable/disable use of this address in the static route configuration. default is "disable".
* `adom` - ADOM name. default is 'root'.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Address name.
* `type` - Type of address.
* `comment` - Comments.
* `fqdn` - Fully Qualified Domain Name address.
* `associated_intf` - Network interface associated with address.
* `subnet` - IPv4 address/mask
* `start_ip` - First IP address (inclusive) in the range for the address.
* `end_ip` - Final IP address (inclusive) in the range for the address.
* `allow_routing` - Enable/disable use of this address in the static route configuration.
* `adom` - ADOM name.
