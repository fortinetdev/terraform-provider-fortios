---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewallservice_custom"
description: |-
  Get information on an fortios firewallservice custom.
---

# Data Source: fortios_firewallservice_custom
Use this data source to get information on an fortios firewallservice custom

## Argument Reference

* `name` - (Required) Specify the name of the desired firewallservice custom.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Custom service name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
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
* `udplite_portrange` - Multiple UDP-Lite port ranges.
* `sctp_portrange` - Multiple SCTP port ranges.
* `tcp_halfclose_timer` - Wait time to close a TCP session waiting for an unanswered FIN packet (1 - 86400 sec, 0 = default).
* `tcp_halfopen_timer` - Wait time to close a TCP session waiting for an unanswered open session packet (1 - 86400 sec, 0 = default).
* `tcp_timewait_timer` - Set the length of the TCP TIME-WAIT state in seconds (1 - 300 sec, 0 = default).
* `tcp_rst_timer` - Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
* `udp_idle_timer` - UDP half close timeout (0 - 86400 sec, 0 = default).
* `session_ttl` - Session TTL (300 - 604800, 0 = default).
* `check_reset_range` - Configure the type of ICMP error message verification.
* `comment` - Comment.
* `color` - Color of icon on the GUI.
* `visibility` - Enable/disable the visibility of the service on the GUI.
* `app_service_type` - Application service type.
* `app_category` - Application category ID. The structure of `app_category` block is documented below.
* `application` - Application ID. The structure of `application` block is documented below.
* `fabric_object` - Security Fabric global object setting.

The `app_category` block contains:

* `id` - Application category id.

The `application` block contains:

* `id` - Application id.

