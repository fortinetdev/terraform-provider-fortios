---
subcategory: "FortiGate System"
layout: "fortios"
page_title: "FortiOS: fortios_system_proberesponse"
description: |-
  Get information on fortios system proberesponse.
---

# Data Source: fortios_system_proberesponse
Use this data source to get information on fortios system proberesponse

## Argument Reference


* `vdomparam` - Specifies the vdom to which the data source will be applied when the FortiGate unit is running in VDOM mode. Only one vdom can be specified. If you want to inherit the vdom configuration of the provider, please do not set this parameter.


## Attribute Reference

The following attributes are exported:

* `port` - Port number to response.
* `http_probe_value` - Value to respond to the monitoring server.
* `ttl_mode` - Mode for TWAMP packet TTL modification.
* `mode` - SLA response mode.
* `security_mode` - Twamp respondor security mode.
* `password` - Twamp respondor password in authentication mode
* `timeout` - An inactivity timer for a twamp test session.

