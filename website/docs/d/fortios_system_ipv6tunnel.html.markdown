---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipv6tunnel"
description: |-
  Get information on an fortios system ipv6tunnel.
---

# Data Source: fortios_system_ipv6tunnel
Use this data source to get information on an fortios system ipv6tunnel

## Argument Reference

* `name` - (Required) Specify the name of the desired system ipv6tunnel.

## Attribute Reference

The following attributes are exported:

* `name` - IPv6 tunnel name.
* `source` - Local IPv6 address of the tunnel.
* `destination` - Remote IPv6 address of the tunnel.
* `interface` - Interface name.

