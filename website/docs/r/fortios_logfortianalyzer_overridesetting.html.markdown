---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzer_overridesetting"
description: |-
  Override FortiAnalyzer settings.
---

# fortios_logfortianalyzer_overridesetting
Override FortiAnalyzer settings.

## Example Usage

```hcl
resource "fortios_logfortianalyzer_overridesetting" "trname" {
  __change_ip                  = 0
  conn_timeout                 = 10
  enc_algorithm                = "high"
  faz_type                     = 4
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

* `override` - Enable/disable overriding FortiAnalyzer settings or use global settings.
* `use_management_vdom` - Enable/disable use of management VDOM IP address as source IP for logs sent to FortiAnalyzer.
* `status` - Enable/disable logging to FortiAnalyzer.
* `ips_archive` - Enable/disable IPS packet archive logging.
* `server` - The remote FortiAnalyzer.
* `certificate_verification` - Enable/disable identity verification of FortiAnalyzer by use of certificate.
* `serial` - Serial numbers of the FortiAnalyzer. The structure of `serial` block is documented below.
* `access_config` - Enable/disable FortiAnalyzer access to configuration and data.
* `hmac_algorithm` - FortiAnalyzer IPsec tunnel HMAC algorithm.
* `enc_algorithm` - Enable/disable sending FortiAnalyzer log data with SSL encryption.
* `ssl_min_proto_version` - Minimum supported protocol version for SSL/TLS connections (default is to follow system global setting).
* `conn_timeout` - FortiAnalyzer connection time-out in seconds (for status and log buffer).
* `monitor_keepalive_period` - Time between OFTP keepalives in seconds (for status and log buffer).
* `monitor_failure_retry_period` - Time between FortiAnalyzer connection retries in seconds (for status and log buffer).
* `mgmt_name` - Hidden management name of FortiAnalyzer.
* `faz_type` - Hidden setting index of FortiAnalyzer.
* `certificate` - Certificate used to communicate with FortiAnalyzer.
* `source_ip` - Source IPv4 or IPv6 address used to communicate with FortiAnalyzer.
* `__change_ip` - Hidden attribute.
* `upload_option` - Enable/disable logging to hard disk and then uploading to FortiAnalyzer.
* `upload_interval` - Frequency to upload log files to FortiAnalyzer.
* `upload_day` - Day of week (month) to upload logs.
* `upload_time` - Time to upload logs (hh:mm).
* `reliable` - Enable/disable reliable logging to FortiAnalyzer.
* `priority` - Set log transmission priority.
* `max_log_rate` - FortiAnalyzer maximum log rate in MBps (0 = unlimited).
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `serial` block supports:

* `name` - Serial Number.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortianalyzer OverrideSetting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logfortianalyzer_overridesetting.labelname LogFortianalyzerOverrideSetting
$ unset "FORTIOS_IMPORT_TABLE"
```
