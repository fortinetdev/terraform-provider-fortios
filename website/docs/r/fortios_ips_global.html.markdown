---
subcategory: "FortiGate IPS"
layout: "fortios"
page_title: "FortiOS: fortios_ips_global"
description: |-
  Configure IPS global parameter.
---

# fortios_ips_global
Configure IPS global parameter.

## Example Usage

```hcl
resource "fortios_ips_global" "trname" {
  anomaly_mode           = "continuous"
  database               = "regular"
  deep_app_insp_db_limit = 0
  deep_app_insp_timeout  = 0
  engine_count           = 0
  exclude_signatures     = "industrial"
  fail_open              = "disable"
  intelligent_mode       = "enable"
  session_limit_mode     = "heuristic"
  socket_size            = 0
  sync_session_ttl       = "enable"
  traffic_submit         = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `fail_open` - Enable to allow traffic if the IPS process crashes. Default is disable and IPS traffic is blocked when the IPS process crashes. Valid values: `enable`, `disable`.
* `database` - Regular or extended IPS database. Regular protects against the latest common and in-the-wild attacks. Extended includes protection from legacy attacks. Valid values: `regular`, `extended`.
* `traffic_submit` - Enable/disable submitting attack data found by this FortiGate to FortiGuard. Valid values: `enable`, `disable`.
* `anomaly_mode` - Global blocking mode for rate-based anomalies. Valid values: `periodical`, `continuous`.
* `session_limit_mode` - Method of counting concurrent sessions used by session limit anomalies. Choose between greater accuracy (accurate) or improved performance (heuristics). Valid values: `accurate`, `heuristic`.
* `intelligent_mode` - Enable/disable IPS adaptive scanning (intelligent mode). Intelligent mode optimizes the scanning method for the type of traffic. Valid values: `enable`, `disable`.
* `socket_size` - IPS socket buffer size (0 - 256 MB). Default depends on available memory. Can be changed to tune performance.
* `engine_count` - Number of IPS engines running. If set to the default value of 0, FortiOS sets the number to optimize performance depending on the number of CPU cores.
* `sync_session_ttl` - Enable/disable use of kernel session TTL for IPS sessions. Valid values: `enable`, `disable`.
* `np_accel_mode` - Acceleration mode for IPS processing by NPx processors. Valid values: `none`, `basic`.
* `ips_reserve_cpu` - Enable/disable IPS daemon's use of CPUs other than CPU 0 Valid values: `disable`, `enable`.
* `cp_accel_mode` - IPS Pattern matching acceleration/offloading to CPx processors. Valid values: `none`, `basic`, `advanced`.
* `skype_client_public_ipaddr` - Public IP addresses of your network that receive Skype sessions. Helps identify Skype sessions. Separate IP addresses with commas.
* `deep_app_insp_timeout` - Timeout for Deep application inspection (1 - 2147483647 sec., 0 = use recommended setting).
* `deep_app_insp_db_limit` - Limit on number of entries in deep application inspection database (1 - 2147483647, 0 = use recommended setting)
* `exclude_signatures` - Excluded signatures. Valid values: `none`, `industrial`.
* `packet_log_queue_depth` - Packet/pcap log queue depth per IPS engine.
* `ngfw_max_scan_range` - NGFW policy-mode app detection threshold.
* `tls_active_probe` - TLS active probe configuration. The structure of `tls_active_probe` block is documented below.
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `tls_active_probe` block supports:

* `interface_select_method` - Specify how to select outgoing interface to reach server. Valid values: `auto`, `sdwan`, `specify`.
* `interface` - Specify outgoing interface to reach server.
* `vdom` - Virtual domain name for TLS active probe.
* `source_ip` - Source IP address used for TLS active probe.
* `source_ip6` - Source IPv6 address used for TLS active probe.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource.

## Import

Ips Global can be imported using any of these accepted formats:
```
$ terraform import fortios_ips_global.labelname IpsGlobal

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_ips_global.labelname IpsGlobal
$ unset "FORTIOS_IMPORT_TABLE"
```
