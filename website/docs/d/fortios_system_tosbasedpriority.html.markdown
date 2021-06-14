---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_tosbasedpriority"
description: |-
  Get information on an fortios system tosbasedpriority.
---

# Data Source: fortios_system_tosbasedpriority
Use this data source to get information on an fortios system tosbasedpriority

## Argument Reference

* `fosid` - (Required) Specify the fosid of the desired system tosbasedpriority.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fosid` - Item ID.
* `tos` - Value of the ToS byte in the IP datagram header (0-15, 8: minimize delay, 4: maximize throughput, 2: maximize reliability, 1: minimize monetary cost, and 0: default service).
* `priority` - ToS based priority level to low, medium or high (these priorities match firewall traffic shaping priorities) (default = medium).

