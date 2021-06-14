---
subcategory: "FortiGate Router"
layout: "fortios"
page_title: "FortiOS: fortios_router_authpath"
description: |-
  Get information on an fortios router authpath.
---

# Data Source: fortios_router_authpath
Use this data source to get information on an fortios router authpath

## Argument Reference

* `name` - (Required) Specify the name of the desired router authpath.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `name` - Name of the entry.
* `device` - Outgoing interface.
* `gateway` - Gateway IP address.

