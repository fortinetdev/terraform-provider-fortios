---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy46"
description: |-
  Configure IPv4 to IPv6 policies.
---

# fortios_firewall_policy46
Configure IPv4 to IPv6 policies.

## Example Usage

```hcl
resource "fortios_firewall_vip46" "trname" {
  arp_reply   = "enable"
  color       = 0
  extip       = "10.1.100.55"
  extport     = "0-65535"
  fosid       = 0
  ldb_method  = "static"
  mappedip    = "2000:172:16:200::55"
  mappedport  = "0-65535"
  name        = "1"
  portforward = "disable"
  protocol    = "tcp"
  type        = "static-nat"
}

resource "fortios_firewall_policy46" "trname" {
  action           = "deny"
  dstintf          = "port3"
  fixedport        = "disable"
  ippool           = "disable"
  logtraffic       = "disable"
  permit_any_host  = "disable"
  policyid         = 2
  schedule         = "always"
  srcintf          = "port2"
  status           = "enable"
  tcp_mss_receiver = 0
  tcp_mss_sender   = 0

  dstaddr {
    name = fortios_firewall_vip46.trname.name
  }

  service {
    name = "ALL"
  }

  srcaddr {
    name = "FIREWALL_AUTH_PORTAL_ADDRESS"
  }
}
```

## Argument Reference


The following arguments are supported:

* `permit_any_host` - Enable/disable allowing any host.
* `policyid` - Policy ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `srcintf` - (Required) Source interface name.
* `dstintf` - (Required) Destination interface name.
* `srcaddr` - (Required) Source address objects. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) Destination address objects. The structure of `dstaddr` block is documented below.
* `action` - Accept or deny traffic matching the policy.
* `status` - Enable/disable this policy.
* `schedule` - (Required) Schedule name.
* `service` - Service name. The structure of `service` block is documented below.
* `logtraffic` - Enable/disable traffic logging for this policy.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per IP traffic shaper.
* `fixedport` - Enable/disable fixed port for this policy.
* `tcp_mss_sender` - TCP Maximum Segment Size value of sender (0 - 65535, default = 0).
* `tcp_mss_receiver` - TCP Maximum Segment Size value of receiver (0 - 65535, default = 0)
* `comments` - Comment.
* `ippool` - Enable/disable use of IP Pools for source NAT.
* `poolname` - IP Pool names. The structure of `poolname` block is documented below.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Service name.

The `poolname` block supports:

* `name` - IP pool name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall Policy46 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_policy46.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
