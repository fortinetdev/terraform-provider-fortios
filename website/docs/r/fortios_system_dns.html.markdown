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
* `domain` - Search suffix list for hostname lookup. The structure of `domain` block is documented below.
* `ip6_primary` - Primary DNS server IPv6 address.
* `ip6_secondary` - Secondary DNS server IPv6 address.
* `timeout` - DNS query timeout interval in seconds (1 - 10).
* `retry` - Number of times to retry (0 - 5).
* `dns_cache_limit` - Maximum number of records in the DNS cache.
* `dns_cache_ttl` - Duration in seconds that the DNS cache retains information.
* `cache_notfound_responses` - Enable/disable response from the DNS server when a record is not in cache.
* `source_ip` - IP address used by the DNS server as its source IP.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `domain` block supports:

* `domain` - DNS search domain list separated by space (maximum 8 domains)


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Dns can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_dns.labelname SystemDns
$ unset "FORTIOS_IMPORT_TABLE"
```
