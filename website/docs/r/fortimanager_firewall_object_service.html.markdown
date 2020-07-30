---
layout: "fortios"
page_title: "FortiOS: fortios_fmg_firewall_object_service"
sidebar_current: "docs-fortios-fortimanager-resource-firewall-object-service"
subcategory: "FortiManager"
description: |-
  Provides a resource to configure firewall object service for FortiManager.
---

# fortios_fmg_firewall_object_service
This resource supports Create/Read/Update/Delete firewall object service for FortiManager.

## Example Usage
```hcl
resource "fortios_fmg_firewall_object_service" "test1" {
  name           = "fos_test"
  comment        = "test obj service"
  protocol       = "TCP/UDP/SCTP"
  category       = "Email"
  iprange        = "1.1.1.1"
  tcp_portrange  = ["100-200:150-250"]
  udp_portrange  = ["100-200:150-250"]
  sctp_portrange = ["100-200:150-250"]
}

resource "fortios_fmg_firewall_object_service" "test2" {
  name      = "fos_test2"
  comment   = "test obj service"
  protocol  = "ICMP"
  category  = "Web Access"
  icmp_type = 2
  icmp_code = 3
}

resource "fortios_fmg_firewall_object_service" "test3" {
  name            = "fos_test3"
  comment         = "test obj service"
  protocol        = "IP"
  category        = "File Access"
  protocol_number = 4
}
```

## Argument Reference
The following arguments are supported:

* `name` - (Required) Custom service name.
* `comment` - Comments.
* `protocol` - Protocol type. Enum: ["TCP/UDP/SCTP", "ICMP", "ICMP6", "IP"]
* `category` - Service category. Enum: ["", "File Access", "Authentication", "Email", "General", "Network Services", "Remote Access", "Tunneling", "VoIP, Messaging & Other Applications", "Web Access", "Web Proxy"]
* `fqdn` - Fully qualified domain name.
* `iprange` - Start and end of the IP range associated with service. Ip or Ip range(eg: 1.1.1.1-1.1.1.100).
* `tcp_portrange` - TCP port ranges. e.g: ["dst-port-range:src-port-range"]
* `udp_portrange` - UDP port ranges. e.g: ["dst-port-range:src-port-range"]
* `sctp_portrange` - SCTP port ranges. e.g: ["dst-port-range:src-port-range"]
* `icmp_type` - ICMP Type.
* `icmp_code` - ICMP Code.
* `protocol_number` - IP protocol number.
* `adom` - ADOM name. default is 'root'.

## Attributes Reference
The following attributes are exported:

* `id` - The resource id.
* `name` - Custom service name.
* `comment` - Comments.
* `protocol` - Protocol type.
* `category` - Service category.
* `fqdn` - Fully qualified domain name.
* `iprange` - Start and end of the IP range associated with service.
* `tcp_portrange` - TCP port ranges.
* `udp_portrange` - UDP port ranges.
* `sctp_portrange` - SCTP port ranges.
* `icmp_type` - ICMP Type.
* `icmp_code` - ICMP Code.
* `protocol_number` - IP protocol number.
* `adom` - ADOM name.
