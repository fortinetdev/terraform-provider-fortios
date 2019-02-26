---
layout: "fortios"
page_title: "FortiOS: fortios_firewall_object_service"
sidebar_current: "docs-fortios-resource-firewall-object-service"
description: |-
  Provides a resource to configure firewall service of FortiOS.
---

# fortios_firewall_object_service
Provides a resource to configure firewall service of FortiOS.

## Example Usage for Fqdn Service
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_firewall_object_service" "v11" {
	name = "a23322"
	category = "General"
	protocol = "TCP/UDP/SCTP"
	fqdn = "fdsafdsa"
	comment = "comment"
}
```

## Example Usage for Iprange Service
```hcl
provider "fortios" {
	hostname = "54.226.179.231"
	token = "jn3t3Nw7qckQzt955Htkfj5hwQ6jdb"	
}

resource "fortios_firewall_object_service" "v13" {
	name = "1fdsafd11a"
	category = "General"
	protocol = "TCP/UDP/SCTP"
	iprange = "1.1.1.1-2.2.2.2"
	comment = "comment"
}
```

## Argument Reference
The following arguments are supported:
* `name` - (Required) Number of minutes before an idle administrator session time out.
* `category` - (Required) Service category.
* `protocol` - Protocol type based on IANA numbers.
* `fqdn` - Fully qualified domain name.
* `iprange` - Start and end of the IP range associated with service.
* `comment` - Comment.

## Attributes Reference
The following attributes are exported:
* `id` - The ID of the firewall service item.
* `name` - Number of minutes before an idle administrator session time out.
* `category` - Service category.
* `protocol` - Protocol type based on IANA numbers.
* `fqdn` - Fully qualified domain name.
* `iprange` - Start and end of the IP range associated with service.
* `comment` - Comment.

