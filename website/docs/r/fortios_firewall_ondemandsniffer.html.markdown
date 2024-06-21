---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ondemandsniffer"
description: |-
  Configure on-demand packet sniffer.
---

# fortios_firewall_ondemandsniffer
Configure on-demand packet sniffer. Applies to FortiOS Version `>= 7.4.4`.

## Argument Reference

The following arguments are supported:

* `name` - On-demand packet sniffer name.
* `interface` - Interface name that on-demand packet sniffer will take place.
* `max_packet_count` - Maximum number of packets to capture per on-demand packet sniffer.
* `hosts` - IPv4 or IPv6 hosts to filter in this traffic sniffer. The structure of `hosts` block is documented below.
* `ports` - Ports to filter for in this traffic sniffer. The structure of `ports` block is documented below.
* `protocols` - Protocols to filter in this traffic sniffer. The structure of `protocols` block is documented below.
* `non_ip_packet` - Include non-IP packets. Valid values: `enable`, `disable`.
* `advanced_filter` - Advanced freeform filter that will be used over existing filter settings if set. Can only be used by super admin.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `hosts` block supports:

* `host` - IPv4 or IPv6 host.

The `ports` block supports:

* `port` - Port to filter in this traffic sniffer.

The `protocols` block supports:

* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

Firewall OnDemandSniffer can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_ondemandsniffer.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_ondemandsniffer.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
