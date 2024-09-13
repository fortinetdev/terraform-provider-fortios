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
* `protocol` - DNS protocols.
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
* `root_servers` - Configure up to two preferred servers that serve the DNS root zone (default uses all 13 root servers).
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.
* `server_select_method` - Specify how configured servers are prioritized.
* `alt_primary` - Alternate primary DNS server. (This is not used as a failover DNS server.)
* `alt_secondary` - Alternate secondary DNS server. (This is not used as a failover DNS server.)
* `log` - Local DNS log setting.
* `fqdn_cache_ttl` - FQDN cache time to live in seconds (0 - 86400, default = 0).
* `fqdn_max_refresh` - FQDN cache maximum refresh time in seconds (3600 - 86400, default = 3600).
* `fqdn_min_refresh` - FQDN cache minimum refresh time in seconds (10 - 3600, default = 60).
* `hostname_ttl` - TTL of hostname table entries (60 - 86400).
* `hostname_limit` - Limit of the number of hostname table entries (0 - 50000).

The `server_hostname` block contains:

* `hostname` - DNS server host name list separated by space (maximum 4 domains).

The `domain` block contains:

* `domain` - DNS search domain list separated by space (maximum 8 domains)

