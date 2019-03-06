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
	name = "servicetest"
	category = "General"
	protocol = "TCP/UDP/SCTP"
	fqdn = "abc.com"
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
	tcp_portrange = "22-33"
	udp_portrange = "44-55"
	sctp_portrange = "66-88"
	comment = "comment"
}
```

## Example Usage for ICMP Service
```hcl
resource "fortios_firewall_object_service" "ICMP" {
	name = "ICMPService"
	category = "General"
	protocol = "ICMP"
	fqdn = "www.baidu.com"
	icmptype = "2"
	icmpcode = "3"
	protocol_number = "1"
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
* `tcp_portrange` - Multiple TCP port ranges.
* `udp_portrange` - Multiple UDP port ranges.
* `sctp_portrange` - Multiple SCTP port ranges.
* `icmptype` - ICMP type.
* `icmpcode` - ICMP code.
* `protocol_number` - IP protocol number.
* `comment` - Comment.

## Attributes Reference
The following attributes are exported:
* `id` - The ID of the firewall service item.
* `name` - Number of minutes before an idle administrator session time out.
* `category` - Service category.
* `protocol` - Protocol type based on IANA numbers.
* `fqdn` - Fully qualified domain name.
* `iprange` - Start and end of the IP range associated with service.
* `tcp_portrange` - Multiple TCP port ranges.
* `udp_portrange` - Multiple UDP port ranges.
* `sctp_portrange` - Multiple SCTP port ranges.
* `icmptype` - ICMP type.
* `icmpcode` - ICMP code.
* `protocol_number` - IP protocol number.
* `comment` - Comment.

