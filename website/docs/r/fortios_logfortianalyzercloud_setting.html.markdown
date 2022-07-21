---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzercloud_setting"
description: |-
  Global FortiAnalyzer Cloud settings.
---

# fortios_logfortianalyzercloud_setting
Global FortiAnalyzer Cloud settings. Applies to FortiOS Version `>= 6.2.4`.

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable logging to FortiAnalyzer. Valid values: `enable`, `disable`.
* `ips_archive` - Enable/disable IPS packet archive logging. Valid values: `enable`, `disable`.
* `certificate_verification` - Enable/disable identity verification of FortiAnalyzer by use of certificate. Valid values: `enable`, `disable`.
* `serial` - Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
* `preshared_key` - Preshared-key used for auto-authorization on FortiAnalyzer.
* `access_config` - Enable/disable FortiAnalyzer access to configuration and data. Valid values: `enable`, `disable`.
* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm. Valid values: `sha256`, `sha1`.
* `enc_algorithm` - Configure the level of SSL protection for secure communication with FortiAnalyzer. Valid values: `high-medium`, `high`, `low`.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `monitor_keepalive_period` - Time between OFTP keepalives in seconds (for status and log buffer).
* `monitor_failure_retry_period` - Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
* `certificate` - Certificate used to communicate with FortiAnalyzer.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer. Valid values: `store-and-upload`, `realtime`, `1-minute`, `5-minute`.
* `upload_interval` - Frequency to upload log files to FortiAnalyzer. Valid values: `daily`, `weekly`, `monthly`.
* `upload_day` - Day of week (month) to upload logs.
* `upload_time` - Time to upload logs (hh:mm).
* `priority` - Set log transmission priority. Valid values: `default`, `low`.
* `max_log_rate` - FortiAnalyzer maximum log rate in MBps (0 = unlimited).
* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `serial` block supports:

* `name` - Serial Number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortianalyzerCloud Setting can be imported using any of these accepted formats:
```
$ terraform import fortios_logfortianalyzercloud_setting.labelname LogFortianalyzerCloudSetting

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_logfortianalyzercloud_setting.labelname LogFortianalyzerCloudSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
