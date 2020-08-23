---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fortiguard"
description: |-
  Configure FortiGuard services.
---

# fortios_system_fortiguard
Configure FortiGuard services.

## Example Usage

```hcl
resource "fortios_system_fortiguard" "trname" {
  antispam_cache                     = "enable"
  antispam_cache_mpercent            = 2
  antispam_cache_ttl                 = 1800
  antispam_expiration                = 1618617600
  antispam_force_off                 = "disable"
  antispam_license                   = 1
  antispam_timeout                   = 7
  auto_join_forticloud               = "enable"
  ddns_server_ip                     = "0.0.0.0"
  ddns_server_port                   = 443
  load_balance_servers               = 1
  outbreak_prevention_cache          = "enable"
  outbreak_prevention_cache_mpercent = 2
  outbreak_prevention_cache_ttl      = 300
  outbreak_prevention_expiration     = 1618617600
  outbreak_prevention_force_off      = "disable"
  outbreak_prevention_license        = 1
  outbreak_prevention_timeout        = 7
  port                               = "8888"
  sdns_server_ip                     = "\"208.91.112.220\" "
  sdns_server_port                   = 53
  source_ip                          = "0.0.0.0"
  source_ip6                         = "::"
  update_server_location             = "usa"
  webfilter_cache                    = "enable"
  webfilter_cache_ttl                = 3600
  webfilter_expiration               = 1618617600
  webfilter_force_off                = "disable"
  webfilter_license                  = 1
  webfilter_timeout                  = 15
}
```

## Argument Reference

The following arguments are supported:

* `port` - Port used to communicate with the FortiGuard servers.
* `service_account_id` - Service account ID.
* `load_balance_servers` - Number of servers to alternate between as first FortiGuard option.
* `auto_join_forticloud` - Automatically connect to and login to FortiCloud.
* `update_server_location` - Signature update server location.
* `sandbox_region` - Cloud sandbox region.
* `antispam_force_off` - Enable/disable turning off the FortiGuard antispam service.
* `antispam_cache` - Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance.
* `antispam_cache_ttl` - Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
* `antispam_cache_mpercent` - Maximum percent of FortiGate memory the antispam cache is allowed to use (1 - 15%).
* `antispam_license` - Interval of time between license checks for the FortiGuard antispam contract.
* `antispam_expiration` - Expiration date of the FortiGuard antispam contract.
* `antispam_timeout` - (Required) Antispam query time out (1 - 30 sec, default = 7).
* `outbreak_prevention_force_off` - Turn off FortiGuard Virus Outbreak Prevention service.
* `outbreak_prevention_cache` - Enable/disable FortiGuard Virus Outbreak Prevention cache.
* `outbreak_prevention_cache_ttl` - Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
* `outbreak_prevention_cache_mpercent` - Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
* `outbreak_prevention_license` - Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_expiration` - Expiration date of FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_timeout` - (Required) FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
* `webfilter_force_off` - Enable/disable turning off the FortiGuard web filtering service.
* `webfilter_cache` - Enable/disable FortiGuard web filter caching.
* `webfilter_cache_ttl` - Time-to-live for web filter cache entries in seconds (300 - 86400).
* `webfilter_license` - Interval of time between license checks for the FortiGuard web filter contract.
* `webfilter_expiration` - Expiration date of the FortiGuard web filter contract.
* `webfilter_timeout` - (Required) Web filter query time out (1 - 30 sec, default = 7).
* `sdns_server_ip` - IP address of the FortiDNS server.
* `sdns_server_port` - Port used to communicate with FortiDNS servers.
* `source_ip` - Source IPv4 address used to communicate with FortiGuard.
* `source_ip6` - Source IPv6 address used to communicate with FortiGuard.
* `ddns_server_ip` - IP address of the FortiDDNS server.
* `ddns_server_port` - Port used to communicate with FortiDDNS servers.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortiguard can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_system_fortiguard.labelname SystemFortiguard
$ unset "FORTIOS_IMPORT_TABLE"
```
