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

* `policyid` - (Required) Policy ID.
* `status` - Enable/disable this policy.
* `comments` - Comments.
* `logtraffic` - Logging type to be used in this policy (Options: all | utm | disable, Default: utm).
* `address_type` - Policy address type (IPv4 or IPv6).
* `interface` - (Required) Monitored interface name from available interfaces.
* `srcaddr` - (Required) Address object to limit traffic monitoring to network traffic sent from the specified address or range. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) Address object to limit traffic monitoring to network traffic sent to the specified address or range. The structure of `dstaddr` block is documented below.
* `service` - (Required) Service object from available options. The structure of `service` block is documented below.
* `application_list_status` - Enable/disable application control.
* `application_list` - Application list name.
* `ips_sensor_status` - Enable/disable IPS.
* `ips_sensor` - IPS sensor name.
* `dsri` - Enable/disable DSRI.
* `av_profile_status` - Enable/disable antivirus.
* `av_profile` - Antivirus profile.
* `webfilter_profile_status` - Enable/disable web filtering.
* `webfilter_profile` - Web filter profile.
* `spamfilter_profile_status` - Enable/disable antispam.
* `spamfilter_profile` - Antispam profile.
* `dlp_sensor_status` - Enable/disable DLP.
* `dlp_sensor` - DLP sensor name.
* `scan_botnet_connections` - Enable/disable scanning for connections to Botnet servers.
* `label` - Label.

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
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_interfacepolicy.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
