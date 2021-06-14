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
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - IPv6 tunnel name.
* `source` - Local IPv6 address of the tunnel.
* `destination` - Remote IPv6 address of the tunnel.
* `interface` - Interface name.

