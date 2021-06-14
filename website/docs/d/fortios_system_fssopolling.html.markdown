---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_fssopolling"
description: |-
  Get information on fortios system fssopolling.
---

# Data Source: fortios_system_fssopolling
Use this data source to get information on fortios system fssopolling

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `status` - Enable/disable FSSO Polling Mode.
* `listening_port` - Listening port to accept clients (1 - 65535).
* `authentication` - Enable/disable FSSO Agent Authentication.
* `auth_password` - Password to connect to FSSO Agent.

