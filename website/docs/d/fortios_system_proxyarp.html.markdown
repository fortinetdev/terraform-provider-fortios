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
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fosid` - Unique integer ID of the entry.
* `interface` - Interface acting proxy-ARP.
* `ip` - IP address or start IP to be proxied.
* `end_ip` - End IP of IP range to be proxied.

