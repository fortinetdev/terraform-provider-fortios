---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_vdomnetflow"
description: |-
  Get information on fortios system vdomnetflow.
---

# Data Source: fortios_system_vdomnetflow
Use this data source to get information on fortios system vdomnetflow

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `vdom_netflow` - Enable/disable NetFlow per VDOM.
* `collectors` - Netflow collectors. The structure of `collectors` block is documented below.
* `collector_ip` - NetFlow collector IP address.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

The `collectors` block contains:

* `id` - ID.
* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `source_ip_interface` - Name of the interface used to determine the source IP for exporting packets.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

