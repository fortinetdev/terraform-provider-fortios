---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_policy64"
description: |-
  Get information on an fortios firewall policy64.
---

# Data Source: fortios_firewall_policy64
Use this data source to get information on an fortios firewall policy64

## Argument Reference

* `policyid` - (Required) Specify the policyid of the desired firewall policy64.

## Attribute Reference

The following attributes are exported:

* `policyid` - Policy ID.
* `uuid` - Universally Unique Identifier (UUID; automatically assigned but can be manually reset).
* `srcintf` - Source interface name.
* `dstintf` - Destination interface name.
* `srcaddr` - Source address name. The structure of `srcaddr` block is documented below.
* `dstaddr` - Destination address name. The structure of `dstaddr` block is documented below.
* `action` - Policy action.
* `status` - Enable/disable policy status.
* `schedule` - Schedule name.
* `service` - Service name. The structure of `service` block is documented below.
* `logtraffic` - Enable/disable policy log traffic.
* `permit_any_host` - Enable/disable permit any host in.
* `traffic_shaper` - Traffic shaper.
* `traffic_shaper_reverse` - Reverse traffic shaper.
* `per_ip_shaper` - Per-IP traffic shaper.
* `fixedport` - Enable/disable policy fixed port.
* `ippool` - Enable/disable policy64 IP pool.
* `poolname` - Policy IP pool names. The structure of `poolname` block is documented below.
* `tcp_mss_sender` - TCP MSS value of sender.
* `tcp_mss_receiver` - TCP MSS value of receiver.
* `comments` - Comment.

The `srcaddr` block contains:

* `name` - Address name.

The `dstaddr` block contains:

* `name` - Address name.

The `service` block contains:

* `name` - Address name.

The `poolname` block contains:

* `name` - IP pool name.

