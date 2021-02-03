---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_proxyarp"
description: |-
  Get information on an fortios system proxyarp.
---

# Data Source: fortios_system_proxyarp
Use this data source to get information on an fortios system proxyarp

## Argument Reference

* `fosid` - (Required) Specify the fosid of the desired system proxyarp.

## Attribute Reference

The following attributes are exported:

* `fosid` - Unique integer ID of the entry.
* `interface` - Interface acting proxy-ARP.
* `ip` - IP address or start IP to be proxied.
* `end_ip` - End IP of IP range to be proxied.

