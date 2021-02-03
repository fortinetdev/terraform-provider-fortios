---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_sflow"
description: |-
  Get information on fortios system sflow.
---

# Data Source: fortios_system_sflow
Use this data source to get information on fortios system sflow

## Argument Reference

No arguments available for the data source.

## Attribute Reference

The following attributes are exported:

* `collector_ip` - IP address of the sFlow collector that sFlow agents added to interfaces in this VDOM send sFlow datagrams to (default = 0.0.0.0).
* `collector_port` - UDP port number used for sending sFlow datagrams (configure only if required by your sFlow collector or your network configuration) (0 - 65535, default = 6343).
* `source_ip` - Source IP address for sFlow agent.

