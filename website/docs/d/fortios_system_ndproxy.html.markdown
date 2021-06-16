---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_ndproxy"
description: |-
  Get information on fortios system ndproxy.
---

# Data Source: fortios_system_ndproxy
Use this data source to get information on fortios system ndproxy

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable neighbor discovery proxy.
* `member` - Interfaces using the neighbor discovery proxy. The structure of `member` block is documented below.

The `member` block contains:

* `interface_name` - Interface name.

