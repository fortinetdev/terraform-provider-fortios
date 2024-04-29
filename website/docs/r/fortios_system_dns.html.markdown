---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dns"
description: |-
  Configure DNS.
---

# fortios_system_dns
Configure DNS.

## Example Usage

```hcl
resource "fortios_system_dns" "trname" {
  cache_notfound_responses = "disable"
  dns_cache_limit          = 5000
  dns_cache_ttl            = 1800
  ip6_primary              = "::"
  ip6_secondary            = "::"
  primary                  = "208.91.112.53"
  retry                    = 2
  secondary                = "208.91.112.51"
  source_ip                = "0.0.0.0"
  timeout                  = 5
}
```

## Argument Reference

The following arguments are supported:

* `primary` - (Required) Primary DNS server IP address.
* `secondary` - Secondary DNS server IP address.
* `protocol` - DNS protocols. Valid values: `cleartext`, `dot`, `doh`.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS. Valid values: `disable`, `enable`, `enforce`.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `server_hostname` - DNS server host name list. The structure of `server_hostname` block is documented below.
* `domain` - Search suffix list for hostname lookup. The structure of `domain` block is documented below.
* `ip6_primary` - Primary DNS server IPv6 address.
* `ip6_secondary` - Secondary DNS server IPv6 address.
* `timeout` - DNS query timeout interval in seconds (1 - 10).
* `retry` - Number of times to retry (0 - 5).
* `dns_cache_limit` - Maximum number of records in the DNS cache.
* `dns_cache_ttl` - Duration in seconds that the DNS cache retains information.
* `cache_notfound_responses` - Enable/disable response from the DNS server when a record is not in cache. Valid values: `disable`, `enable`.
* `source_ip` - IP address used by the DNS server as its source IP.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `server_select_method` - Specify how configured servers are prioritized. Valid values: `least-rtt`, `failover`.
* `alt_primary` - Alternate primary DNS server. This is not used as a failover DNS server.
* `alt_secondary` - Alternate secondary DNS server. This is not used as a failover DNS server.
* `log` - Local DNS log setting. Valid values: `disable`, `error`, `all`.
* `fqdn_cache_ttl` - FQDN cache time to live in seconds (0 - 86400, default = 0).
* `fqdn_max_refresh` - FQDN cache maximum refresh time in seconds (3600 - 86400, default = 3600).
* `fqdn_min_refresh` - FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `server_hostname` block supports:

* `hostname` - DNS server host name list separated by space (maximum 4 domains).

The `domain` block supports:

* `domain` - DNS search domain list separated by space (maximum 8 domains)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Dns can be imported using any of these accepted formats:
```
$ terraform import fortios_system_dns.labelname SystemDns

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_dns.labelname SystemDns
$ unset "FORTIOS_IMPORT_TABLE"
```
