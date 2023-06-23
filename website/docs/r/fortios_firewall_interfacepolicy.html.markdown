---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_interfacepolicy"
description: |-
  Configure IPv4 interface policies.
---

# fortios_firewall_interfacepolicy
Configure IPv4 interface policies.

## Example Usage

```hcl
resource "fortios_firewall_interfacepolicy" "trname" {
  address_type              = "ipv4"
  application_list_status   = "disable"
  av_profile_status         = "disable"
  dlp_sensor_status         = "disable"
  dsri                      = "disable"
  interface                 = "port4"
  ips_sensor_status         = "disable"
  logtraffic                = "all"
  policyid                  = 1
  scan_botnet_connections   = "block"
  spamfilter_profile_status = "disable"
  status                    = "enable"
  webfilter_profile_status  = "disable"

  dstaddr {
    name = "all"
  }

  service {
    name = "ALL"
  }

  srcaddr {
    name = "all"
  }
}
```

## Argument Reference

The following arguments are supported:

* `policyid` - Policy ID.
* `status` - Enable/disable this policy. Valid values: `enable`, `disable`.
* `comments` - Comments.
* `logtraffic` - Logging type to be used in this policy (Options: all | utm | disable, Default: utm). Valid values: `all`, `utm`, `disable`.
* `address_type` - Policy address type (IPv4 or IPv6). Valid values: `ipv4`, `ipv6`.
* `interface` - (Required) Monitored interface name from available interfaces.
* `srcaddr` - (Required) Address object to limit traffic monitoring to network traffic sent from the specified address or range. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) Address object to limit traffic monitoring to network traffic sent to the specified address or range. The structure of `dstaddr` block is documented below.
* `service` - (Required) Service object from available options. The structure of `service` block is documented below.
* `application_list_status` - Enable/disable application control. Valid values: `enable`, `disable`.
* `application_list` - Application list name.
* `ips_sensor_status` - Enable/disable IPS. Valid values: `enable`, `disable`.
* `ips_sensor` - IPS sensor name.
* `dsri` - Enable/disable DSRI. Valid values: `enable`, `disable`.
* `av_profile_status` - Enable/disable antivirus. Valid values: `enable`, `disable`.
* `av_profile` - Antivirus profile.
* `webfilter_profile_status` - Enable/disable web filtering. Valid values: `enable`, `disable`.
* `webfilter_profile` - Web filter profile.
* `emailfilter_profile_status` - Enable/disable email filter. Valid values: `enable`, `disable`.
* `emailfilter_profile` - Email filter profile.
* `dlp_profile_status` - Enable/disable DLP. Valid values: `enable`, `disable`.
* `dlp_profile` - DLP profile name.
* `spamfilter_profile_status` - Enable/disable antispam. Valid values: `enable`, `disable`.
* `spamfilter_profile` - Antispam profile.
* `dlp_sensor_status` - Enable/disable DLP. Valid values: `enable`, `disable`.
* `dlp_sensor` - DLP sensor name.
* `scan_botnet_connections` - Enable/disable scanning for connections to Botnet servers. Valid values: `disable`, `block`, `monitor`.
* `label` - Label.
* `dynamic_sort_subtable` - Sort sub-tables, please do not set this parameter when configuring static sub-tables. Options: [ false, true, natural, alphabetical ]. false: Default value, do not sort tables; true/natural: sort tables in natural order. For example: [ a10, a2 ] --> [ a2, a10 ]; alphabetical: sort tables in alphabetical order. For example: [ a10, a2 ] --> [ a10, a2 ].
* `get_all_tables` - Get all sub-tables including unconfigured tables. Do not set this variable to true if you configure sub-table in another resource, otherwish conflicts and overwrite will occur. Options: [ false, true ]. false: Default value, do not get unconfigured tables; true: get all tables including unconfigured tables. 
* `vdomparam` - Specifies the vdom to which the resource will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Service name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall InterfacePolicy can be imported using any of these accepted formats:
```
$ terraform import fortios_firewall_interfacepolicy.labelname {{policyid}}

If you do not want to import arguments of block:
$ export "FORTIOS_IMPORT_TABLE"="false"
$ terraform import fortios_firewall_interfacepolicy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
