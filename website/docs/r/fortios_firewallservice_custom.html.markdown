---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_custom"
description: |-
  Configure custom services.
---

# fortios_firewallservice_custom
Configure custom services.

## Example Usage

```hcl
resource "fortios_firewallservice_custom" "trname" {
  app_service_type    = "disable"
  category            = "General"
  check_reset_range   = "default"
  color               = 0
  helper              = "auto"
  iprange             = "0.0.0.0"
  name                = "sservice_custom2"
  protocol            = "TCP/UDP/SCTP"
  protocol_number     = 6
  proxy               = "disable"
  tcp_halfclose_timer = 0
  tcp_halfopen_timer  = 0
  tcp_portrange       = "223-332"
  tcp_timewait_timer  = 0
  udp_idle_timer      = 0
  visibility          = "enable"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Custom service name.
* `proxy` - Enable/disable web proxy service.
* `category` - Service category.
* `protocol` - Protocol type based on IANA numbers.
* `helper` - Helper name.
* `iprange` - Start and end of the IP range associated with service.
* `fqdn` - Fully qualified domain name.
* `protocol_number` - IP protocol number.
* `icmptype` - ICMP type.
* `icmpcode` - ICMP code.
* `tcp_portrange` - Multiple TCP port ranges.
* `udp_portrange` - Multiple UDP port ranges.
* `sctp_portrange` - Multiple SCTP port ranges.
* `tcp_halfclose_timer` - Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
* `tcp_halfopen_timer` - Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
* `tcp_timewait_timer` - Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
* `udp_idle_timer` - UDP half close timeout (0 - 86400 sec, 0 = default).
* `session_ttl` - Session TTL (300 - 604800, 0 = default).
* `check_reset_range` - Configure the type of ICMP error message verification.
* `comment` - Comment.
* `color` - Color of icon on the GUI.
* `visibility` - Enable/disable the visibility of the service on the GUI.
* `app_service_type` - Application service type.
* `app_category` - Application category ID. The structure of `app_category` block is documented below.
* `application` - Application ID. The structure of `application` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `app_category` block supports:

* `id` - Application category id.

The `application` block supports:

* `id` - Application id.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

FirewallService Custom can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewallservice_custom.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
