---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortiguard"
description: |-
  Get information on fortios system fortiguard.
---

# Data Source: fortios_system_fortiguard
Use this data source to get information on fortios system fortiguard

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `protocol` - Protocol used to communicate with the FortiGuard servers.
* `port` - Port used to communicate with the FortiGuard servers.
* `service_account_id` - Service account ID.
* `load_balance_servers` - Number of servers to alternate between as first FortiGuard option.
* `auto_join_forticloud` - Automatically connect to and login to FortiCloud.
* `update_server_location` - Signature update server location.
* `sandbox_region` - Cloud sandbox region.
* `fortiguard_anycast` - Enable/disable use of FortiGuard's anycast network.
* `fortiguard_anycast_source` - Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet.
* `antispam_force_off` - Enable/disable turning off the FortiGuard antispam service.
* `antispam_cache` - Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance.
* `antispam_cache_ttl` - Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
* `antispam_cache_mpercent` - Maximum percent of FortiGate memory the antispam cache is allowed to use (1 - 15%).
* `antispam_license` - Interval of time between license checks for the FortiGuard antispam contract.
* `antispam_expiration` - Expiration date of the FortiGuard antispam contract.
* `antispam_timeout` - Antispam query time out (1 - 30 sec, default = 7).
* `outbreak_prevention_force_off` - Turn off FortiGuard Virus Outbreak Prevention service.
* `outbreak_prevention_cache` - Enable/disable FortiGuard Virus Outbreak Prevention cache.
* `outbreak_prevention_cache_ttl` - Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
* `outbreak_prevention_cache_mpercent` - Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
* `outbreak_prevention_license` - Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_expiration` - Expiration date of FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_timeout` - FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
* `webfilter_force_off` - Enable/disable turning off the FortiGuard web filtering service.
* `webfilter_cache` - Enable/disable FortiGuard web filter caching.
* `webfilter_cache_ttl` - Time-to-live for web filter cache entries in seconds (300 - 86400).
* `webfilter_license` - Interval of time between license checks for the FortiGuard web filter contract.
* `webfilter_expiration` - Expiration date of the FortiGuard web filter contract.
* `webfilter_timeout` - Web filter query time out (1 - 30 sec, default = 7).
* `sdns_server_ip` - IP address of the FortiDNS server.
* `sdns_server_port` - Port used to communicate with FortiDNS servers.
* `anycast_sdns_server_ip` - IP address of the FortiGuard anycast DNS rating server.
* `anycast_sdns_server_port` - Port to connect to on the FortiGuard anycast DNS rating server.
* `sdns_options` - Customization options for the FortiGuard DNS service.
* `source_ip` - Source IPv4 address used to communicate with FortiGuard.
* `source_ip6` - Source IPv6 address used to communicate with FortiGuard.
* `proxy_server_ip` - IP address of the proxy server.
* `proxy_server_port` - Port used to communicate with the proxy server.
* `proxy_username` - Proxy user name.
* `proxy_password` - Proxy user password.
* `ddns_server_ip` - IP address of the FortiDDNS server.
* `ddns_server_port` - Port used to communicate with FortiDDNS servers.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

