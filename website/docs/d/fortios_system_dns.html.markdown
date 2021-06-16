---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dns"
description: |-
  Get information on fortios system dns.
---

# Data Source: fortios_system_dns
Use this data source to get information on fortios system dns

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `primary` - Primary DNS server IP address.
* `secondary` - Secondary DNS server IP address.
* `dns_over_tls` - Enable/disable/enforce DNS over TLS.
* `ssl_certificate` - Name of local certificate for SSL connections.
* `server_hostname` - DNS server host name list. The structure of `server_hostname` block is documented below.
* `domain` - Search suffix list for hostname lookup. The structure of `domain` block is documented below.
* `ip6_primary` - Primary DNS server IPv6 address.
* `ip6_secondary` - Secondary DNS server IPv6 address.
* `timeout` - DNS query timeout interval in seconds (1 - 10).
* `retry` - Number of times to retry (0 - 5).
* `dns_cache_limit` - Maximum number of records in the DNS cache.
* `dns_cache_ttl` - Duration in seconds that the DNS cache retains information.
* `cache_notfound_responses` - Enable/disable response from the DNS server when a record is not in cache.
* `source_ip` - IP address used by the DNS server as its source IP.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

The `server_hostname` block contains:

* `hostname` - DNS server host name list separated by space (maximum 4 domains).

The `domain` block contains:

* `domain` - DNS search domain list separated by space (maximum 8 domains)

