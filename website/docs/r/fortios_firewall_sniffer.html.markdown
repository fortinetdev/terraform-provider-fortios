---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_sniffer"
description: |-
  Configure sniffer.
---

# fortios_firewall_sniffer
Configure sniffer.

## Example Usage

```hcl
resource "fortios_firewall_sniffer" "trname" {
  application_list_status   = "disable"
  av_profile_status         = "disable"
  dlp_sensor_status         = "disable"
  dsri                      = "disable"
  fosid                     = 1
  interface                 = "port4"
  ips_dos_status            = "disable"
  ips_sensor_status         = "disable"
  ipv6                      = "disable"
  logtraffic                = "utm"
  max_packet_count          = 4000
  non_ip                    = "enable"
  scan_botnet_connections   = "disable"
  spamfilter_profile_status = "disable"
  status                    = "enable"
  webfilter_profile_status  = "disable"
}
```

## Argument Reference

The following arguments are supported:

* `fosid` - Sniffer ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `status` - Enable/disable the active status of the sniffer. Valid values: `enable`, `disable`.
* `logtraffic` - Either log all sessions, only sessions that have a security profile applied, or disable all logging for this policy. Valid values: `all`, `utm`, `disable`.
* `ipv6` - Enable/disable sniffing IPv6 packets. Valid values: `enable`, `disable`.
* `non_ip` - Enable/disable sniffing non-IP packets. Valid values: `enable`, `disable`.
* `interface` - (Required) Interface name that traffic sniffing will take place on.
* `host` - Hosts to filter for in sniffer traffic (Format examples: 1.1.1.1, 2.2.2.0/24, 3.3.3.3/255.255.255.0, 4.4.4.0-4.4.4.240).
* `port` - Ports to sniff (Format examples: 10, :20, 30:40, 50-, 100-200).
* `protocol` - Integer value for the protocol type as defined by IANA (0 - 255).
* `vlan` - List of VLANs to sniff.
* `application_list_status` - Enable/disable application control profile. Valid values: `enable`, `disable`.
* `application_list` - Name of an existing application list.
* `ips_sensor_status` - Enable/disable IPS sensor. Valid values: `enable`, `disable`.
* `ips_sensor` - Name of an existing IPS sensor.
* `dsri` - Enable/disable DSRI. Valid values: `enable`, `disable`.
* `av_profile_status` - Enable/disable antivirus profile. Valid values: `enable`, `disable`.
* `av_profile` - Name of an existing antivirus profile.
* `casb_profile_status` - Enable/disable CASB profile. Valid values: `enable`, `disable`.
* `casb_profile` - Name of an existing CASB profile.
* `webfilter_profile_status` - Enable/disable web filter profile. Valid values: `enable`, `disable`.
* `webfilter_profile` - Name of an existing web filter profile.
* `emailfilter_profile_status` - Enable/disable emailfilter. Valid values: `enable`, `disable`.
* `emailfilter_profile` - Name of an existing email filter profile.
* `dlp_profile_status` - Enable/disable DLP profile. Valid values: `enable`, `disable`.
* `dlp_profile` - Name of an existing DLP profile.
* `spamfilter_profile_status` - Enable/disable spam filter. Valid values: `enable`, `disable`.
* `spamfilter_profile` - Name of an existing spam filter profile.
* `dlp_sensor_status` - Enable/disable DLP sensor. Valid values: `enable`, `disable`.
* `dlp_sensor` - Name of an existing DLP sensor.
* `ip_threatfeed_status` - Enable/disable IP threat feed. Valid values: `enable`, `disable`.
* `ip_threatfeed` - Name of an existing IP threat feed. The structure of `ip_threatfeed` block is documented below.
* `file_filter_profile_status` - Enable/disable file filter. Valid values: `enable`, `disable`.
* `file_filter_profile` - Name of an existing file-filter profile.
* `ips_dos_status` - Enable/disable IPS DoS anomaly detection. Valid values: `enable`, `disable`.
* `anomaly` - Configuration method to edit Denial of Service (DoS) anomaly settings. The structure of `anomaly` block is documented below.
* `scan_botnet_connections` - Enable/disable scanning of connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
* `max_packet_count` - Maximum packet count. On FortiOS versions 6.2.0: 1 - 1000000, default = 10000. On FortiOS versions 6.2.4-6.4.2, 7.0.0: 1 - 10000, default = 4000. On FortiOS versions 6.4.10-6.4.15, 7.0.1-7.0.17: 1 - 1000000, default = 4000.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwise, conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `ip_threatfeed` block supports:

* `name` - Threat feed name.

The `anomaly` block supports:

* `name` - Anomaly name.
* `status` - Enable/disable this anomaly. Valid values: `disable`, `enable`.
* `log` - Enable/disable anomaly logging. Valid values: `enable`, `disable`.
* `action` - Action taken when the threshold is reached.
* `quarantine` - Quarantine method. Valid values: `none`, `attacker`.
* `quarantine_expiry` - Duration of quarantine. (Format ###d##h##m, minimum 1m, maximum 364d23h59m, default = 5m). Requires quarantine set to attacker.
* `quarantine_log` - Enable/disable quarantine logging. Valid values: `disable`, `enable`.
* `threshold` - Anomaly threshold. Number of detected instances that triggers the anomaly action. On FortiOS versions 6.2.0-6.4.2, 7.0.0-7.0.5, 7.2.0: packets per minute. On FortiOS versions 6.4.10-6.4.15, 7.0.6-7.0.17, >= 7.2.1: packets per second or concurrent session number.
* `thresholddefault` - Number of detected instances (packets per second or concurrent session number) which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall Sniffer can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_sniffer.labelname {{fosid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_sniffer.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
