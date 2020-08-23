---
subcategory: "FortiGate Log"
layout: "fortios"
page_title: "FortiOS: fortios_logfortianalyzer2_setting"
description: |-
  Global FortiAnalyzer settings.
---

# fortios_logfortianalyzer2_setting
Global FortiAnalyzer settings.

## Example Usage

```hcl
resource "fortios_logfortianalyzer2_setting" "trname" {
  __change_ip                  = 0
  conn_timeout                 = 10
  enc_algorithm                = "high"
  faz_type                     = 2
  hmac_algorithm               = "sha256"
  ips_archive                  = "enable"
  mgmt_name                    = "FGh_Log2"
  monitor_failure_retry_period = 5
  monitor_keepalive_period     = 5
  reliable                     = "disable"
  ssl_min_proto_version        = "default"
  status                       = "disable"
  upload_interval              = "daily"
  upload_option                = "5-minute"
  upload_time                  = "00:59"
}
```

## Argument Reference

The following arguments are supported:

* `status` - Enable/disable logging to FortiAnalyzer.
* `ips_archive` - Enable/disable IPS packet archive logging.
* `server` - The remote FortiAnalyzer.
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


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

LogFortianalyzer2 Setting can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_logfortianalyzer2_setting.labelname LogFortianalyzer2Setting
$ unset "FORTIOS_IMPORT_TABLE"
```
