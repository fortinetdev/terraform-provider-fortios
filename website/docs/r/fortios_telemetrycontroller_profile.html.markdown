---
subcategory: "FortiGate Telemetry-Controller"
layout: "fortios"
page_title: "FortiOS: fortios_telemetrycontroller_profile"
description: |-
  Configure FortiTelemetry profiles.
---

# fortios_telemetrycontroller_profile
Configure FortiTelemetry profiles. Applies to FortiOS Version `>= 8.0.0`.

## Argument Reference

The following arguments are supported:

* `name` - Name of the profile.
* `comment` - Comment.
* `application` - Configure applications. The structure of `application` block is documented below.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.
* `update_if_exist` - Equivalent functionality of import the resource. If set to true, will check whether the resource exist, if so, will do the UPDATE operation rather CREATE. Default is false. If you want to inherit the value of the provider, please do not set this parameter.

The `application` block supports:

* `id` - ID.
* `app_name` - Application name.
* `monitor` - Enable/disable monitoring of the application. Valid values: `enable`, `disable`.
* `interval` - Time in milliseconds to check the application (60 * 1000 - 86,400 * 1000, default = 300 * 1000 ms).
* `sla` - Service level agreement (SLA). The structure of `sla` block is documented below.

The `sla` block supports:

* `sla_factor` - Criteria on which metric to SLA threshold list. Valid values: `experience-score`, `failure-rate`, `latency`, `jitter`, `packet-loss`, `ttfb`, `atdt`, `tcp-rtt`, `dns-time`, `tls-time`, `app-throughput`.
* `experience_score_threshold` - Threshold of experience score (0 - 10, default = 6).
* `failure_rate_threshold` - Threshold of failure rate (0 - 100, default = 5 percentage).
* `latency_threshold` - Threshold of latency in milliseconds (0 - 10,000,000, default = 100 ms).
* `jitter_threshold` - Threshold of jitter in milliseconds (0 - 10,000,000, default = 50 ms).
* `packet_loss_threshold` - Threshold of packet loss (0 - 100, default = 5 percentage).
* `ttfb_threshold` - Threshold of time to first byte in milliseconds (0 - 10,000,000, default = 200 ms).
* `atdt_threshold` - Threshold of application total downloading time in milliseconds (0 - 10,000,000, default = 3,000 ms).
* `tcp_rtt_threshold` - Threshold of TCP round-trip time in milliseconds (0 - 10,000,000, default = 100 ms).
* `dns_time_threshold` - Threshold of 95th percentile of DNS resolution time in milliseconds (0 - 10,000,000, default = 100 ms).
* `tls_time_threshold` - Threshold of 95th percentile of TLS handshake time in milliseconds (0 - 10,000,000, default = 200 ms).
* `app_throughput_threshold` - Threshold of application throughput in megabytes (0 - 10,000, default = 2 MB).


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{name}}.

## Import

TelemetryController Profile can be imported using any of these accepted formats:
```
$ terraform import fortios_telemetrycontroller_profile.labelname {{name}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_telemetrycontroller_profile.labelname {{name}}
$ unset "FORTIOS_IMPORT_TABLE"
```
