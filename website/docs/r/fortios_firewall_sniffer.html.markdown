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
* `webfilter_profile_status` - Enable/disable web filter profile. Valid values: `enable`, `disable`.
* `webfilter_profile` - Name of an existing web filter profile.
* `emailfilter_profile_status` - Enable/disable emailfilter. Valid values: `enable`, `disable`.
* `emailfilter_profile` - Name of an existing email filter profile.
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
* `max_packet_count` - Maximum packet count (1 - 1000000, default = 10000).
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.
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
* `threshold` - Anomaly threshold. Number of detected instances per minute that triggers the anomaly action.
* `thresholddefault` - Number of detected instances per minute which triggers action (1 - 2147483647, default = 1000). Note that each anomaly has a different threshold value assigned to it.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{fosid}}.

## Import

Firewall Sniffer can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_sniffer.labelname {{fosid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
