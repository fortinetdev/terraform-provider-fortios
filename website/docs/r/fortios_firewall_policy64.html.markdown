---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy64"
description: |-
  Configure IPv6 to IPv4 policies.
---

# fortios_firewall_policy64
Configure IPv6 to IPv4 policies.

## Example Usage

```hcl
resource "fortios_firewall_policy64" "trname" {
  action           = "accept"
  dstintf          = "port4"
  fixedport        = "disable"
  ippool           = "disable"
  logtraffic       = "disable"
  permit_any_host  = "disable"
  policyid         = 1
  schedule         = "always"
  srcintf          = "port3"
  status           = "enable"
  tcp_mss_receiver = 0
  tcp_mss_sender   = 0

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
* `name` - Policy name.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `srcintf` - (Required) Source interface name.
* `dstintf` - (Required) Destination interface name.
* `srcaddr` - (Required) Source address name. The structure of `srcaddr` block is documented below.
* `dstaddr` - (Required) Destination address name. The structure of `dstaddr` block is documented below.
* `action` - Policy action. Valid values: `accept`, `deny`.
* `status` - Enable/disable policy status. Valid values: `enable`, `disable`.
* `schedule` - (Required) Schedule name.
* `service` - Service name. The structure of `service` block is documented below.
* `logtraffic` - Enable/disable policy log traffic. Valid values: `enable`, `disable`.
* `logtraffic_start` - Record logs when a session starts and ends. Valid values: `enable`, `disable`.
* `permit_any_host` - Enable/disable permit any host in. Valid values: `enable`, `disable`.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `fixedport` - Enable/disable policy fixed port. Valid values: `enable`, `disable`.
* `ippool` - Enable/disable policy64 IP pool. Valid values: `enable`, `disable`.
* `poolname` - Policy IP pool names. The structure of `poolname` block is documented below.
* `tcp_mss_sender` - TCP MSS value of sender.
* `tcp_mss_receiver` - TCP MSS value of receiver.
* `comments` - Comment.
* `dynamic_sort_subtable` - true or false, set this parameter to true when using dynamic for_each + toset to configure and sort sub-tables, please do not set this parameter when configuring static sub-tables.

The `srcaddr` block supports:

* `name` - Address name.

The `dstaddr` block supports:

* `name` - Address name.

The `service` block supports:

* `name` - Address name.

The `poolname` block supports:

* `name` - IP pool name.


## Attribute Reference

In addition to all the above arguments, the following attributes are exported:
* `id` - an identifier for the resource with format {{policyid}}.

## Import

Firewall Policy64 can be imported using any of these accepted formats:
```
$ export "FORTIOS_IMPORT_TABLE"="true"
$ terraform import fortios_firewall_policy64.labelname {{policyid}}
$ unset "FORTIOS_IMPORT_TABLE"
```
