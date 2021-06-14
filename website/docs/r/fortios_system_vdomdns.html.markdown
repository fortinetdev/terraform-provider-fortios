---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomdns"
description: |-
  Configure DNS servers for a non-management VDOM.
---

# fortios_system_vdomdns
Configure DNS servers for a non-management VDOM.

## Argument Reference

The following arguments are supported:

* `vdom_dns` - Enable/disable configuring DNS servers for the current VDOM. Valid values: `enable`, `disable`.
* `primary` - Primary DNS server IP address for the VDOM.
* `secondary` - Secondary DNS server IP address for the VDOM.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `server_hostname` - DNS server host name list. The structure of `server_hostname` block is documented below.
* `ip6_primary` - Primary IPv6 DNS server IP address for the VDOM.
* `ip6_secondary` - Secondary IPv6 DNS server IP address for the VDOM.
* `source_ip` - Source IP for communications with the DNS server.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `server_hostname` block supports:

* `hostname` - DNS server host name list separated by space (maximum 4 domains).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System VdomDns can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_vdomdns.labelname SystemVdomDns
$ unset "FORTIOS_IMPORT_TABLE"
```
