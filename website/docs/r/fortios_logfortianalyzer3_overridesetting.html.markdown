---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzer3_overridesetting"
description: |-
  Override FortiAnalyzer settings.
---

# fortios_logfortianalyzer3_overridesetting
Override FortiAnalyzer settings.

## Example Usage

```hcl
resource "fortios_logfortianalyzer3_overridesetting" "trname" {
  __change_ip                  = 0
  conn_timeout                 = 10
  enc_algorithm                = "high"
  faz_type                     = 6
  hmac_algorithm               = "sha256"
  ips_archive                  = "enable"
  monitor_failure_retry_period = 5
  monitor_keepalive_period     = 5
  override                     = "disable"
  reliable                     = "disable"
  ssl_min_proto_version        = "default"
  status                       = "disable"
  upload_interval              = "daily"
  upload_option                = "5-minute"
  upload_time                  = "00:59"
  use_management_vdom          = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `override` - Enable/disable overriding FortiAnalyzer settings or use global settings. Valid values: `enable`, `disable`.
* `use_management_vdom` - Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer. Valid values: `enable`, `disable`.
* `status` - Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
* `ips_archive` - Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
* `server` - The remote FortiAnalyzer.
* `alt_server` - Alternate FortiAnalyzer.
* `fallback_to_primary` - Enable/disable this FortiGate unit to fallback to the primary FortiAnalyzer when it is available. Valid values: `enable`, `disable`.
* `certificate_verification` - Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
* `serial` - Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
* `server_cert_ca` - Mandatory CA on FortiGate in certificate chain of server.
* `preshared_key` - Preshared-key used for auto-authorization on FortiAnalyzer.
* `access_config` - Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm.
* `enc_algorithm` - Enable/disable sending FortiAnalyzer log data with SSL encryption. Valid values: `high-medium`, `high`, `low`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `monitor_keepalive_period` - Time between OFTP keepalives in seconds (for status and log buffer).
* `monitor_failure_retry_period` - Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
* `mgmt_name` - Hidden management name of FortiAnalyzer.
* `faz_type` - Hidden setting index of FortiAnalyzer.
* `certificate` - Certificate used to communicate with FortiAnalyzer.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `__change_ip` - Hidden attribute.
* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
* `upload_interval` - Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
* `upload_day` - Day of week (month) to upload logs.
* `upload_time` - Time to upload logs (hh:mm).
* `reliable` - Enable/disable reliable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
* `priority` - Set log transmission priority. Valid values: `default`, `low`.
* `max_log_rate` - FortiAnalyzer maximum log rate in MBps (0 = unlimited).
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vrf_select` - VRF ID used for connection to server.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `serial` block supports:

* `name` - Serial Number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortianalyzer3 OverrideSetting can be imported using any of these accepted formats:
```
$ terraform import fortios_logfortianalyzer3_overridesetting.labelname LogFortianalyzer3OverrideSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_logfortianalyzer3_overridesetting.labelname LogFortianalyzer3OverrideSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
