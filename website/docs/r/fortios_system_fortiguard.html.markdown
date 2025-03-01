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

* `protocol` - Protocol used to communicate with the FortiGuard servers. Valid values: `udp`, `http`, `https`.
* `port` - Port used to communicate with the FortiGuard servers.
* `service_account_id` - Service account ID.
* `load_balance_servers` - Number of servers to alternate between as first FortiGuard option.
* `auto_join_forticloud` - Automatically connect to and login to FortiCloud. Valid values: `enable`, `disable`.
* `update_server_location` - Signature update server location.
* `sandbox_region` - Cloud sandbox region.
* `sandbox_inline_scan` - Enable/disable FortiCloud Sandbox inline-scan. Valid values: `enable`, `disable`.
* `update_ffdb` - Enable/disable Internet Service Database update. Valid values: `enable`, `disable`.
* `update_uwdb` - Enable/disable allowlist update. Valid values: `enable`, `disable`.
* `update_dldb` - Enable/disable DLP signature update. Valid values: `enable`, `disable`.
* `update_extdb` - Enable/disable external resource update. Valid values: `enable`, `disable`.
* `update_build_proxy` - Enable/disable proxy dictionary rebuild. Valid values: `enable`, `disable`.
* `persistent_connection` - Enable/disable use of persistent connection to receive update notification from FortiGuard. Valid values: `enable`, `disable`.
* `vdom` - FortiGuard Service virtual domain name.
* `auto_firmware_upgrade` - Enable/disable automatic patch-level firmware upgrade from FortiGuard. The FortiGate unit searches for new patches only in the same major and minor version. Valid values: `enable`, `disable`.
* `auto_firmware_upgrade_day` - Allowed day(s) of the week to install an automatic patch-level firmware upgrade from FortiGuard (default is none). Disallow any day of the week to use auto-firmware-upgrade-delay instead, which waits for designated days before installing an automatic patch-level firmware upgrade. Valid values: `sunday`, `monday`, `tuesday`, `wednesday`, `thursday`, `friday`, `saturday`.
* `auto_firmware_upgrade_delay` - Delay of day(s) before installing an automatic patch-level firmware upgrade from FortiGuard (default = 3). Set it 0 to use auto-firmware-upgrade-day instead, which selects allowed day(s) of the week for installing an automatic patch-level firmware upgrade.
* `auto_firmware_upgrade_start_hour` - Start time in the designated time window for automatic patch-level firmware upgrade from FortiGuard in 24 hour time (0 ~ 23, default = 2). The actual upgrade time is selected randomly within the time window.
* `auto_firmware_upgrade_end_hour` - End time in the designated time window for automatic patch-level firmware upgrade from FortiGuard in 24 hour time (0 ~ 23, default = 4). When the end time is smaller than the start time, the end time is interpreted as the next day. The actual upgrade time is selected randomly within the time window.
* `gui_prompt_auto_upgrade` - Enable/disable prompting of automatic patch-level firmware upgrade recommendation. Valid values: `enable`, `disable`.
* `fds_license_expiring_days` - Threshold for number of days before FortiGuard license expiration to generate license expiring event log (1 - 100 days, default = 15).
* `fortiguard_anycast` - Enable/disable use of FortiGuard's anycast network. Valid values: `enable`, `disable`.
* `fortiguard_anycast_source` - Configure which of Fortinet's servers to provide FortiGuard services in FortiGuard's anycast network. Default is Fortinet. Valid values: `fortinet`, `aws`, `debug`.
* `antispam_force_off` - Enable/disable turning off the FortiGuard antispam service. Valid values: `enable`, `disable`.
* `antispam_cache` - Enable/disable FortiGuard antispam request caching. Uses a small amount of memory but improves performance. Valid values: `enable`, `disable`.
* `antispam_cache_ttl` - Time-to-live for antispam cache entries in seconds (300 - 86400). Lower times reduce the cache size. Higher times may improve performance since the cache will have more entries.
* `antispam_cache_mpermille` - Maximum permille of FortiGate memory the antispam cache is allowed to use (1 - 150).
* `antispam_cache_mpercent` - Maximum percentage of FortiGate memory the antispam cache is allowed to use (1 - 15).
* `antispam_license` - Interval of time between license checks for the FortiGuard antispam contract.
* `antispam_expiration` - Expiration date of the FortiGuard antispam contract.
* `antispam_timeout` - (Required) Antispam query time out (1 - 30 sec, default = 7).
* `outbreak_prevention_force_off` - Turn off FortiGuard Virus Outbreak Prevention service. Valid values: `enable`, `disable`.
* `outbreak_prevention_cache` - Enable/disable FortiGuard Virus Outbreak Prevention cache. Valid values: `enable`, `disable`.
* `outbreak_prevention_cache_ttl` - Time-to-live for FortiGuard Virus Outbreak Prevention cache entries (300 - 86400 sec, default = 300).
* `outbreak_prevention_cache_mpermille` - Maximum permille of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 150 permille, default = 1).
* `outbreak_prevention_cache_mpercent` - Maximum percent of memory FortiGuard Virus Outbreak Prevention cache can use (1 - 15%, default = 2).
* `outbreak_prevention_license` - Interval of time between license checks for FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_expiration` - Expiration date of FortiGuard Virus Outbreak Prevention contract.
* `outbreak_prevention_timeout` - (Required) FortiGuard Virus Outbreak Prevention time out (1 - 30 sec, default = 7).
* `webfilter_force_off` - Enable/disable turning off the FortiGuard web filtering service. Valid values: `enable`, `disable`.
* `webfilter_cache` - Enable/disable FortiGuard web filter caching. Valid values: `enable`, `disable`.
* `webfilter_cache_ttl` - Time-to-live for web filter cache entries in seconds (300 - 86400).
* `webfilter_license` - Interval of time between license checks for the FortiGuard web filter contract.
* `webfilter_expiration` - Expiration date of the FortiGuard web filter contract.
* `webfilter_timeout` - (Required) Web filter query time out, 1 - 30 sec. On FortiOS versions 6.2.0-7.4.0: default = 7. On FortiOS versions >= 7.4.1: default = 15.
* `sdns_server_ip` - IP address of the FortiDNS server.
* `sdns_server_port` - Port used to communicate with FortiDNS servers.
* `anycast_sdns_server_ip` - IP address of the FortiGuard anycast DNS rating server.
* `anycast_sdns_server_port` - Port to connect to on the FortiGuard anycast DNS rating server.
* `sdns_options` - Customization options for the FortiGuard DNS service. Valid values: `include-question-section`.
* `source_ip` - Source IPv4 address used to communicate with FortiGuard.
* `source_ip6` - Source IPv6 address used to communicate with FortiGuard.
* `proxy_server_ip` - IP address of the proxy server.
* `proxy_server_port` - Port used to communicate with the proxy server.
* `proxy_username` - Proxy user name.
* `proxy_password` - Proxy user password.
* `videofilter_license` - Interval of time between license checks for the FortiGuard video filter contract.
* `videofilter_expiration` - Expiration date of the FortiGuard video filter contract.
* `ddns_server_ip` - IP address of the FortiDDNS server.
* `ddns_server_ip6` - IPv6 address of the FortiDDNS server.
* `ddns_server_port` - Port used to communicate with FortiDDNS servers.
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

System Fortiguard can be imported using any of these accepted formats:
```
$ terraform import fortios_system_fortiguard.labelname SystemFortiguard

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_system_fortiguard.labelname SystemFortiguard
$ unset "FORTIOS_IMPORT_TABLE"
```
