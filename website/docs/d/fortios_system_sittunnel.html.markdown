---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sittunnel"
description: |-
  Get information on an fortios system sittunnel.
---

# Data Source: fortios_system_sittunnel
Use this data source to get information on an fortios system sittunnel

## Argument Reference

* `name` - (Required) Specify the name of the desired system sittunnel.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Tunnel name.
* `source` - Source IP address of the tunnel.
* `destination` - Destination IP address of the tunnel.
* `ip6` - IPv6 address of the tunnel.
* `interface` - Interface name.

