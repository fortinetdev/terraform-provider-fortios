---
subcategory: "FortiGate Firewall"
layout: "fortios"
page_title: "FortiOS: fortios_firewall_ipv6ehfilter"
description: |-
  Get information on fortios firewall ipv6ehfilter.
---

# Data Source: fortios_firewall_ipv6ehfilter
Use this data source to get information on fortios firewall ipv6ehfilter

## Argument Reference

No arguments available for the data source.

## Attribute Reference

The following attributes are exported:

* `hop_opt` - Enable/disable blocking packets with the Hop-by-Hop Options header (default = disable).
* `dest_opt` - Enable/disable blocking packets with Destination Options headers (default = disable).
* `hdopt_type` - Block specific Hop-by-Hop and/or Destination Option types (max. 7 types, each between 0 and 255, default = 0).
* `routing` - Enable/disable blocking packets with Routing headers (default = enable).
* `routing_type` - Block specific Routing header types (max. 7 types, each between 0 and 255, default =  0).
* `fragment` - Enable/disable blocking packets with the Fragment header (default = disable).
* `auth` - Enable/disable blocking packets with the Authentication header (default = disable).
* `no_next` - Enable/disable blocking packets with the No Next header (default = disable)

