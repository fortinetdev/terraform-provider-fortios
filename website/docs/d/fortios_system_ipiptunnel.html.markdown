---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ipiptunnel"
description: |-
  Get information on an fortios system ipiptunnel.
---

# Data Source: fortios_system_ipiptunnel
Use this data source to get information on an fortios system ipiptunnel

## Argument Reference

* `name` - (Required) Specify the name of the desired system ipiptunnel.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - IPIP Tunnel name.
* `interface` - Interface name that is associated with the incoming traffic from available options.
* `remote_gw` - IPv4 address for the remote gateway.
* `local_gw` - IPv4 address for the local gateway.
* `use_sdwan` - Enable/disable use of SD-WAN to reach remote gateway.
* `auto_asic_offload` - Enable/disable tunnel ASIC offloading.

