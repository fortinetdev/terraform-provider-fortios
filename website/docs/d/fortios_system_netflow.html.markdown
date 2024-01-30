---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_netflow"
description: |-
  Get information on fortios system netflow.
---

# Data Source: fortios_system_netflow
Use this data source to get information on fortios system netflow

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `active_flow_timeout` - Timeout to report active flows (1 - 60 min, default = 30).
* `inactive_flow_timeout` - Timeout for periodic report of finished flows (10 - 600 sec, default = 15).
* `template_tx_timeout` - Timeout for periodic template flowset transmission (1 - 1440 min, default = 30).
* `template_tx_counter` - Counter of flowset records before resending a template flowset record.
* `collectors` - Netflow collectors. The structure of `collectors` block is documented below.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

The `collectors` block contains:

* `id` - ID.
* `collector_ip` - Collector IP.
* `collector_port` - NetFlow collector port number.
* `source_ip` - Source IP address for communication with the NetFlow agent.
* `interface_select_method` - Specify how to select outgoing interface to reach server.
* `interface` - Specify outgoing interface to reach server.

