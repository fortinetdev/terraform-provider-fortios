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
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `proxy` - Enable/disable web proxy service. Valid values: `enable`, `disable`.
* `category` - Service category.
* `protocol` - Protocol type based on IANA numbers. Valid values: `TCP/UDP/SCTP`, `ICMP`, `ICMP6`, `IP`, `HTTP`, `FTP`, `CONNECT`, `SOCKS-TCP`, `SOCKS-UDP`, `ALL`.
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
* `tcp_rst_timer` - Set the length of the TCP CLOSE state in seconds (5 - 300 sec, 0 = default).
* `udp_idle_timer` - Number of seconds before an idle UDP connection times out (0 - 86400 sec, 0 = default).
* `session_ttl` - Session TTL (300 - 604800, 0 = default).
* `check_reset_range` - Configure the type of ICMP error message verification. Valid values: `disable`, `strict`, `default`.
* `comment` - Comment.
* `color` - Color of icon on the GUI.
* `visibility` - Enable/disable the visibility of the service on the GUI. Valid values: `enable`, `disable`.
* `app_service_type` - Application service type. Valid values: `disable`, `app-id`, `app-category`.
* `app_category` - Application category ID. The structure of `app_category` block is documented below.
* `application` - Application ID. The structure of `application` block is documented below.
* `fabric_object` - Security Fabric global object setting. Valid values: `enable`, `disable`.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

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
$ terraform import fortios_firewallservice_custom.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewallservice_custom.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
