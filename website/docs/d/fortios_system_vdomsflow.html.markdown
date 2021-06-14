---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomsflow"
description: |-
  Get information on fortios system vdomsflow.
---

# Data Source: fortios_system_vdomsflow
Use this data source to get information on fortios system vdomsflow

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `vdom_sflow` - Enable/disable the sFlow configuration for the current VDOM.
* `collector_ip` - IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
* `collector_port` - UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
* `source_ip` - Source IP address for sFlow agent.

