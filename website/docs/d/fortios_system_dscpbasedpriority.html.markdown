---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_dscpbasedpriority"
description: |-
  Get information on an fortios system dscpbasedpriority.
---

# Data Source: fortios_system_dscpbasedpriority
Use this data source to get information on an fortios system dscpbasedpriority

## Argument Reference

* `fosid` - (Required) Specify the fosid of the desired system dscpbasedpriority.
* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `fosid` - Item ID.
* `ds` - DSCP(DiffServ) DS value (0 - 63).
* `priority` - DSCP based priority level.

