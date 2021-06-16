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


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


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

