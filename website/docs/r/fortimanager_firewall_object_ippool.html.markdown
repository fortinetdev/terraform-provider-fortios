---
layout: "fortios"
page_title: "FortiOS: fortios_fortimanager_firewall_object_ippool"
sidebar_current: "docs-fortios-fortimanager-resource-firewall-object-ippool"
description: |-
  Provides a resource to configure firewall object ippool for FortiManager.
---

# fortios_fortimanager_firewall_object_ippool
This resource supports Create/Read/Update/Delete firewall object ippool for FortiManager.

## Example Usage
```hcl
provider "fortios" {
	hostname = "192.168.88.100"
	username = "APIUser"
	passwd = "admin"
	provider = "fortimanager"
}

resource "fortios_fortimanager_firewall_object_ippool" "test1" {
	name = "foi_test"
	startip = "1.1.10.1"
	endip = "1.1.10.100"
	comment = "test obj ippool"
	type = "one-to-one"
	arp_intf = "any"
	arp_reply = "enable"
	associated_intf = "any"
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Ippool name.
* `startip` - (Required) First IPv4 address (inclusive) in the range for the address pool.
* `endip` - (Required) Final IPv4 address (inclusive) in the range for the address pool.
* `comment` - Comments.
* `type` - Ip pool type, Enum: ["overload", "one-to-one"].
* `arp_intf` - Select an interface that will reply to ARP requests.
* `arp_reply` - Enable/disable replying to ARP request, default is "enable".
* `associated_intf` - Associated interface name.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Ippool name.
* `startip` - First IPv4 address (inclusive) in the range for the address pool.
* `endip` -  Final IPv4 address (inclusive) in the range for the address pool.
* `comment` - Comments.
* `type` - Ip pool type, Enum: ["overload", "one-to-one"].
* `arp_intf` - Select an interface that will reply to ARP requests.
* `arp_reply` - Enable/disable replying to ARP request.
* `associated_intf` - Associated interface name.
